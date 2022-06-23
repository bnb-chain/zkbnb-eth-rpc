package legend

import (
	"github.com/bnb-chain/zkbas-eth-rpc/_rpc"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

/*
	LoadZkbasInstance: load zkbas instance if it is already deployed
*/
func LoadZkbasInstance(cli *_rpc.ProviderClient, addr string) (instance *Zkbas, err error) {
	instance, err = NewZkbas(common.HexToAddress(addr), *cli)
	return instance, err
}

/*
	CommitBlocks: commit blocks
*/
func CommitBlocks(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Zkbas,
	lastBlock StorageStoredBlockInfo, commitBlocksInfo []OldZkbasCommitBlockInfo,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.CommitBlocks(transactOpts, lastBlock, commitBlocksInfo)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	VerifyAndExecuteBlocks: verify and execute blocks
*/
func VerifyAndExecuteBlocks(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Zkbas,
	verifyAndExecuteBlocksInfo []OldZkbasVerifyAndExecuteBlockInfo, proofs []*big.Int,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.VerifyAndExecuteBlocks(transactOpts, verifyAndExecuteBlocksInfo, proofs)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}
