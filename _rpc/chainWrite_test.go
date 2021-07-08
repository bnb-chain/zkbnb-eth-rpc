package _rpc

import (
	"Zecrey-eth-rpc/_const"
	"Zecrey-eth-rpc/_utils"
	"fmt"
	"testing"
)

func TestTransfer(t *testing.T) {
	cli, err := NewClient(_const.InfuraRinkebyNetwork)
	defer cli.Close()
	toAddress := _const.ToAddress
	authClient, err := NewAuthClient(cli, _const.RinkebySuperSk)
	if err != nil {
		t.Error(err)
	}
	txHash, err := cli.Transfer(authClient, toAddress, _utils.EtherToWei(0.1), nil, _const.SuggestGasLimit)
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
	authClient, err := NewAuthClient(cli, _const.RinkebySuperSk)
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
	authClient, err := NewAuthClient(cli, _const.RinkebySuperSk)
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
