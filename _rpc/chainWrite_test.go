package _rpc

import (
	"fmt"
	"github.com/bnb-chain/zkbas-eth-rpc/_const"
	"github.com/bnb-chain/zkbas-eth-rpc/_utils"
	"testing"
)

func TestTransfer(t *testing.T) {
	cli, err := NewClient(_const.LocalNetwork)
	defer cli.Close()
	toAddress := "0xE9b15a2D396B349ABF60e53ec66Bcf9af262D449"
	authClient, err := NewAuthClient(cli, "1479a8a2464a43a9ecc8e6e6baf789ef9e35b89a63d864a22e01b3ef2ad6b6ba", RinkebyChainId)
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
	contractAddress, txHash, err := cli.DeployContract(authClient, nil, "../zecrey/_interface/example/Store_sol_Store.abi", "../zecrey/_interface/example/Store_sol_Store.bin", []interface{}{"1.0"})
	if err != nil {
		panic(err)
	}
	fmt.Println("zecrey address:", contractAddress.String())
	fmt.Println("tx hash:", txHash.String())
}

func TestDeployContractUntil(t *testing.T) {
	cli, err := NewClient(_const.InfuraRinkebyNetwork)
	defer cli.Close()
	authClient, err := NewAuthClient(cli, _const.RinkebySuperSk, RinkebyChainId)
	if err != nil {
		panic(err)
	}
	status, contractAddress, txHash, err := cli.DeployContractUntil(authClient, nil, "../zecrey/_interface/example/Store_sol_Store.abi", "../zecrey/_interface/example/Store_sol_Store.bin", []interface{}{"1.0"})
	if err != nil {
		panic(err)
	}
	fmt.Println("zecrey address:", contractAddress.String())
	fmt.Println("tx hash:", txHash.String())
	fmt.Println("status:", status)
}
