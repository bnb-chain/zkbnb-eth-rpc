package rpc

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethawskmssigner "github.com/welthee/go-ethereum-aws-kms-tx-signer/v2"
	"math/big"
)

type KMSKeyClient struct {
	KMSClient *kms.Client
	KeyId     string
	ChainId   *big.Int
}

func NewKMSKeyClient(client *kms.Client, keyId string, chainID *big.Int) (authCli *KMSKeyClient, err error) {
	return &KMSKeyClient{KMSClient: client, KeyId: keyId, ChainId: chainID}, nil
}

func (c *KMSKeyClient) ConstructTransactOpts(cli *ProviderClient, gasPrice *big.Int, gasLimit uint64) (transactOpts *bind.TransactOpts, err error) {
	transactOpts, err = ethawskmssigner.NewAwsKmsTransactorWithChainIDCtx(context.Background(), c.KMSClient, c.KeyId, c.ChainId)
	if err != nil {
		return nil, err
	}
	nonce, err := cli.GetPendingNonce(transactOpts.From.Hex())
	if err != nil {
		return nil, err
	}
	transactOpts.Nonce = big.NewInt(int64(nonce))
	transactOpts.GasPrice = gasPrice
	transactOpts.GasLimit = gasLimit
	transactOpts.Value = big.NewInt(0)
	return transactOpts, nil
}

func (c *KMSKeyClient) ConstructTransactOptsWithNonce(gasPrice *big.Int, gasLimit uint64, nonce uint64) (transactOpts *bind.TransactOpts, err error) {
	transactOpts, err = ethawskmssigner.NewAwsKmsTransactorWithChainIDCtx(context.Background(), c.KMSClient, c.KeyId, c.ChainId)
	if err != nil {
		return nil, err
	}
	transactOpts.Nonce = big.NewInt(int64(nonce))
	transactOpts.GasPrice = gasPrice
	transactOpts.GasLimit = gasLimit
	transactOpts.Value = big.NewInt(0)
	return transactOpts, nil
}
