package core

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/ethereum/go-ethereum/common"
	ethawskmssigner "github.com/welthee/go-ethereum-aws-kms-tx-signer/v2"
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

func ConstructTransactOptsWithKms(cli *rpc.ProviderClient, ctx context.Context, kmsSvc *kms.Client, keyId string, chainID *big.Int, address common.Address,
	gasPrice *big.Int, gasLimit uint64) (transactOpts *bind.TransactOpts, err error) {
	transactOpts, err = ethawskmssigner.NewAwsKmsTransactorWithChainIDCtx(ctx, kmsSvc, keyId, chainID)
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

func ConstructTransactOptsWithNonceAndKms(ctx context.Context, kmsSvc *kms.Client, keyId string, chainID *big.Int, address common.Address,
	gasPrice *big.Int, gasLimit uint64, nonce uint64) (transactOpts *bind.TransactOpts, err error) {
	transactOpts, err = ethawskmssigner.NewAwsKmsTransactorWithChainIDCtx(ctx, kmsSvc, keyId, chainID)
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
