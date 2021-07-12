package _rpc

import (
	"fmt"
	"github.com/zecrey-labs/zecrey-eth-rpc/_const"
	"github.com/zecrey-labs/zecrey-eth-rpc/_utils"
	"testing"
)

func TestTransfer(t *testing.T) {
	cli, err := NewClient(_const.LocalNetwork)
	defer cli.Close()
	toAddress := "0x89D37ea8a0f102D90C424141F897A6a764A291AF"
	authClient, err := NewAuthClient(cli, "3f0b8ba8bf289106b7de0f75a659f4e9a7f17812c568298e8e83e789c64933c2", RinkebyChainId)
	if err != nil {
		t.Error(err)
	}
	txHash, err := cli.Transfer(authClient, toAddress, _utils.EtherToWei(99), nil, _const.SuggestGasLimit)
	if err != nil {
		t.Error(err)
	}
	status, err := cli.WaitingTransactionStatus(txHash)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("tx status:", status)
	fmt.Println("transaction hash:", txHash)
}

func TestDeployContract(t *testing.T) {
	cli, err := NewClient(_const.InfuraRinkebyNetwork)
	defer cli.Close()
	authClient, err := NewAuthClient(cli, _const.RinkebySuperSk, RinkebyChainId)
	if err != nil {
		panic(err)
	}
	contractAddress, txHash, err := cli.DeployContract(authClient, nil, "../contract/_interface/example/Store_sol_Store.abi", "../contract/_interface/example/Store_sol_Store.bin", []interface{}{"1.0"})
	if err != nil {
		panic(err)
	}
	fmt.Println("contract address:", contractAddress.String())
	fmt.Println("tx hash:", txHash.String())
}

func TestDeployContractUntil(t *testing.T) {
	cli, err := NewClient(_const.InfuraRinkebyNetwork)
	defer cli.Close()
	authClient, err := NewAuthClient(cli, _const.RinkebySuperSk, RinkebyChainId)
	if err != nil {
		panic(err)
	}
	status, contractAddress, txHash, err := cli.DeployContractUntil(authClient, nil, "../contract/_interface/example/Store_sol_Store.abi", "../contract/_interface/example/Store_sol_Store.bin", []interface{}{"1.0"})
	if err != nil {
		panic(err)
	}
	fmt.Println("contract address:", contractAddress.String())
	fmt.Println("tx hash:", txHash.String())
	fmt.Println("status:", status)
}
