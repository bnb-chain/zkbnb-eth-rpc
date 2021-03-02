package zksneak

import (
	"ZKSneak-eth-rpc/_rpc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func DeployZKSneakContract(cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, gasPrice *big.Int, gasLimit uint64) (addr string, txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", "", err
	}
	address, tx, _, err := DeployZksneak(transactOpts, *cli)
	if err != nil {
		return "", "", err
	}
	return address.Hex(), tx.Hash().Hex(), nil
}

func LoadZKSneakInstance(cli *_rpc.ProviderClient, addr string) (zksneakInstance *Zksneak, err error) {
	zksneakInstance, err = NewZksneak(common.HexToAddress(addr), *cli)
	return zksneakInstance, err
}

func Deposit(cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Zksneak, value *big.Int, pk [32]byte, gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	transactOpts.Value = value
	tx, err := instance.Deposit(transactOpts, pk)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func VerifyBlock(cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Zksneak, newRoot [32]byte, gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	tx, err := instance.VerifyBlock(transactOpts, newRoot)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func Withdraw(cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Zksneak, addr string, amount *big.Int, gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	tx, err := instance.Withdraw(transactOpts, common.HexToAddress(addr), amount)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func ConstructTransactOpts(cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, gasPrice *big.Int, gasLimit uint64) (transactOpts *bind.TransactOpts, err error) {
	transactOpts, err = bind.NewKeyedTransactorWithChainID(authCli.PrivateKey, authCli.ChainId)
	if err != nil {
		return nil, err
	}
	nonce, err := cli.GetPendingNonce(authCli.Address.Hex())
	if err != nil {
		return nil, err
	}
	transactOpts.Nonce = big.NewInt(int64(nonce))
	transactOpts.GasPrice = gasPrice
	transactOpts.GasLimit = gasLimit
	transactOpts.Value = big.NewInt(0)
	return transactOpts, nil
}
