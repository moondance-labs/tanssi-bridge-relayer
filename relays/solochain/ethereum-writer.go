package solochain

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"golang.org/x/sync/errgroup"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/moondance-labs/tanssi-bridge-relayer/relays/contracts"
	"github.com/snowfork/snowbridge/relayer/chain/ethereum"
	"github.com/snowfork/snowbridge/relayer/crypto/keccak"

	gsrpcTypes "github.com/snowfork/go-substrate-rpc-client/v4/types"

	log "github.com/sirupsen/logrus"
)

type EthereumWriter struct {
	config     *SinkConfig
	conn       *ethereum.Connection
	gateway    *contracts.Gateway
	tasks      <-chan *Task
	gatewayABI abi.ABI
}

func NewEthereumWriter(
	config *SinkConfig,
	conn *ethereum.Connection,
	tasks <-chan *Task,
) (*EthereumWriter, error) {
	return &EthereumWriter{
		config:  config,
		conn:    conn,
		gateway: nil,
		tasks:   tasks,
	}, nil
}

func (wr *EthereumWriter) Start(ctx context.Context, eg *errgroup.Group) error {
	address := common.HexToAddress(wr.config.Contracts.Gateway)
	gateway, err := contracts.NewGateway(address, wr.conn.Client())
	if err != nil {
		return err
	}
	wr.gateway = gateway

	gatewayABI, err := abi.JSON(strings.NewReader(contracts.GatewayABI))
	if err != nil {
		return err
	}
	wr.gatewayABI = gatewayABI

	eg.Go(func() error {
		err := wr.writeMessagesLoop(ctx)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				return nil
			}
			return fmt.Errorf("write message loop: %w", err)
		}
		return nil
	})

	return nil
}

func (wr *EthereumWriter) writeMessagesLoop(ctx context.Context) error {
	options := wr.conn.MakeTxOpts(ctx)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case task, ok := <-wr.tasks:
			if !ok {
				return nil
			}
			err := wr.WriteChannels(ctx, options, task)
			if err != nil {
				return fmt.Errorf("write message: %w", err)
			}
		}
	}
}

func (wr *EthereumWriter) WriteChannels(
	ctx context.Context,
	options *bind.TransactOpts,
	task *Task,
) error {
	for _, proof := range *task.MessageProofs {
		err := wr.WriteChannel(ctx, options, &proof, task.ProofOutput)
		if err != nil {
			return fmt.Errorf("write eth gateway: %w", err)
		}
	}

	return nil
}

// Submit sends a SCALE-encoded message to an application deployed on the Ethereum network
func (wr *EthereumWriter) WriteChannel(
	ctx context.Context,
	options *bind.TransactOpts,
	commitmentProof *MessageProof,
	proof *ProofOutput,
) error {
	message := commitmentProof.Message.IntoInboundMessage()

	var merkleProofItems [][32]byte
	for _, proofItem := range proof.MMRProof.MerkleProofItems {
		merkleProofItems = append(merkleProofItems, proofItem)
	}

	verificationProof := contracts.VerificationProof{
		LeafPartial: contracts.VerificationMMRLeafPartial{
			Version:              uint8(proof.MMRProof.Leaf.Version),
			ParentNumber:         uint32(proof.MMRProof.Leaf.ParentNumberAndHash.ParentNumber),
			ParentHash:           proof.MMRProof.Leaf.ParentNumberAndHash.Hash,
			NextAuthoritySetID:   uint64(proof.MMRProof.Leaf.BeefyNextAuthoritySet.ID),
			NextAuthoritySetLen:  uint32(proof.MMRProof.Leaf.BeefyNextAuthoritySet.Len),
			NextAuthoritySetRoot: proof.MMRProof.Leaf.BeefyNextAuthoritySet.Root,
		},
		LeafProof:          merkleProofItems,
		ParachainHeadsRoot: proof.MMRProof.Leaf.ParachainHeads,
		LeafProofOrder:     new(big.Int).SetUint64(proof.MMRProof.MerkleProofOrder),
	}

	tx, err := wr.gateway.SubmitV1(
		options, message, commitmentProof.Proof.InnerHashes, verificationProof,
	)
	if err != nil {
		return fmt.Errorf("send transaction Gateway.submit: %w", err)
	}

	hasher := &keccak.Keccak256{}
	mmrLeafEncoded, err := gsrpcTypes.EncodeToBytes(proof.MMRProof.Leaf)
	if err != nil {
		return fmt.Errorf("encode MMRLeaf: %w", err)
	}
	log.WithField("txHash", tx.Hash().Hex()).
		WithField("params", wr.logFieldsForSubmission(message, commitmentProof.Proof.InnerHashes, verificationProof)).
		WithFields(log.Fields{
			"commitmentHash": commitmentProof.Proof.Root.Hex(),
			"MMRRoot":        proof.MMRRootHash.Hex(),
			"MMRLeafHash":    Hex(hasher.Hash(mmrLeafEncoded)),
			"beefyBlock":     proof.MMRProof.Blockhash.Hex(),
		}).
		Info("Sent transaction Gateway.submit")

	receipt, err := wr.conn.WatchTransaction(ctx, tx, 1)

	if err != nil {
		log.WithError(err).Error("Failed to SubmitInbound")
		return err
	}

	for _, ev := range receipt.Logs {
		if ev.Topics[0] == wr.gatewayABI.Events["InboundMessageDispatched"].ID {
			var holder contracts.GatewayInboundMessageDispatched
			err = wr.gatewayABI.UnpackIntoInterface(&holder, "InboundMessageDispatched", ev.Data)
			if err != nil {
				return fmt.Errorf("unpack event log: %w", err)
			}
			log.WithFields(log.Fields{
				"channelID": Hex(holder.ChannelID[:]),
				"nonce":     holder.Nonce,
				"success":   holder.Success,
			}).Info("Message dispatched")
		}
	}

	return nil
}
