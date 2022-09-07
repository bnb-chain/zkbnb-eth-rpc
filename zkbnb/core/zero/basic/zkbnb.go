package basic

import (
	"errors"
	"github.com/bnb-chain/zkbnb-eth-rpc/_rpc"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

/*
	DeployZkBNBContract: deploy zkbnb zkbnb
*/
func DeployZkBNBContract(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient,
	gasPrice *big.Int, gasLimit uint64,
) (addr string, txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", "", err
	}
	address, tx, _, err := DeployZkBNB(
		transactOpts, *cli)
	if err != nil {
		return "", "", err
	}
	return address.Hex(), tx.Hash().Hex(), nil
}

/*
	LoadZkBNBInstance: load governance instance if it is already deployed
*/
func LoadZkBNBInstance(cli *_rpc.ProviderClient, addr string) (instance *ZkBNB, err error) {
	instance, err = NewZkBNB(common.HexToAddress(addr), *cli)
	return instance, err
}

func ZkBNBComputeCommitment(hashVal []byte, newBlock ZkBNBCommitBlockInfo) (value []byte, err error) {
	arguments := abi.Arguments{
		{Type: Bytes32Type},
		{Type: Uint32Type},
		{Type: Bytes32Type},
		{Type: Uint256Type},
	}
	value, err = arguments.Pack(
		SetFixed32Bytes(hashVal),
		newBlock.BlockNumber,
		newBlock.NewAccountRoot,
		newBlock.Timestamp,
	)
	if err != nil {
		return nil, err
	}
	return crypto.Keccak256Hash(value).Bytes(), nil
}

func ZkBNBGetStoredBlockHashesByHeight(instance *ZkBNB, height uint32) (res []byte, err error) {
	hashes, err := instance.StoredBlockHeaderHashes(EmptyCallOpts(), height)
	if err != nil {
		return nil, err
	}
	return hashes[:], nil
}

func PackInitializeZkBNBParams(
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
		return nil, errors.New("[PackInitializeZkBNBParams] invalid bytes32")
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
	ZkBNBInitialize: initialize zkbnb zkbnb
*/
func ZkBNBInitialize(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *ZkBNB,
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
		return "", errors.New("[ZkBNBInitialize] invalid address")
	}
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	params, err := PackInitializeZkBNBParams(
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
	ZkBNBDepositOrLockNativeAsset: deposit or lock native asset
*/
func ZkBNBDepositOrLockNativeAsset(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *ZkBNB,
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
	ZkBNBDepositOrLockERC20: deposit or lock erc20
*/
func ZkBNBDepositOrLockERC20(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *ZkBNB,
	opType uint8, assetAddr string, amount *big.Int, accountName string,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	if !IsValidAddress(assetAddr) {
		return "", errors.New("[ZkBNBDepositOrLockERC20] invalid asset address")
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
	ZkBNBGetPendingBalance: get pending balance
*/
func ZkBNBGetPendingBalance(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *ZkBNB,
	owner string, assetAddr string,
) (*big.Int, error) {
	if !IsValidAddress(owner) || !IsValidAddress(assetAddr) {
		return nil, errors.New("[ZkBNBGetPendingBalance] invalid address")
	}
	balance, err := instance.GetPendingBalance(EmptyCallOpts(), common.HexToAddress(owner), common.HexToAddress(assetAddr))
	return balance, err
}

/*
	ZkBNBWithdrawPendingBalance: withdraw pending balance
*/
func ZkBNBWithdrawPendingBalance(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *ZkBNB,
	owner string, assetAddr string, amount *big.Int,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	if !IsValidAddress(owner) || !IsValidAddress(assetAddr) {
		return "", errors.New("[ZkBNBWithdrawPendingBalance] invalid address")
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
	ZkBNBComputeOnchainOpsHash: compute onchain ops hash
*/
func ZkBNBComputeOnchainOpsHash(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *ZkBNB,
	onchainOpsPubData []byte,
) (hashVal []byte, err error) {
	val, err := instance.ComputeOnchainOpsHash(EmptyCallOpts(), onchainOpsPubData)
	return val[:], err
}

/*
	ZkBNBCommitBlocks: commit blocks
*/
func ZkBNBCommitBlocks(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *ZkBNB,
	lastBlock StorageBlockHeader, commitBlocksInfo []ZkBNBCommitBlockInfo,
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
	ZkBNBVerifyBlocks: verify blocks
*/
func ZkBNBVerifyBlocks(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *ZkBNB,
	blockInfos []StorageBlockHeader,
	proofs []*big.Int,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.VerifyBlocks(transactOpts, blockInfos, proofs)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	ZkBNBExecuteBlocks: execute blocks
*/
func ZkBNBExecuteBlocks(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *ZkBNB,
	blocks []ZkBNBExecuteBlockInfo,
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
