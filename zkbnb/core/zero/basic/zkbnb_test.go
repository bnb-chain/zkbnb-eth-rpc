package basic

import (
	"github.com/bnb-chain/zkbnb-eth-rpc/zkbnb/core/zero/localchain"
	"testing"
)

func TestZkBNBComputeCommitment(t *testing.T) {
	cli, _, err := ConstructCliAndAuthCli(localchain.ChainId, localchain.NetworkEndPoint, localchain.SuperSk)
	if err != nil {
		t.Fatal(err)
	}
	instance, err := LoadZkBNBInstance(cli, "0x66F06D938F90b8cc1604F6f2C4b4520cBDD23DCf")
	if err != nil {
		t.Fatal(err)
	}

	initHash, err := ZkBNBGetStoredBlockHashesByHeight(instance, 0)
	if err != nil {
		t.Fatal(err)
	}
	ZkBNBComputeCommitment(initHash, ZkBNBCommitBlockInfo{
		BlockNumber:           0,
		OnchainOpsRoot:        [32]byte{},
		NewAccountRoot:        [32]byte{},
		Timestamp:             nil,
		Commitment:            [32]byte{},
		OnchainOpsPubData:     nil,
		OnchainOpsCount:       0,
		OnchainOpsMerkleProof: [7][32]byte{},
	})
}
