package basic

import (
	"github.com/bnb-chain/zkbnb-eth-rpc/_rpc"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

/*
	DeployVerifierContract: deploy verifier
*/
func DeployVerifierContract(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient,
	gasPrice *big.Int, gasLimit uint64,
) (addr string, txHash string, err error) {
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
	LoadVerifierInstance: load verifier instance if it is already deployed
*/
func LoadVerifierInstance(cli *_rpc.ProviderClient, addr string) (instance *Verifier, err error) {
	instance, err = NewVerifier(common.HexToAddress(addr), *cli)
	return instance, err
}
