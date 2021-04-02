package zksneak

import (
	"context"
	"fmt"
	"golang.org/x/crypto/sha3"
	"testing"
	"PrivaL-eth-rpc/_const"
	"PrivaL-eth-rpc/_rpc"
	"PrivaL-eth-rpc/_utils"
)

func TestDeployZKSneakContract(t *testing.T) {
	cli, _ := _rpc.NewClient(_const.LocalNetwork)
	defer cli.Close()
	authCli, _ := _rpc.NewAuthClient(cli, _const.LocalSuperSk)
	gasPrice, _ := cli.SuggestGasPrice(context.Background())
	addr, txHash, err := DeployZKSneakContract(cli, authCli, gasPrice, _const.SuggestHighGasLimit)
	if err != nil {
		panic(err)
	}
	fmt.Println("contract address:", addr)
	fmt.Println("deploy contract tx hash:", txHash)
}

func TestDeposit(t *testing.T) {
	cli, _ := _rpc.NewClient(_const.LocalNetwork)
	defer cli.Close()
	authCli, _ := _rpc.NewAuthClient(cli, _const.LocalSuperSk)
	gasPrice, _ := cli.SuggestGasPrice(context.Background())
	instance, err := LoadZKSneakInstance(cli, _const.ZKSneakAddress)
	if err != nil {
		panic(err)
	}
	var pk [32]byte
	copy(pk[:], []byte("test"))
	value := _utils.EtherToWei(10)
	hash, err := Deposit(cli, authCli, instance, value, pk, gasPrice, _const.SuggestContractGasLimit)
	if err != nil {
		panic(err)
	}
	fmt.Println("deposit tx hash:", hash)
}

func TestGetRoot(t *testing.T) {
	cli, _ := _rpc.NewClient(_const.LocalNetwork)
	defer cli.Close()
	instance, err := LoadZKSneakInstance(cli, _const.ZKSneakAddress)
	if err != nil {
		panic(err)
	}
	accountRoot, err := instance.AccountRoot(nil)
	fmt.Println("root:", accountRoot)
}

func TestVerifyBlock(t *testing.T) {
	cli, _ := _rpc.NewClient(_const.LocalNetwork)
	defer cli.Close()
	authCli, _ := _rpc.NewAuthClient(cli, _const.LocalSuperSk)
	gasPrice, _ := cli.SuggestGasPrice(context.Background())
	instance, err := LoadZKSneakInstance(cli, _const.ZKSneakAddress)
	if err != nil {
		panic(err)
	}
	var newRoot [32]byte
	hashVal, _ := _utils.CalHash([]byte("Lin"), sha3.NewLegacyKeccak256)
	copy(newRoot[:], hashVal)
	fmt.Println("new root:", newRoot)
	hash, err := VerifyBlock(cli, authCli, instance, newRoot, gasPrice, _const.SuggestContractGasLimit)
	if err != nil {
		panic(err)
	}
	fmt.Println("verifyBlock tx hash:", hash)
}

func TestWithdraw(t *testing.T) {
	cli, _ := _rpc.NewClient(_const.LocalNetwork)
	defer cli.Close()
	authCli, _ := _rpc.NewAuthClient(cli, _const.LocalSuperSk)
	gasPrice, _ := cli.SuggestGasPrice(context.Background())
	instance, err := LoadZKSneakInstance(cli, _const.ZKSneakAddress)
	if err != nil {
		panic(err)
	}
	txHash, err := Withdraw(cli, authCli, instance, _const.ToAddress, _utils.EtherToWei(0.5), gasPrice, _const.SuggestContractGasLimit)
	if err != nil {
		panic(err)
	}
	fmt.Println("withdraw tx hash:", txHash)
}
