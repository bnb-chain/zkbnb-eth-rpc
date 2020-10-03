package contract

import (
	"eva-go-rpc/_const"
	"eva-go-rpc/_rpc"
	"eva-go-rpc/_rpc/chain"
	"eva-go-rpc/_utils"
	"fmt"
	"testing"
)

func TestGetErc20Balance(t *testing.T) {
	balance, err := GetErc20EtherBalance(_const.SuperAddress, _const.EVATokenAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("eva token balance:", balance)
}

func TestTransferErc20(t *testing.T) {
	toAddress := _const.ToAddress
	beforeBalance, _ := GetErc20EtherBalance(toAddress, _const.EVATokenAddress)
	fmt.Println("balance before transfer:", beforeBalance)
	authCli, _ := _rpc.NewAuthClient(_const.SuperSk)
	txHash, err := TransferErc20(authCli, toAddress, _utils.EtherToWei(100.0), nil, _const.EVATokenAddress)
	if err != nil {
		t.Error(err)
	}
	status, err := chain.WaitingTransactionStatus(txHash)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("transaction hash:", txHash)
	fmt.Println("transaction status:", status)
	afterBalance, _ := GetErc20EtherBalance(toAddress, _const.EVATokenAddress)
	fmt.Println("balance after transfer:", afterBalance)
}
