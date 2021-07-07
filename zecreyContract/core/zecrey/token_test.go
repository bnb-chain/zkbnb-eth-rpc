package zecrey

import (
	"Zecrey-eth-rpc/_const"
	"fmt"
	"math/big"
	"testing"
	"time"
)

const ZecreyTokenAddr = "0xdd542d4e6AA059D4235169E0e1e019827d12196e"

func TestDeployZecreyTokenContract(t *testing.T) {
	elapse := time.Now()
	addr, txHash, err := DeployZecreyTokenContract(cli, authCli, InitialTokenSupply, gasPrice, _const.SuggestHighGasLimit)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(time.Since(elapse))
	fmt.Println("contract address:", addr)
	fmt.Println("deploy contract tx hash:", txHash)
}

func LoadZecreyToken() *ZecreyToken {
	instance, _ := LoadZecreyTokenInstance(cli, ZecreyTokenAddr)
	return instance
}

func TestGetREYBalance(t *testing.T) {
	instance := LoadZecreyToken()
	balance, err := GetREYBalance(cli, instance, _const.LocalSuperAddress)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(balance.String())
}

func TestGetAllowance(t *testing.T) {
	instance := LoadZecreyToken()
	toAddr := "0xbE008826B14b8E65Dee9104fB67D47a09Cdbbd0E"
	txHash, err := ApproveREY(cli, authCli, instance, toAddr, InitialTokenSupply, gasPrice, _const.SuggestFunctionGasLimit)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(txHash)
	allowance, err := GetAllowance(cli, instance, _const.LocalSuperAddress, toAddr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(allowance.String())
}

func TestTransferREY(t *testing.T) {
	instance := LoadZecreyToken()
	toAddr := "0xbE008826B14b8E65Dee9104fB67D47a09Cdbbd0E"
	txHash, err := TransferREY(cli, authCli, instance, toAddr, big.NewInt(1000000), gasPrice, _const.SuggestHighGasLimit)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(txHash)
	balance1, err := GetREYBalance(cli, instance, _const.LocalSuperAddress)
	if err != nil {
		t.Fatal(err)
	}
	balance2, err := GetREYBalance(cli, instance, toAddr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(balance1.String())
	fmt.Println(balance2.String())
}
