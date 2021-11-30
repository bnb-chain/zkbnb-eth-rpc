package localchain

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zecrey-labs/zecrey-eth-rpc/_const"
	"github.com/zecrey-labs/zecrey-eth-rpc/zecreyContract/core/zecrey/basic"
	"io/fs"
	"io/ioutil"
	"math/big"
	"testing"
)

func TestFirstDeployment(t *testing.T) {
	governor := "0x2F3b007175A6E9FC25e68bD97CF17ACac1e6ec9A"
	l2ChainId := NativeChainId
	nativeAssetId := NativeAssetId
	maxPendingBlocks := MaxPendingBlocks
	listingFeeTokenAddr := "0x2F3b007175A6E9FC25e68bD97CF17ACac1e6ec9A"
	listingFee := big.NewInt(10000000000)
	listingCap := uint16(10000)
	treasuryAddr := "0x2F3b007175A6E9FC25e68bD97CF17ACac1e6ec9A"
	genesisBlockNumber := uint32(0)
	genesisOnchainOpsRoot := common.Hex2Bytes("003064dc24ffc11d9b89087dcc6a4a8aa1f11951c5c8fcd77d44d649c8cad658")
	genesisStateRoot := common.Hex2Bytes("0046cc1da34dce9856372bea41b7fecdad02c7b73848fe7d73b2872b8ab931b4")
	genesisTimestamp := big.NewInt(0)
	genesisCommitment := common.Hex2Bytes("0046cc1da34dce9856372bea41b7fecdad02c7b73848fe7d73b2872b8ab931b4")
	genesisMerkleHelper := [OnChainOpsTreeHelperDepth]bool{false, false, false, false}
	gasPrice := SuggestGasPrice
	gasLimit := _const.SuggestHighGasLimit
	zecreyContracts, err := FirstDeployment(
		governor,
		l2ChainId, nativeAssetId, maxPendingBlocks,
		listingFeeTokenAddr, listingFee, listingCap,
		treasuryAddr,
		genesisBlockNumber, genesisOnchainOpsRoot, genesisStateRoot, genesisTimestamp,
		genesisCommitment, genesisMerkleHelper,
		gasPrice, gasLimit,
	)
	if err != nil {
		t.Fatal(err)
	}
	contractsBytes, err := json.Marshal(zecreyContracts)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile(AddrFileName, contractsBytes, fs.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	addrBytes, err := ioutil.ReadFile(AddrFileName)
	if err != nil {
		t.Fatal(err)
	}
	var addrs basic.ZecreyContracts
	err = json.Unmarshal(addrBytes, &addrs)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(addrs.ZecreyAddr)
}
