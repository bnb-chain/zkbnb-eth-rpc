package zecrey

import (
	"Zecrey-eth-rpc/_const"
	"Zecrey-eth-rpc/_utils"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
	"time"
)

const ZecreyAddr = "0x18766657eC1Bf9DbFc0d229258f205d88D736c38"
const RinkebyZecreyAddr = "0x8D5ad3bCc10492286bDeDC0A992bEd99F1e5729c"

func TestDeployZecreyContract(t *testing.T) {
	elapse := time.Now()
	addr, txHash, err := DeployZecreyContract(cli, authCli, RinkebyZecreyVerifierAddr, RinkebyGovernanceAddr, gasPrice, _const.SuggestHighGasLimit)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(time.Since(elapse))
	fmt.Println("contract address:", addr)
	fmt.Println("deploy contract tx hash:", txHash)
}

func LoadZecrey() *Zecrey {
	instance, _ := LoadZecreyInstance(cli, ZecreyAddr)
	return instance
}

func TestDeposit(t *testing.T) {
	instance := LoadZecrey()
	depositAmount, _ := new(big.Int).SetString("1000000000000000000", 10)
	oldBalance, err := cli.GetBalance(_const.LocalSuperAddress)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(_utils.WeiToEther(oldBalance).String())
	txHash, err := Deposit(cli, authCli, instance, 0, "sher", depositAmount, gasPrice, _const.SuggestHighGasLimit)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(txHash)
	newBalance, err := cli.GetBalance(_const.LocalSuperAddress)
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
	txHash, err := CommitBlocks(cli, authCli, instance, block, commitBlocks, gasPrice, _const.SuggestHighGasLimit)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(txHash)
	committedCount, err := GetTotalBlocksCommitted(cli, instance)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(committedCount)
}

func TestProcessBlocks(t *testing.T) {
	instance := LoadZecrey()
	oldBalance, err := cli.GetBalance(_const.LocalSuperAddress)
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
		[]*big.Int{big.NewInt(0), big.NewInt(1)},
		[2]*big.Int{big.NewInt(0), big.NewInt(0)},
		b,
		[2]*big.Int{big.NewInt(0),
			big.NewInt(0)})
	processBlocks := []ZecreyProcessBlock{processBlock}
	txHash, err := ProcessBlocks(cli, authCli, instance, processBlocks, gasPrice, _const.SuggestHighGasLimit)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(txHash)
	processedCount, err := GetTotalBlocksProcessed(cli, instance)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(processedCount)
	newBalance, err := cli.GetBalance(_const.LocalSuperAddress)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(_utils.WeiToEther(newBalance).String())
}

func TestLoadLogs(t *testing.T) {
	instance := LoadZecrey()
	depositLogs, err := instance.ZecreyFilterer.FilterDeposit(&bind.FilterOpts{Start: 535, End: nil, Context: context.Background()})
	if err != nil {
		t.Fatal(err)
	}
	for depositLogs.Next() {
		fmt.Println(depositLogs.Event.DepositAmount)
	}
}
