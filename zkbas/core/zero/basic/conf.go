package basic

import (
	"github.com/bnb-chain/zkbas-eth-rpc/_rpc"
	"math/big"
)

func ConstructCliAndAuthCli(chainId *big.Int, networkEndpoint string, sk string) (cli *ProviderClient, authCli *AuthClient, err error) {
	cli, err = _rpc.NewClient(networkEndpoint)
	if err != nil {
		return nil, nil, err
	}
	authCli, err = _rpc.NewAuthClient(cli, sk, chainId)
	if err != nil {
		return nil, nil, err
	}
	return cli, authCli, nil
}
