package core

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/bnb-chain/zkbnb-eth-rpc/rpc"
)

func ConstructTransactOpts(cli *rpc.ProviderClient, authCli *rpc.AuthClient, gasPrice *big.Int, gasLimit uint64) (transactOpts *bind.TransactOpts, err error) {
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
	transactOpts.From = authCli.Address
	transactOpts.Value = big.NewInt(0)
	return transactOpts, nil
}

func ConstructTransactOptsWithSigner(cli *rpc.ProviderClient, signer bind.SignerFn, address common.Address,
	gasPrice *big.Int, gasLimit uint64) (transactOpts *bind.TransactOpts, err error) {
	transactOpts, err = NewKeyedTransactorWithSigner(signer, address)
	if err != nil {
		return nil, err
	}
	nonce, err := cli.GetPendingNonce(address.Hex())
	if err != nil {
		return nil, err
	}
	transactOpts.Nonce = big.NewInt(int64(nonce))
	transactOpts.GasPrice = gasPrice
	transactOpts.GasLimit = gasLimit
	transactOpts.From = address
	transactOpts.Value = big.NewInt(0)
	return transactOpts, nil
}

func ConstructTransactOptsWithNonce(authCli *rpc.AuthClient, gasPrice *big.Int, gasLimit uint64, nonce uint64) (transactOpts *bind.TransactOpts, err error) {
	transactOpts, err = bind.NewKeyedTransactorWithChainID(authCli.PrivateKey, authCli.ChainId)
	if err != nil {
		return nil, err
	}
	transactOpts.Nonce = big.NewInt(int64(nonce))
	transactOpts.GasPrice = gasPrice
	transactOpts.GasLimit = gasLimit
	transactOpts.From = authCli.Address
	transactOpts.Value = big.NewInt(0)
	return transactOpts, nil
}

func ConstructTransactOptsWithNonceAndSigner(signer bind.SignerFn, address common.Address,
	gasPrice *big.Int, gasLimit uint64, nonce uint64) (transactOpts *bind.TransactOpts, err error) {
	transactOpts, err = NewKeyedTransactorWithSigner(signer, address)
	if err != nil {
		return nil, err
	}
	transactOpts.Nonce = big.NewInt(int64(nonce))
	transactOpts.GasPrice = gasPrice
	transactOpts.GasLimit = gasLimit
	transactOpts.From = address
	transactOpts.Value = big.NewInt(0)
	return transactOpts, nil
}

func NewKeyedTransactorWithSigner(signer bind.SignerFn, address common.Address) (*bind.TransactOpts, error) {
	return &bind.TransactOpts{
		From:    address,
		Signer:  signer,
		Context: context.Background(),
	}, nil
}
