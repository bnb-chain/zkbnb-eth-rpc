package contract

import (
	"fmt"
	"testing"
	"Zecrey-eth-rpc/_const"
	"Zecrey-eth-rpc/_rpc"
	"Zecrey-eth-rpc/_utils"
)

func TestGetErc20Balance(t *testing.T) {
	cli, err := _rpc.NewClient(_const.InfuraRinkebyNetwork)
	defer cli.Close()
	balance, err := GetErc20EtherBalance(cli, _const.RinkebySuperAddress, _const.EVATokenAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("eva token balance:", balance)
}

func TestTransferErc20(t *testing.T) {
	cli, err := _rpc.NewClient(_const.InfuraRinkebyNetwork)
	defer cli.Close()
	toAddress := _const.ToAddress
	beforeBalance, _ := GetErc20EtherBalance(cli, toAddress, _const.EVATokenAddress)
	fmt.Println("balance before transfer:", beforeBalance)
	authCli, _ := _rpc.NewAuthClient(cli, _const.RinkebySuperSk)
	txHash, err := TransferErc20(cli, authCli, toAddress, _utils.EtherToWei(100.0), nil, _const.EVATokenAddress)
	if err != nil {
		t.Error(err)
	}
	status, err := cli.WaitingTransactionStatus(txHash)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("transaction hash:", txHash)
	fmt.Println("transaction status:", status)
	afterBalance, _ := GetErc20EtherBalance(cli, toAddress, _const.EVATokenAddress)
	fmt.Println("balance after transfer:", afterBalance)
}
