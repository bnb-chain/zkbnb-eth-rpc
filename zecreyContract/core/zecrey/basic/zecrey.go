package basic

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
	"math/big"
)

/*
	DeployZecreyContract: deploy zecrey contract
*/
func DeployZecreyContract(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient,
	l2ChainId uint8, nativeAssetId uint16, maxPendingBlocks uint16,
	gasPrice *big.Int, gasLimit uint64,
) (addr string, txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", "", err
	}
	address, tx, _, err := DeployZecrey(
		transactOpts, *cli,
		l2ChainId, nativeAssetId, maxPendingBlocks)
	if err != nil {
		return "", "", err
	}
	return address.Hex(), tx.Hash().Hex(), nil
}

/*
	LoadGovernanceInstance: load governance instance if it is already deployed
*/
func LoadZecreyInstance(cli *_rpc.ProviderClient, addr string) (instance *Zecrey, err error) {
	instance, err = NewZecrey(common.HexToAddress(addr), *cli)
	return instance, err
}

func PackInitializeZecreyParams(
	governanceAddr string,
	verifierAddr string,
	genesisBlockNumber uint32,
	genesisOnchainOpsRoot []byte,
	genesisStateRoot []byte,
	genesisTimestamp *big.Int,
	genesisCommitment []byte,
	merkleHelper [OnChainOpsTreeHelperDepth]bool,
) ([]byte, error) {
	if len(genesisOnchainOpsRoot) != Bytes32Len || len(genesisStateRoot) != Bytes32Len || len(genesisCommitment) != Bytes32Len {
		return nil, errors.New("[PackInitializeZecreyParams] invalid bytes32")
	}
	arguments := abi.Arguments{
		{Type: AddressType},
		{Type: AddressType},
		{Type: Uint32Type},
		{Type: Bytes32Type},
		{Type: Bytes32Type},
		{Type: Uint256Type},
		{Type: Bytes32Type},
		{Type: MerkleHelperType},
	}
	hashBytes, err := arguments.Pack(
		common.HexToAddress(governanceAddr),
		common.HexToAddress(verifierAddr),
		genesisBlockNumber,
		SetFixed32Bytes(genesisOnchainOpsRoot),
		SetFixed32Bytes(genesisStateRoot),
		genesisTimestamp,
		SetFixed32Bytes(genesisCommitment),
		merkleHelper,
	)
	if err != nil {
		return nil, err
	}
	return hashBytes, nil
}

/*
	ZecreyInitialize: initialize zecrey contract
*/
func ZecreyInitialize(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Zecrey,
	governanceAddr string,
	verifierAddr string,
	genesisBlockNumber uint32,
	genesisOnchainOpsRoot []byte,
	genesisStateRoot []byte,
	genesisTimestamp *big.Int,
	genesisCommitment []byte,
	genesisMerkleHelper [OnChainOpsTreeHelperDepth]bool,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	if !IsValidAddress(governanceAddr) || !IsValidAddress(verifierAddr) {
		return "", errors.New("[ZecreyInitialize] invalid address")
	}
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	params, err := PackInitializeZecreyParams(
		governanceAddr,
		verifierAddr,
		genesisBlockNumber,
		genesisOnchainOpsRoot,
		genesisStateRoot,
		genesisTimestamp,
		genesisCommitment,
		genesisMerkleHelper,
	)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.Initialize(transactOpts, params)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	ZecreyDepositOrLockNativeAsset: deposit or lock native asset
*/
func ZecreyDepositOrLockNativeAsset(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Zecrey,
	opType uint8, accountName string,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.DepositOrLockNativeAsset(transactOpts, opType, SetFixed32Bytes([]byte(accountName)))
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	ZecreyDepositOrLockERC20: deposit or lock erc20
*/
func ZecreyDepositOrLockERC20(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Zecrey,
	opType uint8, assetAddr string, amount *big.Int, accountName string,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	if !IsValidAddress(assetAddr) {
		return "", errors.New("[ZecreyDepositOrLockERC20] invalid asset address")
	}
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.DepositOrLockERC20(transactOpts, opType, common.HexToAddress(assetAddr), amount, SetFixed32Bytes([]byte(accountName)))
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	ZecreyGetPendingBalance: get pending balance
*/
func ZecreyGetPendingBalance(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Zecrey,
	owner string, assetAddr string,
) (*big.Int, error) {
	if !IsValidAddress(owner) || !IsValidAddress(assetAddr) {
		return nil, errors.New("[ZecreyGetPendingBalance] invalid address")
	}
	balance, err := instance.GetPendingBalance(EmptyCallOpts(), common.HexToAddress(owner), common.HexToAddress(assetAddr))
	return balance, err
}

/*
	ZecreyWithdrawPendingBalance: withdraw pending balance
*/
func ZecreyWithdrawPendingBalance(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Zecrey,
	owner string, assetAddr string, amount *big.Int,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	if !IsValidAddress(owner) || !IsValidAddress(assetAddr) {
		return "", errors.New("[ZecreyWithdrawPendingBalance] invalid address")
	}
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.WithdrawPendingBalance(transactOpts, common.HexToAddress(owner), common.HexToAddress(assetAddr), amount)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	ZecreyComputeOnchainOpsHash: compute onchain ops hash
*/
func ZecreyComputeOnchainOpsHash(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Zecrey,
	onchainOpsPubData []byte,
	gasPrice *big.Int, gasLimit uint64,
) (hashVal []byte, err error) {
	val, err := instance.ComputeOnchainOpsHash(EmptyCallOpts(), onchainOpsPubData)
	return val[:], err
}

/*
	ZecreyCommitBlocks: commit blocks
*/
func ZecreyCommitBlocks(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Zecrey,
	lastBlock StorageBlockHeader, commitBlocksInfo []ZecreyCommitBlockInfo,
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
	ZecreyVerifyBlocks: verify blocks
*/
func ZecreyVerifyBlocks(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Zecrey,
	verifyBlocksInfo []ZecreyVerifyBlockInfo,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.VerifyBlocks(transactOpts, verifyBlocksInfo)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	ZecreyExecuteBlocks: execute blocks
*/
func ZecreyExecuteBlocks(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Zecrey,
	blocks []StorageBlockHeader,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.ExecuteBlocks(transactOpts, blocks)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	ZecreyRevertBlocks: revert blocks
*/
func ZecreyRevertBlocks(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Zecrey,
	blocks []StorageBlockHeader,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.RevertBlocks(transactOpts, blocks)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}
