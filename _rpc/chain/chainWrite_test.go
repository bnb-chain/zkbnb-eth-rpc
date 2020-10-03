package chain

import (
	"eva-go-rpc/_const"
	"eva-go-rpc/_rpc"
	"eva-go-rpc/_utils"
	"fmt"
	"testing"
)

func TestTransfer(t *testing.T) {
	toAddress := "0xE9b15a2D396B349ABF60e53ec66Bcf9af262D449"
	beginBalance, _ := GetEtherBalance(toAddress)
	fmt.Println("balance before transfer:", beginBalance)
	authClient, err := _rpc.NewAuthClient(_const.SuperSk)
	if err != nil {
		t.Error(err)
	}
	txHash, err := Transfer(authClient, toAddress, _utils.EtherToWei(0.1), nil, _const.SuggestGasLimit)
	if err != nil {
		t.Error(err)
	}
	status, err := WaitingTransactionStatus(txHash)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("tx status:", status)
	endBalance, _ := GetEtherBalance(toAddress)
	fmt.Println("balance after transfer:", endBalance)
	fmt.Println("transaction hash:", txHash)
}

func TestDeployContract(t *testing.T) {
	authClient, err := _rpc.NewAuthClient(_const.SuperSk)
	if err != nil {
		panic(err)
	}
	contractAddress, txHash, err := DeployContract(authClient, nil, "../contract/_interface/example/Store_sol_Store.abi", "../contract/_interface/example/Store_sol_Store.bin", []interface{}{"1.0"})
	if err != nil {
		panic(err)
	}
	fmt.Println("contract address:", contractAddress.String())
	fmt.Println("tx hash:", txHash.String())
}

func TestDeployContractUntil(t *testing.T) {
	authClient, err := _rpc.NewAuthClient(_const.SuperSk)
	if err != nil {
		panic(err)
	}
	status, contractAddress, txHash, err := DeployContractUntil(authClient, nil, "../contract/_interface/example/Store_sol_Store.abi", "../contract/_interface/example/Store_sol_Store.bin", []interface{}{"1.0"})
	if err != nil {
		panic(err)
	}
	fmt.Println("contract address:", contractAddress.String())
	fmt.Println("tx hash:", txHash.String())
	fmt.Println("status:", status)
}
