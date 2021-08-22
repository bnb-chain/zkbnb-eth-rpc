package zecrey

import (
	"fmt"
	"github.com/zecrey-labs/zecrey-eth-rpc/_const"
	"math/big"
	"testing"
	"time"
)

const ZecreyTokenAddr = "0xCd44483447D3dbe166e1D8c74c6DC0705dFcDAc2"
const RinkebyZecreyTokenAddr = "0xB2aA4c363DCF5e5838a90A49670755a8Ec4BaF3E"

func TestDeployZecreyTokenContract(t *testing.T) {
	elapse := time.Now()
	addr, txHash, err := DeployZecreyTokenContract(localCli, localAuthCli, InitialTokenSupply, localGasPrice, _const.SuggestHighGasLimit)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(time.Since(elapse))
	fmt.Println("contract address:", addr)
	fmt.Println("deploy contract tx hash:", txHash)
}

func LoadZecreyToken() *ZecreyToken {
	instance, _ := LoadZecreyTokenInstance(rinkebyCli, RinkebyZecreyTokenAddr)
	return instance
}

func TestGetREYBalance(t *testing.T) {
	instance := LoadZecreyToken()
	balance, err := GetREYBalance(localCli, instance, _const.LocalSuperAddress)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(balance.String())
}

func TestGetAllowance(t *testing.T) {
	instance := LoadZecreyToken()
	toAddr := "0xbE008826B14b8E65Dee9104fB67D47a09Cdbbd0E"
	txHash, err := ApproveREY(localCli, localAuthCli, instance, toAddr, InitialTokenSupply, localGasPrice, _const.SuggestFunctionGasLimit)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(txHash)
	allowance, err := GetAllowance(localCli, instance, _const.LocalSuperAddress, toAddr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(allowance.String())
}

func TestTransferREY(t *testing.T) {
	instance := LoadZecreyToken()
	toAddr := "0xbE008826B14b8E65Dee9104fB67D47a09Cdbbd0E"
	txHash, err := TransferREY(localCli, localAuthCli, instance, toAddr, big.NewInt(1000000), localGasPrice, _const.SuggestHighGasLimit)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(txHash)
	balance1, err := GetREYBalance(localCli, instance, _const.LocalSuperAddress)
	if err != nil {
		t.Fatal(err)
	}
	balance2, err := GetREYBalance(localCli, instance, toAddr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(balance1.String())
	fmt.Println(balance2.String())
}
