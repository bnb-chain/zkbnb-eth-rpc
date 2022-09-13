package rpc

import (
	"fmt"
	"testing"

	"github.com/bnb-chain/zkbnb-eth-rpc/constants"
	"github.com/bnb-chain/zkbnb-eth-rpc/utils"
)

func TestTransfer(t *testing.T) {
	cli, err := NewClient(constants.LocalNetwork)
	defer cli.Close()
	toAddress := "0xE9b15a2D396B349ABF60e53ec66Bcf9af262D449"
	authClient, err := NewAuthClient("1479a8a2464a43a9ecc8e6e6baf789ef9e35b89a63d864a22e01b3ef2ad6b6ba", RinkebyChainId)
	if err != nil {
		t.Error(err)
	}
	txHash, err := cli.Transfer(authClient, toAddress, utils.EtherToWei(99), nil, constants.SuggestGasLimit)
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
	cli, err := NewClient(constants.InfuraRinkebyNetwork)
	defer cli.Close()
	authClient, err := NewAuthClient(constants.RinkebySuperSk, RinkebyChainId)
	if err != nil {
		panic(err)
	}
	contractAddress, txHash, err := cli.DeployContract(authClient, nil, "../zkbnb/_interface/example/Store_sol_Store.abi", "../zkbnb/_interface/example/Store_sol_Store.bin", []interface{}{"1.0"})
	if err != nil {
		panic(err)
	}
	fmt.Println("zkbnb address:", contractAddress.String())
	fmt.Println("tx hash:", txHash.String())
}

func TestDeployContractUntil(t *testing.T) {
	cli, err := NewClient(constants.InfuraRinkebyNetwork)
	defer cli.Close()
	authClient, err := NewAuthClient(constants.RinkebySuperSk, RinkebyChainId)
	if err != nil {
		panic(err)
	}
	status, contractAddress, txHash, err := cli.DeployContractUntil(authClient, nil, "../zkbnb/_interface/example/Store_sol_Store.abi", "../zkbnb/_interface/example/Store_sol_Store.bin", []interface{}{"1.0"})
	if err != nil {
		panic(err)
	}
	fmt.Println("zkbnb address:", contractAddress.String())
	fmt.Println("tx hash:", txHash.String())
	fmt.Println("status:", status)
}
