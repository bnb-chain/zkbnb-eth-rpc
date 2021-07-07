package zecrey

import (
	"Zecrey-eth-rpc/_rpc"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

/*
	DeployZecreyTokenContract: deploy Zecrey Token
*/
func DeployZecreyTokenContract(cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, initialSupply *big.Int, gasPrice *big.Int, gasLimit uint64) (addr string, txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", "", err
	}
	address, tx, _, err := DeployZecreyToken(transactOpts, *cli, initialSupply)
	if err != nil {
		return "", "", err
	}
	return address.Hex(), tx.Hash().Hex(), nil
}

/*
	LoadZecreyTokenInstance: load zecrey token instance if it is already deployed
*/
func LoadZecreyTokenInstance(cli *_rpc.ProviderClient, addr string) (instance *ZecreyToken, err error) {
	instance, err = NewZecreyToken(common.HexToAddress(addr), *cli)
	return instance, err
}

/*
	TransferREY: ERC20 transferFrom
*/
func TransferREY(cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *ZecreyToken, recipientStr string, amount *big.Int, gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	recipient := common.HexToAddress(recipientStr)
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	tx, err := instance.Transfer(transactOpts, recipient, amount)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

/*
	ApproveREY: approve method for ERC20
*/
func ApproveREY(cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *ZecreyToken, recipientStr string, amount *big.Int, gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	recipient := common.HexToAddress(recipientStr)
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	tx, err := instance.Approve(transactOpts, recipient, amount)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

/*
	GetREYBalance: get ERC20 balance
*/
func GetREYBalance(cli *_rpc.ProviderClient, instance *ZecreyToken, accountStr string) (*big.Int, error) {
	account := common.HexToAddress(accountStr)
	balance, err := instance.BalanceOf(EmptyCallOpts(), account)
	return balance, err
}

func GetAllowance(cli *_rpc.ProviderClient, instance *ZecreyToken, fromStr string, toStr string) (*big.Int, error) {
	from := common.HexToAddress(fromStr)
	to := common.HexToAddress(toStr)
	allowance, err := instance.Allowance(EmptyCallOpts(), from, to)
	return allowance, err
}
