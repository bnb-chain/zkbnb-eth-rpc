package basic

import (
	"github.com/zecrey-labs/zecrey-eth-rpc/zecreyContract/core/zecrey/localchain"
	"testing"
)

func TestZecreyComputeCommitment(t *testing.T) {
	cli, _, err := ConstructCliAndAuthCli(localchain.ChainId, localchain.NetworkEndPoint, localchain.SuperSk)
	if err != nil {
		t.Fatal(err)
	}
	instance, err := LoadZecreyInstance(cli, "0x66F06D938F90b8cc1604F6f2C4b4520cBDD23DCf")
	if err != nil {
		t.Fatal(err)
	}

	initHash, err := ZecreyGetStoredBlockHashesByHeight(instance, 0)
	if err != nil {
		t.Fatal(err)
	}
	ZecreyComputeCommitment(initHash, ZecreyCommitBlockInfo{
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
