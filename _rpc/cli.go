package _rpc

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type ProviderClient struct {
	*ethclient.Client
}

func NewClient(provider string) (cli *ProviderClient, err error) {
	client, err := ethclient.Dial(provider)
	if err != nil {
		return nil, err
	}
	return &ProviderClient{client}, nil
}
