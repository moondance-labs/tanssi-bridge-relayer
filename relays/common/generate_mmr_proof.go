package common

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/snowfork/go-substrate-rpc-client/v4/client"
	"github.com/snowfork/go-substrate-rpc-client/v4/types"
)

type GenerateMMRProofResponseForSolochain struct {
	BlockHash types.H256
	Leaf      MMRLeafSolochain
	Proof     types.MMRProof
}

// UnmarshalJSON fills d with the JSON encoded byte array given by b
func (d *GenerateMMRProofResponseForSolochain) UnmarshalJSON(bz []byte) error {
	var tmp struct {
		BlockHash string `json:"blockHash"`
		Leaves    string `json:"leaves"`
		Proof     string `json:"proof"`
	}
	if err := json.Unmarshal(bz, &tmp); err != nil {
		return err
	}
	err := types.DecodeFromHexString(tmp.BlockHash, &d.BlockHash)
	if err != nil {
		return err
	}
	var encodedLeaf []types.MMREncodableOpaqueLeaf
	err = types.DecodeFromHexString(tmp.Leaves, &encodedLeaf)
	if err != nil {
		return err
	}
	if len(encodedLeaf) == 0 {
		return fmt.Errorf("decode leaf error")
	}

	err = types.DecodeFromBytes(encodedLeaf[0], &d.Leaf)
	if err != nil {
		return err
	}
	var proof types.MultiMMRProof
	err = types.DecodeFromHexString(tmp.Proof, &proof)
	if err != nil {
		return err
	}
	if proof.LeafIndices == nil || len(proof.LeafIndices) == 0 {
		return fmt.Errorf("decode proof LeafIndices error")
	}
	d.Proof.LeafCount = proof.LeafCount
	d.Proof.Items = proof.Items
	d.Proof.LeafIndex = proof.LeafIndices[0]
	return nil
}

type MMRLeafSolochain struct {
	types.MMRLeaf
	CommitmentRoot types.H256
}

func GenerateMMRProofForSolochain(c client.Client, blockNumber uint32, blockHash types.Hash) (GenerateMMRProofResponseForSolochain, error) {
	var proofResponse GenerateMMRProofResponseForSolochain
	blocks := [1]uint32{blockNumber}
	err := client.CallWithBlockHash(c, &proofResponse, "mmr_generateProof", &blockHash, blocks, nil)
	if err != nil {
		return GenerateMMRProofResponseForSolochain{}, err
	}

	return proofResponse, nil
}

func GenerateProofForBlock(
	c client.Client,
	blockNumber uint64,
	latestBeefyBlockHash types.Hash,
) (GenerateMMRProofResponseForSolochain, error) {
	log.WithFields(log.Fields{
		"blockNumber": blockNumber,
		"blockHash":   latestBeefyBlockHash.Hex(),
	}).Debug("Getting MMR Leaf for block...")

	proofResponse, err := GenerateMMRProofForSolochain(c, uint32(blockNumber), latestBeefyBlockHash)
	if err != nil {
		return GenerateMMRProofResponseForSolochain{}, err
	}

	var proofItemsHex = []string{}
	for _, item := range proofResponse.Proof.Items {
		proofItemsHex = append(proofItemsHex, item.Hex())
	}

	log.WithFields(log.Fields{
		"BlockHash": proofResponse.BlockHash.Hex(),
		"Leaf": log.Fields{
			"ParentNumber":   proofResponse.Leaf.ParentNumberAndHash.ParentNumber,
			"ParentHash":     proofResponse.Leaf.ParentNumberAndHash.Hash.Hex(),
			"ParachainHeads": proofResponse.Leaf.ParachainHeads.Hex(),
			"NextAuthoritySet": log.Fields{
				"Id":   proofResponse.Leaf.BeefyNextAuthoritySet.ID,
				"Len":  proofResponse.Leaf.BeefyNextAuthoritySet.Len,
				"Root": proofResponse.Leaf.BeefyNextAuthoritySet.Root.Hex(),
			},
		},
		"Proof": log.Fields{
			"LeafIndex": proofResponse.Proof.LeafIndex,
			"LeafCount": proofResponse.Proof.LeafCount,
			"Items":     proofItemsHex,
		},
	}).Debug("Generated MMR proof")

	return proofResponse, nil
}
