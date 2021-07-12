package zecrey

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
	"math/big"
)

func DeployZecreyVerifierContract(cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, gasPrice *big.Int, gasLimit uint64) (addr string, txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", "", err
	}
	address, tx, _, err := DeployVerifier(transactOpts, *cli)
	if err != nil {
		return "", "", err
	}
	return address.Hex(), tx.Hash().Hex(), nil
}

/*
	LoadZecreyVerifierInstance: load zecrey verifier instance if it is already deployed
*/
func LoadZecreyVerifierInstance(cli *_rpc.ProviderClient, addr string) (instance *Verifier, err error) {
	instance, err = NewVerifier(common.HexToAddress(addr), *cli)
	return instance, err
}
