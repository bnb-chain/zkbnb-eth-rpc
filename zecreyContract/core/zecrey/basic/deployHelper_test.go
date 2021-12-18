package basic

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zecrey-labs/zecrey-eth-rpc/_const"
	"github.com/zecrey-labs/zecrey-eth-rpc/zecreyContract/core/zecrey/aurora"
	"github.com/zecrey-labs/zecrey-eth-rpc/zecreyContract/core/zecrey/avalanche"
	"github.com/zecrey-labs/zecrey-eth-rpc/zecreyContract/core/zecrey/bsc"
	"github.com/zecrey-labs/zecrey-eth-rpc/zecreyContract/core/zecrey/eth"
	"github.com/zecrey-labs/zecrey-eth-rpc/zecreyContract/core/zecrey/localchain"
	"github.com/zecrey-labs/zecrey-eth-rpc/zecreyContract/core/zecrey/polygon"
	"io/fs"
	"io/ioutil"
	"math/big"
	"testing"
	"time"
)

func TestFirstDeployment_Local(t *testing.T) {
	cli, authCli, err := ConstructCliAndAuthCli(localchain.ChainId, localchain.NetworkEndPoint, SuperSk)
	if err != nil {
		t.Fatal(err)
	}
	waitTime := 2 * time.Second
	governor := SuperAddress
	l2ChainId := localchain.NativeChainId
	nativeAssetId := localchain.NativeAssetId
	maxPendingBlocks := localchain.MaxPendingBlocks
	listingFeeTokenAddr := governor
	listingFee := big.NewInt(10000000000)
	listingCap := uint16(10000)
	treasuryAddr := governor
	genesisBlockNumber := uint32(0)
	genesisOnchainOpsRoot := common.Hex2Bytes("003064dc24ffc11d9b89087dcc6a4a8aa1f11951c5c8fcd77d44d649c8cad658")
	genesisStateRoot := common.Hex2Bytes("0046cc1da34dce9856372bea41b7fecdad02c7b73848fe7d73b2872b8ab931b4")
	genesisTimestamp := big.NewInt(0)
	genesisCommitment := common.Hex2Bytes("0046cc1da34dce9856372bea41b7fecdad02c7b73848fe7d73b2872b8ab931b4")
	genesisMerkleHelper := localchain.OnChainOpsMerkleHelper
	SuggestGasPrice, _ := cli.SuggestGasPrice(context.Background())
	gasPrice := SuggestGasPrice
	gasLimit := _const.SuggestHighGasLimit
	contracts, err := FirstDeployment(
		cli, authCli, waitTime,
		governor,
		l2ChainId, nativeAssetId, maxPendingBlocks,
		listingFeeTokenAddr, listingFee, listingCap,
		treasuryAddr,
		genesisBlockNumber,
		genesisOnchainOpsRoot,
		genesisStateRoot,
		genesisTimestamp,
		genesisCommitment, genesisMerkleHelper,
		gasPrice, gasLimit,
	)
	if err != nil {
		t.Fatal(err)
	}
	contractsBytes, err := json.Marshal(contracts)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile(localchain.Dir+ContractInfoFileName, contractsBytes, fs.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFirstDeployment_Ethereum(t *testing.T) {
	cli, authCli, err := ConstructCliAndAuthCli(eth.ChainId, eth.NetworkEndPoint, SuperSk)
	if err != nil {
		t.Fatal(err)
	}
	waitTime := 30 * time.Second
	governor := SuperAddress
	l2ChainId := eth.NativeChainId
	nativeAssetId := eth.NativeAssetId
	maxPendingBlocks := eth.MaxPendingBlocks
	listingFeeTokenAddr := governor
	listingFee := big.NewInt(10000000000)
	listingCap := uint16(10000)
	treasuryAddr := governor
	genesisBlockNumber := uint32(0)
	genesisOnchainOpsRoot := common.Hex2Bytes("003064dc24ffc11d9b89087dcc6a4a8aa1f11951c5c8fcd77d44d649c8cad658")
	genesisStateRoot := common.Hex2Bytes("0046cc1da34dce9856372bea41b7fecdad02c7b73848fe7d73b2872b8ab931b4")
	genesisTimestamp := big.NewInt(0)
	genesisCommitment := common.Hex2Bytes("0046cc1da34dce9856372bea41b7fecdad02c7b73848fe7d73b2872b8ab931b4")
	genesisMerkleHelper := eth.OnChainOpsMerkleHelper
	SuggestGasPrice, _ := cli.SuggestGasPrice(context.Background())
	gasPrice := SuggestGasPrice
	gasLimit := _const.SuggestHighGasLimit
	contracts, err := FirstDeployment(
		cli, authCli, waitTime,
		governor,
		l2ChainId, nativeAssetId, maxPendingBlocks,
		listingFeeTokenAddr, listingFee, listingCap,
		treasuryAddr,
		genesisBlockNumber,
		genesisOnchainOpsRoot,
		genesisStateRoot,
		genesisTimestamp,
		genesisCommitment, genesisMerkleHelper,
		gasPrice, gasLimit,
	)
	if err != nil {
		t.Fatal(err)
	}
	contractsBytes, err := json.Marshal(contracts)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile(eth.Dir+ContractInfoFileName, contractsBytes, fs.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}


func TestFirstDeployment_Avalanche(t *testing.T) {
	cli, authCli, err := ConstructCliAndAuthCli(avalanche.ChainId, avalanche.NetworkEndPoint, SuperSk)
	if err != nil {
		t.Fatal(err)
	}
	waitTime := 10 * time.Second
	governor := SuperAddress
	l2ChainId := avalanche.NativeChainId
	nativeAssetId := avalanche.NativeAssetId
	maxPendingBlocks := avalanche.MaxPendingBlocks
	listingFeeTokenAddr := governor
	listingFee := big.NewInt(10000000000)
	listingCap := uint16(10000)
	treasuryAddr := governor
	genesisBlockNumber := uint32(0)
	genesisOnchainOpsRoot := common.Hex2Bytes("003064dc24ffc11d9b89087dcc6a4a8aa1f11951c5c8fcd77d44d649c8cad658")
	genesisStateRoot := common.Hex2Bytes("0046cc1da34dce9856372bea41b7fecdad02c7b73848fe7d73b2872b8ab931b4")
	genesisTimestamp := big.NewInt(0)
	genesisCommitment := common.Hex2Bytes("0046cc1da34dce9856372bea41b7fecdad02c7b73848fe7d73b2872b8ab931b4")
	genesisMerkleHelper := avalanche.OnChainOpsMerkleHelper
	SuggestGasPrice, _ := cli.SuggestGasPrice(context.Background())
	gasPrice := SuggestGasPrice
	gasLimit := _const.SuggestHighGasLimit
	contracts, err := FirstDeployment(
		cli, authCli, waitTime,
		governor,
		l2ChainId, nativeAssetId, maxPendingBlocks,
		listingFeeTokenAddr, listingFee, listingCap,
		treasuryAddr,
		genesisBlockNumber,
		genesisOnchainOpsRoot,
		genesisStateRoot,
		genesisTimestamp,
		genesisCommitment, genesisMerkleHelper,
		gasPrice, gasLimit,
	)
	if err != nil {
		t.Fatal(err)
	}
	contractsBytes, err := json.Marshal(contracts)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile(avalanche.Dir+ContractInfoFileName, contractsBytes, fs.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFirstDeployment_BSC(t *testing.T) {
	cli, authCli, err := ConstructCliAndAuthCli(bsc.ChainId, bsc.NetworkEndPoint, SuperSk)
	if err != nil {
		t.Fatal(err)
	}
	waitTime := 10 * time.Second
	governor := SuperAddress
	l2ChainId := bsc.NativeChainId
	nativeAssetId := bsc.NativeAssetId
	maxPendingBlocks := bsc.MaxPendingBlocks
	listingFeeTokenAddr := governor
	listingFee := big.NewInt(10000000000)
	listingCap := uint16(10000)
	treasuryAddr := governor
	genesisBlockNumber := uint32(0)
	genesisOnchainOpsRoot := common.Hex2Bytes("003064dc24ffc11d9b89087dcc6a4a8aa1f11951c5c8fcd77d44d649c8cad658")
	genesisStateRoot := common.Hex2Bytes("0046cc1da34dce9856372bea41b7fecdad02c7b73848fe7d73b2872b8ab931b4")
	genesisTimestamp := big.NewInt(0)
	genesisCommitment := common.Hex2Bytes("0046cc1da34dce9856372bea41b7fecdad02c7b73848fe7d73b2872b8ab931b4")
	genesisMerkleHelper := bsc.OnChainOpsMerkleHelper
	SuggestGasPrice, _ := cli.SuggestGasPrice(context.Background())
	gasPrice := SuggestGasPrice
	gasLimit := _const.SuggestHighGasLimit
	contracts, err := FirstDeployment(
		cli, authCli, waitTime,
		governor,
		l2ChainId, nativeAssetId, maxPendingBlocks,
		listingFeeTokenAddr, listingFee, listingCap,
		treasuryAddr,
		genesisBlockNumber,
		genesisOnchainOpsRoot,
		genesisStateRoot,
		genesisTimestamp,
		genesisCommitment, genesisMerkleHelper,
		gasPrice, gasLimit,
	)
	if err != nil {
		t.Fatal(err)
	}
	contractsBytes, err := json.Marshal(contracts)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile(bsc.Dir+ContractInfoFileName, contractsBytes, fs.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFirstDeployment_Polygon(t *testing.T) {
	cli, authCli, err := ConstructCliAndAuthCli(polygon.ChainId, polygon.NetworkEndPoint, SuperSk)
	if err != nil {
		t.Fatal(err)
	}
	waitTime := 5 * time.Second
	governor := SuperAddress
	l2ChainId := polygon.NativeChainId
	nativeAssetId := polygon.NativeAssetId
	maxPendingBlocks := polygon.MaxPendingBlocks
	listingFeeTokenAddr := governor
	listingFee := big.NewInt(10000000000)
	listingCap := uint16(10000)
	treasuryAddr := governor
	genesisBlockNumber := uint32(0)
	genesisOnchainOpsRoot := common.Hex2Bytes("003064dc24ffc11d9b89087dcc6a4a8aa1f11951c5c8fcd77d44d649c8cad658")
	genesisStateRoot := common.Hex2Bytes("0046cc1da34dce9856372bea41b7fecdad02c7b73848fe7d73b2872b8ab931b4")
	genesisTimestamp := big.NewInt(0)
	genesisCommitment := common.Hex2Bytes("0046cc1da34dce9856372bea41b7fecdad02c7b73848fe7d73b2872b8ab931b4")
	genesisMerkleHelper := polygon.OnChainOpsMerkleHelper
	SuggestGasPrice, _ := cli.SuggestGasPrice(context.Background())
	gasPrice := SuggestGasPrice
	gasLimit := _const.SuggestHighGasLimit
	contracts, err := FirstDeployment(
		cli, authCli, waitTime,
		governor,
		l2ChainId, nativeAssetId, maxPendingBlocks,
		listingFeeTokenAddr, listingFee, listingCap,
		treasuryAddr,
		genesisBlockNumber,
		genesisOnchainOpsRoot,
		genesisStateRoot,
		genesisTimestamp,
		genesisCommitment, genesisMerkleHelper,
		gasPrice, gasLimit,
	)
	if err != nil {
		t.Fatal(err)
	}
	contractsBytes, err := json.Marshal(contracts)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile(polygon.Dir+ContractInfoFileName, contractsBytes, fs.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFirstDeployment_Aurora(t *testing.T) {
	cli, authCli, err := ConstructCliAndAuthCli(aurora.ChainId, aurora.NetworkEndPoint, SuperSk)
	if err != nil {
		t.Fatal(err)
	}
	waitTime := 10 * time.Second
	governor := SuperAddress
	l2ChainId := aurora.NativeChainId
	nativeAssetId := aurora.NativeAssetId
	maxPendingBlocks := aurora.MaxPendingBlocks
	listingFeeTokenAddr := governor
	listingFee := big.NewInt(10000000000)
	listingCap := uint16(10000)
	treasuryAddr := governor
	genesisBlockNumber := uint32(0)
	genesisOnchainOpsRoot := common.Hex2Bytes("003064dc24ffc11d9b89087dcc6a4a8aa1f11951c5c8fcd77d44d649c8cad658")
	genesisStateRoot := common.Hex2Bytes("0046cc1da34dce9856372bea41b7fecdad02c7b73848fe7d73b2872b8ab931b4")
	genesisTimestamp := big.NewInt(0)
	genesisCommitment := common.Hex2Bytes("0046cc1da34dce9856372bea41b7fecdad02c7b73848fe7d73b2872b8ab931b4")
	genesisMerkleHelper := aurora.OnChainOpsMerkleHelper
	SuggestGasPrice, _ := cli.SuggestGasPrice(context.Background())
	gasPrice := SuggestGasPrice
	gasLimit := _const.SuggestHighGasLimit
	contracts, err := FirstDeployment(
		cli, authCli, waitTime,
		governor,
		l2ChainId, nativeAssetId, maxPendingBlocks,
		listingFeeTokenAddr, listingFee, listingCap,
		treasuryAddr,
		genesisBlockNumber,
		genesisOnchainOpsRoot,
		genesisStateRoot,
		genesisTimestamp,
		genesisCommitment, genesisMerkleHelper,
		gasPrice, gasLimit,
	)
	if err != nil {
		t.Fatal(err)
	}
	contractsBytes, err := json.Marshal(contracts)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile(aurora.Dir+ContractInfoFileName, contractsBytes, fs.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}
