package zecreyLegend

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
	"math/big"
)

/*
	LoadZecreyLegendInstance: load zecrey legend instance if it is already deployed
*/
func LoadZecreyLegendInstance(cli *_rpc.ProviderClient, addr string) (instance *ZecreyLegend, err error) {
	instance, err = NewZecreyLegend(common.HexToAddress(addr), *cli)
	return instance, err
}

/*
	CommitBlocks: commit blocks
*/
func CommitBlocks(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *ZecreyLegend,
	lastBlock StorageStoredBlockInfo, commitBlocksInfo []OldZecreyLegendCommitBlockInfo,
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
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *ZecreyLegend,
	verifyAndExecuteBlocksInfo []OldZecreyLegendVerifyAndExecuteBlockInfo, proofs []*big.Int,
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
