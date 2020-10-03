package chain

import (
	"eva-go-rpc/econst"
	"eva-go-rpc/erpc"
	"eva-go-rpc/eutils"
	"fmt"
	"testing"
)

func TestTransfer(t *testing.T) {
	toAddress := "0xE9b15a2D396B349ABF60e53ec66Bcf9af262D449"
	beginBalance, _ := GetEtherBalance(toAddress)
	fmt.Println("balance before transfer:", beginBalance)
	authClient, err := erpc.NewAuthClient(econst.SuperSk)
	if err != nil {
		t.Error(err)
	}
	txHash, err := Transfer(authClient, toAddress, eutils.EtherToWei(0.1), nil, econst.SuggestGasLimit)
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
	authClient, err := erpc.NewAuthClient(econst.SuperSk)
	if err != nil {
		panic(err)
	}
	contractAddress, txHash, err := DeployContract(authClient, nil, "../contract/einterface/example/Store_sol_Store.abi", "../contract/einterface/example/Store_sol_Store.bin", []interface{}{"1.0"})
	if err != nil {
		panic(err)
	}
	fmt.Println("contract address:", contractAddress.String())
	fmt.Println("tx hash:", txHash.String())
}

func TestDeployContractUntil(t *testing.T) {
	authClient, err := erpc.NewAuthClient(econst.SuperSk)
	if err != nil {
		panic(err)
	}
	status, contractAddress, txHash, err := DeployContractUntil(authClient, nil, "../contract/einterface/example/Store_sol_Store.abi", "../contract/einterface/example/Store_sol_Store.bin", []interface{}{"1.0"})
	if err != nil {
		panic(err)
	}
	fmt.Println("contract address:", contractAddress.String())
	fmt.Println("tx hash:", txHash.String())
	fmt.Println("status:", status)
}
