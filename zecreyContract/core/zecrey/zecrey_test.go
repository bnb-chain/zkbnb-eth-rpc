package zecrey

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zecrey-labs/zecrey-eth-rpc/_const"
	"github.com/zecrey-labs/zecrey-eth-rpc/_utils"
	"math/big"
	"testing"
	"time"
)

const ZecreyAddr = "0x6d551E6b91f94c422Db90786A8c9Cd9eDBcbb911"
const RinkebyZecreyAddr = "0x95bBA8D93C794717eFF37d37d5F5d6573e713321"
const AuroraZecreyAddr = "0x56C4e158658551d9C20AB430A5d09fb4d4f0DDa8"
const ArbitrumZecreyAddr = "0x56C4e158658551d9C20AB430A5d09fb4d4f0DDa8"
const PolyTestZecreyAddr = "0x56C4e158658551d9C20AB430A5d09fb4d4f0DDa8"

func TestDeployZecreyContract(t *testing.T) {
	elapse := time.Now()
	//addr, txHash, err := DeployZecreyContract(localCli, localAuthCli, ZecreyVerifierAddr, GovernanceAddr, localGasPrice, _const.SuggestHighGasLimit)
	addr, txHash, err := DeployZecreyContract(rinkebyCli, rinkebyAuthCli, RinkebyZecreyVerifierAddr, RinkebyGovernanceAddr, localGasPrice, _const.SuggestHighGasLimit)
	//addr, txHash, err := DeployZecreyContract(polyTestCli, polyTestAuthCli, PolyTestVerifierAddr, PolyTestNetGovernanceAddr, big.NewInt(1000000000), _const.SuggestHighGasLimit)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(time.Since(elapse))
	fmt.Println("contract address:", addr)
	fmt.Println("deploy contract tx hash:", txHash)
}

func LoadZecrey() *Zecrey {
	instance, _ := LoadZecreyInstance(polyTestCli, PolyTestZecreyAddr)
	return instance
}

func TestDeposit(t *testing.T) {
	instance := LoadZecrey()
	depositAmount, _ := new(big.Int).SetString("10000000000000000", 10)
	oldBalance, err := polyTestCli.GetBalance(_const.RinkebySuperAddress)
	if err != nil {
		t.Fatal(err)
	}
	info, err := polyTestCli.GetHeight()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("current height:", info)
	fmt.Println(_utils.WeiToEther(oldBalance).String())
	txHash, err := Deposit(polyTestCli, polyTestAuthCli, instance, 0, "name4.zecrey", depositAmount, polyGasPrice, _const.SuggestHighGasLimit)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(txHash)
	newBalance, err := polyTestCli.GetBalance(_const.RinkebySuperAddress)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(_utils.WeiToEther(newBalance).String())
}

func TestCommitBlocks(t *testing.T) {
	instance := LoadZecrey()
	block := ConstructStorageBlock(0,
		common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000000"),
		common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000000"),
		0)
	withdrawAmount, _ := new(big.Int).SetString("1000000000000000000", 10)
	withdrawOperation := ConstructOperation(5, 0, "0x89D37ea8a0f102D90C424141F897A6a764A291AF", "sher", withdrawAmount)
	withdrawOperations := []ZecreyOperation{withdrawOperation}
	commitBlock := ConstructCommitBlock(1,
		common.FromHex("0x9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"),
		common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000000"),
		withdrawOperations,
		1625562034)
	commitBlocks := []ZecreyCommitBlock{commitBlock}
	txHash, err := CommitBlocks(localCli, localAuthCli, instance, block, commitBlocks, localGasPrice, _const.SuggestHighGasLimit)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(txHash)
	committedCount, err := GetTotalBlocksCommitted(localCli, instance)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(committedCount)
}

func TestProcessBlocks(t *testing.T) {
	instance := LoadZecrey()
	oldBalance, err := localCli.GetBalance(_const.LocalSuperAddress)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(_utils.WeiToEther(oldBalance).String())
	block := ConstructStorageBlock(1,
		common.FromHex("0x9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"),
		common.FromHex("0x19891c8615b4fc36ddb9ce3ef00e5440bb4849c6617cd1bd5083987dc248cdde"),
		1625562034)
	b := [2][2]*big.Int{}
	b[0][0] = big.NewInt(0)
	b[0][1] = big.NewInt(0)
	b[1][0] = big.NewInt(0)
	b[1][1] = big.NewInt(0)
	processBlock := ConstructProcessBlock(block,
		[2]*big.Int{big.NewInt(0), big.NewInt(0)},
		b,
		[2]*big.Int{big.NewInt(0),
			big.NewInt(0)})
	processBlocks := []ZecreyProcessBlock{processBlock}
	txHash, err := ProcessBlocks(localCli, localAuthCli, instance, processBlocks, localGasPrice, _const.SuggestHighGasLimit)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(txHash)
	processedCount, err := GetTotalBlocksProcessed(localCli, instance)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(processedCount)
	newBalance, err := localCli.GetBalance(_const.LocalSuperAddress)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(_utils.WeiToEther(newBalance).String())
}

func TestLoadLogs(t *testing.T) {
	instance := LoadZecrey()
	depositLogs, err := instance.ZecreyFilterer.FilterDeposit(&bind.FilterOpts{Start: 641, End: nil, Context: context.Background()})
	if err != nil {
		t.Fatal(err)
	}
	for depositLogs.Next() {
		fmt.Println(depositLogs.Event.ZecreyAddr)
	}
}

func TestGetTotalBlocksCommitted(t *testing.T) {
	instance := LoadZecrey()
	blocksCommitted, err := GetTotalBlocksCommitted(localCli, instance)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(blocksCommitted)
}

func TestCommittedHash(t *testing.T) {
	instance := LoadZecrey()
	blockHashes, err := instance.StoredBlockHashes(EmptyCallOpts(), 6)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(common.BytesToHash(blockHashes[:]))
}

func TestZecreyCaller_PendingOperationsCount(t *testing.T) {
	instance := LoadZecrey()
	count, err := instance.PendingOperationsCount(EmptyCallOpts())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(count.String())
}

func TestZecreyCaller_OnchainOperationsCount(t *testing.T) {
	instance := LoadZecrey()
	count, err := instance.OnchainOperationsCount(EmptyCallOpts())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(count.String())
}

func TestZecreyCaller_OnchainOperations(t *testing.T) {
	instance := LoadZecrey()
	operations, err := instance.OnchainOperations(EmptyCallOpts(), big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(operations.AssetId)
}

func TestConstructDepositSignedTx(t *testing.T) {
	nonce, err := localCli.GetPendingNonce(_const.LocalSuperAddress)
	if err != nil {
		t.Fatal(err)
	}
	amount := _utils.EtherToWei(0.1)
	gasPrice, err := localCli.SuggestGasPrice(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	txStr, err := ConstructDepositSignedTx(_const.LocalSuperSk, localAuthCli.ChainId.Int64(), nonce, ZecreyAddr, 0, "sher", amount, gasPrice, _const.SuggestHighGasLimit)
	if err != nil {
		t.Fatal(err)
	}
	signedTx, err := DecodeSignedTx(txStr)
	if err != nil {
		t.Fatal(err)
	}
	err = localCli.SendTransaction(context.Background(), signedTx)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("txHash:", signedTx.Hash().Hex())
}

func TestGetBlockHeight(t *testing.T) {
	height, err := arbitrumCli.GetHeight()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(height)
}
