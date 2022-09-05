package basic

import (
	"github.com/bnb-chain/zkbnb-eth-rpc/_rpc"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

/*
	DeployGovernanceContract: deploy governance
*/
func DeployErc20Contract(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient,
	name string, symbol string,
	gasPrice *big.Int, gasLimit uint64,
) (addr string, txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", "", err
	}
	address, tx, _, err := DeployErc20(transactOpts, *cli, InitialSupply, name, symbol)
	if err != nil {
		return "", "", err
	}
	return address.Hex(), tx.Hash().Hex(), nil
}

/*
	LoadGovernanceInstance: load governance instance if it is already deployed
*/
func LoadErc20Instance(cli *_rpc.ProviderClient, addr string) (instance *Erc20, err error) {
	instance, err = NewErc20(common.HexToAddress(addr), *cli)
	return instance, err
}
