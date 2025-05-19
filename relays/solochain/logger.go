package solochain

import (
	"fmt"

    // TODO: This reference should go to snowbridge/relayer/contracts
	"github.com/moondance-labs/tanssi-bridge-relayer/relays/contracts"
	log "github.com/sirupsen/logrus"
	"github.com/snowfork/go-substrate-rpc-client/v4/types"
)

func Hex(b []byte) string {
	return types.HexEncodeToString(b)
}

func (wr *EthereumWriter) logFieldsForSubmission(
	message contracts.InboundMessage,
	messageProof [][32]byte,
	proof contracts.VerificationProof,
) log.Fields {
	messageProofHexes := make([]string, len(messageProof))
	for i, proof := range messageProof {
		messageProofHexes[i] = Hex(proof[:])
	}

	mmrLeafProofHexes := make([]string, len(proof.LeafProof))
	for i, proof := range proof.LeafProof {
		mmrLeafProofHexes[i] = Hex(proof[:])
	}

	params := log.Fields{
		"message": log.Fields{
			"channelID": Hex(message.ChannelID[:]),
			"nonce":     message.Nonce,
			"command":   message.Command,
			"params":    Hex(message.Params),
		},
		"messageProof": messageProofHexes,
		"proof": log.Fields{
			"leafPartial": log.Fields{
				"version":              proof.LeafPartial.Version,
				"parentNumber":         proof.LeafPartial.ParentNumber,
				"parentHash":           Hex(proof.LeafPartial.ParentHash[:]),
				"nextAuthoritySetID":   proof.LeafPartial.NextAuthoritySetID,
				"nextAuthoritySetLen":  proof.LeafPartial.NextAuthoritySetLen,
				"nextAuthoritySetRoot": Hex(proof.LeafPartial.NextAuthoritySetRoot[:]),
			},
			"parachainHeadsroot": proof.ParachainHeadsRoot,
			"leafProof":          mmrLeafProofHexes,
			"leafProofOrder":     fmt.Sprintf("%b", proof.LeafProofOrder),
		},
	}

	return params
}
