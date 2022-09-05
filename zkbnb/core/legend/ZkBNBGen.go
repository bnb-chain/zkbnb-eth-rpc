// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package legend

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OldZkBNBCommitBlockInfo is an auto generated low-level Go binding around an user-defined struct.
type OldZkBNBCommitBlockInfo struct {
	NewStateRoot      [32]byte
	PublicData        []byte
	Timestamp         *big.Int
	PublicDataOffsets []uint32
	BlockNumber       uint32
	BlockSize         uint16
}

// OldZkBNBPairInfo is an auto generated low-level Go binding around an user-defined struct.
type OldZkBNBPairInfo struct {
	TokenA               common.Address
	TokenB               common.Address
	FeeRate              uint16
	TreasuryAccountIndex uint32
	TreasuryRate         uint16
}

// OldZkBNBVerifyAndExecuteBlockInfo is an auto generated low-level Go binding around an user-defined struct.
type OldZkBNBVerifyAndExecuteBlockInfo struct {
	BlockHeader              StorageStoredBlockInfo
	PendingOnchainOpsPubData [][]byte
}

// StorageStoredBlockInfo is an auto generated low-level Go binding around an user-defined struct.
type StorageStoredBlockInfo struct {
	BlockSize                    uint16
	BlockNumber                  uint32
	PriorityOperations           uint64
	PendingOnchainOperationsHash [32]byte
	Timestamp                    *big.Int
	StateRoot                    [32]byte
	Commitment                   [32]byte
}

// ZkBNBMetaData contains all meta data concerning the ZkBNB contract.
var ZkBNBMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"BlockCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"BlockVerification\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"totalBlocksVerified\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"totalBlocksCommitted\",\"type\":\"uint32\"}],\"name\":\"BlocksRevert\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"pairIndex\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"asset0Id\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"asset1Id\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"feeRate\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"treasuryAccountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"treasuryRate\",\"type\":\"uint16\"}],\"name\":\"CreateTokenPair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"assetId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accountName\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"zkbnbBlockNumber\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accountName\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"assetId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"DepositCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accountNameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nftContentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"creatorTreasuryRate\",\"type\":\"uint16\"}],\"name\":\"DepositNft\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DesertMode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"zkbnbBlockId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tokenId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"FullExitCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"NewDefaultNFTFactory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_creatorAccountNameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"_collectionId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_factoryAddress\",\"type\":\"address\"}],\"name\":\"NewNFTFactory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serialId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumTxTypes.TxType\",\"name\":\"txType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expirationBlock\",\"type\":\"uint256\"}],\"name\":\"NewPriorityRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNoticePeriod\",\"type\":\"uint256\"}],\"name\":\"NoticePeriodChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"zkbnbPubKeyX\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"zkbnbPubKeyY\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"}],\"name\":\"RegisterZNS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"pairIndex\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"feeRate\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"treasuryAccountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"treasuryRate\",\"type\":\"uint16\"}],\"name\":\"UpdateTokenPair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftL1Address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftL1TokenId\",\"type\":\"uint256\"}],\"name\":\"WithdrawNft\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"assetId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"nftIndex\",\"type\":\"uint40\"}],\"name\":\"WithdrawalNFTPending\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_ACCOUNT_INDEX\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_AMOUNT_OF_REGISTERED_ASSETS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_DEPOSIT_AMOUNT\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FUNGIBLE_ASSET_ID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_NFT_INDEX\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECURITY_COUNCIL_MEMBERS_NUMBER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHORTEST_UPGRADE_NOTICE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SPECIAL_ACCOUNT_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SPECIAL_ACCOUNT_ID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TX_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_NOTICE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWAL_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateDesertMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"blockSize\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"priorityOperations\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"pendingOnchainOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structStorage.StoredBlockInfo\",\"name\":\"_lastCommittedBlockData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint32[]\",\"name\":\"publicDataOffsets\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"blockSize\",\"type\":\"uint16\"}],\"internalType\":\"structOldZkBNB.CommitBlockInfo[]\",\"name\":\"_newBlocksData\",\"type\":\"tuple[]\"}],\"name\":\"commitBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cutUpgradeNoticePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultNFTFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint104\",\"name\":\"_amount\",\"type\":\"uint104\"},{\"internalType\":\"string\",\"name\":\"_accountName\",\"type\":\"string\"}],\"name\":\"depositBEP20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accountName\",\"type\":\"string\"}],\"name\":\"depositBNB\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accountName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_nftL1Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nftL1TokenId\",\"type\":\"uint256\"}],\"name\":\"depositNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"desertMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstPriorityRequestId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"accountNameHash\",\"type\":\"bytes32\"}],\"name\":\"getAddressByAccountNameHash\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_creatorAccountNameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_collectionId\",\"type\":\"uint32\"}],\"name\":\"getNFTFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNoticePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_assetAddr\",\"type\":\"address\"}],\"name\":\"getPendingBalance\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initializationParameters\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isReadyForUpgrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"nftFactories\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_zkbnbPubKeyX\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_zkbnbPubKeyY\",\"type\":\"bytes32\"}],\"name\":\"registerZNS\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accountName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"requestFullExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accountName\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"_nftIndex\",\"type\":\"uint32\"}],\"name\":\"requestFullExitNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"blockSize\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"priorityOperations\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"pendingOnchainOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structStorage.StoredBlockInfo[]\",\"name\":\"_blocksToRevert\",\"type\":\"tuple[]\"}],\"name\":\"revertBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractNFTFactory\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"setDefaultNFTFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"storedBlockHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBlocksCommitted\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBlocksVerified\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalOpenPriorityRequests\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTokenPairs\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_maxAmount\",\"type\":\"uint128\"}],\"name\":\"transferERC20\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"withdrawnAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"feeRate\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"treasuryAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"treasuryRate\",\"type\":\"uint16\"}],\"internalType\":\"structOldZkBNB.PairInfo\",\"name\":\"_pairInfo\",\"type\":\"tuple\"}],\"name\":\"updatePairRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newVerifierAddress\",\"type\":\"address\"}],\"name\":\"updateZkBNBVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"upgradeParameters\",\"type\":\"bytes\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeCanceled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeFinishes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeNoticePeriodStarted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradePreparationStarted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"blockSize\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"priorityOperations\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"pendingOnchainOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structStorage.StoredBlockInfo\",\"name\":\"blockHeader\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"pendingOnchainOpsPubData\",\"type\":\"bytes[]\"}],\"internalType\":\"structOldZkBNB.VerifyAndExecuteBlockInfo[]\",\"name\":\"_blocks\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_proofs\",\"type\":\"uint256[]\"}],\"name\":\"verifyAndExecuteBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_amount\",\"type\":\"uint128\"}],\"name\":\"withdrawPendingBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_nftIndex\",\"type\":\"uint40\"}],\"name\":\"withdrawPendingNFTBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ZkBNBABI is the input ABI used to generate the binding from.
// Deprecated: Use ZkBNBMetaData.ABI instead.
var ZkBNBABI = ZkBNBMetaData.ABI

// ZkBNB is an auto generated Go binding around an Ethereum contract.
type ZkBNB struct {
	ZkBNBCaller     // Read-only binding to the contract
	ZkBNBTransactor // Write-only binding to the contract
	ZkBNBFilterer   // Log filterer for contract events
}

// ZkBNBCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZkBNBCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkBNBTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZkBNBTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkBNBFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZkBNBFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkBNBSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZkBNBSession struct {
	Contract     *ZkBNB            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZkBNBCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZkBNBCallerSession struct {
	Contract *ZkBNBCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ZkBNBTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZkBNBTransactorSession struct {
	Contract     *ZkBNBTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZkBNBRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZkBNBRaw struct {
	Contract *ZkBNB // Generic contract binding to access the raw methods on
}

// ZkBNBCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZkBNBCallerRaw struct {
	Contract *ZkBNBCaller // Generic read-only contract binding to access the raw methods on
}

// ZkBNBTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZkBNBTransactorRaw struct {
	Contract *ZkBNBTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZkBNB creates a new instance of ZkBNB, bound to a specific deployed contract.
func NewZkBNB(address common.Address, backend bind.ContractBackend) (*ZkBNB, error) {
	contract, err := bindZkBNB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZkBNB{ZkBNBCaller: ZkBNBCaller{contract: contract}, ZkBNBTransactor: ZkBNBTransactor{contract: contract}, ZkBNBFilterer: ZkBNBFilterer{contract: contract}}, nil
}

// NewZkBNBCaller creates a new read-only instance of ZkBNB, bound to a specific deployed contract.
func NewZkBNBCaller(address common.Address, caller bind.ContractCaller) (*ZkBNBCaller, error) {
	contract, err := bindZkBNB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZkBNBCaller{contract: contract}, nil
}

// NewZkBNBTransactor creates a new write-only instance of ZkBNB, bound to a specific deployed contract.
func NewZkBNBTransactor(address common.Address, transactor bind.ContractTransactor) (*ZkBNBTransactor, error) {
	contract, err := bindZkBNB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZkBNBTransactor{contract: contract}, nil
}

// NewZkBNBFilterer creates a new log filterer instance of ZkBNB, bound to a specific deployed contract.
func NewZkBNBFilterer(address common.Address, filterer bind.ContractFilterer) (*ZkBNBFilterer, error) {
	contract, err := bindZkBNB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZkBNBFilterer{contract: contract}, nil
}

// bindZkBNB binds a generic wrapper to an already deployed contract.
func bindZkBNB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZkBNBABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZkBNB *ZkBNBRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZkBNB.Contract.ZkBNBCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZkBNB *ZkBNBRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkBNB.Contract.ZkBNBTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZkBNB *ZkBNBRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZkBNB.Contract.ZkBNBTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZkBNB *ZkBNBCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZkBNB.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZkBNB *ZkBNBTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkBNB.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZkBNB *ZkBNBTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZkBNB.Contract.contract.Transact(opts, method, params...)
}

// MAXACCOUNTINDEX is a free data retrieval call binding the contract method 0x437545f9.
//
// Solidity: function MAX_ACCOUNT_INDEX() view returns(uint32)
func (_ZkBNB *ZkBNBCaller) MAXACCOUNTINDEX(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "MAX_ACCOUNT_INDEX")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXACCOUNTINDEX is a free data retrieval call binding the contract method 0x437545f9.
//
// Solidity: function MAX_ACCOUNT_INDEX() view returns(uint32)
func (_ZkBNB *ZkBNBSession) MAXACCOUNTINDEX() (uint32, error) {
	return _ZkBNB.Contract.MAXACCOUNTINDEX(&_ZkBNB.CallOpts)
}

// MAXACCOUNTINDEX is a free data retrieval call binding the contract method 0x437545f9.
//
// Solidity: function MAX_ACCOUNT_INDEX() view returns(uint32)
func (_ZkBNB *ZkBNBCallerSession) MAXACCOUNTINDEX() (uint32, error) {
	return _ZkBNB.Contract.MAXACCOUNTINDEX(&_ZkBNB.CallOpts)
}

// MAXAMOUNTOFREGISTEREDASSETS is a free data retrieval call binding the contract method 0x0d360b7f.
//
// Solidity: function MAX_AMOUNT_OF_REGISTERED_ASSETS() view returns(uint16)
func (_ZkBNB *ZkBNBCaller) MAXAMOUNTOFREGISTEREDASSETS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "MAX_AMOUNT_OF_REGISTERED_ASSETS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXAMOUNTOFREGISTEREDASSETS is a free data retrieval call binding the contract method 0x0d360b7f.
//
// Solidity: function MAX_AMOUNT_OF_REGISTERED_ASSETS() view returns(uint16)
func (_ZkBNB *ZkBNBSession) MAXAMOUNTOFREGISTEREDASSETS() (uint16, error) {
	return _ZkBNB.Contract.MAXAMOUNTOFREGISTEREDASSETS(&_ZkBNB.CallOpts)
}

// MAXAMOUNTOFREGISTEREDASSETS is a free data retrieval call binding the contract method 0x0d360b7f.
//
// Solidity: function MAX_AMOUNT_OF_REGISTERED_ASSETS() view returns(uint16)
func (_ZkBNB *ZkBNBCallerSession) MAXAMOUNTOFREGISTEREDASSETS() (uint16, error) {
	return _ZkBNB.Contract.MAXAMOUNTOFREGISTEREDASSETS(&_ZkBNB.CallOpts)
}

// MAXDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x4c34a982.
//
// Solidity: function MAX_DEPOSIT_AMOUNT() view returns(uint128)
func (_ZkBNB *ZkBNBCaller) MAXDEPOSITAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "MAX_DEPOSIT_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x4c34a982.
//
// Solidity: function MAX_DEPOSIT_AMOUNT() view returns(uint128)
func (_ZkBNB *ZkBNBSession) MAXDEPOSITAMOUNT() (*big.Int, error) {
	return _ZkBNB.Contract.MAXDEPOSITAMOUNT(&_ZkBNB.CallOpts)
}

// MAXDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x4c34a982.
//
// Solidity: function MAX_DEPOSIT_AMOUNT() view returns(uint128)
func (_ZkBNB *ZkBNBCallerSession) MAXDEPOSITAMOUNT() (*big.Int, error) {
	return _ZkBNB.Contract.MAXDEPOSITAMOUNT(&_ZkBNB.CallOpts)
}

// MAXFUNGIBLEASSETID is a free data retrieval call binding the contract method 0x437da02f.
//
// Solidity: function MAX_FUNGIBLE_ASSET_ID() view returns(uint32)
func (_ZkBNB *ZkBNBCaller) MAXFUNGIBLEASSETID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "MAX_FUNGIBLE_ASSET_ID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXFUNGIBLEASSETID is a free data retrieval call binding the contract method 0x437da02f.
//
// Solidity: function MAX_FUNGIBLE_ASSET_ID() view returns(uint32)
func (_ZkBNB *ZkBNBSession) MAXFUNGIBLEASSETID() (uint32, error) {
	return _ZkBNB.Contract.MAXFUNGIBLEASSETID(&_ZkBNB.CallOpts)
}

// MAXFUNGIBLEASSETID is a free data retrieval call binding the contract method 0x437da02f.
//
// Solidity: function MAX_FUNGIBLE_ASSET_ID() view returns(uint32)
func (_ZkBNB *ZkBNBCallerSession) MAXFUNGIBLEASSETID() (uint32, error) {
	return _ZkBNB.Contract.MAXFUNGIBLEASSETID(&_ZkBNB.CallOpts)
}

// MAXNFTINDEX is a free data retrieval call binding the contract method 0x14791ad2.
//
// Solidity: function MAX_NFT_INDEX() view returns(uint40)
func (_ZkBNB *ZkBNBCaller) MAXNFTINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "MAX_NFT_INDEX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXNFTINDEX is a free data retrieval call binding the contract method 0x14791ad2.
//
// Solidity: function MAX_NFT_INDEX() view returns(uint40)
func (_ZkBNB *ZkBNBSession) MAXNFTINDEX() (*big.Int, error) {
	return _ZkBNB.Contract.MAXNFTINDEX(&_ZkBNB.CallOpts)
}

// MAXNFTINDEX is a free data retrieval call binding the contract method 0x14791ad2.
//
// Solidity: function MAX_NFT_INDEX() view returns(uint40)
func (_ZkBNB *ZkBNBCallerSession) MAXNFTINDEX() (*big.Int, error) {
	return _ZkBNB.Contract.MAXNFTINDEX(&_ZkBNB.CallOpts)
}

// SECURITYCOUNCILMEMBERSNUMBER is a free data retrieval call binding the contract method 0x4a51a71f.
//
// Solidity: function SECURITY_COUNCIL_MEMBERS_NUMBER() view returns(uint256)
func (_ZkBNB *ZkBNBCaller) SECURITYCOUNCILMEMBERSNUMBER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "SECURITY_COUNCIL_MEMBERS_NUMBER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECURITYCOUNCILMEMBERSNUMBER is a free data retrieval call binding the contract method 0x4a51a71f.
//
// Solidity: function SECURITY_COUNCIL_MEMBERS_NUMBER() view returns(uint256)
func (_ZkBNB *ZkBNBSession) SECURITYCOUNCILMEMBERSNUMBER() (*big.Int, error) {
	return _ZkBNB.Contract.SECURITYCOUNCILMEMBERSNUMBER(&_ZkBNB.CallOpts)
}

// SECURITYCOUNCILMEMBERSNUMBER is a free data retrieval call binding the contract method 0x4a51a71f.
//
// Solidity: function SECURITY_COUNCIL_MEMBERS_NUMBER() view returns(uint256)
func (_ZkBNB *ZkBNBCallerSession) SECURITYCOUNCILMEMBERSNUMBER() (*big.Int, error) {
	return _ZkBNB.Contract.SECURITYCOUNCILMEMBERSNUMBER(&_ZkBNB.CallOpts)
}

// SHORTESTUPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0x85053581.
//
// Solidity: function SHORTEST_UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_ZkBNB *ZkBNBCaller) SHORTESTUPGRADENOTICEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "SHORTEST_UPGRADE_NOTICE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SHORTESTUPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0x85053581.
//
// Solidity: function SHORTEST_UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_ZkBNB *ZkBNBSession) SHORTESTUPGRADENOTICEPERIOD() (*big.Int, error) {
	return _ZkBNB.Contract.SHORTESTUPGRADENOTICEPERIOD(&_ZkBNB.CallOpts)
}

// SHORTESTUPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0x85053581.
//
// Solidity: function SHORTEST_UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_ZkBNB *ZkBNBCallerSession) SHORTESTUPGRADENOTICEPERIOD() (*big.Int, error) {
	return _ZkBNB.Contract.SHORTESTUPGRADENOTICEPERIOD(&_ZkBNB.CallOpts)
}

// SPECIALACCOUNTADDRESS is a free data retrieval call binding the contract method 0x7ea399c1.
//
// Solidity: function SPECIAL_ACCOUNT_ADDRESS() view returns(address)
func (_ZkBNB *ZkBNBCaller) SPECIALACCOUNTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "SPECIAL_ACCOUNT_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SPECIALACCOUNTADDRESS is a free data retrieval call binding the contract method 0x7ea399c1.
//
// Solidity: function SPECIAL_ACCOUNT_ADDRESS() view returns(address)
func (_ZkBNB *ZkBNBSession) SPECIALACCOUNTADDRESS() (common.Address, error) {
	return _ZkBNB.Contract.SPECIALACCOUNTADDRESS(&_ZkBNB.CallOpts)
}

// SPECIALACCOUNTADDRESS is a free data retrieval call binding the contract method 0x7ea399c1.
//
// Solidity: function SPECIAL_ACCOUNT_ADDRESS() view returns(address)
func (_ZkBNB *ZkBNBCallerSession) SPECIALACCOUNTADDRESS() (common.Address, error) {
	return _ZkBNB.Contract.SPECIALACCOUNTADDRESS(&_ZkBNB.CallOpts)
}

// SPECIALACCOUNTID is a free data retrieval call binding the contract method 0x4242d5b3.
//
// Solidity: function SPECIAL_ACCOUNT_ID() view returns(uint32)
func (_ZkBNB *ZkBNBCaller) SPECIALACCOUNTID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "SPECIAL_ACCOUNT_ID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SPECIALACCOUNTID is a free data retrieval call binding the contract method 0x4242d5b3.
//
// Solidity: function SPECIAL_ACCOUNT_ID() view returns(uint32)
func (_ZkBNB *ZkBNBSession) SPECIALACCOUNTID() (uint32, error) {
	return _ZkBNB.Contract.SPECIALACCOUNTID(&_ZkBNB.CallOpts)
}

// SPECIALACCOUNTID is a free data retrieval call binding the contract method 0x4242d5b3.
//
// Solidity: function SPECIAL_ACCOUNT_ID() view returns(uint32)
func (_ZkBNB *ZkBNBCallerSession) SPECIALACCOUNTID() (uint32, error) {
	return _ZkBNB.Contract.SPECIALACCOUNTID(&_ZkBNB.CallOpts)
}

// TXSIZE is a free data retrieval call binding the contract method 0xe6e3c012.
//
// Solidity: function TX_SIZE() view returns(uint256)
func (_ZkBNB *ZkBNBCaller) TXSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "TX_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TXSIZE is a free data retrieval call binding the contract method 0xe6e3c012.
//
// Solidity: function TX_SIZE() view returns(uint256)
func (_ZkBNB *ZkBNBSession) TXSIZE() (*big.Int, error) {
	return _ZkBNB.Contract.TXSIZE(&_ZkBNB.CallOpts)
}

// TXSIZE is a free data retrieval call binding the contract method 0xe6e3c012.
//
// Solidity: function TX_SIZE() view returns(uint256)
func (_ZkBNB *ZkBNBCallerSession) TXSIZE() (*big.Int, error) {
	return _ZkBNB.Contract.TXSIZE(&_ZkBNB.CallOpts)
}

// UPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0xcc375fb7.
//
// Solidity: function UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_ZkBNB *ZkBNBCaller) UPGRADENOTICEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "UPGRADE_NOTICE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0xcc375fb7.
//
// Solidity: function UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_ZkBNB *ZkBNBSession) UPGRADENOTICEPERIOD() (*big.Int, error) {
	return _ZkBNB.Contract.UPGRADENOTICEPERIOD(&_ZkBNB.CallOpts)
}

// UPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0xcc375fb7.
//
// Solidity: function UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_ZkBNB *ZkBNBCallerSession) UPGRADENOTICEPERIOD() (*big.Int, error) {
	return _ZkBNB.Contract.UPGRADENOTICEPERIOD(&_ZkBNB.CallOpts)
}

// WITHDRAWALGASLIMIT is a free data retrieval call binding the contract method 0xc701f955.
//
// Solidity: function WITHDRAWAL_GAS_LIMIT() view returns(uint256)
func (_ZkBNB *ZkBNBCaller) WITHDRAWALGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "WITHDRAWAL_GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WITHDRAWALGASLIMIT is a free data retrieval call binding the contract method 0xc701f955.
//
// Solidity: function WITHDRAWAL_GAS_LIMIT() view returns(uint256)
func (_ZkBNB *ZkBNBSession) WITHDRAWALGASLIMIT() (*big.Int, error) {
	return _ZkBNB.Contract.WITHDRAWALGASLIMIT(&_ZkBNB.CallOpts)
}

// WITHDRAWALGASLIMIT is a free data retrieval call binding the contract method 0xc701f955.
//
// Solidity: function WITHDRAWAL_GAS_LIMIT() view returns(uint256)
func (_ZkBNB *ZkBNBCallerSession) WITHDRAWALGASLIMIT() (*big.Int, error) {
	return _ZkBNB.Contract.WITHDRAWALGASLIMIT(&_ZkBNB.CallOpts)
}

// DefaultNFTFactory is a free data retrieval call binding the contract method 0x940f19c0.
//
// Solidity: function defaultNFTFactory() view returns(address)
func (_ZkBNB *ZkBNBCaller) DefaultNFTFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "defaultNFTFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultNFTFactory is a free data retrieval call binding the contract method 0x940f19c0.
//
// Solidity: function defaultNFTFactory() view returns(address)
func (_ZkBNB *ZkBNBSession) DefaultNFTFactory() (common.Address, error) {
	return _ZkBNB.Contract.DefaultNFTFactory(&_ZkBNB.CallOpts)
}

// DefaultNFTFactory is a free data retrieval call binding the contract method 0x940f19c0.
//
// Solidity: function defaultNFTFactory() view returns(address)
func (_ZkBNB *ZkBNBCallerSession) DefaultNFTFactory() (common.Address, error) {
	return _ZkBNB.Contract.DefaultNFTFactory(&_ZkBNB.CallOpts)
}

// DesertMode is a free data retrieval call binding the contract method 0x02cfb563.
//
// Solidity: function desertMode() view returns(bool)
func (_ZkBNB *ZkBNBCaller) DesertMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "desertMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DesertMode is a free data retrieval call binding the contract method 0x02cfb563.
//
// Solidity: function desertMode() view returns(bool)
func (_ZkBNB *ZkBNBSession) DesertMode() (bool, error) {
	return _ZkBNB.Contract.DesertMode(&_ZkBNB.CallOpts)
}

// DesertMode is a free data retrieval call binding the contract method 0x02cfb563.
//
// Solidity: function desertMode() view returns(bool)
func (_ZkBNB *ZkBNBCallerSession) DesertMode() (bool, error) {
	return _ZkBNB.Contract.DesertMode(&_ZkBNB.CallOpts)
}

// FirstPriorityRequestId is a free data retrieval call binding the contract method 0x67708dae.
//
// Solidity: function firstPriorityRequestId() view returns(uint64)
func (_ZkBNB *ZkBNBCaller) FirstPriorityRequestId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "firstPriorityRequestId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// FirstPriorityRequestId is a free data retrieval call binding the contract method 0x67708dae.
//
// Solidity: function firstPriorityRequestId() view returns(uint64)
func (_ZkBNB *ZkBNBSession) FirstPriorityRequestId() (uint64, error) {
	return _ZkBNB.Contract.FirstPriorityRequestId(&_ZkBNB.CallOpts)
}

// FirstPriorityRequestId is a free data retrieval call binding the contract method 0x67708dae.
//
// Solidity: function firstPriorityRequestId() view returns(uint64)
func (_ZkBNB *ZkBNBCallerSession) FirstPriorityRequestId() (uint64, error) {
	return _ZkBNB.Contract.FirstPriorityRequestId(&_ZkBNB.CallOpts)
}

// GetAddressByAccountNameHash is a free data retrieval call binding the contract method 0x0b8f1c0c.
//
// Solidity: function getAddressByAccountNameHash(bytes32 accountNameHash) view returns(address)
func (_ZkBNB *ZkBNBCaller) GetAddressByAccountNameHash(opts *bind.CallOpts, accountNameHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "getAddressByAccountNameHash", accountNameHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressByAccountNameHash is a free data retrieval call binding the contract method 0x0b8f1c0c.
//
// Solidity: function getAddressByAccountNameHash(bytes32 accountNameHash) view returns(address)
func (_ZkBNB *ZkBNBSession) GetAddressByAccountNameHash(accountNameHash [32]byte) (common.Address, error) {
	return _ZkBNB.Contract.GetAddressByAccountNameHash(&_ZkBNB.CallOpts, accountNameHash)
}

// GetAddressByAccountNameHash is a free data retrieval call binding the contract method 0x0b8f1c0c.
//
// Solidity: function getAddressByAccountNameHash(bytes32 accountNameHash) view returns(address)
func (_ZkBNB *ZkBNBCallerSession) GetAddressByAccountNameHash(accountNameHash [32]byte) (common.Address, error) {
	return _ZkBNB.Contract.GetAddressByAccountNameHash(&_ZkBNB.CallOpts, accountNameHash)
}

// GetNFTFactory is a free data retrieval call binding the contract method 0x76b4da60.
//
// Solidity: function getNFTFactory(bytes32 _creatorAccountNameHash, uint32 _collectionId) view returns(address)
func (_ZkBNB *ZkBNBCaller) GetNFTFactory(opts *bind.CallOpts, _creatorAccountNameHash [32]byte, _collectionId uint32) (common.Address, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "getNFTFactory", _creatorAccountNameHash, _collectionId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNFTFactory is a free data retrieval call binding the contract method 0x76b4da60.
//
// Solidity: function getNFTFactory(bytes32 _creatorAccountNameHash, uint32 _collectionId) view returns(address)
func (_ZkBNB *ZkBNBSession) GetNFTFactory(_creatorAccountNameHash [32]byte, _collectionId uint32) (common.Address, error) {
	return _ZkBNB.Contract.GetNFTFactory(&_ZkBNB.CallOpts, _creatorAccountNameHash, _collectionId)
}

// GetNFTFactory is a free data retrieval call binding the contract method 0x76b4da60.
//
// Solidity: function getNFTFactory(bytes32 _creatorAccountNameHash, uint32 _collectionId) view returns(address)
func (_ZkBNB *ZkBNBCallerSession) GetNFTFactory(_creatorAccountNameHash [32]byte, _collectionId uint32) (common.Address, error) {
	return _ZkBNB.Contract.GetNFTFactory(&_ZkBNB.CallOpts, _creatorAccountNameHash, _collectionId)
}

// GetNoticePeriod is a free data retrieval call binding the contract method 0x2a3174f4.
//
// Solidity: function getNoticePeriod() pure returns(uint256)
func (_ZkBNB *ZkBNBCaller) GetNoticePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "getNoticePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNoticePeriod is a free data retrieval call binding the contract method 0x2a3174f4.
//
// Solidity: function getNoticePeriod() pure returns(uint256)
func (_ZkBNB *ZkBNBSession) GetNoticePeriod() (*big.Int, error) {
	return _ZkBNB.Contract.GetNoticePeriod(&_ZkBNB.CallOpts)
}

// GetNoticePeriod is a free data retrieval call binding the contract method 0x2a3174f4.
//
// Solidity: function getNoticePeriod() pure returns(uint256)
func (_ZkBNB *ZkBNBCallerSession) GetNoticePeriod() (*big.Int, error) {
	return _ZkBNB.Contract.GetNoticePeriod(&_ZkBNB.CallOpts)
}

// GetPendingBalance is a free data retrieval call binding the contract method 0x5aca41f6.
//
// Solidity: function getPendingBalance(address _address, address _assetAddr) view returns(uint128)
func (_ZkBNB *ZkBNBCaller) GetPendingBalance(opts *bind.CallOpts, _address common.Address, _assetAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "getPendingBalance", _address, _assetAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingBalance is a free data retrieval call binding the contract method 0x5aca41f6.
//
// Solidity: function getPendingBalance(address _address, address _assetAddr) view returns(uint128)
func (_ZkBNB *ZkBNBSession) GetPendingBalance(_address common.Address, _assetAddr common.Address) (*big.Int, error) {
	return _ZkBNB.Contract.GetPendingBalance(&_ZkBNB.CallOpts, _address, _assetAddr)
}

// GetPendingBalance is a free data retrieval call binding the contract method 0x5aca41f6.
//
// Solidity: function getPendingBalance(address _address, address _assetAddr) view returns(uint128)
func (_ZkBNB *ZkBNBCallerSession) GetPendingBalance(_address common.Address, _assetAddr common.Address) (*big.Int, error) {
	return _ZkBNB.Contract.GetPendingBalance(&_ZkBNB.CallOpts, _address, _assetAddr)
}

// IsReadyForUpgrade is a free data retrieval call binding the contract method 0x8773334c.
//
// Solidity: function isReadyForUpgrade() view returns(bool)
func (_ZkBNB *ZkBNBCaller) IsReadyForUpgrade(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "isReadyForUpgrade")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReadyForUpgrade is a free data retrieval call binding the contract method 0x8773334c.
//
// Solidity: function isReadyForUpgrade() view returns(bool)
func (_ZkBNB *ZkBNBSession) IsReadyForUpgrade() (bool, error) {
	return _ZkBNB.Contract.IsReadyForUpgrade(&_ZkBNB.CallOpts)
}

// IsReadyForUpgrade is a free data retrieval call binding the contract method 0x8773334c.
//
// Solidity: function isReadyForUpgrade() view returns(bool)
func (_ZkBNB *ZkBNBCallerSession) IsReadyForUpgrade() (bool, error) {
	return _ZkBNB.Contract.IsReadyForUpgrade(&_ZkBNB.CallOpts)
}

// NftFactories is a free data retrieval call binding the contract method 0x16335e06.
//
// Solidity: function nftFactories(bytes32 , uint32 ) view returns(address)
func (_ZkBNB *ZkBNBCaller) NftFactories(opts *bind.CallOpts, arg0 [32]byte, arg1 uint32) (common.Address, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "nftFactories", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftFactories is a free data retrieval call binding the contract method 0x16335e06.
//
// Solidity: function nftFactories(bytes32 , uint32 ) view returns(address)
func (_ZkBNB *ZkBNBSession) NftFactories(arg0 [32]byte, arg1 uint32) (common.Address, error) {
	return _ZkBNB.Contract.NftFactories(&_ZkBNB.CallOpts, arg0, arg1)
}

// NftFactories is a free data retrieval call binding the contract method 0x16335e06.
//
// Solidity: function nftFactories(bytes32 , uint32 ) view returns(address)
func (_ZkBNB *ZkBNBCallerSession) NftFactories(arg0 [32]byte, arg1 uint32) (common.Address, error) {
	return _ZkBNB.Contract.NftFactories(&_ZkBNB.CallOpts, arg0, arg1)
}

// StateRoot is a free data retrieval call binding the contract method 0x9588eca2.
//
// Solidity: function stateRoot() view returns(bytes32)
func (_ZkBNB *ZkBNBCaller) StateRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "stateRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateRoot is a free data retrieval call binding the contract method 0x9588eca2.
//
// Solidity: function stateRoot() view returns(bytes32)
func (_ZkBNB *ZkBNBSession) StateRoot() ([32]byte, error) {
	return _ZkBNB.Contract.StateRoot(&_ZkBNB.CallOpts)
}

// StateRoot is a free data retrieval call binding the contract method 0x9588eca2.
//
// Solidity: function stateRoot() view returns(bytes32)
func (_ZkBNB *ZkBNBCallerSession) StateRoot() ([32]byte, error) {
	return _ZkBNB.Contract.StateRoot(&_ZkBNB.CallOpts)
}

// StoredBlockHashes is a free data retrieval call binding the contract method 0x9ba0d146.
//
// Solidity: function storedBlockHashes(uint32 ) view returns(bytes32)
func (_ZkBNB *ZkBNBCaller) StoredBlockHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "storedBlockHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StoredBlockHashes is a free data retrieval call binding the contract method 0x9ba0d146.
//
// Solidity: function storedBlockHashes(uint32 ) view returns(bytes32)
func (_ZkBNB *ZkBNBSession) StoredBlockHashes(arg0 uint32) ([32]byte, error) {
	return _ZkBNB.Contract.StoredBlockHashes(&_ZkBNB.CallOpts, arg0)
}

// StoredBlockHashes is a free data retrieval call binding the contract method 0x9ba0d146.
//
// Solidity: function storedBlockHashes(uint32 ) view returns(bytes32)
func (_ZkBNB *ZkBNBCallerSession) StoredBlockHashes(arg0 uint32) ([32]byte, error) {
	return _ZkBNB.Contract.StoredBlockHashes(&_ZkBNB.CallOpts, arg0)
}

// TotalBlocksCommitted is a free data retrieval call binding the contract method 0xfaf4d8cb.
//
// Solidity: function totalBlocksCommitted() view returns(uint32)
func (_ZkBNB *ZkBNBCaller) TotalBlocksCommitted(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "totalBlocksCommitted")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TotalBlocksCommitted is a free data retrieval call binding the contract method 0xfaf4d8cb.
//
// Solidity: function totalBlocksCommitted() view returns(uint32)
func (_ZkBNB *ZkBNBSession) TotalBlocksCommitted() (uint32, error) {
	return _ZkBNB.Contract.TotalBlocksCommitted(&_ZkBNB.CallOpts)
}

// TotalBlocksCommitted is a free data retrieval call binding the contract method 0xfaf4d8cb.
//
// Solidity: function totalBlocksCommitted() view returns(uint32)
func (_ZkBNB *ZkBNBCallerSession) TotalBlocksCommitted() (uint32, error) {
	return _ZkBNB.Contract.TotalBlocksCommitted(&_ZkBNB.CallOpts)
}

// TotalBlocksVerified is a free data retrieval call binding the contract method 0x2d24006c.
//
// Solidity: function totalBlocksVerified() view returns(uint32)
func (_ZkBNB *ZkBNBCaller) TotalBlocksVerified(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "totalBlocksVerified")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TotalBlocksVerified is a free data retrieval call binding the contract method 0x2d24006c.
//
// Solidity: function totalBlocksVerified() view returns(uint32)
func (_ZkBNB *ZkBNBSession) TotalBlocksVerified() (uint32, error) {
	return _ZkBNB.Contract.TotalBlocksVerified(&_ZkBNB.CallOpts)
}

// TotalBlocksVerified is a free data retrieval call binding the contract method 0x2d24006c.
//
// Solidity: function totalBlocksVerified() view returns(uint32)
func (_ZkBNB *ZkBNBCallerSession) TotalBlocksVerified() (uint32, error) {
	return _ZkBNB.Contract.TotalBlocksVerified(&_ZkBNB.CallOpts)
}

// TotalOpenPriorityRequests is a free data retrieval call binding the contract method 0xc57b22be.
//
// Solidity: function totalOpenPriorityRequests() view returns(uint64)
func (_ZkBNB *ZkBNBCaller) TotalOpenPriorityRequests(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "totalOpenPriorityRequests")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalOpenPriorityRequests is a free data retrieval call binding the contract method 0xc57b22be.
//
// Solidity: function totalOpenPriorityRequests() view returns(uint64)
func (_ZkBNB *ZkBNBSession) TotalOpenPriorityRequests() (uint64, error) {
	return _ZkBNB.Contract.TotalOpenPriorityRequests(&_ZkBNB.CallOpts)
}

// TotalOpenPriorityRequests is a free data retrieval call binding the contract method 0xc57b22be.
//
// Solidity: function totalOpenPriorityRequests() view returns(uint64)
func (_ZkBNB *ZkBNBCallerSession) TotalOpenPriorityRequests() (uint64, error) {
	return _ZkBNB.Contract.TotalOpenPriorityRequests(&_ZkBNB.CallOpts)
}

// TotalTokenPairs is a free data retrieval call binding the contract method 0xd0ad718d.
//
// Solidity: function totalTokenPairs() view returns(uint16)
func (_ZkBNB *ZkBNBCaller) TotalTokenPairs(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "totalTokenPairs")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// TotalTokenPairs is a free data retrieval call binding the contract method 0xd0ad718d.
//
// Solidity: function totalTokenPairs() view returns(uint16)
func (_ZkBNB *ZkBNBSession) TotalTokenPairs() (uint16, error) {
	return _ZkBNB.Contract.TotalTokenPairs(&_ZkBNB.CallOpts)
}

// TotalTokenPairs is a free data retrieval call binding the contract method 0xd0ad718d.
//
// Solidity: function totalTokenPairs() view returns(uint16)
func (_ZkBNB *ZkBNBCallerSession) TotalTokenPairs() (uint16, error) {
	return _ZkBNB.Contract.TotalTokenPairs(&_ZkBNB.CallOpts)
}

// ActivateDesertMode is a paid mutator transaction binding the contract method 0x22b22256.
//
// Solidity: function activateDesertMode() returns(bool)
func (_ZkBNB *ZkBNBTransactor) ActivateDesertMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "activateDesertMode")
}

// ActivateDesertMode is a paid mutator transaction binding the contract method 0x22b22256.
//
// Solidity: function activateDesertMode() returns(bool)
func (_ZkBNB *ZkBNBSession) ActivateDesertMode() (*types.Transaction, error) {
	return _ZkBNB.Contract.ActivateDesertMode(&_ZkBNB.TransactOpts)
}

// ActivateDesertMode is a paid mutator transaction binding the contract method 0x22b22256.
//
// Solidity: function activateDesertMode() returns(bool)
func (_ZkBNB *ZkBNBTransactorSession) ActivateDesertMode() (*types.Transaction, error) {
	return _ZkBNB.Contract.ActivateDesertMode(&_ZkBNB.TransactOpts)
}

// CommitBlocks is a paid mutator transaction binding the contract method 0x2a54f972.
//
// Solidity: function commitBlocks((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32) _lastCommittedBlockData, (bytes32,bytes,uint256,uint32[],uint32,uint16)[] _newBlocksData) returns()
func (_ZkBNB *ZkBNBTransactor) CommitBlocks(opts *bind.TransactOpts, _lastCommittedBlockData StorageStoredBlockInfo, _newBlocksData []OldZkBNBCommitBlockInfo) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "commitBlocks", _lastCommittedBlockData, _newBlocksData)
}

// CommitBlocks is a paid mutator transaction binding the contract method 0x2a54f972.
//
// Solidity: function commitBlocks((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32) _lastCommittedBlockData, (bytes32,bytes,uint256,uint32[],uint32,uint16)[] _newBlocksData) returns()
func (_ZkBNB *ZkBNBSession) CommitBlocks(_lastCommittedBlockData StorageStoredBlockInfo, _newBlocksData []OldZkBNBCommitBlockInfo) (*types.Transaction, error) {
	return _ZkBNB.Contract.CommitBlocks(&_ZkBNB.TransactOpts, _lastCommittedBlockData, _newBlocksData)
}

// CommitBlocks is a paid mutator transaction binding the contract method 0x2a54f972.
//
// Solidity: function commitBlocks((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32) _lastCommittedBlockData, (bytes32,bytes,uint256,uint32[],uint32,uint16)[] _newBlocksData) returns()
func (_ZkBNB *ZkBNBTransactorSession) CommitBlocks(_lastCommittedBlockData StorageStoredBlockInfo, _newBlocksData []OldZkBNBCommitBlockInfo) (*types.Transaction, error) {
	return _ZkBNB.Contract.CommitBlocks(&_ZkBNB.TransactOpts, _lastCommittedBlockData, _newBlocksData)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address _tokenA, address _tokenB) returns()
func (_ZkBNB *ZkBNBTransactor) CreatePair(opts *bind.TransactOpts, _tokenA common.Address, _tokenB common.Address) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "createPair", _tokenA, _tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address _tokenA, address _tokenB) returns()
func (_ZkBNB *ZkBNBSession) CreatePair(_tokenA common.Address, _tokenB common.Address) (*types.Transaction, error) {
	return _ZkBNB.Contract.CreatePair(&_ZkBNB.TransactOpts, _tokenA, _tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address _tokenA, address _tokenB) returns()
func (_ZkBNB *ZkBNBTransactorSession) CreatePair(_tokenA common.Address, _tokenB common.Address) (*types.Transaction, error) {
	return _ZkBNB.Contract.CreatePair(&_ZkBNB.TransactOpts, _tokenA, _tokenB)
}

// CutUpgradeNoticePeriod is a paid mutator transaction binding the contract method 0x3e71e1e7.
//
// Solidity: function cutUpgradeNoticePeriod() returns()
func (_ZkBNB *ZkBNBTransactor) CutUpgradeNoticePeriod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "cutUpgradeNoticePeriod")
}

// CutUpgradeNoticePeriod is a paid mutator transaction binding the contract method 0x3e71e1e7.
//
// Solidity: function cutUpgradeNoticePeriod() returns()
func (_ZkBNB *ZkBNBSession) CutUpgradeNoticePeriod() (*types.Transaction, error) {
	return _ZkBNB.Contract.CutUpgradeNoticePeriod(&_ZkBNB.TransactOpts)
}

// CutUpgradeNoticePeriod is a paid mutator transaction binding the contract method 0x3e71e1e7.
//
// Solidity: function cutUpgradeNoticePeriod() returns()
func (_ZkBNB *ZkBNBTransactorSession) CutUpgradeNoticePeriod() (*types.Transaction, error) {
	return _ZkBNB.Contract.CutUpgradeNoticePeriod(&_ZkBNB.TransactOpts)
}

// DepositBEP20 is a paid mutator transaction binding the contract method 0x1caf5d25.
//
// Solidity: function depositBEP20(address _token, uint104 _amount, string _accountName) returns()
func (_ZkBNB *ZkBNBTransactor) DepositBEP20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _accountName string) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "depositBEP20", _token, _amount, _accountName)
}

// DepositBEP20 is a paid mutator transaction binding the contract method 0x1caf5d25.
//
// Solidity: function depositBEP20(address _token, uint104 _amount, string _accountName) returns()
func (_ZkBNB *ZkBNBSession) DepositBEP20(_token common.Address, _amount *big.Int, _accountName string) (*types.Transaction, error) {
	return _ZkBNB.Contract.DepositBEP20(&_ZkBNB.TransactOpts, _token, _amount, _accountName)
}

// DepositBEP20 is a paid mutator transaction binding the contract method 0x1caf5d25.
//
// Solidity: function depositBEP20(address _token, uint104 _amount, string _accountName) returns()
func (_ZkBNB *ZkBNBTransactorSession) DepositBEP20(_token common.Address, _amount *big.Int, _accountName string) (*types.Transaction, error) {
	return _ZkBNB.Contract.DepositBEP20(&_ZkBNB.TransactOpts, _token, _amount, _accountName)
}

// DepositBNB is a paid mutator transaction binding the contract method 0x51344683.
//
// Solidity: function depositBNB(string _accountName) payable returns()
func (_ZkBNB *ZkBNBTransactor) DepositBNB(opts *bind.TransactOpts, _accountName string) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "depositBNB", _accountName)
}

// DepositBNB is a paid mutator transaction binding the contract method 0x51344683.
//
// Solidity: function depositBNB(string _accountName) payable returns()
func (_ZkBNB *ZkBNBSession) DepositBNB(_accountName string) (*types.Transaction, error) {
	return _ZkBNB.Contract.DepositBNB(&_ZkBNB.TransactOpts, _accountName)
}

// DepositBNB is a paid mutator transaction binding the contract method 0x51344683.
//
// Solidity: function depositBNB(string _accountName) payable returns()
func (_ZkBNB *ZkBNBTransactorSession) DepositBNB(_accountName string) (*types.Transaction, error) {
	return _ZkBNB.Contract.DepositBNB(&_ZkBNB.TransactOpts, _accountName)
}

// DepositNft is a paid mutator transaction binding the contract method 0xfb99514b.
//
// Solidity: function depositNft(string _accountName, address _nftL1Address, uint256 _nftL1TokenId) returns()
func (_ZkBNB *ZkBNBTransactor) DepositNft(opts *bind.TransactOpts, _accountName string, _nftL1Address common.Address, _nftL1TokenId *big.Int) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "depositNft", _accountName, _nftL1Address, _nftL1TokenId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xfb99514b.
//
// Solidity: function depositNft(string _accountName, address _nftL1Address, uint256 _nftL1TokenId) returns()
func (_ZkBNB *ZkBNBSession) DepositNft(_accountName string, _nftL1Address common.Address, _nftL1TokenId *big.Int) (*types.Transaction, error) {
	return _ZkBNB.Contract.DepositNft(&_ZkBNB.TransactOpts, _accountName, _nftL1Address, _nftL1TokenId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xfb99514b.
//
// Solidity: function depositNft(string _accountName, address _nftL1Address, uint256 _nftL1TokenId) returns()
func (_ZkBNB *ZkBNBTransactorSession) DepositNft(_accountName string, _nftL1Address common.Address, _nftL1TokenId *big.Int) (*types.Transaction, error) {
	return _ZkBNB.Contract.DepositNft(&_ZkBNB.TransactOpts, _accountName, _nftL1Address, _nftL1TokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializationParameters) returns()
func (_ZkBNB *ZkBNBTransactor) Initialize(opts *bind.TransactOpts, initializationParameters []byte) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "initialize", initializationParameters)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializationParameters) returns()
func (_ZkBNB *ZkBNBSession) Initialize(initializationParameters []byte) (*types.Transaction, error) {
	return _ZkBNB.Contract.Initialize(&_ZkBNB.TransactOpts, initializationParameters)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializationParameters) returns()
func (_ZkBNB *ZkBNBTransactorSession) Initialize(initializationParameters []byte) (*types.Transaction, error) {
	return _ZkBNB.Contract.Initialize(&_ZkBNB.TransactOpts, initializationParameters)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_ZkBNB *ZkBNBTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_ZkBNB *ZkBNBSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ZkBNB.Contract.OnERC721Received(&_ZkBNB.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_ZkBNB *ZkBNBTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ZkBNB.Contract.OnERC721Received(&_ZkBNB.TransactOpts, operator, from, tokenId, data)
}

// RegisterZNS is a paid mutator transaction binding the contract method 0x3fdeb67d.
//
// Solidity: function registerZNS(string _name, address _owner, bytes32 _zkbnbPubKeyX, bytes32 _zkbnbPubKeyY) payable returns()
func (_ZkBNB *ZkBNBTransactor) RegisterZNS(opts *bind.TransactOpts, _name string, _owner common.Address, _zkbnbPubKeyX [32]byte, _zkbnbPubKeyY [32]byte) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "registerZNS", _name, _owner, _zkbnbPubKeyX, _zkbnbPubKeyY)
}

// RegisterZNS is a paid mutator transaction binding the contract method 0x3fdeb67d.
//
// Solidity: function registerZNS(string _name, address _owner, bytes32 _zkbnbPubKeyX, bytes32 _zkbnbPubKeyY) payable returns()
func (_ZkBNB *ZkBNBSession) RegisterZNS(_name string, _owner common.Address, _zkbnbPubKeyX [32]byte, _zkbnbPubKeyY [32]byte) (*types.Transaction, error) {
	return _ZkBNB.Contract.RegisterZNS(&_ZkBNB.TransactOpts, _name, _owner, _zkbnbPubKeyX, _zkbnbPubKeyY)
}

// RegisterZNS is a paid mutator transaction binding the contract method 0x3fdeb67d.
//
// Solidity: function registerZNS(string _name, address _owner, bytes32 _zkbnbPubKeyX, bytes32 _zkbnbPubKeyY) payable returns()
func (_ZkBNB *ZkBNBTransactorSession) RegisterZNS(_name string, _owner common.Address, _zkbnbPubKeyX [32]byte, _zkbnbPubKeyY [32]byte) (*types.Transaction, error) {
	return _ZkBNB.Contract.RegisterZNS(&_ZkBNB.TransactOpts, _name, _owner, _zkbnbPubKeyX, _zkbnbPubKeyY)
}

// RequestFullExit is a paid mutator transaction binding the contract method 0x8da64c7f.
//
// Solidity: function requestFullExit(string _accountName, address _asset) returns()
func (_ZkBNB *ZkBNBTransactor) RequestFullExit(opts *bind.TransactOpts, _accountName string, _asset common.Address) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "requestFullExit", _accountName, _asset)
}

// RequestFullExit is a paid mutator transaction binding the contract method 0x8da64c7f.
//
// Solidity: function requestFullExit(string _accountName, address _asset) returns()
func (_ZkBNB *ZkBNBSession) RequestFullExit(_accountName string, _asset common.Address) (*types.Transaction, error) {
	return _ZkBNB.Contract.RequestFullExit(&_ZkBNB.TransactOpts, _accountName, _asset)
}

// RequestFullExit is a paid mutator transaction binding the contract method 0x8da64c7f.
//
// Solidity: function requestFullExit(string _accountName, address _asset) returns()
func (_ZkBNB *ZkBNBTransactorSession) RequestFullExit(_accountName string, _asset common.Address) (*types.Transaction, error) {
	return _ZkBNB.Contract.RequestFullExit(&_ZkBNB.TransactOpts, _accountName, _asset)
}

// RequestFullExitNft is a paid mutator transaction binding the contract method 0x1bd24317.
//
// Solidity: function requestFullExitNft(string _accountName, uint32 _nftIndex) returns()
func (_ZkBNB *ZkBNBTransactor) RequestFullExitNft(opts *bind.TransactOpts, _accountName string, _nftIndex uint32) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "requestFullExitNft", _accountName, _nftIndex)
}

// RequestFullExitNft is a paid mutator transaction binding the contract method 0x1bd24317.
//
// Solidity: function requestFullExitNft(string _accountName, uint32 _nftIndex) returns()
func (_ZkBNB *ZkBNBSession) RequestFullExitNft(_accountName string, _nftIndex uint32) (*types.Transaction, error) {
	return _ZkBNB.Contract.RequestFullExitNft(&_ZkBNB.TransactOpts, _accountName, _nftIndex)
}

// RequestFullExitNft is a paid mutator transaction binding the contract method 0x1bd24317.
//
// Solidity: function requestFullExitNft(string _accountName, uint32 _nftIndex) returns()
func (_ZkBNB *ZkBNBTransactorSession) RequestFullExitNft(_accountName string, _nftIndex uint32) (*types.Transaction, error) {
	return _ZkBNB.Contract.RequestFullExitNft(&_ZkBNB.TransactOpts, _accountName, _nftIndex)
}

// RevertBlocks is a paid mutator transaction binding the contract method 0x97a2eabe.
//
// Solidity: function revertBlocks((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32)[] _blocksToRevert) returns()
func (_ZkBNB *ZkBNBTransactor) RevertBlocks(opts *bind.TransactOpts, _blocksToRevert []StorageStoredBlockInfo) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "revertBlocks", _blocksToRevert)
}

// RevertBlocks is a paid mutator transaction binding the contract method 0x97a2eabe.
//
// Solidity: function revertBlocks((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32)[] _blocksToRevert) returns()
func (_ZkBNB *ZkBNBSession) RevertBlocks(_blocksToRevert []StorageStoredBlockInfo) (*types.Transaction, error) {
	return _ZkBNB.Contract.RevertBlocks(&_ZkBNB.TransactOpts, _blocksToRevert)
}

// RevertBlocks is a paid mutator transaction binding the contract method 0x97a2eabe.
//
// Solidity: function revertBlocks((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32)[] _blocksToRevert) returns()
func (_ZkBNB *ZkBNBTransactorSession) RevertBlocks(_blocksToRevert []StorageStoredBlockInfo) (*types.Transaction, error) {
	return _ZkBNB.Contract.RevertBlocks(&_ZkBNB.TransactOpts, _blocksToRevert)
}

// SetDefaultNFTFactory is a paid mutator transaction binding the contract method 0xce09e20d.
//
// Solidity: function setDefaultNFTFactory(address _factory) returns()
func (_ZkBNB *ZkBNBTransactor) SetDefaultNFTFactory(opts *bind.TransactOpts, _factory common.Address) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "setDefaultNFTFactory", _factory)
}

// SetDefaultNFTFactory is a paid mutator transaction binding the contract method 0xce09e20d.
//
// Solidity: function setDefaultNFTFactory(address _factory) returns()
func (_ZkBNB *ZkBNBSession) SetDefaultNFTFactory(_factory common.Address) (*types.Transaction, error) {
	return _ZkBNB.Contract.SetDefaultNFTFactory(&_ZkBNB.TransactOpts, _factory)
}

// SetDefaultNFTFactory is a paid mutator transaction binding the contract method 0xce09e20d.
//
// Solidity: function setDefaultNFTFactory(address _factory) returns()
func (_ZkBNB *ZkBNBTransactorSession) SetDefaultNFTFactory(_factory common.Address) (*types.Transaction, error) {
	return _ZkBNB.Contract.SetDefaultNFTFactory(&_ZkBNB.TransactOpts, _factory)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x68809a23.
//
// Solidity: function transferERC20(address _token, address _to, uint128 _amount, uint128 _maxAmount) returns(uint128 withdrawnAmount)
func (_ZkBNB *ZkBNBTransactor) TransferERC20(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _maxAmount *big.Int) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "transferERC20", _token, _to, _amount, _maxAmount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x68809a23.
//
// Solidity: function transferERC20(address _token, address _to, uint128 _amount, uint128 _maxAmount) returns(uint128 withdrawnAmount)
func (_ZkBNB *ZkBNBSession) TransferERC20(_token common.Address, _to common.Address, _amount *big.Int, _maxAmount *big.Int) (*types.Transaction, error) {
	return _ZkBNB.Contract.TransferERC20(&_ZkBNB.TransactOpts, _token, _to, _amount, _maxAmount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x68809a23.
//
// Solidity: function transferERC20(address _token, address _to, uint128 _amount, uint128 _maxAmount) returns(uint128 withdrawnAmount)
func (_ZkBNB *ZkBNBTransactorSession) TransferERC20(_token common.Address, _to common.Address, _amount *big.Int, _maxAmount *big.Int) (*types.Transaction, error) {
	return _ZkBNB.Contract.TransferERC20(&_ZkBNB.TransactOpts, _token, _to, _amount, _maxAmount)
}

// UpdatePairRate is a paid mutator transaction binding the contract method 0x13a05e23.
//
// Solidity: function updatePairRate((address,address,uint16,uint32,uint16) _pairInfo) returns()
func (_ZkBNB *ZkBNBTransactor) UpdatePairRate(opts *bind.TransactOpts, _pairInfo OldZkBNBPairInfo) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "updatePairRate", _pairInfo)
}

// UpdatePairRate is a paid mutator transaction binding the contract method 0x13a05e23.
//
// Solidity: function updatePairRate((address,address,uint16,uint32,uint16) _pairInfo) returns()
func (_ZkBNB *ZkBNBSession) UpdatePairRate(_pairInfo OldZkBNBPairInfo) (*types.Transaction, error) {
	return _ZkBNB.Contract.UpdatePairRate(&_ZkBNB.TransactOpts, _pairInfo)
}

// UpdatePairRate is a paid mutator transaction binding the contract method 0x13a05e23.
//
// Solidity: function updatePairRate((address,address,uint16,uint32,uint16) _pairInfo) returns()
func (_ZkBNB *ZkBNBTransactorSession) UpdatePairRate(_pairInfo OldZkBNBPairInfo) (*types.Transaction, error) {
	return _ZkBNB.Contract.UpdatePairRate(&_ZkBNB.TransactOpts, _pairInfo)
}

// UpdateZkBNBVerifier is a paid mutator transaction binding the contract method 0x6f88b156.
//
// Solidity: function updateZkBNBVerifier(address _newVerifierAddress) returns()
func (_ZkBNB *ZkBNBTransactor) UpdateZkBNBVerifier(opts *bind.TransactOpts, _newVerifierAddress common.Address) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "updateZkBNBVerifier", _newVerifierAddress)
}

// UpdateZkBNBVerifier is a paid mutator transaction binding the contract method 0x6f88b156.
//
// Solidity: function updateZkBNBVerifier(address _newVerifierAddress) returns()
func (_ZkBNB *ZkBNBSession) UpdateZkBNBVerifier(_newVerifierAddress common.Address) (*types.Transaction, error) {
	return _ZkBNB.Contract.UpdateZkBNBVerifier(&_ZkBNB.TransactOpts, _newVerifierAddress)
}

// UpdateZkBNBVerifier is a paid mutator transaction binding the contract method 0x6f88b156.
//
// Solidity: function updateZkBNBVerifier(address _newVerifierAddress) returns()
func (_ZkBNB *ZkBNBTransactorSession) UpdateZkBNBVerifier(_newVerifierAddress common.Address) (*types.Transaction, error) {
	return _ZkBNB.Contract.UpdateZkBNBVerifier(&_ZkBNB.TransactOpts, _newVerifierAddress)
}

// Upgrade is a paid mutator transaction binding the contract method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_ZkBNB *ZkBNBTransactor) Upgrade(opts *bind.TransactOpts, upgradeParameters []byte) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "upgrade", upgradeParameters)
}

// Upgrade is a paid mutator transaction binding the contract method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_ZkBNB *ZkBNBSession) Upgrade(upgradeParameters []byte) (*types.Transaction, error) {
	return _ZkBNB.Contract.Upgrade(&_ZkBNB.TransactOpts, upgradeParameters)
}

// Upgrade is a paid mutator transaction binding the contract method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_ZkBNB *ZkBNBTransactorSession) Upgrade(upgradeParameters []byte) (*types.Transaction, error) {
	return _ZkBNB.Contract.Upgrade(&_ZkBNB.TransactOpts, upgradeParameters)
}

// UpgradeCanceled is a paid mutator transaction binding the contract method 0x871b8ff1.
//
// Solidity: function upgradeCanceled() returns()
func (_ZkBNB *ZkBNBTransactor) UpgradeCanceled(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "upgradeCanceled")
}

// UpgradeCanceled is a paid mutator transaction binding the contract method 0x871b8ff1.
//
// Solidity: function upgradeCanceled() returns()
func (_ZkBNB *ZkBNBSession) UpgradeCanceled() (*types.Transaction, error) {
	return _ZkBNB.Contract.UpgradeCanceled(&_ZkBNB.TransactOpts)
}

// UpgradeCanceled is a paid mutator transaction binding the contract method 0x871b8ff1.
//
// Solidity: function upgradeCanceled() returns()
func (_ZkBNB *ZkBNBTransactorSession) UpgradeCanceled() (*types.Transaction, error) {
	return _ZkBNB.Contract.UpgradeCanceled(&_ZkBNB.TransactOpts)
}

// UpgradeFinishes is a paid mutator transaction binding the contract method 0xb269b9ae.
//
// Solidity: function upgradeFinishes() returns()
func (_ZkBNB *ZkBNBTransactor) UpgradeFinishes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "upgradeFinishes")
}

// UpgradeFinishes is a paid mutator transaction binding the contract method 0xb269b9ae.
//
// Solidity: function upgradeFinishes() returns()
func (_ZkBNB *ZkBNBSession) UpgradeFinishes() (*types.Transaction, error) {
	return _ZkBNB.Contract.UpgradeFinishes(&_ZkBNB.TransactOpts)
}

// UpgradeFinishes is a paid mutator transaction binding the contract method 0xb269b9ae.
//
// Solidity: function upgradeFinishes() returns()
func (_ZkBNB *ZkBNBTransactorSession) UpgradeFinishes() (*types.Transaction, error) {
	return _ZkBNB.Contract.UpgradeFinishes(&_ZkBNB.TransactOpts)
}

// UpgradeNoticePeriodStarted is a paid mutator transaction binding the contract method 0x3b154b73.
//
// Solidity: function upgradeNoticePeriodStarted() returns()
func (_ZkBNB *ZkBNBTransactor) UpgradeNoticePeriodStarted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "upgradeNoticePeriodStarted")
}

// UpgradeNoticePeriodStarted is a paid mutator transaction binding the contract method 0x3b154b73.
//
// Solidity: function upgradeNoticePeriodStarted() returns()
func (_ZkBNB *ZkBNBSession) UpgradeNoticePeriodStarted() (*types.Transaction, error) {
	return _ZkBNB.Contract.UpgradeNoticePeriodStarted(&_ZkBNB.TransactOpts)
}

// UpgradeNoticePeriodStarted is a paid mutator transaction binding the contract method 0x3b154b73.
//
// Solidity: function upgradeNoticePeriodStarted() returns()
func (_ZkBNB *ZkBNBTransactorSession) UpgradeNoticePeriodStarted() (*types.Transaction, error) {
	return _ZkBNB.Contract.UpgradeNoticePeriodStarted(&_ZkBNB.TransactOpts)
}

// UpgradePreparationStarted is a paid mutator transaction binding the contract method 0x78b91e70.
//
// Solidity: function upgradePreparationStarted() returns()
func (_ZkBNB *ZkBNBTransactor) UpgradePreparationStarted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "upgradePreparationStarted")
}

// UpgradePreparationStarted is a paid mutator transaction binding the contract method 0x78b91e70.
//
// Solidity: function upgradePreparationStarted() returns()
func (_ZkBNB *ZkBNBSession) UpgradePreparationStarted() (*types.Transaction, error) {
	return _ZkBNB.Contract.UpgradePreparationStarted(&_ZkBNB.TransactOpts)
}

// UpgradePreparationStarted is a paid mutator transaction binding the contract method 0x78b91e70.
//
// Solidity: function upgradePreparationStarted() returns()
func (_ZkBNB *ZkBNBTransactorSession) UpgradePreparationStarted() (*types.Transaction, error) {
	return _ZkBNB.Contract.UpgradePreparationStarted(&_ZkBNB.TransactOpts)
}

// VerifyAndExecuteBlocks is a paid mutator transaction binding the contract method 0xf9ea2e71.
//
// Solidity: function verifyAndExecuteBlocks(((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32),bytes[])[] _blocks, uint256[] _proofs) returns()
func (_ZkBNB *ZkBNBTransactor) VerifyAndExecuteBlocks(opts *bind.TransactOpts, _blocks []OldZkBNBVerifyAndExecuteBlockInfo, _proofs []*big.Int) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "verifyAndExecuteBlocks", _blocks, _proofs)
}

// VerifyAndExecuteBlocks is a paid mutator transaction binding the contract method 0xf9ea2e71.
//
// Solidity: function verifyAndExecuteBlocks(((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32),bytes[])[] _blocks, uint256[] _proofs) returns()
func (_ZkBNB *ZkBNBSession) VerifyAndExecuteBlocks(_blocks []OldZkBNBVerifyAndExecuteBlockInfo, _proofs []*big.Int) (*types.Transaction, error) {
	return _ZkBNB.Contract.VerifyAndExecuteBlocks(&_ZkBNB.TransactOpts, _blocks, _proofs)
}

// VerifyAndExecuteBlocks is a paid mutator transaction binding the contract method 0xf9ea2e71.
//
// Solidity: function verifyAndExecuteBlocks(((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32),bytes[])[] _blocks, uint256[] _proofs) returns()
func (_ZkBNB *ZkBNBTransactorSession) VerifyAndExecuteBlocks(_blocks []OldZkBNBVerifyAndExecuteBlockInfo, _proofs []*big.Int) (*types.Transaction, error) {
	return _ZkBNB.Contract.VerifyAndExecuteBlocks(&_ZkBNB.TransactOpts, _blocks, _proofs)
}

// WithdrawPendingBalance is a paid mutator transaction binding the contract method 0xd514da50.
//
// Solidity: function withdrawPendingBalance(address _owner, address _token, uint128 _amount) returns()
func (_ZkBNB *ZkBNBTransactor) WithdrawPendingBalance(opts *bind.TransactOpts, _owner common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "withdrawPendingBalance", _owner, _token, _amount)
}

// WithdrawPendingBalance is a paid mutator transaction binding the contract method 0xd514da50.
//
// Solidity: function withdrawPendingBalance(address _owner, address _token, uint128 _amount) returns()
func (_ZkBNB *ZkBNBSession) WithdrawPendingBalance(_owner common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ZkBNB.Contract.WithdrawPendingBalance(&_ZkBNB.TransactOpts, _owner, _token, _amount)
}

// WithdrawPendingBalance is a paid mutator transaction binding the contract method 0xd514da50.
//
// Solidity: function withdrawPendingBalance(address _owner, address _token, uint128 _amount) returns()
func (_ZkBNB *ZkBNBTransactorSession) WithdrawPendingBalance(_owner common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ZkBNB.Contract.WithdrawPendingBalance(&_ZkBNB.TransactOpts, _owner, _token, _amount)
}

// WithdrawPendingNFTBalance is a paid mutator transaction binding the contract method 0x7ce1017d.
//
// Solidity: function withdrawPendingNFTBalance(uint40 _nftIndex) returns()
func (_ZkBNB *ZkBNBTransactor) WithdrawPendingNFTBalance(opts *bind.TransactOpts, _nftIndex *big.Int) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "withdrawPendingNFTBalance", _nftIndex)
}

// WithdrawPendingNFTBalance is a paid mutator transaction binding the contract method 0x7ce1017d.
//
// Solidity: function withdrawPendingNFTBalance(uint40 _nftIndex) returns()
func (_ZkBNB *ZkBNBSession) WithdrawPendingNFTBalance(_nftIndex *big.Int) (*types.Transaction, error) {
	return _ZkBNB.Contract.WithdrawPendingNFTBalance(&_ZkBNB.TransactOpts, _nftIndex)
}

// WithdrawPendingNFTBalance is a paid mutator transaction binding the contract method 0x7ce1017d.
//
// Solidity: function withdrawPendingNFTBalance(uint40 _nftIndex) returns()
func (_ZkBNB *ZkBNBTransactorSession) WithdrawPendingNFTBalance(_nftIndex *big.Int) (*types.Transaction, error) {
	return _ZkBNB.Contract.WithdrawPendingNFTBalance(&_ZkBNB.TransactOpts, _nftIndex)
}

// ZkBNBBlockCommitIterator is returned from FilterBlockCommit and is used to iterate over the raw logs and unpacked data for BlockCommit events raised by the ZkBNB contract.
type ZkBNBBlockCommitIterator struct {
	Event *ZkBNBBlockCommit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkBNBBlockCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkBNBBlockCommit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkBNBBlockCommit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkBNBBlockCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkBNBBlockCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkBNBBlockCommit represents a BlockCommit event raised by the ZkBNB contract.
type ZkBNBBlockCommit struct {
	BlockNumber uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockCommit is a free log retrieval operation binding the contract event 0x81a92942d0f9c33b897a438384c9c3d88be397776138efa3ba1a4fc8b6268424.
//
// Solidity: event BlockCommit(uint32 blockNumber)
func (_ZkBNB *ZkBNBFilterer) FilterBlockCommit(opts *bind.FilterOpts) (*ZkBNBBlockCommitIterator, error) {

	logs, sub, err := _ZkBNB.contract.FilterLogs(opts, "BlockCommit")
	if err != nil {
		return nil, err
	}
	return &ZkBNBBlockCommitIterator{contract: _ZkBNB.contract, event: "BlockCommit", logs: logs, sub: sub}, nil
}

// WatchBlockCommit is a free log subscription operation binding the contract event 0x81a92942d0f9c33b897a438384c9c3d88be397776138efa3ba1a4fc8b6268424.
//
// Solidity: event BlockCommit(uint32 blockNumber)
func (_ZkBNB *ZkBNBFilterer) WatchBlockCommit(opts *bind.WatchOpts, sink chan<- *ZkBNBBlockCommit) (event.Subscription, error) {

	logs, sub, err := _ZkBNB.contract.WatchLogs(opts, "BlockCommit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkBNBBlockCommit)
				if err := _ZkBNB.contract.UnpackLog(event, "BlockCommit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlockCommit is a log parse operation binding the contract event 0x81a92942d0f9c33b897a438384c9c3d88be397776138efa3ba1a4fc8b6268424.
//
// Solidity: event BlockCommit(uint32 blockNumber)
func (_ZkBNB *ZkBNBFilterer) ParseBlockCommit(log types.Log) (*ZkBNBBlockCommit, error) {
	event := new(ZkBNBBlockCommit)
	if err := _ZkBNB.contract.UnpackLog(event, "BlockCommit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkBNBBlockVerificationIterator is returned from FilterBlockVerification and is used to iterate over the raw logs and unpacked data for BlockVerification events raised by the ZkBNB contract.
type ZkBNBBlockVerificationIterator struct {
	Event *ZkBNBBlockVerification // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkBNBBlockVerificationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkBNBBlockVerification)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkBNBBlockVerification)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkBNBBlockVerificationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkBNBBlockVerificationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkBNBBlockVerification represents a BlockVerification event raised by the ZkBNB contract.
type ZkBNBBlockVerification struct {
	BlockNumber uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockVerification is a free log retrieval operation binding the contract event 0x0cdbd8bd7813095001c5fe7917bd69d834dc01db7c1dfcf52ca135bd20384413.
//
// Solidity: event BlockVerification(uint32 blockNumber)
func (_ZkBNB *ZkBNBFilterer) FilterBlockVerification(opts *bind.FilterOpts) (*ZkBNBBlockVerificationIterator, error) {

	logs, sub, err := _ZkBNB.contract.FilterLogs(opts, "BlockVerification")
	if err != nil {
		return nil, err
	}
	return &ZkBNBBlockVerificationIterator{contract: _ZkBNB.contract, event: "BlockVerification", logs: logs, sub: sub}, nil
}

// WatchBlockVerification is a free log subscription operation binding the contract event 0x0cdbd8bd7813095001c5fe7917bd69d834dc01db7c1dfcf52ca135bd20384413.
//
// Solidity: event BlockVerification(uint32 blockNumber)
func (_ZkBNB *ZkBNBFilterer) WatchBlockVerification(opts *bind.WatchOpts, sink chan<- *ZkBNBBlockVerification) (event.Subscription, error) {

	logs, sub, err := _ZkBNB.contract.WatchLogs(opts, "BlockVerification")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkBNBBlockVerification)
				if err := _ZkBNB.contract.UnpackLog(event, "BlockVerification", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlockVerification is a log parse operation binding the contract event 0x0cdbd8bd7813095001c5fe7917bd69d834dc01db7c1dfcf52ca135bd20384413.
//
// Solidity: event BlockVerification(uint32 blockNumber)
func (_ZkBNB *ZkBNBFilterer) ParseBlockVerification(log types.Log) (*ZkBNBBlockVerification, error) {
	event := new(ZkBNBBlockVerification)
	if err := _ZkBNB.contract.UnpackLog(event, "BlockVerification", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkBNBBlocksRevertIterator is returned from FilterBlocksRevert and is used to iterate over the raw logs and unpacked data for BlocksRevert events raised by the ZkBNB contract.
type ZkBNBBlocksRevertIterator struct {
	Event *ZkBNBBlocksRevert // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkBNBBlocksRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkBNBBlocksRevert)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkBNBBlocksRevert)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkBNBBlocksRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkBNBBlocksRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkBNBBlocksRevert represents a BlocksRevert event raised by the ZkBNB contract.
type ZkBNBBlocksRevert struct {
	TotalBlocksVerified  uint32
	TotalBlocksCommitted uint32
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterBlocksRevert is a free log retrieval operation binding the contract event 0x6f3a8259cce1ea2680115053d21c971aa1764295a45850f520525f2bfdf3c9d3.
//
// Solidity: event BlocksRevert(uint32 totalBlocksVerified, uint32 totalBlocksCommitted)
func (_ZkBNB *ZkBNBFilterer) FilterBlocksRevert(opts *bind.FilterOpts) (*ZkBNBBlocksRevertIterator, error) {

	logs, sub, err := _ZkBNB.contract.FilterLogs(opts, "BlocksRevert")
	if err != nil {
		return nil, err
	}
	return &ZkBNBBlocksRevertIterator{contract: _ZkBNB.contract, event: "BlocksRevert", logs: logs, sub: sub}, nil
}

// WatchBlocksRevert is a free log subscription operation binding the contract event 0x6f3a8259cce1ea2680115053d21c971aa1764295a45850f520525f2bfdf3c9d3.
//
// Solidity: event BlocksRevert(uint32 totalBlocksVerified, uint32 totalBlocksCommitted)
func (_ZkBNB *ZkBNBFilterer) WatchBlocksRevert(opts *bind.WatchOpts, sink chan<- *ZkBNBBlocksRevert) (event.Subscription, error) {

	logs, sub, err := _ZkBNB.contract.WatchLogs(opts, "BlocksRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkBNBBlocksRevert)
				if err := _ZkBNB.contract.UnpackLog(event, "BlocksRevert", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlocksRevert is a log parse operation binding the contract event 0x6f3a8259cce1ea2680115053d21c971aa1764295a45850f520525f2bfdf3c9d3.
//
// Solidity: event BlocksRevert(uint32 totalBlocksVerified, uint32 totalBlocksCommitted)
func (_ZkBNB *ZkBNBFilterer) ParseBlocksRevert(log types.Log) (*ZkBNBBlocksRevert, error) {
	event := new(ZkBNBBlocksRevert)
	if err := _ZkBNB.contract.UnpackLog(event, "BlocksRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkBNBCreateTokenPairIterator is returned from FilterCreateTokenPair and is used to iterate over the raw logs and unpacked data for CreateTokenPair events raised by the ZkBNB contract.
type ZkBNBCreateTokenPairIterator struct {
	Event *ZkBNBCreateTokenPair // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkBNBCreateTokenPairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkBNBCreateTokenPair)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkBNBCreateTokenPair)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkBNBCreateTokenPairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkBNBCreateTokenPairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkBNBCreateTokenPair represents a CreateTokenPair event raised by the ZkBNB contract.
type ZkBNBCreateTokenPair struct {
	PairIndex            uint16
	Asset0Id             uint16
	Asset1Id             uint16
	FeeRate              uint16
	TreasuryAccountIndex uint32
	TreasuryRate         uint16
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterCreateTokenPair is a free log retrieval operation binding the contract event 0x82e21fd4ced46b510ee1bf5a34eea3dc34aa1f5460f53a25dd22828cde30e087.
//
// Solidity: event CreateTokenPair(uint16 pairIndex, uint16 asset0Id, uint16 asset1Id, uint16 feeRate, uint32 treasuryAccountIndex, uint16 treasuryRate)
func (_ZkBNB *ZkBNBFilterer) FilterCreateTokenPair(opts *bind.FilterOpts) (*ZkBNBCreateTokenPairIterator, error) {

	logs, sub, err := _ZkBNB.contract.FilterLogs(opts, "CreateTokenPair")
	if err != nil {
		return nil, err
	}
	return &ZkBNBCreateTokenPairIterator{contract: _ZkBNB.contract, event: "CreateTokenPair", logs: logs, sub: sub}, nil
}

// WatchCreateTokenPair is a free log subscription operation binding the contract event 0x82e21fd4ced46b510ee1bf5a34eea3dc34aa1f5460f53a25dd22828cde30e087.
//
// Solidity: event CreateTokenPair(uint16 pairIndex, uint16 asset0Id, uint16 asset1Id, uint16 feeRate, uint32 treasuryAccountIndex, uint16 treasuryRate)
func (_ZkBNB *ZkBNBFilterer) WatchCreateTokenPair(opts *bind.WatchOpts, sink chan<- *ZkBNBCreateTokenPair) (event.Subscription, error) {

	logs, sub, err := _ZkBNB.contract.WatchLogs(opts, "CreateTokenPair")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkBNBCreateTokenPair)
				if err := _ZkBNB.contract.UnpackLog(event, "CreateTokenPair", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreateTokenPair is a log parse operation binding the contract event 0x82e21fd4ced46b510ee1bf5a34eea3dc34aa1f5460f53a25dd22828cde30e087.
//
// Solidity: event CreateTokenPair(uint16 pairIndex, uint16 asset0Id, uint16 asset1Id, uint16 feeRate, uint32 treasuryAccountIndex, uint16 treasuryRate)
func (_ZkBNB *ZkBNBFilterer) ParseCreateTokenPair(log types.Log) (*ZkBNBCreateTokenPair, error) {
	event := new(ZkBNBCreateTokenPair)
	if err := _ZkBNB.contract.UnpackLog(event, "CreateTokenPair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkBNBDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ZkBNB contract.
type ZkBNBDepositIterator struct {
	Event *ZkBNBDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkBNBDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkBNBDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkBNBDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkBNBDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkBNBDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkBNBDeposit represents a Deposit event raised by the ZkBNB contract.
type ZkBNBDeposit struct {
	AssetId     uint16
	AccountName [32]byte
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xaa46d46658f805449ac7eaaf11d3eaebb6f508930128d276dfefcc8dc02a13c7.
//
// Solidity: event Deposit(uint16 assetId, bytes32 accountName, uint128 amount)
func (_ZkBNB *ZkBNBFilterer) FilterDeposit(opts *bind.FilterOpts) (*ZkBNBDepositIterator, error) {

	logs, sub, err := _ZkBNB.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &ZkBNBDepositIterator{contract: _ZkBNB.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xaa46d46658f805449ac7eaaf11d3eaebb6f508930128d276dfefcc8dc02a13c7.
//
// Solidity: event Deposit(uint16 assetId, bytes32 accountName, uint128 amount)
func (_ZkBNB *ZkBNBFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ZkBNBDeposit) (event.Subscription, error) {

	logs, sub, err := _ZkBNB.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkBNBDeposit)
				if err := _ZkBNB.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0xaa46d46658f805449ac7eaaf11d3eaebb6f508930128d276dfefcc8dc02a13c7.
//
// Solidity: event Deposit(uint16 assetId, bytes32 accountName, uint128 amount)
func (_ZkBNB *ZkBNBFilterer) ParseDeposit(log types.Log) (*ZkBNBDeposit, error) {
	event := new(ZkBNBDeposit)
	if err := _ZkBNB.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkBNBDepositCommitIterator is returned from FilterDepositCommit and is used to iterate over the raw logs and unpacked data for DepositCommit events raised by the ZkBNB contract.
type ZkBNBDepositCommitIterator struct {
	Event *ZkBNBDepositCommit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkBNBDepositCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkBNBDepositCommit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkBNBDepositCommit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkBNBDepositCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkBNBDepositCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkBNBDepositCommit represents a DepositCommit event raised by the ZkBNB contract.
type ZkBNBDepositCommit struct {
	ZkBNBBlockNumber uint32
	AccountIndex     uint32
	AccountName      [32]byte
	AssetId          uint16
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDepositCommit is a free log retrieval operation binding the contract event 0x9b3e77f05cee61edb7f3e0e85d71a29e89cb18cb4236a8b8cc3e4d66640d8f71.
//
// Solidity: event DepositCommit(uint32 indexed zkbnbBlockNumber, uint32 indexed accountIndex, bytes32 accountName, uint16 indexed assetId, uint128 amount)
func (_ZkBNB *ZkBNBFilterer) FilterDepositCommit(opts *bind.FilterOpts, zkbnbBlockNumber []uint32, accountIndex []uint32, assetId []uint16) (*ZkBNBDepositCommitIterator, error) {

	var zkbnbBlockNumberRule []interface{}
	for _, zkbnbBlockNumberItem := range zkbnbBlockNumber {
		zkbnbBlockNumberRule = append(zkbnbBlockNumberRule, zkbnbBlockNumberItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _ZkBNB.contract.FilterLogs(opts, "DepositCommit", zkbnbBlockNumberRule, accountIndexRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &ZkBNBDepositCommitIterator{contract: _ZkBNB.contract, event: "DepositCommit", logs: logs, sub: sub}, nil
}

// WatchDepositCommit is a free log subscription operation binding the contract event 0x9b3e77f05cee61edb7f3e0e85d71a29e89cb18cb4236a8b8cc3e4d66640d8f71.
//
// Solidity: event DepositCommit(uint32 indexed zkbnbBlockNumber, uint32 indexed accountIndex, bytes32 accountName, uint16 indexed assetId, uint128 amount)
func (_ZkBNB *ZkBNBFilterer) WatchDepositCommit(opts *bind.WatchOpts, sink chan<- *ZkBNBDepositCommit, zkbnbBlockNumber []uint32, accountIndex []uint32, assetId []uint16) (event.Subscription, error) {

	var zkbnbBlockNumberRule []interface{}
	for _, zkbnbBlockNumberItem := range zkbnbBlockNumber {
		zkbnbBlockNumberRule = append(zkbnbBlockNumberRule, zkbnbBlockNumberItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _ZkBNB.contract.WatchLogs(opts, "DepositCommit", zkbnbBlockNumberRule, accountIndexRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkBNBDepositCommit)
				if err := _ZkBNB.contract.UnpackLog(event, "DepositCommit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositCommit is a log parse operation binding the contract event 0x9b3e77f05cee61edb7f3e0e85d71a29e89cb18cb4236a8b8cc3e4d66640d8f71.
//
// Solidity: event DepositCommit(uint32 indexed zkbnbBlockNumber, uint32 indexed accountIndex, bytes32 accountName, uint16 indexed assetId, uint128 amount)
func (_ZkBNB *ZkBNBFilterer) ParseDepositCommit(log types.Log) (*ZkBNBDepositCommit, error) {
	event := new(ZkBNBDepositCommit)
	if err := _ZkBNB.contract.UnpackLog(event, "DepositCommit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkBNBDepositNftIterator is returned from FilterDepositNft and is used to iterate over the raw logs and unpacked data for DepositNft events raised by the ZkBNB contract.
type ZkBNBDepositNftIterator struct {
	Event *ZkBNBDepositNft // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkBNBDepositNftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkBNBDepositNft)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkBNBDepositNft)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkBNBDepositNftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkBNBDepositNftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkBNBDepositNft represents a DepositNft event raised by the ZkBNB contract.
type ZkBNBDepositNft struct {
	AccountNameHash     [32]byte
	NftContentHash      [32]byte
	TokenAddress        common.Address
	NftTokenId          *big.Int
	CreatorTreasuryRate uint16
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDepositNft is a free log retrieval operation binding the contract event 0x42f9fbd77faf066fde386943aaafcc841fba6b755dc5da7939a046cbc493c24f.
//
// Solidity: event DepositNft(bytes32 accountNameHash, bytes32 nftContentHash, address tokenAddress, uint256 nftTokenId, uint16 creatorTreasuryRate)
func (_ZkBNB *ZkBNBFilterer) FilterDepositNft(opts *bind.FilterOpts) (*ZkBNBDepositNftIterator, error) {

	logs, sub, err := _ZkBNB.contract.FilterLogs(opts, "DepositNft")
	if err != nil {
		return nil, err
	}
	return &ZkBNBDepositNftIterator{contract: _ZkBNB.contract, event: "DepositNft", logs: logs, sub: sub}, nil
}

// WatchDepositNft is a free log subscription operation binding the contract event 0x42f9fbd77faf066fde386943aaafcc841fba6b755dc5da7939a046cbc493c24f.
//
// Solidity: event DepositNft(bytes32 accountNameHash, bytes32 nftContentHash, address tokenAddress, uint256 nftTokenId, uint16 creatorTreasuryRate)
func (_ZkBNB *ZkBNBFilterer) WatchDepositNft(opts *bind.WatchOpts, sink chan<- *ZkBNBDepositNft) (event.Subscription, error) {

	logs, sub, err := _ZkBNB.contract.WatchLogs(opts, "DepositNft")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkBNBDepositNft)
				if err := _ZkBNB.contract.UnpackLog(event, "DepositNft", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositNft is a log parse operation binding the contract event 0x42f9fbd77faf066fde386943aaafcc841fba6b755dc5da7939a046cbc493c24f.
//
// Solidity: event DepositNft(bytes32 accountNameHash, bytes32 nftContentHash, address tokenAddress, uint256 nftTokenId, uint16 creatorTreasuryRate)
func (_ZkBNB *ZkBNBFilterer) ParseDepositNft(log types.Log) (*ZkBNBDepositNft, error) {
	event := new(ZkBNBDepositNft)
	if err := _ZkBNB.contract.UnpackLog(event, "DepositNft", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkBNBDesertModeIterator is returned from FilterDesertMode and is used to iterate over the raw logs and unpacked data for DesertMode events raised by the ZkBNB contract.
type ZkBNBDesertModeIterator struct {
	Event *ZkBNBDesertMode // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkBNBDesertModeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkBNBDesertMode)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkBNBDesertMode)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkBNBDesertModeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkBNBDesertModeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkBNBDesertMode represents a DesertMode event raised by the ZkBNB contract.
type ZkBNBDesertMode struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDesertMode is a free log retrieval operation binding the contract event 0x9f7e400a81dddbf1c18b1c37f82aa303d166295ca4b577eb2a7c23d4b704ba89.
//
// Solidity: event DesertMode()
func (_ZkBNB *ZkBNBFilterer) FilterDesertMode(opts *bind.FilterOpts) (*ZkBNBDesertModeIterator, error) {

	logs, sub, err := _ZkBNB.contract.FilterLogs(opts, "DesertMode")
	if err != nil {
		return nil, err
	}
	return &ZkBNBDesertModeIterator{contract: _ZkBNB.contract, event: "DesertMode", logs: logs, sub: sub}, nil
}

// WatchDesertMode is a free log subscription operation binding the contract event 0x9f7e400a81dddbf1c18b1c37f82aa303d166295ca4b577eb2a7c23d4b704ba89.
//
// Solidity: event DesertMode()
func (_ZkBNB *ZkBNBFilterer) WatchDesertMode(opts *bind.WatchOpts, sink chan<- *ZkBNBDesertMode) (event.Subscription, error) {

	logs, sub, err := _ZkBNB.contract.WatchLogs(opts, "DesertMode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkBNBDesertMode)
				if err := _ZkBNB.contract.UnpackLog(event, "DesertMode", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDesertMode is a log parse operation binding the contract event 0x9f7e400a81dddbf1c18b1c37f82aa303d166295ca4b577eb2a7c23d4b704ba89.
//
// Solidity: event DesertMode()
func (_ZkBNB *ZkBNBFilterer) ParseDesertMode(log types.Log) (*ZkBNBDesertMode, error) {
	event := new(ZkBNBDesertMode)
	if err := _ZkBNB.contract.UnpackLog(event, "DesertMode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkBNBFullExitCommitIterator is returned from FilterFullExitCommit and is used to iterate over the raw logs and unpacked data for FullExitCommit events raised by the ZkBNB contract.
type ZkBNBFullExitCommitIterator struct {
	Event *ZkBNBFullExitCommit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkBNBFullExitCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkBNBFullExitCommit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkBNBFullExitCommit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkBNBFullExitCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkBNBFullExitCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkBNBFullExitCommit represents a FullExitCommit event raised by the ZkBNB contract.
type ZkBNBFullExitCommit struct {
	ZkBNBBlockId uint32
	AccountId    uint32
	Owner        common.Address
	TokenId      uint16
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFullExitCommit is a free log retrieval operation binding the contract event 0x66fc63d751ecbefca61d4e2e7c534e4f29c61aed8ece23ed635277a7ea6f9bc4.
//
// Solidity: event FullExitCommit(uint32 indexed zkbnbBlockId, uint32 indexed accountId, address owner, uint16 indexed tokenId, uint128 amount)
func (_ZkBNB *ZkBNBFilterer) FilterFullExitCommit(opts *bind.FilterOpts, zkbnbBlockId []uint32, accountId []uint32, tokenId []uint16) (*ZkBNBFullExitCommitIterator, error) {

	var zkbnbBlockIdRule []interface{}
	for _, zkbnbBlockIdItem := range zkbnbBlockId {
		zkbnbBlockIdRule = append(zkbnbBlockIdRule, zkbnbBlockIdItem)
	}
	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ZkBNB.contract.FilterLogs(opts, "FullExitCommit", zkbnbBlockIdRule, accountIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ZkBNBFullExitCommitIterator{contract: _ZkBNB.contract, event: "FullExitCommit", logs: logs, sub: sub}, nil
}

// WatchFullExitCommit is a free log subscription operation binding the contract event 0x66fc63d751ecbefca61d4e2e7c534e4f29c61aed8ece23ed635277a7ea6f9bc4.
//
// Solidity: event FullExitCommit(uint32 indexed zkbnbBlockId, uint32 indexed accountId, address owner, uint16 indexed tokenId, uint128 amount)
func (_ZkBNB *ZkBNBFilterer) WatchFullExitCommit(opts *bind.WatchOpts, sink chan<- *ZkBNBFullExitCommit, zkbnbBlockId []uint32, accountId []uint32, tokenId []uint16) (event.Subscription, error) {

	var zkbnbBlockIdRule []interface{}
	for _, zkbnbBlockIdItem := range zkbnbBlockId {
		zkbnbBlockIdRule = append(zkbnbBlockIdRule, zkbnbBlockIdItem)
	}
	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ZkBNB.contract.WatchLogs(opts, "FullExitCommit", zkbnbBlockIdRule, accountIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkBNBFullExitCommit)
				if err := _ZkBNB.contract.UnpackLog(event, "FullExitCommit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFullExitCommit is a log parse operation binding the contract event 0x66fc63d751ecbefca61d4e2e7c534e4f29c61aed8ece23ed635277a7ea6f9bc4.
//
// Solidity: event FullExitCommit(uint32 indexed zkbnbBlockId, uint32 indexed accountId, address owner, uint16 indexed tokenId, uint128 amount)
func (_ZkBNB *ZkBNBFilterer) ParseFullExitCommit(log types.Log) (*ZkBNBFullExitCommit, error) {
	event := new(ZkBNBFullExitCommit)
	if err := _ZkBNB.contract.UnpackLog(event, "FullExitCommit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkBNBNewDefaultNFTFactoryIterator is returned from FilterNewDefaultNFTFactory and is used to iterate over the raw logs and unpacked data for NewDefaultNFTFactory events raised by the ZkBNB contract.
type ZkBNBNewDefaultNFTFactoryIterator struct {
	Event *ZkBNBNewDefaultNFTFactory // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkBNBNewDefaultNFTFactoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkBNBNewDefaultNFTFactory)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkBNBNewDefaultNFTFactory)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkBNBNewDefaultNFTFactoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkBNBNewDefaultNFTFactoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkBNBNewDefaultNFTFactory represents a NewDefaultNFTFactory event raised by the ZkBNB contract.
type ZkBNBNewDefaultNFTFactory struct {
	Factory common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewDefaultNFTFactory is a free log retrieval operation binding the contract event 0xadf2ec6c092b7101c958de6b55caf4818f26a946e94fa556cbe3b3f29ff9f2ee.
//
// Solidity: event NewDefaultNFTFactory(address indexed factory)
func (_ZkBNB *ZkBNBFilterer) FilterNewDefaultNFTFactory(opts *bind.FilterOpts, factory []common.Address) (*ZkBNBNewDefaultNFTFactoryIterator, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _ZkBNB.contract.FilterLogs(opts, "NewDefaultNFTFactory", factoryRule)
	if err != nil {
		return nil, err
	}
	return &ZkBNBNewDefaultNFTFactoryIterator{contract: _ZkBNB.contract, event: "NewDefaultNFTFactory", logs: logs, sub: sub}, nil
}

// WatchNewDefaultNFTFactory is a free log subscription operation binding the contract event 0xadf2ec6c092b7101c958de6b55caf4818f26a946e94fa556cbe3b3f29ff9f2ee.
//
// Solidity: event NewDefaultNFTFactory(address indexed factory)
func (_ZkBNB *ZkBNBFilterer) WatchNewDefaultNFTFactory(opts *bind.WatchOpts, sink chan<- *ZkBNBNewDefaultNFTFactory, factory []common.Address) (event.Subscription, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _ZkBNB.contract.WatchLogs(opts, "NewDefaultNFTFactory", factoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkBNBNewDefaultNFTFactory)
				if err := _ZkBNB.contract.UnpackLog(event, "NewDefaultNFTFactory", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewDefaultNFTFactory is a log parse operation binding the contract event 0xadf2ec6c092b7101c958de6b55caf4818f26a946e94fa556cbe3b3f29ff9f2ee.
//
// Solidity: event NewDefaultNFTFactory(address indexed factory)
func (_ZkBNB *ZkBNBFilterer) ParseNewDefaultNFTFactory(log types.Log) (*ZkBNBNewDefaultNFTFactory, error) {
	event := new(ZkBNBNewDefaultNFTFactory)
	if err := _ZkBNB.contract.UnpackLog(event, "NewDefaultNFTFactory", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkBNBNewNFTFactoryIterator is returned from FilterNewNFTFactory and is used to iterate over the raw logs and unpacked data for NewNFTFactory events raised by the ZkBNB contract.
type ZkBNBNewNFTFactoryIterator struct {
	Event *ZkBNBNewNFTFactory // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkBNBNewNFTFactoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkBNBNewNFTFactory)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkBNBNewNFTFactory)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkBNBNewNFTFactoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkBNBNewNFTFactoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkBNBNewNFTFactory represents a NewNFTFactory event raised by the ZkBNB contract.
type ZkBNBNewNFTFactory struct {
	CreatorAccountNameHash [32]byte
	CollectionId           uint32
	FactoryAddress         common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterNewNFTFactory is a free log retrieval operation binding the contract event 0xaa31b08b26b4240f7d2837e3fb6359824b509a4e4c5a5766af80f63c319add7e.
//
// Solidity: event NewNFTFactory(bytes32 indexed _creatorAccountNameHash, uint32 _collectionId, address _factoryAddress)
func (_ZkBNB *ZkBNBFilterer) FilterNewNFTFactory(opts *bind.FilterOpts, _creatorAccountNameHash [][32]byte) (*ZkBNBNewNFTFactoryIterator, error) {

	var _creatorAccountNameHashRule []interface{}
	for _, _creatorAccountNameHashItem := range _creatorAccountNameHash {
		_creatorAccountNameHashRule = append(_creatorAccountNameHashRule, _creatorAccountNameHashItem)
	}

	logs, sub, err := _ZkBNB.contract.FilterLogs(opts, "NewNFTFactory", _creatorAccountNameHashRule)
	if err != nil {
		return nil, err
	}
	return &ZkBNBNewNFTFactoryIterator{contract: _ZkBNB.contract, event: "NewNFTFactory", logs: logs, sub: sub}, nil
}

// WatchNewNFTFactory is a free log subscription operation binding the contract event 0xaa31b08b26b4240f7d2837e3fb6359824b509a4e4c5a5766af80f63c319add7e.
//
// Solidity: event NewNFTFactory(bytes32 indexed _creatorAccountNameHash, uint32 _collectionId, address _factoryAddress)
func (_ZkBNB *ZkBNBFilterer) WatchNewNFTFactory(opts *bind.WatchOpts, sink chan<- *ZkBNBNewNFTFactory, _creatorAccountNameHash [][32]byte) (event.Subscription, error) {

	var _creatorAccountNameHashRule []interface{}
	for _, _creatorAccountNameHashItem := range _creatorAccountNameHash {
		_creatorAccountNameHashRule = append(_creatorAccountNameHashRule, _creatorAccountNameHashItem)
	}

	logs, sub, err := _ZkBNB.contract.WatchLogs(opts, "NewNFTFactory", _creatorAccountNameHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkBNBNewNFTFactory)
				if err := _ZkBNB.contract.UnpackLog(event, "NewNFTFactory", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewNFTFactory is a log parse operation binding the contract event 0xaa31b08b26b4240f7d2837e3fb6359824b509a4e4c5a5766af80f63c319add7e.
//
// Solidity: event NewNFTFactory(bytes32 indexed _creatorAccountNameHash, uint32 _collectionId, address _factoryAddress)
func (_ZkBNB *ZkBNBFilterer) ParseNewNFTFactory(log types.Log) (*ZkBNBNewNFTFactory, error) {
	event := new(ZkBNBNewNFTFactory)
	if err := _ZkBNB.contract.UnpackLog(event, "NewNFTFactory", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkBNBNewPriorityRequestIterator is returned from FilterNewPriorityRequest and is used to iterate over the raw logs and unpacked data for NewPriorityRequest events raised by the ZkBNB contract.
type ZkBNBNewPriorityRequestIterator struct {
	Event *ZkBNBNewPriorityRequest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkBNBNewPriorityRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkBNBNewPriorityRequest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkBNBNewPriorityRequest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkBNBNewPriorityRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkBNBNewPriorityRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkBNBNewPriorityRequest represents a NewPriorityRequest event raised by the ZkBNB contract.
type ZkBNBNewPriorityRequest struct {
	Sender          common.Address
	SerialId        uint64
	TxType          uint8
	PubData         []byte
	ExpirationBlock *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPriorityRequest is a free log retrieval operation binding the contract event 0xd0943372c08b438a88d4b39d77216901079eda9ca59d45349841c099083b6830.
//
// Solidity: event NewPriorityRequest(address sender, uint64 serialId, uint8 txType, bytes pubData, uint256 expirationBlock)
func (_ZkBNB *ZkBNBFilterer) FilterNewPriorityRequest(opts *bind.FilterOpts) (*ZkBNBNewPriorityRequestIterator, error) {

	logs, sub, err := _ZkBNB.contract.FilterLogs(opts, "NewPriorityRequest")
	if err != nil {
		return nil, err
	}
	return &ZkBNBNewPriorityRequestIterator{contract: _ZkBNB.contract, event: "NewPriorityRequest", logs: logs, sub: sub}, nil
}

// WatchNewPriorityRequest is a free log subscription operation binding the contract event 0xd0943372c08b438a88d4b39d77216901079eda9ca59d45349841c099083b6830.
//
// Solidity: event NewPriorityRequest(address sender, uint64 serialId, uint8 txType, bytes pubData, uint256 expirationBlock)
func (_ZkBNB *ZkBNBFilterer) WatchNewPriorityRequest(opts *bind.WatchOpts, sink chan<- *ZkBNBNewPriorityRequest) (event.Subscription, error) {

	logs, sub, err := _ZkBNB.contract.WatchLogs(opts, "NewPriorityRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkBNBNewPriorityRequest)
				if err := _ZkBNB.contract.UnpackLog(event, "NewPriorityRequest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewPriorityRequest is a log parse operation binding the contract event 0xd0943372c08b438a88d4b39d77216901079eda9ca59d45349841c099083b6830.
//
// Solidity: event NewPriorityRequest(address sender, uint64 serialId, uint8 txType, bytes pubData, uint256 expirationBlock)
func (_ZkBNB *ZkBNBFilterer) ParseNewPriorityRequest(log types.Log) (*ZkBNBNewPriorityRequest, error) {
	event := new(ZkBNBNewPriorityRequest)
	if err := _ZkBNB.contract.UnpackLog(event, "NewPriorityRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkBNBNoticePeriodChangeIterator is returned from FilterNoticePeriodChange and is used to iterate over the raw logs and unpacked data for NoticePeriodChange events raised by the ZkBNB contract.
type ZkBNBNoticePeriodChangeIterator struct {
	Event *ZkBNBNoticePeriodChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkBNBNoticePeriodChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkBNBNoticePeriodChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkBNBNoticePeriodChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkBNBNoticePeriodChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkBNBNoticePeriodChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkBNBNoticePeriodChange represents a NoticePeriodChange event raised by the ZkBNB contract.
type ZkBNBNoticePeriodChange struct {
	NewNoticePeriod *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNoticePeriodChange is a free log retrieval operation binding the contract event 0xf2b18f8abbd8a0d0c1fb8245146eedf5304887b12f6395b548ca238e054a1483.
//
// Solidity: event NoticePeriodChange(uint256 newNoticePeriod)
func (_ZkBNB *ZkBNBFilterer) FilterNoticePeriodChange(opts *bind.FilterOpts) (*ZkBNBNoticePeriodChangeIterator, error) {

	logs, sub, err := _ZkBNB.contract.FilterLogs(opts, "NoticePeriodChange")
	if err != nil {
		return nil, err
	}
	return &ZkBNBNoticePeriodChangeIterator{contract: _ZkBNB.contract, event: "NoticePeriodChange", logs: logs, sub: sub}, nil
}

// WatchNoticePeriodChange is a free log subscription operation binding the contract event 0xf2b18f8abbd8a0d0c1fb8245146eedf5304887b12f6395b548ca238e054a1483.
//
// Solidity: event NoticePeriodChange(uint256 newNoticePeriod)
func (_ZkBNB *ZkBNBFilterer) WatchNoticePeriodChange(opts *bind.WatchOpts, sink chan<- *ZkBNBNoticePeriodChange) (event.Subscription, error) {

	logs, sub, err := _ZkBNB.contract.WatchLogs(opts, "NoticePeriodChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkBNBNoticePeriodChange)
				if err := _ZkBNB.contract.UnpackLog(event, "NoticePeriodChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNoticePeriodChange is a log parse operation binding the contract event 0xf2b18f8abbd8a0d0c1fb8245146eedf5304887b12f6395b548ca238e054a1483.
//
// Solidity: event NoticePeriodChange(uint256 newNoticePeriod)
func (_ZkBNB *ZkBNBFilterer) ParseNoticePeriodChange(log types.Log) (*ZkBNBNoticePeriodChange, error) {
	event := new(ZkBNBNoticePeriodChange)
	if err := _ZkBNB.contract.UnpackLog(event, "NoticePeriodChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkBNBRegisterZNSIterator is returned from FilterRegisterZNS and is used to iterate over the raw logs and unpacked data for RegisterZNS events raised by the ZkBNB contract.
type ZkBNBRegisterZNSIterator struct {
	Event *ZkBNBRegisterZNS // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkBNBRegisterZNSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkBNBRegisterZNS)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkBNBRegisterZNS)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkBNBRegisterZNSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkBNBRegisterZNSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkBNBRegisterZNS represents a RegisterZNS event raised by the ZkBNB contract.
type ZkBNBRegisterZNS struct {
	Name         string
	NameHash     [32]byte
	Owner        common.Address
	ZkBNBPubKeyX [32]byte
	ZkBNBPubKeyY [32]byte
	AccountIndex uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRegisterZNS is a free log retrieval operation binding the contract event 0xebe5b00e9e9052d45c93d72544961460aa0fa30e0142c07d7dd81025cf9dab6c.
//
// Solidity: event RegisterZNS(string name, bytes32 nameHash, address owner, bytes32 zkbnbPubKeyX, bytes32 zkbnbPubKeyY, uint32 accountIndex)
func (_ZkBNB *ZkBNBFilterer) FilterRegisterZNS(opts *bind.FilterOpts) (*ZkBNBRegisterZNSIterator, error) {

	logs, sub, err := _ZkBNB.contract.FilterLogs(opts, "RegisterZNS")
	if err != nil {
		return nil, err
	}
	return &ZkBNBRegisterZNSIterator{contract: _ZkBNB.contract, event: "RegisterZNS", logs: logs, sub: sub}, nil
}

// WatchRegisterZNS is a free log subscription operation binding the contract event 0xebe5b00e9e9052d45c93d72544961460aa0fa30e0142c07d7dd81025cf9dab6c.
//
// Solidity: event RegisterZNS(string name, bytes32 nameHash, address owner, bytes32 zkbnbPubKeyX, bytes32 zkbnbPubKeyY, uint32 accountIndex)
func (_ZkBNB *ZkBNBFilterer) WatchRegisterZNS(opts *bind.WatchOpts, sink chan<- *ZkBNBRegisterZNS) (event.Subscription, error) {

	logs, sub, err := _ZkBNB.contract.WatchLogs(opts, "RegisterZNS")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkBNBRegisterZNS)
				if err := _ZkBNB.contract.UnpackLog(event, "RegisterZNS", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRegisterZNS is a log parse operation binding the contract event 0xebe5b00e9e9052d45c93d72544961460aa0fa30e0142c07d7dd81025cf9dab6c.
//
// Solidity: event RegisterZNS(string name, bytes32 nameHash, address owner, bytes32 zkbnbPubKeyX, bytes32 zkbnbPubKeyY, uint32 accountIndex)
func (_ZkBNB *ZkBNBFilterer) ParseRegisterZNS(log types.Log) (*ZkBNBRegisterZNS, error) {
	event := new(ZkBNBRegisterZNS)
	if err := _ZkBNB.contract.UnpackLog(event, "RegisterZNS", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkBNBUpdateTokenPairIterator is returned from FilterUpdateTokenPair and is used to iterate over the raw logs and unpacked data for UpdateTokenPair events raised by the ZkBNB contract.
type ZkBNBUpdateTokenPairIterator struct {
	Event *ZkBNBUpdateTokenPair // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkBNBUpdateTokenPairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkBNBUpdateTokenPair)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkBNBUpdateTokenPair)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkBNBUpdateTokenPairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkBNBUpdateTokenPairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkBNBUpdateTokenPair represents a UpdateTokenPair event raised by the ZkBNB contract.
type ZkBNBUpdateTokenPair struct {
	PairIndex            uint16
	FeeRate              uint16
	TreasuryAccountIndex uint32
	TreasuryRate         uint16
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateTokenPair is a free log retrieval operation binding the contract event 0x8727712867f019c156b0c25f308498ea1524b635902129b0acd3e9ff2c96f1c8.
//
// Solidity: event UpdateTokenPair(uint16 pairIndex, uint16 feeRate, uint32 treasuryAccountIndex, uint16 treasuryRate)
func (_ZkBNB *ZkBNBFilterer) FilterUpdateTokenPair(opts *bind.FilterOpts) (*ZkBNBUpdateTokenPairIterator, error) {

	logs, sub, err := _ZkBNB.contract.FilterLogs(opts, "UpdateTokenPair")
	if err != nil {
		return nil, err
	}
	return &ZkBNBUpdateTokenPairIterator{contract: _ZkBNB.contract, event: "UpdateTokenPair", logs: logs, sub: sub}, nil
}

// WatchUpdateTokenPair is a free log subscription operation binding the contract event 0x8727712867f019c156b0c25f308498ea1524b635902129b0acd3e9ff2c96f1c8.
//
// Solidity: event UpdateTokenPair(uint16 pairIndex, uint16 feeRate, uint32 treasuryAccountIndex, uint16 treasuryRate)
func (_ZkBNB *ZkBNBFilterer) WatchUpdateTokenPair(opts *bind.WatchOpts, sink chan<- *ZkBNBUpdateTokenPair) (event.Subscription, error) {

	logs, sub, err := _ZkBNB.contract.WatchLogs(opts, "UpdateTokenPair")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkBNBUpdateTokenPair)
				if err := _ZkBNB.contract.UnpackLog(event, "UpdateTokenPair", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateTokenPair is a log parse operation binding the contract event 0x8727712867f019c156b0c25f308498ea1524b635902129b0acd3e9ff2c96f1c8.
//
// Solidity: event UpdateTokenPair(uint16 pairIndex, uint16 feeRate, uint32 treasuryAccountIndex, uint16 treasuryRate)
func (_ZkBNB *ZkBNBFilterer) ParseUpdateTokenPair(log types.Log) (*ZkBNBUpdateTokenPair, error) {
	event := new(ZkBNBUpdateTokenPair)
	if err := _ZkBNB.contract.UnpackLog(event, "UpdateTokenPair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkBNBWithdrawNftIterator is returned from FilterWithdrawNft and is used to iterate over the raw logs and unpacked data for WithdrawNft events raised by the ZkBNB contract.
type ZkBNBWithdrawNftIterator struct {
	Event *ZkBNBWithdrawNft // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkBNBWithdrawNftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkBNBWithdrawNft)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkBNBWithdrawNft)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkBNBWithdrawNftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkBNBWithdrawNftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkBNBWithdrawNft represents a WithdrawNft event raised by the ZkBNB contract.
type ZkBNBWithdrawNft struct {
	AccountIndex uint32
	NftL1Address common.Address
	ToAddress    common.Address
	NftL1TokenId *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawNft is a free log retrieval operation binding the contract event 0x001c1fcac2c66b39f6b0a825948e9f8892b2d4a17edc60aaf61ca474e08402ec.
//
// Solidity: event WithdrawNft(uint32 accountIndex, address nftL1Address, address toAddress, uint256 nftL1TokenId)
func (_ZkBNB *ZkBNBFilterer) FilterWithdrawNft(opts *bind.FilterOpts) (*ZkBNBWithdrawNftIterator, error) {

	logs, sub, err := _ZkBNB.contract.FilterLogs(opts, "WithdrawNft")
	if err != nil {
		return nil, err
	}
	return &ZkBNBWithdrawNftIterator{contract: _ZkBNB.contract, event: "WithdrawNft", logs: logs, sub: sub}, nil
}

// WatchWithdrawNft is a free log subscription operation binding the contract event 0x001c1fcac2c66b39f6b0a825948e9f8892b2d4a17edc60aaf61ca474e08402ec.
//
// Solidity: event WithdrawNft(uint32 accountIndex, address nftL1Address, address toAddress, uint256 nftL1TokenId)
func (_ZkBNB *ZkBNBFilterer) WatchWithdrawNft(opts *bind.WatchOpts, sink chan<- *ZkBNBWithdrawNft) (event.Subscription, error) {

	logs, sub, err := _ZkBNB.contract.WatchLogs(opts, "WithdrawNft")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkBNBWithdrawNft)
				if err := _ZkBNB.contract.UnpackLog(event, "WithdrawNft", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawNft is a log parse operation binding the contract event 0x001c1fcac2c66b39f6b0a825948e9f8892b2d4a17edc60aaf61ca474e08402ec.
//
// Solidity: event WithdrawNft(uint32 accountIndex, address nftL1Address, address toAddress, uint256 nftL1TokenId)
func (_ZkBNB *ZkBNBFilterer) ParseWithdrawNft(log types.Log) (*ZkBNBWithdrawNft, error) {
	event := new(ZkBNBWithdrawNft)
	if err := _ZkBNB.contract.UnpackLog(event, "WithdrawNft", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkBNBWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the ZkBNB contract.
type ZkBNBWithdrawalIterator struct {
	Event *ZkBNBWithdrawal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkBNBWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkBNBWithdrawal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkBNBWithdrawal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkBNBWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkBNBWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkBNBWithdrawal represents a Withdrawal event raised by the ZkBNB contract.
type ZkBNBWithdrawal struct {
	AssetId uint16
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0xf4bf32c167ee6e782944cd1db8174729b46adcd3bc732e282cc4a80793933154.
//
// Solidity: event Withdrawal(uint16 assetId, uint128 amount)
func (_ZkBNB *ZkBNBFilterer) FilterWithdrawal(opts *bind.FilterOpts) (*ZkBNBWithdrawalIterator, error) {

	logs, sub, err := _ZkBNB.contract.FilterLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return &ZkBNBWithdrawalIterator{contract: _ZkBNB.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0xf4bf32c167ee6e782944cd1db8174729b46adcd3bc732e282cc4a80793933154.
//
// Solidity: event Withdrawal(uint16 assetId, uint128 amount)
func (_ZkBNB *ZkBNBFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *ZkBNBWithdrawal) (event.Subscription, error) {

	logs, sub, err := _ZkBNB.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkBNBWithdrawal)
				if err := _ZkBNB.contract.UnpackLog(event, "Withdrawal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawal is a log parse operation binding the contract event 0xf4bf32c167ee6e782944cd1db8174729b46adcd3bc732e282cc4a80793933154.
//
// Solidity: event Withdrawal(uint16 assetId, uint128 amount)
func (_ZkBNB *ZkBNBFilterer) ParseWithdrawal(log types.Log) (*ZkBNBWithdrawal, error) {
	event := new(ZkBNBWithdrawal)
	if err := _ZkBNB.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkBNBWithdrawalNFTPendingIterator is returned from FilterWithdrawalNFTPending and is used to iterate over the raw logs and unpacked data for WithdrawalNFTPending events raised by the ZkBNB contract.
type ZkBNBWithdrawalNFTPendingIterator struct {
	Event *ZkBNBWithdrawalNFTPending // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkBNBWithdrawalNFTPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkBNBWithdrawalNFTPending)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkBNBWithdrawalNFTPending)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkBNBWithdrawalNFTPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkBNBWithdrawalNFTPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkBNBWithdrawalNFTPending represents a WithdrawalNFTPending event raised by the ZkBNB contract.
type ZkBNBWithdrawalNFTPending struct {
	NftIndex *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalNFTPending is a free log retrieval operation binding the contract event 0x71aea045e572384e5b53812f175c12adb074001c9c13d379730d15188d3729bc.
//
// Solidity: event WithdrawalNFTPending(uint40 indexed nftIndex)
func (_ZkBNB *ZkBNBFilterer) FilterWithdrawalNFTPending(opts *bind.FilterOpts, nftIndex []*big.Int) (*ZkBNBWithdrawalNFTPendingIterator, error) {

	var nftIndexRule []interface{}
	for _, nftIndexItem := range nftIndex {
		nftIndexRule = append(nftIndexRule, nftIndexItem)
	}

	logs, sub, err := _ZkBNB.contract.FilterLogs(opts, "WithdrawalNFTPending", nftIndexRule)
	if err != nil {
		return nil, err
	}
	return &ZkBNBWithdrawalNFTPendingIterator{contract: _ZkBNB.contract, event: "WithdrawalNFTPending", logs: logs, sub: sub}, nil
}

// WatchWithdrawalNFTPending is a free log subscription operation binding the contract event 0x71aea045e572384e5b53812f175c12adb074001c9c13d379730d15188d3729bc.
//
// Solidity: event WithdrawalNFTPending(uint40 indexed nftIndex)
func (_ZkBNB *ZkBNBFilterer) WatchWithdrawalNFTPending(opts *bind.WatchOpts, sink chan<- *ZkBNBWithdrawalNFTPending, nftIndex []*big.Int) (event.Subscription, error) {

	var nftIndexRule []interface{}
	for _, nftIndexItem := range nftIndex {
		nftIndexRule = append(nftIndexRule, nftIndexItem)
	}

	logs, sub, err := _ZkBNB.contract.WatchLogs(opts, "WithdrawalNFTPending", nftIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkBNBWithdrawalNFTPending)
				if err := _ZkBNB.contract.UnpackLog(event, "WithdrawalNFTPending", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawalNFTPending is a log parse operation binding the contract event 0x71aea045e572384e5b53812f175c12adb074001c9c13d379730d15188d3729bc.
//
// Solidity: event WithdrawalNFTPending(uint40 indexed nftIndex)
func (_ZkBNB *ZkBNBFilterer) ParseWithdrawalNFTPending(log types.Log) (*ZkBNBWithdrawalNFTPending, error) {
	event := new(ZkBNBWithdrawalNFTPending)
	if err := _ZkBNB.contract.UnpackLog(event, "WithdrawalNFTPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
