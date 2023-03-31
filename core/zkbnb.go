package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/bnb-chain/zkbnb-eth-rpc/rpc"
)

/*
	LoadZkBNBInstance: load zkbnb instance if it is already deployed
*/
func LoadZkBNBInstance(cli *rpc.ProviderClient, addr string) (instance *ZkBNB, err error) {
	instance, err = NewZkBNB(common.HexToAddress(addr), *cli)
	return instance, err
}

/*
	CommitBlocks: commit blocks
*/
func CommitBlocks(
	cli *rpc.ProviderClient, constructor TransactOptsConstructor, instance *ZkBNB,
	lastBlock StorageStoredBlockInfo, commitBlocksInfo []ZkBNBCommitBlockInfo,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := constructor.ConstructTransactOpts(cli, gasPrice, gasLimit)
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
	CommitBlocks: commit blocks
*/
func CommitBlocksWithNonce(constructor TransactOptsConstructor, instance *ZkBNB,
	lastBlock StorageStoredBlockInfo, commitBlocksInfo []ZkBNBCommitBlockInfo,
	gasPrice *big.Int, gasLimit uint64, nonce uint64,
) (txHash string, err error) {
	transactOpts, err := constructor.ConstructTransactOptsWithNonce(gasPrice, gasLimit, nonce)
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
	Estimate Gas for commit blocks operation
*/
func EstimateCommitGasWithNonce(constructor TransactOptsConstructor, instance *ZkBNB,
	lastBlock StorageStoredBlockInfo, commitBlocksInfo []ZkBNBCommitBlockInfo,
	gasPrice *big.Int, gasLimit uint64, nonce uint64,
) (gas uint64, err error) {
	transactOpts, err := constructor.ConstructTransactOptsWithNonce(gasPrice, gasLimit, nonce)
	if err != nil {
		return 0, err
	}
	// Only for gas estimation and set NoSend = true
	transactOpts.NoSend = true
	// call initialize
	tx, err := instance.CommitBlocks(transactOpts, lastBlock, commitBlocksInfo)
	if err != nil {
		return 0, err
	}
	return tx.Gas(), nil
}

/*
	VerifyAndExecuteBlocks: verify and execute blocks
*/
func VerifyAndExecuteBlocks(cli *rpc.ProviderClient, constructor TransactOptsConstructor, instance *ZkBNB,
	verifyAndExecuteBlocksInfo []ZkBNBVerifyAndExecuteBlockInfo, proofs []*big.Int,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := constructor.ConstructTransactOpts(cli, gasPrice, gasLimit)
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

/*
	VerifyAndExecuteBlocks: verify and execute blocks
*/
func VerifyAndExecuteBlocksWithNonce(constructor TransactOptsConstructor, instance *ZkBNB,
	verifyAndExecuteBlocksInfo []ZkBNBVerifyAndExecuteBlockInfo, proofs []*big.Int,
	gasPrice *big.Int, gasLimit uint64, nonce uint64,
) (txHash string, err error) {
	transactOpts, err := constructor.ConstructTransactOptsWithNonce(gasPrice, gasLimit, nonce)
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

/*
	Estimate Gas for verifying and executing blocks with kms signature facility
*/
func EstimateVerifyAndExecuteWithNonce(constructor TransactOptsConstructor, instance *ZkBNB,
	verifyAndExecuteBlocksInfo []ZkBNBVerifyAndExecuteBlockInfo, proofs []*big.Int,
	gasPrice *big.Int, gasLimit uint64, nonce uint64,
) (gas uint64, err error) {
	transactOpts, err := constructor.ConstructTransactOptsWithNonce(gasPrice, gasLimit, nonce)
	if err != nil {
		return 0, err
	}
	// Only for gas estimation and set NoSend = true
	transactOpts.NoSend = true
	// call initialize
	tx, err := instance.VerifyAndExecuteBlocks(transactOpts, verifyAndExecuteBlocksInfo, proofs)
	if err != nil {
		return 0, err
	}
	return tx.Gas(), nil
}

/*
	RevertBlocks: revert blocks
*/
func RevertBlocks(
	cli *rpc.ProviderClient, constructor TransactOptsConstructor, instance *ZkBNB,
	revertBlocks []StorageStoredBlockInfo,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := constructor.ConstructTransactOpts(cli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	tx, err := instance.RevertBlocks(transactOpts, revertBlocks)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	PerformDesert: perform desert
*/
func PerformDesert(cli *rpc.ProviderClient, constructor TransactOptsConstructor, instance *ZkBNB, storedBlockInfo StorageStoredBlockInfo, nftRoot *big.Int, assetExitData DesertVerifierAssetExitData, accountExitData DesertVerifierAccountExitData,
	assetMerkleProof [16]*big.Int, accountMerkleProof [32]*big.Int,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := constructor.ConstructTransactOpts(cli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.PerformDesert(transactOpts, storedBlockInfo, nftRoot, assetExitData, accountExitData, assetMerkleProof, accountMerkleProof)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	PerformDesertNft: perform desert nft
*/
func PerformDesertNft(cli *rpc.ProviderClient, constructor TransactOptsConstructor, instance *ZkBNB,
	storedBlockInfo StorageStoredBlockInfo, assetRoot *big.Int, accountExitData DesertVerifierAccountExitData, exitNfts []DesertVerifierNftExitData, accountMerkleProof [32]*big.Int, nftMerkleProofs [][40]*big.Int,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := constructor.ConstructTransactOpts(cli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.PerformDesertNft(transactOpts, storedBlockInfo, assetRoot, accountExitData, exitNfts, accountMerkleProof, nftMerkleProofs)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	WithdrawPendingBalance: withdraw pending balance
*/
func WithdrawPendingBalance(cli *rpc.ProviderClient, constructor TransactOptsConstructor, instance *ZkBNB, owner common.Address, token common.Address, amount *big.Int, gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	transactOpts, err := constructor.ConstructTransactOpts(cli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.WithdrawPendingBalance(transactOpts, owner, token, amount)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	WithdrawPendingNFTBalance: withdraw pending nft balance
*/
func WithdrawPendingNFTBalance(cli *rpc.ProviderClient, constructor TransactOptsConstructor, instance *ZkBNB, nftIndex *big.Int, gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	transactOpts, err := constructor.ConstructTransactOpts(cli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.WithdrawPendingNFTBalance(transactOpts, nftIndex)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	CancelOutstandingDepositsForExodusMode: cancel outstanding deposit
*/
func CancelOutstandingDepositsForExodusMode(cli *rpc.ProviderClient, constructor TransactOptsConstructor, instance *ZkBNB, priorityRequestId uint64, depositsPubData [][]byte, gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	transactOpts, err := constructor.ConstructTransactOpts(cli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.CancelOutstandingDepositsForDesertMode(transactOpts, priorityRequestId, depositsPubData)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	ActivateDesertMode: activate desert mode
*/
func ActivateDesertMode(cli *rpc.ProviderClient, constructor TransactOptsConstructor, instance *ZkBNB, gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	transactOpts, err := constructor.ConstructTransactOpts(cli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.ActivateDesertMode(transactOpts)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}
