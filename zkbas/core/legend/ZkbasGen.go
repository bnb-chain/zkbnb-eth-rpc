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

// OldZkbasCommitBlockInfo is an auto generated low-level Go binding around an user-defined struct.
type OldZkbasCommitBlockInfo struct {
	NewStateRoot      [32]byte
	PublicData        []byte
	Timestamp         *big.Int
	PublicDataOffsets []uint32
	BlockNumber       uint32
	BlockSize         uint16
}

// OldZkbasPairInfo is an auto generated low-level Go binding around an user-defined struct.
type OldZkbasPairInfo struct {
	TokenA               common.Address
	TokenB               common.Address
	FeeRate              uint16
	TreasuryAccountIndex uint32
	TreasuryRate         uint16
}

// OldZkbasVerifyAndExecuteBlockInfo is an auto generated low-level Go binding around an user-defined struct.
type OldZkbasVerifyAndExecuteBlockInfo struct {
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

// ZkbasMetaData contains all meta data concerning the Zkbas contract.
var ZkbasMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"BlockCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"BlockVerification\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"totalBlocksVerified\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"totalBlocksCommitted\",\"type\":\"uint32\"}],\"name\":\"BlocksRevert\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"pairIndex\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"asset0Id\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"asset1Id\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"feeRate\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"treasuryAccountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"treasuryRate\",\"type\":\"uint16\"}],\"name\":\"CreateTokenPair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"assetId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accountName\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"zkbasBlockNumber\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accountName\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"assetId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"DepositCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accountNameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nftContentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"creatorTreasuryRate\",\"type\":\"uint16\"}],\"name\":\"DepositNft\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DesertMode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"zkbasBlockId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tokenId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"FullExitCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"NewDefaultNFTFactory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_creatorAccountNameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"_collectionId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_factoryAddress\",\"type\":\"address\"}],\"name\":\"NewNFTFactory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serialId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumTxTypes.TxType\",\"name\":\"txType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expirationBlock\",\"type\":\"uint256\"}],\"name\":\"NewPriorityRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNoticePeriod\",\"type\":\"uint256\"}],\"name\":\"NoticePeriodChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"zkbasPubKeyX\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"zkbasPubKeyY\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"}],\"name\":\"RegisterZNS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"pairIndex\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"feeRate\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"treasuryAccountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"treasuryRate\",\"type\":\"uint16\"}],\"name\":\"UpdateTokenPair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftL1Address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftL1TokenId\",\"type\":\"uint256\"}],\"name\":\"WithdrawNft\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"assetId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"nftIndex\",\"type\":\"uint40\"}],\"name\":\"WithdrawalNFTPending\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_ACCOUNT_INDEX\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_AMOUNT_OF_REGISTERED_ASSETS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_DEPOSIT_AMOUNT\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FUNGIBLE_ASSET_ID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_NFT_INDEX\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECURITY_COUNCIL_MEMBERS_NUMBER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHORTEST_UPGRADE_NOTICE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SPECIAL_ACCOUNT_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SPECIAL_ACCOUNT_ID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TX_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_NOTICE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWAL_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateDesertMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"blockSize\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"priorityOperations\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"pendingOnchainOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structStorage.StoredBlockInfo\",\"name\":\"_lastCommittedBlockData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint32[]\",\"name\":\"publicDataOffsets\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"blockSize\",\"type\":\"uint16\"}],\"internalType\":\"structOldZkbas.CommitBlockInfo[]\",\"name\":\"_newBlocksData\",\"type\":\"tuple[]\"}],\"name\":\"commitBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cutUpgradeNoticePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultNFTFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint104\",\"name\":\"_amount\",\"type\":\"uint104\"},{\"internalType\":\"string\",\"name\":\"_accountName\",\"type\":\"string\"}],\"name\":\"depositBEP20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accountName\",\"type\":\"string\"}],\"name\":\"depositBNB\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accountName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_nftL1Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nftL1TokenId\",\"type\":\"uint256\"}],\"name\":\"depositNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"desertMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstPriorityRequestId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"accountNameHash\",\"type\":\"bytes32\"}],\"name\":\"getAddressByAccountNameHash\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_creatorAccountNameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_collectionId\",\"type\":\"uint32\"}],\"name\":\"getNFTFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNoticePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_assetAddr\",\"type\":\"address\"}],\"name\":\"getPendingBalance\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initializationParameters\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isReadyForUpgrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"nftFactories\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_zkbasPubKeyX\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_zkbasPubKeyY\",\"type\":\"bytes32\"}],\"name\":\"registerZNS\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accountName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"requestFullExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accountName\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"_nftIndex\",\"type\":\"uint32\"}],\"name\":\"requestFullExitNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"blockSize\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"priorityOperations\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"pendingOnchainOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structStorage.StoredBlockInfo[]\",\"name\":\"_blocksToRevert\",\"type\":\"tuple[]\"}],\"name\":\"revertBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractNFTFactory\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"setDefaultNFTFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"storedBlockHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBlocksCommitted\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBlocksVerified\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalOpenPriorityRequests\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTokenPairs\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_maxAmount\",\"type\":\"uint128\"}],\"name\":\"transferERC20\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"withdrawnAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"feeRate\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"treasuryAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"treasuryRate\",\"type\":\"uint16\"}],\"internalType\":\"structOldZkbas.PairInfo\",\"name\":\"_pairInfo\",\"type\":\"tuple\"}],\"name\":\"updatePairRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newVerifierAddress\",\"type\":\"address\"}],\"name\":\"updateZkbasVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"upgradeParameters\",\"type\":\"bytes\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeCanceled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeFinishes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeNoticePeriodStarted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradePreparationStarted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"blockSize\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"priorityOperations\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"pendingOnchainOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structStorage.StoredBlockInfo\",\"name\":\"blockHeader\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"pendingOnchainOpsPubData\",\"type\":\"bytes[]\"}],\"internalType\":\"structOldZkbas.VerifyAndExecuteBlockInfo[]\",\"name\":\"_blocks\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_proofs\",\"type\":\"uint256[]\"}],\"name\":\"verifyAndExecuteBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_amount\",\"type\":\"uint128\"}],\"name\":\"withdrawPendingBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_nftIndex\",\"type\":\"uint40\"}],\"name\":\"withdrawPendingNFTBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ZkbasABI is the input ABI used to generate the binding from.
// Deprecated: Use ZkbasMetaData.ABI instead.
var ZkbasABI = ZkbasMetaData.ABI

// Zkbas is an auto generated Go binding around an Ethereum contract.
type Zkbas struct {
	ZkbasCaller     // Read-only binding to the contract
	ZkbasTransactor // Write-only binding to the contract
	ZkbasFilterer   // Log filterer for contract events
}

// ZkbasCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZkbasCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkbasTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZkbasTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkbasFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZkbasFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkbasSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZkbasSession struct {
	Contract     *Zkbas            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZkbasCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZkbasCallerSession struct {
	Contract *ZkbasCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ZkbasTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZkbasTransactorSession struct {
	Contract     *ZkbasTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZkbasRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZkbasRaw struct {
	Contract *Zkbas // Generic contract binding to access the raw methods on
}

// ZkbasCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZkbasCallerRaw struct {
	Contract *ZkbasCaller // Generic read-only contract binding to access the raw methods on
}

// ZkbasTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZkbasTransactorRaw struct {
	Contract *ZkbasTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZkbas creates a new instance of Zkbas, bound to a specific deployed contract.
func NewZkbas(address common.Address, backend bind.ContractBackend) (*Zkbas, error) {
	contract, err := bindZkbas(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Zkbas{ZkbasCaller: ZkbasCaller{contract: contract}, ZkbasTransactor: ZkbasTransactor{contract: contract}, ZkbasFilterer: ZkbasFilterer{contract: contract}}, nil
}

// NewZkbasCaller creates a new read-only instance of Zkbas, bound to a specific deployed contract.
func NewZkbasCaller(address common.Address, caller bind.ContractCaller) (*ZkbasCaller, error) {
	contract, err := bindZkbas(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZkbasCaller{contract: contract}, nil
}

// NewZkbasTransactor creates a new write-only instance of Zkbas, bound to a specific deployed contract.
func NewZkbasTransactor(address common.Address, transactor bind.ContractTransactor) (*ZkbasTransactor, error) {
	contract, err := bindZkbas(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZkbasTransactor{contract: contract}, nil
}

// NewZkbasFilterer creates a new log filterer instance of Zkbas, bound to a specific deployed contract.
func NewZkbasFilterer(address common.Address, filterer bind.ContractFilterer) (*ZkbasFilterer, error) {
	contract, err := bindZkbas(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZkbasFilterer{contract: contract}, nil
}

// bindZkbas binds a generic wrapper to an already deployed contract.
func bindZkbas(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZkbasABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zkbas *ZkbasRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zkbas.Contract.ZkbasCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zkbas *ZkbasRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zkbas.Contract.ZkbasTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zkbas *ZkbasRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zkbas.Contract.ZkbasTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zkbas *ZkbasCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zkbas.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zkbas *ZkbasTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zkbas.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zkbas *ZkbasTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zkbas.Contract.contract.Transact(opts, method, params...)
}

// MAXACCOUNTINDEX is a free data retrieval call binding the contract method 0x437545f9.
//
// Solidity: function MAX_ACCOUNT_INDEX() view returns(uint32)
func (_Zkbas *ZkbasCaller) MAXACCOUNTINDEX(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "MAX_ACCOUNT_INDEX")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXACCOUNTINDEX is a free data retrieval call binding the contract method 0x437545f9.
//
// Solidity: function MAX_ACCOUNT_INDEX() view returns(uint32)
func (_Zkbas *ZkbasSession) MAXACCOUNTINDEX() (uint32, error) {
	return _Zkbas.Contract.MAXACCOUNTINDEX(&_Zkbas.CallOpts)
}

// MAXACCOUNTINDEX is a free data retrieval call binding the contract method 0x437545f9.
//
// Solidity: function MAX_ACCOUNT_INDEX() view returns(uint32)
func (_Zkbas *ZkbasCallerSession) MAXACCOUNTINDEX() (uint32, error) {
	return _Zkbas.Contract.MAXACCOUNTINDEX(&_Zkbas.CallOpts)
}

// MAXAMOUNTOFREGISTEREDASSETS is a free data retrieval call binding the contract method 0x0d360b7f.
//
// Solidity: function MAX_AMOUNT_OF_REGISTERED_ASSETS() view returns(uint16)
func (_Zkbas *ZkbasCaller) MAXAMOUNTOFREGISTEREDASSETS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "MAX_AMOUNT_OF_REGISTERED_ASSETS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXAMOUNTOFREGISTEREDASSETS is a free data retrieval call binding the contract method 0x0d360b7f.
//
// Solidity: function MAX_AMOUNT_OF_REGISTERED_ASSETS() view returns(uint16)
func (_Zkbas *ZkbasSession) MAXAMOUNTOFREGISTEREDASSETS() (uint16, error) {
	return _Zkbas.Contract.MAXAMOUNTOFREGISTEREDASSETS(&_Zkbas.CallOpts)
}

// MAXAMOUNTOFREGISTEREDASSETS is a free data retrieval call binding the contract method 0x0d360b7f.
//
// Solidity: function MAX_AMOUNT_OF_REGISTERED_ASSETS() view returns(uint16)
func (_Zkbas *ZkbasCallerSession) MAXAMOUNTOFREGISTEREDASSETS() (uint16, error) {
	return _Zkbas.Contract.MAXAMOUNTOFREGISTEREDASSETS(&_Zkbas.CallOpts)
}

// MAXDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x4c34a982.
//
// Solidity: function MAX_DEPOSIT_AMOUNT() view returns(uint128)
func (_Zkbas *ZkbasCaller) MAXDEPOSITAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "MAX_DEPOSIT_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x4c34a982.
//
// Solidity: function MAX_DEPOSIT_AMOUNT() view returns(uint128)
func (_Zkbas *ZkbasSession) MAXDEPOSITAMOUNT() (*big.Int, error) {
	return _Zkbas.Contract.MAXDEPOSITAMOUNT(&_Zkbas.CallOpts)
}

// MAXDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x4c34a982.
//
// Solidity: function MAX_DEPOSIT_AMOUNT() view returns(uint128)
func (_Zkbas *ZkbasCallerSession) MAXDEPOSITAMOUNT() (*big.Int, error) {
	return _Zkbas.Contract.MAXDEPOSITAMOUNT(&_Zkbas.CallOpts)
}

// MAXFUNGIBLEASSETID is a free data retrieval call binding the contract method 0x437da02f.
//
// Solidity: function MAX_FUNGIBLE_ASSET_ID() view returns(uint32)
func (_Zkbas *ZkbasCaller) MAXFUNGIBLEASSETID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "MAX_FUNGIBLE_ASSET_ID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXFUNGIBLEASSETID is a free data retrieval call binding the contract method 0x437da02f.
//
// Solidity: function MAX_FUNGIBLE_ASSET_ID() view returns(uint32)
func (_Zkbas *ZkbasSession) MAXFUNGIBLEASSETID() (uint32, error) {
	return _Zkbas.Contract.MAXFUNGIBLEASSETID(&_Zkbas.CallOpts)
}

// MAXFUNGIBLEASSETID is a free data retrieval call binding the contract method 0x437da02f.
//
// Solidity: function MAX_FUNGIBLE_ASSET_ID() view returns(uint32)
func (_Zkbas *ZkbasCallerSession) MAXFUNGIBLEASSETID() (uint32, error) {
	return _Zkbas.Contract.MAXFUNGIBLEASSETID(&_Zkbas.CallOpts)
}

// MAXNFTINDEX is a free data retrieval call binding the contract method 0x14791ad2.
//
// Solidity: function MAX_NFT_INDEX() view returns(uint40)
func (_Zkbas *ZkbasCaller) MAXNFTINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "MAX_NFT_INDEX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXNFTINDEX is a free data retrieval call binding the contract method 0x14791ad2.
//
// Solidity: function MAX_NFT_INDEX() view returns(uint40)
func (_Zkbas *ZkbasSession) MAXNFTINDEX() (*big.Int, error) {
	return _Zkbas.Contract.MAXNFTINDEX(&_Zkbas.CallOpts)
}

// MAXNFTINDEX is a free data retrieval call binding the contract method 0x14791ad2.
//
// Solidity: function MAX_NFT_INDEX() view returns(uint40)
func (_Zkbas *ZkbasCallerSession) MAXNFTINDEX() (*big.Int, error) {
	return _Zkbas.Contract.MAXNFTINDEX(&_Zkbas.CallOpts)
}

// SECURITYCOUNCILMEMBERSNUMBER is a free data retrieval call binding the contract method 0x4a51a71f.
//
// Solidity: function SECURITY_COUNCIL_MEMBERS_NUMBER() view returns(uint256)
func (_Zkbas *ZkbasCaller) SECURITYCOUNCILMEMBERSNUMBER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "SECURITY_COUNCIL_MEMBERS_NUMBER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECURITYCOUNCILMEMBERSNUMBER is a free data retrieval call binding the contract method 0x4a51a71f.
//
// Solidity: function SECURITY_COUNCIL_MEMBERS_NUMBER() view returns(uint256)
func (_Zkbas *ZkbasSession) SECURITYCOUNCILMEMBERSNUMBER() (*big.Int, error) {
	return _Zkbas.Contract.SECURITYCOUNCILMEMBERSNUMBER(&_Zkbas.CallOpts)
}

// SECURITYCOUNCILMEMBERSNUMBER is a free data retrieval call binding the contract method 0x4a51a71f.
//
// Solidity: function SECURITY_COUNCIL_MEMBERS_NUMBER() view returns(uint256)
func (_Zkbas *ZkbasCallerSession) SECURITYCOUNCILMEMBERSNUMBER() (*big.Int, error) {
	return _Zkbas.Contract.SECURITYCOUNCILMEMBERSNUMBER(&_Zkbas.CallOpts)
}

// SHORTESTUPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0x85053581.
//
// Solidity: function SHORTEST_UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_Zkbas *ZkbasCaller) SHORTESTUPGRADENOTICEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "SHORTEST_UPGRADE_NOTICE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SHORTESTUPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0x85053581.
//
// Solidity: function SHORTEST_UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_Zkbas *ZkbasSession) SHORTESTUPGRADENOTICEPERIOD() (*big.Int, error) {
	return _Zkbas.Contract.SHORTESTUPGRADENOTICEPERIOD(&_Zkbas.CallOpts)
}

// SHORTESTUPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0x85053581.
//
// Solidity: function SHORTEST_UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_Zkbas *ZkbasCallerSession) SHORTESTUPGRADENOTICEPERIOD() (*big.Int, error) {
	return _Zkbas.Contract.SHORTESTUPGRADENOTICEPERIOD(&_Zkbas.CallOpts)
}

// SPECIALACCOUNTADDRESS is a free data retrieval call binding the contract method 0x7ea399c1.
//
// Solidity: function SPECIAL_ACCOUNT_ADDRESS() view returns(address)
func (_Zkbas *ZkbasCaller) SPECIALACCOUNTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "SPECIAL_ACCOUNT_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SPECIALACCOUNTADDRESS is a free data retrieval call binding the contract method 0x7ea399c1.
//
// Solidity: function SPECIAL_ACCOUNT_ADDRESS() view returns(address)
func (_Zkbas *ZkbasSession) SPECIALACCOUNTADDRESS() (common.Address, error) {
	return _Zkbas.Contract.SPECIALACCOUNTADDRESS(&_Zkbas.CallOpts)
}

// SPECIALACCOUNTADDRESS is a free data retrieval call binding the contract method 0x7ea399c1.
//
// Solidity: function SPECIAL_ACCOUNT_ADDRESS() view returns(address)
func (_Zkbas *ZkbasCallerSession) SPECIALACCOUNTADDRESS() (common.Address, error) {
	return _Zkbas.Contract.SPECIALACCOUNTADDRESS(&_Zkbas.CallOpts)
}

// SPECIALACCOUNTID is a free data retrieval call binding the contract method 0x4242d5b3.
//
// Solidity: function SPECIAL_ACCOUNT_ID() view returns(uint32)
func (_Zkbas *ZkbasCaller) SPECIALACCOUNTID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "SPECIAL_ACCOUNT_ID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SPECIALACCOUNTID is a free data retrieval call binding the contract method 0x4242d5b3.
//
// Solidity: function SPECIAL_ACCOUNT_ID() view returns(uint32)
func (_Zkbas *ZkbasSession) SPECIALACCOUNTID() (uint32, error) {
	return _Zkbas.Contract.SPECIALACCOUNTID(&_Zkbas.CallOpts)
}

// SPECIALACCOUNTID is a free data retrieval call binding the contract method 0x4242d5b3.
//
// Solidity: function SPECIAL_ACCOUNT_ID() view returns(uint32)
func (_Zkbas *ZkbasCallerSession) SPECIALACCOUNTID() (uint32, error) {
	return _Zkbas.Contract.SPECIALACCOUNTID(&_Zkbas.CallOpts)
}

// TXSIZE is a free data retrieval call binding the contract method 0xe6e3c012.
//
// Solidity: function TX_SIZE() view returns(uint256)
func (_Zkbas *ZkbasCaller) TXSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "TX_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TXSIZE is a free data retrieval call binding the contract method 0xe6e3c012.
//
// Solidity: function TX_SIZE() view returns(uint256)
func (_Zkbas *ZkbasSession) TXSIZE() (*big.Int, error) {
	return _Zkbas.Contract.TXSIZE(&_Zkbas.CallOpts)
}

// TXSIZE is a free data retrieval call binding the contract method 0xe6e3c012.
//
// Solidity: function TX_SIZE() view returns(uint256)
func (_Zkbas *ZkbasCallerSession) TXSIZE() (*big.Int, error) {
	return _Zkbas.Contract.TXSIZE(&_Zkbas.CallOpts)
}

// UPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0xcc375fb7.
//
// Solidity: function UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_Zkbas *ZkbasCaller) UPGRADENOTICEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "UPGRADE_NOTICE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0xcc375fb7.
//
// Solidity: function UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_Zkbas *ZkbasSession) UPGRADENOTICEPERIOD() (*big.Int, error) {
	return _Zkbas.Contract.UPGRADENOTICEPERIOD(&_Zkbas.CallOpts)
}

// UPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0xcc375fb7.
//
// Solidity: function UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_Zkbas *ZkbasCallerSession) UPGRADENOTICEPERIOD() (*big.Int, error) {
	return _Zkbas.Contract.UPGRADENOTICEPERIOD(&_Zkbas.CallOpts)
}

// WITHDRAWALGASLIMIT is a free data retrieval call binding the contract method 0xc701f955.
//
// Solidity: function WITHDRAWAL_GAS_LIMIT() view returns(uint256)
func (_Zkbas *ZkbasCaller) WITHDRAWALGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "WITHDRAWAL_GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WITHDRAWALGASLIMIT is a free data retrieval call binding the contract method 0xc701f955.
//
// Solidity: function WITHDRAWAL_GAS_LIMIT() view returns(uint256)
func (_Zkbas *ZkbasSession) WITHDRAWALGASLIMIT() (*big.Int, error) {
	return _Zkbas.Contract.WITHDRAWALGASLIMIT(&_Zkbas.CallOpts)
}

// WITHDRAWALGASLIMIT is a free data retrieval call binding the contract method 0xc701f955.
//
// Solidity: function WITHDRAWAL_GAS_LIMIT() view returns(uint256)
func (_Zkbas *ZkbasCallerSession) WITHDRAWALGASLIMIT() (*big.Int, error) {
	return _Zkbas.Contract.WITHDRAWALGASLIMIT(&_Zkbas.CallOpts)
}

// DefaultNFTFactory is a free data retrieval call binding the contract method 0x940f19c0.
//
// Solidity: function defaultNFTFactory() view returns(address)
func (_Zkbas *ZkbasCaller) DefaultNFTFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "defaultNFTFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultNFTFactory is a free data retrieval call binding the contract method 0x940f19c0.
//
// Solidity: function defaultNFTFactory() view returns(address)
func (_Zkbas *ZkbasSession) DefaultNFTFactory() (common.Address, error) {
	return _Zkbas.Contract.DefaultNFTFactory(&_Zkbas.CallOpts)
}

// DefaultNFTFactory is a free data retrieval call binding the contract method 0x940f19c0.
//
// Solidity: function defaultNFTFactory() view returns(address)
func (_Zkbas *ZkbasCallerSession) DefaultNFTFactory() (common.Address, error) {
	return _Zkbas.Contract.DefaultNFTFactory(&_Zkbas.CallOpts)
}

// DesertMode is a free data retrieval call binding the contract method 0x02cfb563.
//
// Solidity: function desertMode() view returns(bool)
func (_Zkbas *ZkbasCaller) DesertMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "desertMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DesertMode is a free data retrieval call binding the contract method 0x02cfb563.
//
// Solidity: function desertMode() view returns(bool)
func (_Zkbas *ZkbasSession) DesertMode() (bool, error) {
	return _Zkbas.Contract.DesertMode(&_Zkbas.CallOpts)
}

// DesertMode is a free data retrieval call binding the contract method 0x02cfb563.
//
// Solidity: function desertMode() view returns(bool)
func (_Zkbas *ZkbasCallerSession) DesertMode() (bool, error) {
	return _Zkbas.Contract.DesertMode(&_Zkbas.CallOpts)
}

// FirstPriorityRequestId is a free data retrieval call binding the contract method 0x67708dae.
//
// Solidity: function firstPriorityRequestId() view returns(uint64)
func (_Zkbas *ZkbasCaller) FirstPriorityRequestId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "firstPriorityRequestId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// FirstPriorityRequestId is a free data retrieval call binding the contract method 0x67708dae.
//
// Solidity: function firstPriorityRequestId() view returns(uint64)
func (_Zkbas *ZkbasSession) FirstPriorityRequestId() (uint64, error) {
	return _Zkbas.Contract.FirstPriorityRequestId(&_Zkbas.CallOpts)
}

// FirstPriorityRequestId is a free data retrieval call binding the contract method 0x67708dae.
//
// Solidity: function firstPriorityRequestId() view returns(uint64)
func (_Zkbas *ZkbasCallerSession) FirstPriorityRequestId() (uint64, error) {
	return _Zkbas.Contract.FirstPriorityRequestId(&_Zkbas.CallOpts)
}

// GetAddressByAccountNameHash is a free data retrieval call binding the contract method 0x0b8f1c0c.
//
// Solidity: function getAddressByAccountNameHash(bytes32 accountNameHash) view returns(address)
func (_Zkbas *ZkbasCaller) GetAddressByAccountNameHash(opts *bind.CallOpts, accountNameHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "getAddressByAccountNameHash", accountNameHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressByAccountNameHash is a free data retrieval call binding the contract method 0x0b8f1c0c.
//
// Solidity: function getAddressByAccountNameHash(bytes32 accountNameHash) view returns(address)
func (_Zkbas *ZkbasSession) GetAddressByAccountNameHash(accountNameHash [32]byte) (common.Address, error) {
	return _Zkbas.Contract.GetAddressByAccountNameHash(&_Zkbas.CallOpts, accountNameHash)
}

// GetAddressByAccountNameHash is a free data retrieval call binding the contract method 0x0b8f1c0c.
//
// Solidity: function getAddressByAccountNameHash(bytes32 accountNameHash) view returns(address)
func (_Zkbas *ZkbasCallerSession) GetAddressByAccountNameHash(accountNameHash [32]byte) (common.Address, error) {
	return _Zkbas.Contract.GetAddressByAccountNameHash(&_Zkbas.CallOpts, accountNameHash)
}

// GetNFTFactory is a free data retrieval call binding the contract method 0x76b4da60.
//
// Solidity: function getNFTFactory(bytes32 _creatorAccountNameHash, uint32 _collectionId) view returns(address)
func (_Zkbas *ZkbasCaller) GetNFTFactory(opts *bind.CallOpts, _creatorAccountNameHash [32]byte, _collectionId uint32) (common.Address, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "getNFTFactory", _creatorAccountNameHash, _collectionId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNFTFactory is a free data retrieval call binding the contract method 0x76b4da60.
//
// Solidity: function getNFTFactory(bytes32 _creatorAccountNameHash, uint32 _collectionId) view returns(address)
func (_Zkbas *ZkbasSession) GetNFTFactory(_creatorAccountNameHash [32]byte, _collectionId uint32) (common.Address, error) {
	return _Zkbas.Contract.GetNFTFactory(&_Zkbas.CallOpts, _creatorAccountNameHash, _collectionId)
}

// GetNFTFactory is a free data retrieval call binding the contract method 0x76b4da60.
//
// Solidity: function getNFTFactory(bytes32 _creatorAccountNameHash, uint32 _collectionId) view returns(address)
func (_Zkbas *ZkbasCallerSession) GetNFTFactory(_creatorAccountNameHash [32]byte, _collectionId uint32) (common.Address, error) {
	return _Zkbas.Contract.GetNFTFactory(&_Zkbas.CallOpts, _creatorAccountNameHash, _collectionId)
}

// GetNoticePeriod is a free data retrieval call binding the contract method 0x2a3174f4.
//
// Solidity: function getNoticePeriod() pure returns(uint256)
func (_Zkbas *ZkbasCaller) GetNoticePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "getNoticePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNoticePeriod is a free data retrieval call binding the contract method 0x2a3174f4.
//
// Solidity: function getNoticePeriod() pure returns(uint256)
func (_Zkbas *ZkbasSession) GetNoticePeriod() (*big.Int, error) {
	return _Zkbas.Contract.GetNoticePeriod(&_Zkbas.CallOpts)
}

// GetNoticePeriod is a free data retrieval call binding the contract method 0x2a3174f4.
//
// Solidity: function getNoticePeriod() pure returns(uint256)
func (_Zkbas *ZkbasCallerSession) GetNoticePeriod() (*big.Int, error) {
	return _Zkbas.Contract.GetNoticePeriod(&_Zkbas.CallOpts)
}

// GetPendingBalance is a free data retrieval call binding the contract method 0x5aca41f6.
//
// Solidity: function getPendingBalance(address _address, address _assetAddr) view returns(uint128)
func (_Zkbas *ZkbasCaller) GetPendingBalance(opts *bind.CallOpts, _address common.Address, _assetAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "getPendingBalance", _address, _assetAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingBalance is a free data retrieval call binding the contract method 0x5aca41f6.
//
// Solidity: function getPendingBalance(address _address, address _assetAddr) view returns(uint128)
func (_Zkbas *ZkbasSession) GetPendingBalance(_address common.Address, _assetAddr common.Address) (*big.Int, error) {
	return _Zkbas.Contract.GetPendingBalance(&_Zkbas.CallOpts, _address, _assetAddr)
}

// GetPendingBalance is a free data retrieval call binding the contract method 0x5aca41f6.
//
// Solidity: function getPendingBalance(address _address, address _assetAddr) view returns(uint128)
func (_Zkbas *ZkbasCallerSession) GetPendingBalance(_address common.Address, _assetAddr common.Address) (*big.Int, error) {
	return _Zkbas.Contract.GetPendingBalance(&_Zkbas.CallOpts, _address, _assetAddr)
}

// IsReadyForUpgrade is a free data retrieval call binding the contract method 0x8773334c.
//
// Solidity: function isReadyForUpgrade() view returns(bool)
func (_Zkbas *ZkbasCaller) IsReadyForUpgrade(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "isReadyForUpgrade")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReadyForUpgrade is a free data retrieval call binding the contract method 0x8773334c.
//
// Solidity: function isReadyForUpgrade() view returns(bool)
func (_Zkbas *ZkbasSession) IsReadyForUpgrade() (bool, error) {
	return _Zkbas.Contract.IsReadyForUpgrade(&_Zkbas.CallOpts)
}

// IsReadyForUpgrade is a free data retrieval call binding the contract method 0x8773334c.
//
// Solidity: function isReadyForUpgrade() view returns(bool)
func (_Zkbas *ZkbasCallerSession) IsReadyForUpgrade() (bool, error) {
	return _Zkbas.Contract.IsReadyForUpgrade(&_Zkbas.CallOpts)
}

// NftFactories is a free data retrieval call binding the contract method 0x16335e06.
//
// Solidity: function nftFactories(bytes32 , uint32 ) view returns(address)
func (_Zkbas *ZkbasCaller) NftFactories(opts *bind.CallOpts, arg0 [32]byte, arg1 uint32) (common.Address, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "nftFactories", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftFactories is a free data retrieval call binding the contract method 0x16335e06.
//
// Solidity: function nftFactories(bytes32 , uint32 ) view returns(address)
func (_Zkbas *ZkbasSession) NftFactories(arg0 [32]byte, arg1 uint32) (common.Address, error) {
	return _Zkbas.Contract.NftFactories(&_Zkbas.CallOpts, arg0, arg1)
}

// NftFactories is a free data retrieval call binding the contract method 0x16335e06.
//
// Solidity: function nftFactories(bytes32 , uint32 ) view returns(address)
func (_Zkbas *ZkbasCallerSession) NftFactories(arg0 [32]byte, arg1 uint32) (common.Address, error) {
	return _Zkbas.Contract.NftFactories(&_Zkbas.CallOpts, arg0, arg1)
}

// StateRoot is a free data retrieval call binding the contract method 0x9588eca2.
//
// Solidity: function stateRoot() view returns(bytes32)
func (_Zkbas *ZkbasCaller) StateRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "stateRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateRoot is a free data retrieval call binding the contract method 0x9588eca2.
//
// Solidity: function stateRoot() view returns(bytes32)
func (_Zkbas *ZkbasSession) StateRoot() ([32]byte, error) {
	return _Zkbas.Contract.StateRoot(&_Zkbas.CallOpts)
}

// StateRoot is a free data retrieval call binding the contract method 0x9588eca2.
//
// Solidity: function stateRoot() view returns(bytes32)
func (_Zkbas *ZkbasCallerSession) StateRoot() ([32]byte, error) {
	return _Zkbas.Contract.StateRoot(&_Zkbas.CallOpts)
}

// StoredBlockHashes is a free data retrieval call binding the contract method 0x9ba0d146.
//
// Solidity: function storedBlockHashes(uint32 ) view returns(bytes32)
func (_Zkbas *ZkbasCaller) StoredBlockHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "storedBlockHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StoredBlockHashes is a free data retrieval call binding the contract method 0x9ba0d146.
//
// Solidity: function storedBlockHashes(uint32 ) view returns(bytes32)
func (_Zkbas *ZkbasSession) StoredBlockHashes(arg0 uint32) ([32]byte, error) {
	return _Zkbas.Contract.StoredBlockHashes(&_Zkbas.CallOpts, arg0)
}

// StoredBlockHashes is a free data retrieval call binding the contract method 0x9ba0d146.
//
// Solidity: function storedBlockHashes(uint32 ) view returns(bytes32)
func (_Zkbas *ZkbasCallerSession) StoredBlockHashes(arg0 uint32) ([32]byte, error) {
	return _Zkbas.Contract.StoredBlockHashes(&_Zkbas.CallOpts, arg0)
}

// TotalBlocksCommitted is a free data retrieval call binding the contract method 0xfaf4d8cb.
//
// Solidity: function totalBlocksCommitted() view returns(uint32)
func (_Zkbas *ZkbasCaller) TotalBlocksCommitted(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "totalBlocksCommitted")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TotalBlocksCommitted is a free data retrieval call binding the contract method 0xfaf4d8cb.
//
// Solidity: function totalBlocksCommitted() view returns(uint32)
func (_Zkbas *ZkbasSession) TotalBlocksCommitted() (uint32, error) {
	return _Zkbas.Contract.TotalBlocksCommitted(&_Zkbas.CallOpts)
}

// TotalBlocksCommitted is a free data retrieval call binding the contract method 0xfaf4d8cb.
//
// Solidity: function totalBlocksCommitted() view returns(uint32)
func (_Zkbas *ZkbasCallerSession) TotalBlocksCommitted() (uint32, error) {
	return _Zkbas.Contract.TotalBlocksCommitted(&_Zkbas.CallOpts)
}

// TotalBlocksVerified is a free data retrieval call binding the contract method 0x2d24006c.
//
// Solidity: function totalBlocksVerified() view returns(uint32)
func (_Zkbas *ZkbasCaller) TotalBlocksVerified(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "totalBlocksVerified")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TotalBlocksVerified is a free data retrieval call binding the contract method 0x2d24006c.
//
// Solidity: function totalBlocksVerified() view returns(uint32)
func (_Zkbas *ZkbasSession) TotalBlocksVerified() (uint32, error) {
	return _Zkbas.Contract.TotalBlocksVerified(&_Zkbas.CallOpts)
}

// TotalBlocksVerified is a free data retrieval call binding the contract method 0x2d24006c.
//
// Solidity: function totalBlocksVerified() view returns(uint32)
func (_Zkbas *ZkbasCallerSession) TotalBlocksVerified() (uint32, error) {
	return _Zkbas.Contract.TotalBlocksVerified(&_Zkbas.CallOpts)
}

// TotalOpenPriorityRequests is a free data retrieval call binding the contract method 0xc57b22be.
//
// Solidity: function totalOpenPriorityRequests() view returns(uint64)
func (_Zkbas *ZkbasCaller) TotalOpenPriorityRequests(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "totalOpenPriorityRequests")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalOpenPriorityRequests is a free data retrieval call binding the contract method 0xc57b22be.
//
// Solidity: function totalOpenPriorityRequests() view returns(uint64)
func (_Zkbas *ZkbasSession) TotalOpenPriorityRequests() (uint64, error) {
	return _Zkbas.Contract.TotalOpenPriorityRequests(&_Zkbas.CallOpts)
}

// TotalOpenPriorityRequests is a free data retrieval call binding the contract method 0xc57b22be.
//
// Solidity: function totalOpenPriorityRequests() view returns(uint64)
func (_Zkbas *ZkbasCallerSession) TotalOpenPriorityRequests() (uint64, error) {
	return _Zkbas.Contract.TotalOpenPriorityRequests(&_Zkbas.CallOpts)
}

// TotalTokenPairs is a free data retrieval call binding the contract method 0xd0ad718d.
//
// Solidity: function totalTokenPairs() view returns(uint16)
func (_Zkbas *ZkbasCaller) TotalTokenPairs(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Zkbas.contract.Call(opts, &out, "totalTokenPairs")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// TotalTokenPairs is a free data retrieval call binding the contract method 0xd0ad718d.
//
// Solidity: function totalTokenPairs() view returns(uint16)
func (_Zkbas *ZkbasSession) TotalTokenPairs() (uint16, error) {
	return _Zkbas.Contract.TotalTokenPairs(&_Zkbas.CallOpts)
}

// TotalTokenPairs is a free data retrieval call binding the contract method 0xd0ad718d.
//
// Solidity: function totalTokenPairs() view returns(uint16)
func (_Zkbas *ZkbasCallerSession) TotalTokenPairs() (uint16, error) {
	return _Zkbas.Contract.TotalTokenPairs(&_Zkbas.CallOpts)
}

// ActivateDesertMode is a paid mutator transaction binding the contract method 0x22b22256.
//
// Solidity: function activateDesertMode() returns(bool)
func (_Zkbas *ZkbasTransactor) ActivateDesertMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "activateDesertMode")
}

// ActivateDesertMode is a paid mutator transaction binding the contract method 0x22b22256.
//
// Solidity: function activateDesertMode() returns(bool)
func (_Zkbas *ZkbasSession) ActivateDesertMode() (*types.Transaction, error) {
	return _Zkbas.Contract.ActivateDesertMode(&_Zkbas.TransactOpts)
}

// ActivateDesertMode is a paid mutator transaction binding the contract method 0x22b22256.
//
// Solidity: function activateDesertMode() returns(bool)
func (_Zkbas *ZkbasTransactorSession) ActivateDesertMode() (*types.Transaction, error) {
	return _Zkbas.Contract.ActivateDesertMode(&_Zkbas.TransactOpts)
}

// CommitBlocks is a paid mutator transaction binding the contract method 0x2a54f972.
//
// Solidity: function commitBlocks((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32) _lastCommittedBlockData, (bytes32,bytes,uint256,uint32[],uint32,uint16)[] _newBlocksData) returns()
func (_Zkbas *ZkbasTransactor) CommitBlocks(opts *bind.TransactOpts, _lastCommittedBlockData StorageStoredBlockInfo, _newBlocksData []OldZkbasCommitBlockInfo) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "commitBlocks", _lastCommittedBlockData, _newBlocksData)
}

// CommitBlocks is a paid mutator transaction binding the contract method 0x2a54f972.
//
// Solidity: function commitBlocks((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32) _lastCommittedBlockData, (bytes32,bytes,uint256,uint32[],uint32,uint16)[] _newBlocksData) returns()
func (_Zkbas *ZkbasSession) CommitBlocks(_lastCommittedBlockData StorageStoredBlockInfo, _newBlocksData []OldZkbasCommitBlockInfo) (*types.Transaction, error) {
	return _Zkbas.Contract.CommitBlocks(&_Zkbas.TransactOpts, _lastCommittedBlockData, _newBlocksData)
}

// CommitBlocks is a paid mutator transaction binding the contract method 0x2a54f972.
//
// Solidity: function commitBlocks((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32) _lastCommittedBlockData, (bytes32,bytes,uint256,uint32[],uint32,uint16)[] _newBlocksData) returns()
func (_Zkbas *ZkbasTransactorSession) CommitBlocks(_lastCommittedBlockData StorageStoredBlockInfo, _newBlocksData []OldZkbasCommitBlockInfo) (*types.Transaction, error) {
	return _Zkbas.Contract.CommitBlocks(&_Zkbas.TransactOpts, _lastCommittedBlockData, _newBlocksData)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address _tokenA, address _tokenB) returns()
func (_Zkbas *ZkbasTransactor) CreatePair(opts *bind.TransactOpts, _tokenA common.Address, _tokenB common.Address) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "createPair", _tokenA, _tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address _tokenA, address _tokenB) returns()
func (_Zkbas *ZkbasSession) CreatePair(_tokenA common.Address, _tokenB common.Address) (*types.Transaction, error) {
	return _Zkbas.Contract.CreatePair(&_Zkbas.TransactOpts, _tokenA, _tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address _tokenA, address _tokenB) returns()
func (_Zkbas *ZkbasTransactorSession) CreatePair(_tokenA common.Address, _tokenB common.Address) (*types.Transaction, error) {
	return _Zkbas.Contract.CreatePair(&_Zkbas.TransactOpts, _tokenA, _tokenB)
}

// CutUpgradeNoticePeriod is a paid mutator transaction binding the contract method 0x3e71e1e7.
//
// Solidity: function cutUpgradeNoticePeriod() returns()
func (_Zkbas *ZkbasTransactor) CutUpgradeNoticePeriod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "cutUpgradeNoticePeriod")
}

// CutUpgradeNoticePeriod is a paid mutator transaction binding the contract method 0x3e71e1e7.
//
// Solidity: function cutUpgradeNoticePeriod() returns()
func (_Zkbas *ZkbasSession) CutUpgradeNoticePeriod() (*types.Transaction, error) {
	return _Zkbas.Contract.CutUpgradeNoticePeriod(&_Zkbas.TransactOpts)
}

// CutUpgradeNoticePeriod is a paid mutator transaction binding the contract method 0x3e71e1e7.
//
// Solidity: function cutUpgradeNoticePeriod() returns()
func (_Zkbas *ZkbasTransactorSession) CutUpgradeNoticePeriod() (*types.Transaction, error) {
	return _Zkbas.Contract.CutUpgradeNoticePeriod(&_Zkbas.TransactOpts)
}

// DepositBEP20 is a paid mutator transaction binding the contract method 0x1caf5d25.
//
// Solidity: function depositBEP20(address _token, uint104 _amount, string _accountName) returns()
func (_Zkbas *ZkbasTransactor) DepositBEP20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _accountName string) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "depositBEP20", _token, _amount, _accountName)
}

// DepositBEP20 is a paid mutator transaction binding the contract method 0x1caf5d25.
//
// Solidity: function depositBEP20(address _token, uint104 _amount, string _accountName) returns()
func (_Zkbas *ZkbasSession) DepositBEP20(_token common.Address, _amount *big.Int, _accountName string) (*types.Transaction, error) {
	return _Zkbas.Contract.DepositBEP20(&_Zkbas.TransactOpts, _token, _amount, _accountName)
}

// DepositBEP20 is a paid mutator transaction binding the contract method 0x1caf5d25.
//
// Solidity: function depositBEP20(address _token, uint104 _amount, string _accountName) returns()
func (_Zkbas *ZkbasTransactorSession) DepositBEP20(_token common.Address, _amount *big.Int, _accountName string) (*types.Transaction, error) {
	return _Zkbas.Contract.DepositBEP20(&_Zkbas.TransactOpts, _token, _amount, _accountName)
}

// DepositBNB is a paid mutator transaction binding the contract method 0x51344683.
//
// Solidity: function depositBNB(string _accountName) payable returns()
func (_Zkbas *ZkbasTransactor) DepositBNB(opts *bind.TransactOpts, _accountName string) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "depositBNB", _accountName)
}

// DepositBNB is a paid mutator transaction binding the contract method 0x51344683.
//
// Solidity: function depositBNB(string _accountName) payable returns()
func (_Zkbas *ZkbasSession) DepositBNB(_accountName string) (*types.Transaction, error) {
	return _Zkbas.Contract.DepositBNB(&_Zkbas.TransactOpts, _accountName)
}

// DepositBNB is a paid mutator transaction binding the contract method 0x51344683.
//
// Solidity: function depositBNB(string _accountName) payable returns()
func (_Zkbas *ZkbasTransactorSession) DepositBNB(_accountName string) (*types.Transaction, error) {
	return _Zkbas.Contract.DepositBNB(&_Zkbas.TransactOpts, _accountName)
}

// DepositNft is a paid mutator transaction binding the contract method 0xfb99514b.
//
// Solidity: function depositNft(string _accountName, address _nftL1Address, uint256 _nftL1TokenId) returns()
func (_Zkbas *ZkbasTransactor) DepositNft(opts *bind.TransactOpts, _accountName string, _nftL1Address common.Address, _nftL1TokenId *big.Int) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "depositNft", _accountName, _nftL1Address, _nftL1TokenId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xfb99514b.
//
// Solidity: function depositNft(string _accountName, address _nftL1Address, uint256 _nftL1TokenId) returns()
func (_Zkbas *ZkbasSession) DepositNft(_accountName string, _nftL1Address common.Address, _nftL1TokenId *big.Int) (*types.Transaction, error) {
	return _Zkbas.Contract.DepositNft(&_Zkbas.TransactOpts, _accountName, _nftL1Address, _nftL1TokenId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xfb99514b.
//
// Solidity: function depositNft(string _accountName, address _nftL1Address, uint256 _nftL1TokenId) returns()
func (_Zkbas *ZkbasTransactorSession) DepositNft(_accountName string, _nftL1Address common.Address, _nftL1TokenId *big.Int) (*types.Transaction, error) {
	return _Zkbas.Contract.DepositNft(&_Zkbas.TransactOpts, _accountName, _nftL1Address, _nftL1TokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializationParameters) returns()
func (_Zkbas *ZkbasTransactor) Initialize(opts *bind.TransactOpts, initializationParameters []byte) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "initialize", initializationParameters)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializationParameters) returns()
func (_Zkbas *ZkbasSession) Initialize(initializationParameters []byte) (*types.Transaction, error) {
	return _Zkbas.Contract.Initialize(&_Zkbas.TransactOpts, initializationParameters)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializationParameters) returns()
func (_Zkbas *ZkbasTransactorSession) Initialize(initializationParameters []byte) (*types.Transaction, error) {
	return _Zkbas.Contract.Initialize(&_Zkbas.TransactOpts, initializationParameters)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Zkbas *ZkbasTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Zkbas *ZkbasSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Zkbas.Contract.OnERC721Received(&_Zkbas.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Zkbas *ZkbasTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Zkbas.Contract.OnERC721Received(&_Zkbas.TransactOpts, operator, from, tokenId, data)
}

// RegisterZNS is a paid mutator transaction binding the contract method 0x3fdeb67d.
//
// Solidity: function registerZNS(string _name, address _owner, bytes32 _zkbasPubKeyX, bytes32 _zkbasPubKeyY) payable returns()
func (_Zkbas *ZkbasTransactor) RegisterZNS(opts *bind.TransactOpts, _name string, _owner common.Address, _zkbasPubKeyX [32]byte, _zkbasPubKeyY [32]byte) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "registerZNS", _name, _owner, _zkbasPubKeyX, _zkbasPubKeyY)
}

// RegisterZNS is a paid mutator transaction binding the contract method 0x3fdeb67d.
//
// Solidity: function registerZNS(string _name, address _owner, bytes32 _zkbasPubKeyX, bytes32 _zkbasPubKeyY) payable returns()
func (_Zkbas *ZkbasSession) RegisterZNS(_name string, _owner common.Address, _zkbasPubKeyX [32]byte, _zkbasPubKeyY [32]byte) (*types.Transaction, error) {
	return _Zkbas.Contract.RegisterZNS(&_Zkbas.TransactOpts, _name, _owner, _zkbasPubKeyX, _zkbasPubKeyY)
}

// RegisterZNS is a paid mutator transaction binding the contract method 0x3fdeb67d.
//
// Solidity: function registerZNS(string _name, address _owner, bytes32 _zkbasPubKeyX, bytes32 _zkbasPubKeyY) payable returns()
func (_Zkbas *ZkbasTransactorSession) RegisterZNS(_name string, _owner common.Address, _zkbasPubKeyX [32]byte, _zkbasPubKeyY [32]byte) (*types.Transaction, error) {
	return _Zkbas.Contract.RegisterZNS(&_Zkbas.TransactOpts, _name, _owner, _zkbasPubKeyX, _zkbasPubKeyY)
}

// RequestFullExit is a paid mutator transaction binding the contract method 0x8da64c7f.
//
// Solidity: function requestFullExit(string _accountName, address _asset) returns()
func (_Zkbas *ZkbasTransactor) RequestFullExit(opts *bind.TransactOpts, _accountName string, _asset common.Address) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "requestFullExit", _accountName, _asset)
}

// RequestFullExit is a paid mutator transaction binding the contract method 0x8da64c7f.
//
// Solidity: function requestFullExit(string _accountName, address _asset) returns()
func (_Zkbas *ZkbasSession) RequestFullExit(_accountName string, _asset common.Address) (*types.Transaction, error) {
	return _Zkbas.Contract.RequestFullExit(&_Zkbas.TransactOpts, _accountName, _asset)
}

// RequestFullExit is a paid mutator transaction binding the contract method 0x8da64c7f.
//
// Solidity: function requestFullExit(string _accountName, address _asset) returns()
func (_Zkbas *ZkbasTransactorSession) RequestFullExit(_accountName string, _asset common.Address) (*types.Transaction, error) {
	return _Zkbas.Contract.RequestFullExit(&_Zkbas.TransactOpts, _accountName, _asset)
}

// RequestFullExitNft is a paid mutator transaction binding the contract method 0x1bd24317.
//
// Solidity: function requestFullExitNft(string _accountName, uint32 _nftIndex) returns()
func (_Zkbas *ZkbasTransactor) RequestFullExitNft(opts *bind.TransactOpts, _accountName string, _nftIndex uint32) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "requestFullExitNft", _accountName, _nftIndex)
}

// RequestFullExitNft is a paid mutator transaction binding the contract method 0x1bd24317.
//
// Solidity: function requestFullExitNft(string _accountName, uint32 _nftIndex) returns()
func (_Zkbas *ZkbasSession) RequestFullExitNft(_accountName string, _nftIndex uint32) (*types.Transaction, error) {
	return _Zkbas.Contract.RequestFullExitNft(&_Zkbas.TransactOpts, _accountName, _nftIndex)
}

// RequestFullExitNft is a paid mutator transaction binding the contract method 0x1bd24317.
//
// Solidity: function requestFullExitNft(string _accountName, uint32 _nftIndex) returns()
func (_Zkbas *ZkbasTransactorSession) RequestFullExitNft(_accountName string, _nftIndex uint32) (*types.Transaction, error) {
	return _Zkbas.Contract.RequestFullExitNft(&_Zkbas.TransactOpts, _accountName, _nftIndex)
}

// RevertBlocks is a paid mutator transaction binding the contract method 0x97a2eabe.
//
// Solidity: function revertBlocks((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32)[] _blocksToRevert) returns()
func (_Zkbas *ZkbasTransactor) RevertBlocks(opts *bind.TransactOpts, _blocksToRevert []StorageStoredBlockInfo) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "revertBlocks", _blocksToRevert)
}

// RevertBlocks is a paid mutator transaction binding the contract method 0x97a2eabe.
//
// Solidity: function revertBlocks((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32)[] _blocksToRevert) returns()
func (_Zkbas *ZkbasSession) RevertBlocks(_blocksToRevert []StorageStoredBlockInfo) (*types.Transaction, error) {
	return _Zkbas.Contract.RevertBlocks(&_Zkbas.TransactOpts, _blocksToRevert)
}

// RevertBlocks is a paid mutator transaction binding the contract method 0x97a2eabe.
//
// Solidity: function revertBlocks((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32)[] _blocksToRevert) returns()
func (_Zkbas *ZkbasTransactorSession) RevertBlocks(_blocksToRevert []StorageStoredBlockInfo) (*types.Transaction, error) {
	return _Zkbas.Contract.RevertBlocks(&_Zkbas.TransactOpts, _blocksToRevert)
}

// SetDefaultNFTFactory is a paid mutator transaction binding the contract method 0xce09e20d.
//
// Solidity: function setDefaultNFTFactory(address _factory) returns()
func (_Zkbas *ZkbasTransactor) SetDefaultNFTFactory(opts *bind.TransactOpts, _factory common.Address) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "setDefaultNFTFactory", _factory)
}

// SetDefaultNFTFactory is a paid mutator transaction binding the contract method 0xce09e20d.
//
// Solidity: function setDefaultNFTFactory(address _factory) returns()
func (_Zkbas *ZkbasSession) SetDefaultNFTFactory(_factory common.Address) (*types.Transaction, error) {
	return _Zkbas.Contract.SetDefaultNFTFactory(&_Zkbas.TransactOpts, _factory)
}

// SetDefaultNFTFactory is a paid mutator transaction binding the contract method 0xce09e20d.
//
// Solidity: function setDefaultNFTFactory(address _factory) returns()
func (_Zkbas *ZkbasTransactorSession) SetDefaultNFTFactory(_factory common.Address) (*types.Transaction, error) {
	return _Zkbas.Contract.SetDefaultNFTFactory(&_Zkbas.TransactOpts, _factory)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x68809a23.
//
// Solidity: function transferERC20(address _token, address _to, uint128 _amount, uint128 _maxAmount) returns(uint128 withdrawnAmount)
func (_Zkbas *ZkbasTransactor) TransferERC20(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _maxAmount *big.Int) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "transferERC20", _token, _to, _amount, _maxAmount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x68809a23.
//
// Solidity: function transferERC20(address _token, address _to, uint128 _amount, uint128 _maxAmount) returns(uint128 withdrawnAmount)
func (_Zkbas *ZkbasSession) TransferERC20(_token common.Address, _to common.Address, _amount *big.Int, _maxAmount *big.Int) (*types.Transaction, error) {
	return _Zkbas.Contract.TransferERC20(&_Zkbas.TransactOpts, _token, _to, _amount, _maxAmount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x68809a23.
//
// Solidity: function transferERC20(address _token, address _to, uint128 _amount, uint128 _maxAmount) returns(uint128 withdrawnAmount)
func (_Zkbas *ZkbasTransactorSession) TransferERC20(_token common.Address, _to common.Address, _amount *big.Int, _maxAmount *big.Int) (*types.Transaction, error) {
	return _Zkbas.Contract.TransferERC20(&_Zkbas.TransactOpts, _token, _to, _amount, _maxAmount)
}

// UpdatePairRate is a paid mutator transaction binding the contract method 0x13a05e23.
//
// Solidity: function updatePairRate((address,address,uint16,uint32,uint16) _pairInfo) returns()
func (_Zkbas *ZkbasTransactor) UpdatePairRate(opts *bind.TransactOpts, _pairInfo OldZkbasPairInfo) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "updatePairRate", _pairInfo)
}

// UpdatePairRate is a paid mutator transaction binding the contract method 0x13a05e23.
//
// Solidity: function updatePairRate((address,address,uint16,uint32,uint16) _pairInfo) returns()
func (_Zkbas *ZkbasSession) UpdatePairRate(_pairInfo OldZkbasPairInfo) (*types.Transaction, error) {
	return _Zkbas.Contract.UpdatePairRate(&_Zkbas.TransactOpts, _pairInfo)
}

// UpdatePairRate is a paid mutator transaction binding the contract method 0x13a05e23.
//
// Solidity: function updatePairRate((address,address,uint16,uint32,uint16) _pairInfo) returns()
func (_Zkbas *ZkbasTransactorSession) UpdatePairRate(_pairInfo OldZkbasPairInfo) (*types.Transaction, error) {
	return _Zkbas.Contract.UpdatePairRate(&_Zkbas.TransactOpts, _pairInfo)
}

// UpdateZkbasVerifier is a paid mutator transaction binding the contract method 0x6f88b156.
//
// Solidity: function updateZkbasVerifier(address _newVerifierAddress) returns()
func (_Zkbas *ZkbasTransactor) UpdateZkbasVerifier(opts *bind.TransactOpts, _newVerifierAddress common.Address) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "updateZkbasVerifier", _newVerifierAddress)
}

// UpdateZkbasVerifier is a paid mutator transaction binding the contract method 0x6f88b156.
//
// Solidity: function updateZkbasVerifier(address _newVerifierAddress) returns()
func (_Zkbas *ZkbasSession) UpdateZkbasVerifier(_newVerifierAddress common.Address) (*types.Transaction, error) {
	return _Zkbas.Contract.UpdateZkbasVerifier(&_Zkbas.TransactOpts, _newVerifierAddress)
}

// UpdateZkbasVerifier is a paid mutator transaction binding the contract method 0x6f88b156.
//
// Solidity: function updateZkbasVerifier(address _newVerifierAddress) returns()
func (_Zkbas *ZkbasTransactorSession) UpdateZkbasVerifier(_newVerifierAddress common.Address) (*types.Transaction, error) {
	return _Zkbas.Contract.UpdateZkbasVerifier(&_Zkbas.TransactOpts, _newVerifierAddress)
}

// Upgrade is a paid mutator transaction binding the contract method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_Zkbas *ZkbasTransactor) Upgrade(opts *bind.TransactOpts, upgradeParameters []byte) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "upgrade", upgradeParameters)
}

// Upgrade is a paid mutator transaction binding the contract method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_Zkbas *ZkbasSession) Upgrade(upgradeParameters []byte) (*types.Transaction, error) {
	return _Zkbas.Contract.Upgrade(&_Zkbas.TransactOpts, upgradeParameters)
}

// Upgrade is a paid mutator transaction binding the contract method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_Zkbas *ZkbasTransactorSession) Upgrade(upgradeParameters []byte) (*types.Transaction, error) {
	return _Zkbas.Contract.Upgrade(&_Zkbas.TransactOpts, upgradeParameters)
}

// UpgradeCanceled is a paid mutator transaction binding the contract method 0x871b8ff1.
//
// Solidity: function upgradeCanceled() returns()
func (_Zkbas *ZkbasTransactor) UpgradeCanceled(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "upgradeCanceled")
}

// UpgradeCanceled is a paid mutator transaction binding the contract method 0x871b8ff1.
//
// Solidity: function upgradeCanceled() returns()
func (_Zkbas *ZkbasSession) UpgradeCanceled() (*types.Transaction, error) {
	return _Zkbas.Contract.UpgradeCanceled(&_Zkbas.TransactOpts)
}

// UpgradeCanceled is a paid mutator transaction binding the contract method 0x871b8ff1.
//
// Solidity: function upgradeCanceled() returns()
func (_Zkbas *ZkbasTransactorSession) UpgradeCanceled() (*types.Transaction, error) {
	return _Zkbas.Contract.UpgradeCanceled(&_Zkbas.TransactOpts)
}

// UpgradeFinishes is a paid mutator transaction binding the contract method 0xb269b9ae.
//
// Solidity: function upgradeFinishes() returns()
func (_Zkbas *ZkbasTransactor) UpgradeFinishes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "upgradeFinishes")
}

// UpgradeFinishes is a paid mutator transaction binding the contract method 0xb269b9ae.
//
// Solidity: function upgradeFinishes() returns()
func (_Zkbas *ZkbasSession) UpgradeFinishes() (*types.Transaction, error) {
	return _Zkbas.Contract.UpgradeFinishes(&_Zkbas.TransactOpts)
}

// UpgradeFinishes is a paid mutator transaction binding the contract method 0xb269b9ae.
//
// Solidity: function upgradeFinishes() returns()
func (_Zkbas *ZkbasTransactorSession) UpgradeFinishes() (*types.Transaction, error) {
	return _Zkbas.Contract.UpgradeFinishes(&_Zkbas.TransactOpts)
}

// UpgradeNoticePeriodStarted is a paid mutator transaction binding the contract method 0x3b154b73.
//
// Solidity: function upgradeNoticePeriodStarted() returns()
func (_Zkbas *ZkbasTransactor) UpgradeNoticePeriodStarted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "upgradeNoticePeriodStarted")
}

// UpgradeNoticePeriodStarted is a paid mutator transaction binding the contract method 0x3b154b73.
//
// Solidity: function upgradeNoticePeriodStarted() returns()
func (_Zkbas *ZkbasSession) UpgradeNoticePeriodStarted() (*types.Transaction, error) {
	return _Zkbas.Contract.UpgradeNoticePeriodStarted(&_Zkbas.TransactOpts)
}

// UpgradeNoticePeriodStarted is a paid mutator transaction binding the contract method 0x3b154b73.
//
// Solidity: function upgradeNoticePeriodStarted() returns()
func (_Zkbas *ZkbasTransactorSession) UpgradeNoticePeriodStarted() (*types.Transaction, error) {
	return _Zkbas.Contract.UpgradeNoticePeriodStarted(&_Zkbas.TransactOpts)
}

// UpgradePreparationStarted is a paid mutator transaction binding the contract method 0x78b91e70.
//
// Solidity: function upgradePreparationStarted() returns()
func (_Zkbas *ZkbasTransactor) UpgradePreparationStarted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "upgradePreparationStarted")
}

// UpgradePreparationStarted is a paid mutator transaction binding the contract method 0x78b91e70.
//
// Solidity: function upgradePreparationStarted() returns()
func (_Zkbas *ZkbasSession) UpgradePreparationStarted() (*types.Transaction, error) {
	return _Zkbas.Contract.UpgradePreparationStarted(&_Zkbas.TransactOpts)
}

// UpgradePreparationStarted is a paid mutator transaction binding the contract method 0x78b91e70.
//
// Solidity: function upgradePreparationStarted() returns()
func (_Zkbas *ZkbasTransactorSession) UpgradePreparationStarted() (*types.Transaction, error) {
	return _Zkbas.Contract.UpgradePreparationStarted(&_Zkbas.TransactOpts)
}

// VerifyAndExecuteBlocks is a paid mutator transaction binding the contract method 0xf9ea2e71.
//
// Solidity: function verifyAndExecuteBlocks(((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32),bytes[])[] _blocks, uint256[] _proofs) returns()
func (_Zkbas *ZkbasTransactor) VerifyAndExecuteBlocks(opts *bind.TransactOpts, _blocks []OldZkbasVerifyAndExecuteBlockInfo, _proofs []*big.Int) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "verifyAndExecuteBlocks", _blocks, _proofs)
}

// VerifyAndExecuteBlocks is a paid mutator transaction binding the contract method 0xf9ea2e71.
//
// Solidity: function verifyAndExecuteBlocks(((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32),bytes[])[] _blocks, uint256[] _proofs) returns()
func (_Zkbas *ZkbasSession) VerifyAndExecuteBlocks(_blocks []OldZkbasVerifyAndExecuteBlockInfo, _proofs []*big.Int) (*types.Transaction, error) {
	return _Zkbas.Contract.VerifyAndExecuteBlocks(&_Zkbas.TransactOpts, _blocks, _proofs)
}

// VerifyAndExecuteBlocks is a paid mutator transaction binding the contract method 0xf9ea2e71.
//
// Solidity: function verifyAndExecuteBlocks(((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32),bytes[])[] _blocks, uint256[] _proofs) returns()
func (_Zkbas *ZkbasTransactorSession) VerifyAndExecuteBlocks(_blocks []OldZkbasVerifyAndExecuteBlockInfo, _proofs []*big.Int) (*types.Transaction, error) {
	return _Zkbas.Contract.VerifyAndExecuteBlocks(&_Zkbas.TransactOpts, _blocks, _proofs)
}

// WithdrawPendingBalance is a paid mutator transaction binding the contract method 0xd514da50.
//
// Solidity: function withdrawPendingBalance(address _owner, address _token, uint128 _amount) returns()
func (_Zkbas *ZkbasTransactor) WithdrawPendingBalance(opts *bind.TransactOpts, _owner common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "withdrawPendingBalance", _owner, _token, _amount)
}

// WithdrawPendingBalance is a paid mutator transaction binding the contract method 0xd514da50.
//
// Solidity: function withdrawPendingBalance(address _owner, address _token, uint128 _amount) returns()
func (_Zkbas *ZkbasSession) WithdrawPendingBalance(_owner common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zkbas.Contract.WithdrawPendingBalance(&_Zkbas.TransactOpts, _owner, _token, _amount)
}

// WithdrawPendingBalance is a paid mutator transaction binding the contract method 0xd514da50.
//
// Solidity: function withdrawPendingBalance(address _owner, address _token, uint128 _amount) returns()
func (_Zkbas *ZkbasTransactorSession) WithdrawPendingBalance(_owner common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zkbas.Contract.WithdrawPendingBalance(&_Zkbas.TransactOpts, _owner, _token, _amount)
}

// WithdrawPendingNFTBalance is a paid mutator transaction binding the contract method 0x7ce1017d.
//
// Solidity: function withdrawPendingNFTBalance(uint40 _nftIndex) returns()
func (_Zkbas *ZkbasTransactor) WithdrawPendingNFTBalance(opts *bind.TransactOpts, _nftIndex *big.Int) (*types.Transaction, error) {
	return _Zkbas.contract.Transact(opts, "withdrawPendingNFTBalance", _nftIndex)
}

// WithdrawPendingNFTBalance is a paid mutator transaction binding the contract method 0x7ce1017d.
//
// Solidity: function withdrawPendingNFTBalance(uint40 _nftIndex) returns()
func (_Zkbas *ZkbasSession) WithdrawPendingNFTBalance(_nftIndex *big.Int) (*types.Transaction, error) {
	return _Zkbas.Contract.WithdrawPendingNFTBalance(&_Zkbas.TransactOpts, _nftIndex)
}

// WithdrawPendingNFTBalance is a paid mutator transaction binding the contract method 0x7ce1017d.
//
// Solidity: function withdrawPendingNFTBalance(uint40 _nftIndex) returns()
func (_Zkbas *ZkbasTransactorSession) WithdrawPendingNFTBalance(_nftIndex *big.Int) (*types.Transaction, error) {
	return _Zkbas.Contract.WithdrawPendingNFTBalance(&_Zkbas.TransactOpts, _nftIndex)
}

// ZkbasBlockCommitIterator is returned from FilterBlockCommit and is used to iterate over the raw logs and unpacked data for BlockCommit events raised by the Zkbas contract.
type ZkbasBlockCommitIterator struct {
	Event *ZkbasBlockCommit // Event containing the contract specifics and raw log

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
func (it *ZkbasBlockCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkbasBlockCommit)
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
		it.Event = new(ZkbasBlockCommit)
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
func (it *ZkbasBlockCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkbasBlockCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkbasBlockCommit represents a BlockCommit event raised by the Zkbas contract.
type ZkbasBlockCommit struct {
	BlockNumber uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockCommit is a free log retrieval operation binding the contract event 0x81a92942d0f9c33b897a438384c9c3d88be397776138efa3ba1a4fc8b6268424.
//
// Solidity: event BlockCommit(uint32 blockNumber)
func (_Zkbas *ZkbasFilterer) FilterBlockCommit(opts *bind.FilterOpts) (*ZkbasBlockCommitIterator, error) {

	logs, sub, err := _Zkbas.contract.FilterLogs(opts, "BlockCommit")
	if err != nil {
		return nil, err
	}
	return &ZkbasBlockCommitIterator{contract: _Zkbas.contract, event: "BlockCommit", logs: logs, sub: sub}, nil
}

// WatchBlockCommit is a free log subscription operation binding the contract event 0x81a92942d0f9c33b897a438384c9c3d88be397776138efa3ba1a4fc8b6268424.
//
// Solidity: event BlockCommit(uint32 blockNumber)
func (_Zkbas *ZkbasFilterer) WatchBlockCommit(opts *bind.WatchOpts, sink chan<- *ZkbasBlockCommit) (event.Subscription, error) {

	logs, sub, err := _Zkbas.contract.WatchLogs(opts, "BlockCommit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkbasBlockCommit)
				if err := _Zkbas.contract.UnpackLog(event, "BlockCommit", log); err != nil {
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
func (_Zkbas *ZkbasFilterer) ParseBlockCommit(log types.Log) (*ZkbasBlockCommit, error) {
	event := new(ZkbasBlockCommit)
	if err := _Zkbas.contract.UnpackLog(event, "BlockCommit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkbasBlockVerificationIterator is returned from FilterBlockVerification and is used to iterate over the raw logs and unpacked data for BlockVerification events raised by the Zkbas contract.
type ZkbasBlockVerificationIterator struct {
	Event *ZkbasBlockVerification // Event containing the contract specifics and raw log

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
func (it *ZkbasBlockVerificationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkbasBlockVerification)
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
		it.Event = new(ZkbasBlockVerification)
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
func (it *ZkbasBlockVerificationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkbasBlockVerificationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkbasBlockVerification represents a BlockVerification event raised by the Zkbas contract.
type ZkbasBlockVerification struct {
	BlockNumber uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockVerification is a free log retrieval operation binding the contract event 0x0cdbd8bd7813095001c5fe7917bd69d834dc01db7c1dfcf52ca135bd20384413.
//
// Solidity: event BlockVerification(uint32 blockNumber)
func (_Zkbas *ZkbasFilterer) FilterBlockVerification(opts *bind.FilterOpts) (*ZkbasBlockVerificationIterator, error) {

	logs, sub, err := _Zkbas.contract.FilterLogs(opts, "BlockVerification")
	if err != nil {
		return nil, err
	}
	return &ZkbasBlockVerificationIterator{contract: _Zkbas.contract, event: "BlockVerification", logs: logs, sub: sub}, nil
}

// WatchBlockVerification is a free log subscription operation binding the contract event 0x0cdbd8bd7813095001c5fe7917bd69d834dc01db7c1dfcf52ca135bd20384413.
//
// Solidity: event BlockVerification(uint32 blockNumber)
func (_Zkbas *ZkbasFilterer) WatchBlockVerification(opts *bind.WatchOpts, sink chan<- *ZkbasBlockVerification) (event.Subscription, error) {

	logs, sub, err := _Zkbas.contract.WatchLogs(opts, "BlockVerification")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkbasBlockVerification)
				if err := _Zkbas.contract.UnpackLog(event, "BlockVerification", log); err != nil {
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
func (_Zkbas *ZkbasFilterer) ParseBlockVerification(log types.Log) (*ZkbasBlockVerification, error) {
	event := new(ZkbasBlockVerification)
	if err := _Zkbas.contract.UnpackLog(event, "BlockVerification", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkbasBlocksRevertIterator is returned from FilterBlocksRevert and is used to iterate over the raw logs and unpacked data for BlocksRevert events raised by the Zkbas contract.
type ZkbasBlocksRevertIterator struct {
	Event *ZkbasBlocksRevert // Event containing the contract specifics and raw log

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
func (it *ZkbasBlocksRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkbasBlocksRevert)
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
		it.Event = new(ZkbasBlocksRevert)
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
func (it *ZkbasBlocksRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkbasBlocksRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkbasBlocksRevert represents a BlocksRevert event raised by the Zkbas contract.
type ZkbasBlocksRevert struct {
	TotalBlocksVerified  uint32
	TotalBlocksCommitted uint32
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterBlocksRevert is a free log retrieval operation binding the contract event 0x6f3a8259cce1ea2680115053d21c971aa1764295a45850f520525f2bfdf3c9d3.
//
// Solidity: event BlocksRevert(uint32 totalBlocksVerified, uint32 totalBlocksCommitted)
func (_Zkbas *ZkbasFilterer) FilterBlocksRevert(opts *bind.FilterOpts) (*ZkbasBlocksRevertIterator, error) {

	logs, sub, err := _Zkbas.contract.FilterLogs(opts, "BlocksRevert")
	if err != nil {
		return nil, err
	}
	return &ZkbasBlocksRevertIterator{contract: _Zkbas.contract, event: "BlocksRevert", logs: logs, sub: sub}, nil
}

// WatchBlocksRevert is a free log subscription operation binding the contract event 0x6f3a8259cce1ea2680115053d21c971aa1764295a45850f520525f2bfdf3c9d3.
//
// Solidity: event BlocksRevert(uint32 totalBlocksVerified, uint32 totalBlocksCommitted)
func (_Zkbas *ZkbasFilterer) WatchBlocksRevert(opts *bind.WatchOpts, sink chan<- *ZkbasBlocksRevert) (event.Subscription, error) {

	logs, sub, err := _Zkbas.contract.WatchLogs(opts, "BlocksRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkbasBlocksRevert)
				if err := _Zkbas.contract.UnpackLog(event, "BlocksRevert", log); err != nil {
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
func (_Zkbas *ZkbasFilterer) ParseBlocksRevert(log types.Log) (*ZkbasBlocksRevert, error) {
	event := new(ZkbasBlocksRevert)
	if err := _Zkbas.contract.UnpackLog(event, "BlocksRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkbasCreateTokenPairIterator is returned from FilterCreateTokenPair and is used to iterate over the raw logs and unpacked data for CreateTokenPair events raised by the Zkbas contract.
type ZkbasCreateTokenPairIterator struct {
	Event *ZkbasCreateTokenPair // Event containing the contract specifics and raw log

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
func (it *ZkbasCreateTokenPairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkbasCreateTokenPair)
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
		it.Event = new(ZkbasCreateTokenPair)
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
func (it *ZkbasCreateTokenPairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkbasCreateTokenPairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkbasCreateTokenPair represents a CreateTokenPair event raised by the Zkbas contract.
type ZkbasCreateTokenPair struct {
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
func (_Zkbas *ZkbasFilterer) FilterCreateTokenPair(opts *bind.FilterOpts) (*ZkbasCreateTokenPairIterator, error) {

	logs, sub, err := _Zkbas.contract.FilterLogs(opts, "CreateTokenPair")
	if err != nil {
		return nil, err
	}
	return &ZkbasCreateTokenPairIterator{contract: _Zkbas.contract, event: "CreateTokenPair", logs: logs, sub: sub}, nil
}

// WatchCreateTokenPair is a free log subscription operation binding the contract event 0x82e21fd4ced46b510ee1bf5a34eea3dc34aa1f5460f53a25dd22828cde30e087.
//
// Solidity: event CreateTokenPair(uint16 pairIndex, uint16 asset0Id, uint16 asset1Id, uint16 feeRate, uint32 treasuryAccountIndex, uint16 treasuryRate)
func (_Zkbas *ZkbasFilterer) WatchCreateTokenPair(opts *bind.WatchOpts, sink chan<- *ZkbasCreateTokenPair) (event.Subscription, error) {

	logs, sub, err := _Zkbas.contract.WatchLogs(opts, "CreateTokenPair")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkbasCreateTokenPair)
				if err := _Zkbas.contract.UnpackLog(event, "CreateTokenPair", log); err != nil {
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
func (_Zkbas *ZkbasFilterer) ParseCreateTokenPair(log types.Log) (*ZkbasCreateTokenPair, error) {
	event := new(ZkbasCreateTokenPair)
	if err := _Zkbas.contract.UnpackLog(event, "CreateTokenPair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkbasDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Zkbas contract.
type ZkbasDepositIterator struct {
	Event *ZkbasDeposit // Event containing the contract specifics and raw log

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
func (it *ZkbasDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkbasDeposit)
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
		it.Event = new(ZkbasDeposit)
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
func (it *ZkbasDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkbasDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkbasDeposit represents a Deposit event raised by the Zkbas contract.
type ZkbasDeposit struct {
	AssetId     uint16
	AccountName [32]byte
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xaa46d46658f805449ac7eaaf11d3eaebb6f508930128d276dfefcc8dc02a13c7.
//
// Solidity: event Deposit(uint16 assetId, bytes32 accountName, uint128 amount)
func (_Zkbas *ZkbasFilterer) FilterDeposit(opts *bind.FilterOpts) (*ZkbasDepositIterator, error) {

	logs, sub, err := _Zkbas.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &ZkbasDepositIterator{contract: _Zkbas.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xaa46d46658f805449ac7eaaf11d3eaebb6f508930128d276dfefcc8dc02a13c7.
//
// Solidity: event Deposit(uint16 assetId, bytes32 accountName, uint128 amount)
func (_Zkbas *ZkbasFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ZkbasDeposit) (event.Subscription, error) {

	logs, sub, err := _Zkbas.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkbasDeposit)
				if err := _Zkbas.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_Zkbas *ZkbasFilterer) ParseDeposit(log types.Log) (*ZkbasDeposit, error) {
	event := new(ZkbasDeposit)
	if err := _Zkbas.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkbasDepositCommitIterator is returned from FilterDepositCommit and is used to iterate over the raw logs and unpacked data for DepositCommit events raised by the Zkbas contract.
type ZkbasDepositCommitIterator struct {
	Event *ZkbasDepositCommit // Event containing the contract specifics and raw log

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
func (it *ZkbasDepositCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkbasDepositCommit)
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
		it.Event = new(ZkbasDepositCommit)
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
func (it *ZkbasDepositCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkbasDepositCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkbasDepositCommit represents a DepositCommit event raised by the Zkbas contract.
type ZkbasDepositCommit struct {
	ZkbasBlockNumber uint32
	AccountIndex     uint32
	AccountName      [32]byte
	AssetId          uint16
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDepositCommit is a free log retrieval operation binding the contract event 0x9b3e77f05cee61edb7f3e0e85d71a29e89cb18cb4236a8b8cc3e4d66640d8f71.
//
// Solidity: event DepositCommit(uint32 indexed zkbasBlockNumber, uint32 indexed accountIndex, bytes32 accountName, uint16 indexed assetId, uint128 amount)
func (_Zkbas *ZkbasFilterer) FilterDepositCommit(opts *bind.FilterOpts, zkbasBlockNumber []uint32, accountIndex []uint32, assetId []uint16) (*ZkbasDepositCommitIterator, error) {

	var zkbasBlockNumberRule []interface{}
	for _, zkbasBlockNumberItem := range zkbasBlockNumber {
		zkbasBlockNumberRule = append(zkbasBlockNumberRule, zkbasBlockNumberItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _Zkbas.contract.FilterLogs(opts, "DepositCommit", zkbasBlockNumberRule, accountIndexRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &ZkbasDepositCommitIterator{contract: _Zkbas.contract, event: "DepositCommit", logs: logs, sub: sub}, nil
}

// WatchDepositCommit is a free log subscription operation binding the contract event 0x9b3e77f05cee61edb7f3e0e85d71a29e89cb18cb4236a8b8cc3e4d66640d8f71.
//
// Solidity: event DepositCommit(uint32 indexed zkbasBlockNumber, uint32 indexed accountIndex, bytes32 accountName, uint16 indexed assetId, uint128 amount)
func (_Zkbas *ZkbasFilterer) WatchDepositCommit(opts *bind.WatchOpts, sink chan<- *ZkbasDepositCommit, zkbasBlockNumber []uint32, accountIndex []uint32, assetId []uint16) (event.Subscription, error) {

	var zkbasBlockNumberRule []interface{}
	for _, zkbasBlockNumberItem := range zkbasBlockNumber {
		zkbasBlockNumberRule = append(zkbasBlockNumberRule, zkbasBlockNumberItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _Zkbas.contract.WatchLogs(opts, "DepositCommit", zkbasBlockNumberRule, accountIndexRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkbasDepositCommit)
				if err := _Zkbas.contract.UnpackLog(event, "DepositCommit", log); err != nil {
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
// Solidity: event DepositCommit(uint32 indexed zkbasBlockNumber, uint32 indexed accountIndex, bytes32 accountName, uint16 indexed assetId, uint128 amount)
func (_Zkbas *ZkbasFilterer) ParseDepositCommit(log types.Log) (*ZkbasDepositCommit, error) {
	event := new(ZkbasDepositCommit)
	if err := _Zkbas.contract.UnpackLog(event, "DepositCommit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkbasDepositNftIterator is returned from FilterDepositNft and is used to iterate over the raw logs and unpacked data for DepositNft events raised by the Zkbas contract.
type ZkbasDepositNftIterator struct {
	Event *ZkbasDepositNft // Event containing the contract specifics and raw log

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
func (it *ZkbasDepositNftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkbasDepositNft)
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
		it.Event = new(ZkbasDepositNft)
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
func (it *ZkbasDepositNftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkbasDepositNftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkbasDepositNft represents a DepositNft event raised by the Zkbas contract.
type ZkbasDepositNft struct {
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
func (_Zkbas *ZkbasFilterer) FilterDepositNft(opts *bind.FilterOpts) (*ZkbasDepositNftIterator, error) {

	logs, sub, err := _Zkbas.contract.FilterLogs(opts, "DepositNft")
	if err != nil {
		return nil, err
	}
	return &ZkbasDepositNftIterator{contract: _Zkbas.contract, event: "DepositNft", logs: logs, sub: sub}, nil
}

// WatchDepositNft is a free log subscription operation binding the contract event 0x42f9fbd77faf066fde386943aaafcc841fba6b755dc5da7939a046cbc493c24f.
//
// Solidity: event DepositNft(bytes32 accountNameHash, bytes32 nftContentHash, address tokenAddress, uint256 nftTokenId, uint16 creatorTreasuryRate)
func (_Zkbas *ZkbasFilterer) WatchDepositNft(opts *bind.WatchOpts, sink chan<- *ZkbasDepositNft) (event.Subscription, error) {

	logs, sub, err := _Zkbas.contract.WatchLogs(opts, "DepositNft")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkbasDepositNft)
				if err := _Zkbas.contract.UnpackLog(event, "DepositNft", log); err != nil {
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
func (_Zkbas *ZkbasFilterer) ParseDepositNft(log types.Log) (*ZkbasDepositNft, error) {
	event := new(ZkbasDepositNft)
	if err := _Zkbas.contract.UnpackLog(event, "DepositNft", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkbasDesertModeIterator is returned from FilterDesertMode and is used to iterate over the raw logs and unpacked data for DesertMode events raised by the Zkbas contract.
type ZkbasDesertModeIterator struct {
	Event *ZkbasDesertMode // Event containing the contract specifics and raw log

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
func (it *ZkbasDesertModeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkbasDesertMode)
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
		it.Event = new(ZkbasDesertMode)
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
func (it *ZkbasDesertModeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkbasDesertModeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkbasDesertMode represents a DesertMode event raised by the Zkbas contract.
type ZkbasDesertMode struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDesertMode is a free log retrieval operation binding the contract event 0x9f7e400a81dddbf1c18b1c37f82aa303d166295ca4b577eb2a7c23d4b704ba89.
//
// Solidity: event DesertMode()
func (_Zkbas *ZkbasFilterer) FilterDesertMode(opts *bind.FilterOpts) (*ZkbasDesertModeIterator, error) {

	logs, sub, err := _Zkbas.contract.FilterLogs(opts, "DesertMode")
	if err != nil {
		return nil, err
	}
	return &ZkbasDesertModeIterator{contract: _Zkbas.contract, event: "DesertMode", logs: logs, sub: sub}, nil
}

// WatchDesertMode is a free log subscription operation binding the contract event 0x9f7e400a81dddbf1c18b1c37f82aa303d166295ca4b577eb2a7c23d4b704ba89.
//
// Solidity: event DesertMode()
func (_Zkbas *ZkbasFilterer) WatchDesertMode(opts *bind.WatchOpts, sink chan<- *ZkbasDesertMode) (event.Subscription, error) {

	logs, sub, err := _Zkbas.contract.WatchLogs(opts, "DesertMode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkbasDesertMode)
				if err := _Zkbas.contract.UnpackLog(event, "DesertMode", log); err != nil {
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
func (_Zkbas *ZkbasFilterer) ParseDesertMode(log types.Log) (*ZkbasDesertMode, error) {
	event := new(ZkbasDesertMode)
	if err := _Zkbas.contract.UnpackLog(event, "DesertMode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkbasFullExitCommitIterator is returned from FilterFullExitCommit and is used to iterate over the raw logs and unpacked data for FullExitCommit events raised by the Zkbas contract.
type ZkbasFullExitCommitIterator struct {
	Event *ZkbasFullExitCommit // Event containing the contract specifics and raw log

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
func (it *ZkbasFullExitCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkbasFullExitCommit)
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
		it.Event = new(ZkbasFullExitCommit)
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
func (it *ZkbasFullExitCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkbasFullExitCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkbasFullExitCommit represents a FullExitCommit event raised by the Zkbas contract.
type ZkbasFullExitCommit struct {
	ZkbasBlockId uint32
	AccountId    uint32
	Owner        common.Address
	TokenId      uint16
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFullExitCommit is a free log retrieval operation binding the contract event 0x66fc63d751ecbefca61d4e2e7c534e4f29c61aed8ece23ed635277a7ea6f9bc4.
//
// Solidity: event FullExitCommit(uint32 indexed zkbasBlockId, uint32 indexed accountId, address owner, uint16 indexed tokenId, uint128 amount)
func (_Zkbas *ZkbasFilterer) FilterFullExitCommit(opts *bind.FilterOpts, zkbasBlockId []uint32, accountId []uint32, tokenId []uint16) (*ZkbasFullExitCommitIterator, error) {

	var zkbasBlockIdRule []interface{}
	for _, zkbasBlockIdItem := range zkbasBlockId {
		zkbasBlockIdRule = append(zkbasBlockIdRule, zkbasBlockIdItem)
	}
	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Zkbas.contract.FilterLogs(opts, "FullExitCommit", zkbasBlockIdRule, accountIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ZkbasFullExitCommitIterator{contract: _Zkbas.contract, event: "FullExitCommit", logs: logs, sub: sub}, nil
}

// WatchFullExitCommit is a free log subscription operation binding the contract event 0x66fc63d751ecbefca61d4e2e7c534e4f29c61aed8ece23ed635277a7ea6f9bc4.
//
// Solidity: event FullExitCommit(uint32 indexed zkbasBlockId, uint32 indexed accountId, address owner, uint16 indexed tokenId, uint128 amount)
func (_Zkbas *ZkbasFilterer) WatchFullExitCommit(opts *bind.WatchOpts, sink chan<- *ZkbasFullExitCommit, zkbasBlockId []uint32, accountId []uint32, tokenId []uint16) (event.Subscription, error) {

	var zkbasBlockIdRule []interface{}
	for _, zkbasBlockIdItem := range zkbasBlockId {
		zkbasBlockIdRule = append(zkbasBlockIdRule, zkbasBlockIdItem)
	}
	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Zkbas.contract.WatchLogs(opts, "FullExitCommit", zkbasBlockIdRule, accountIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkbasFullExitCommit)
				if err := _Zkbas.contract.UnpackLog(event, "FullExitCommit", log); err != nil {
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
// Solidity: event FullExitCommit(uint32 indexed zkbasBlockId, uint32 indexed accountId, address owner, uint16 indexed tokenId, uint128 amount)
func (_Zkbas *ZkbasFilterer) ParseFullExitCommit(log types.Log) (*ZkbasFullExitCommit, error) {
	event := new(ZkbasFullExitCommit)
	if err := _Zkbas.contract.UnpackLog(event, "FullExitCommit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkbasNewDefaultNFTFactoryIterator is returned from FilterNewDefaultNFTFactory and is used to iterate over the raw logs and unpacked data for NewDefaultNFTFactory events raised by the Zkbas contract.
type ZkbasNewDefaultNFTFactoryIterator struct {
	Event *ZkbasNewDefaultNFTFactory // Event containing the contract specifics and raw log

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
func (it *ZkbasNewDefaultNFTFactoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkbasNewDefaultNFTFactory)
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
		it.Event = new(ZkbasNewDefaultNFTFactory)
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
func (it *ZkbasNewDefaultNFTFactoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkbasNewDefaultNFTFactoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkbasNewDefaultNFTFactory represents a NewDefaultNFTFactory event raised by the Zkbas contract.
type ZkbasNewDefaultNFTFactory struct {
	Factory common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewDefaultNFTFactory is a free log retrieval operation binding the contract event 0xadf2ec6c092b7101c958de6b55caf4818f26a946e94fa556cbe3b3f29ff9f2ee.
//
// Solidity: event NewDefaultNFTFactory(address indexed factory)
func (_Zkbas *ZkbasFilterer) FilterNewDefaultNFTFactory(opts *bind.FilterOpts, factory []common.Address) (*ZkbasNewDefaultNFTFactoryIterator, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _Zkbas.contract.FilterLogs(opts, "NewDefaultNFTFactory", factoryRule)
	if err != nil {
		return nil, err
	}
	return &ZkbasNewDefaultNFTFactoryIterator{contract: _Zkbas.contract, event: "NewDefaultNFTFactory", logs: logs, sub: sub}, nil
}

// WatchNewDefaultNFTFactory is a free log subscription operation binding the contract event 0xadf2ec6c092b7101c958de6b55caf4818f26a946e94fa556cbe3b3f29ff9f2ee.
//
// Solidity: event NewDefaultNFTFactory(address indexed factory)
func (_Zkbas *ZkbasFilterer) WatchNewDefaultNFTFactory(opts *bind.WatchOpts, sink chan<- *ZkbasNewDefaultNFTFactory, factory []common.Address) (event.Subscription, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _Zkbas.contract.WatchLogs(opts, "NewDefaultNFTFactory", factoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkbasNewDefaultNFTFactory)
				if err := _Zkbas.contract.UnpackLog(event, "NewDefaultNFTFactory", log); err != nil {
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
func (_Zkbas *ZkbasFilterer) ParseNewDefaultNFTFactory(log types.Log) (*ZkbasNewDefaultNFTFactory, error) {
	event := new(ZkbasNewDefaultNFTFactory)
	if err := _Zkbas.contract.UnpackLog(event, "NewDefaultNFTFactory", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkbasNewNFTFactoryIterator is returned from FilterNewNFTFactory and is used to iterate over the raw logs and unpacked data for NewNFTFactory events raised by the Zkbas contract.
type ZkbasNewNFTFactoryIterator struct {
	Event *ZkbasNewNFTFactory // Event containing the contract specifics and raw log

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
func (it *ZkbasNewNFTFactoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkbasNewNFTFactory)
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
		it.Event = new(ZkbasNewNFTFactory)
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
func (it *ZkbasNewNFTFactoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkbasNewNFTFactoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkbasNewNFTFactory represents a NewNFTFactory event raised by the Zkbas contract.
type ZkbasNewNFTFactory struct {
	CreatorAccountNameHash [32]byte
	CollectionId           uint32
	FactoryAddress         common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterNewNFTFactory is a free log retrieval operation binding the contract event 0xaa31b08b26b4240f7d2837e3fb6359824b509a4e4c5a5766af80f63c319add7e.
//
// Solidity: event NewNFTFactory(bytes32 indexed _creatorAccountNameHash, uint32 _collectionId, address _factoryAddress)
func (_Zkbas *ZkbasFilterer) FilterNewNFTFactory(opts *bind.FilterOpts, _creatorAccountNameHash [][32]byte) (*ZkbasNewNFTFactoryIterator, error) {

	var _creatorAccountNameHashRule []interface{}
	for _, _creatorAccountNameHashItem := range _creatorAccountNameHash {
		_creatorAccountNameHashRule = append(_creatorAccountNameHashRule, _creatorAccountNameHashItem)
	}

	logs, sub, err := _Zkbas.contract.FilterLogs(opts, "NewNFTFactory", _creatorAccountNameHashRule)
	if err != nil {
		return nil, err
	}
	return &ZkbasNewNFTFactoryIterator{contract: _Zkbas.contract, event: "NewNFTFactory", logs: logs, sub: sub}, nil
}

// WatchNewNFTFactory is a free log subscription operation binding the contract event 0xaa31b08b26b4240f7d2837e3fb6359824b509a4e4c5a5766af80f63c319add7e.
//
// Solidity: event NewNFTFactory(bytes32 indexed _creatorAccountNameHash, uint32 _collectionId, address _factoryAddress)
func (_Zkbas *ZkbasFilterer) WatchNewNFTFactory(opts *bind.WatchOpts, sink chan<- *ZkbasNewNFTFactory, _creatorAccountNameHash [][32]byte) (event.Subscription, error) {

	var _creatorAccountNameHashRule []interface{}
	for _, _creatorAccountNameHashItem := range _creatorAccountNameHash {
		_creatorAccountNameHashRule = append(_creatorAccountNameHashRule, _creatorAccountNameHashItem)
	}

	logs, sub, err := _Zkbas.contract.WatchLogs(opts, "NewNFTFactory", _creatorAccountNameHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkbasNewNFTFactory)
				if err := _Zkbas.contract.UnpackLog(event, "NewNFTFactory", log); err != nil {
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
func (_Zkbas *ZkbasFilterer) ParseNewNFTFactory(log types.Log) (*ZkbasNewNFTFactory, error) {
	event := new(ZkbasNewNFTFactory)
	if err := _Zkbas.contract.UnpackLog(event, "NewNFTFactory", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkbasNewPriorityRequestIterator is returned from FilterNewPriorityRequest and is used to iterate over the raw logs and unpacked data for NewPriorityRequest events raised by the Zkbas contract.
type ZkbasNewPriorityRequestIterator struct {
	Event *ZkbasNewPriorityRequest // Event containing the contract specifics and raw log

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
func (it *ZkbasNewPriorityRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkbasNewPriorityRequest)
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
		it.Event = new(ZkbasNewPriorityRequest)
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
func (it *ZkbasNewPriorityRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkbasNewPriorityRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkbasNewPriorityRequest represents a NewPriorityRequest event raised by the Zkbas contract.
type ZkbasNewPriorityRequest struct {
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
func (_Zkbas *ZkbasFilterer) FilterNewPriorityRequest(opts *bind.FilterOpts) (*ZkbasNewPriorityRequestIterator, error) {

	logs, sub, err := _Zkbas.contract.FilterLogs(opts, "NewPriorityRequest")
	if err != nil {
		return nil, err
	}
	return &ZkbasNewPriorityRequestIterator{contract: _Zkbas.contract, event: "NewPriorityRequest", logs: logs, sub: sub}, nil
}

// WatchNewPriorityRequest is a free log subscription operation binding the contract event 0xd0943372c08b438a88d4b39d77216901079eda9ca59d45349841c099083b6830.
//
// Solidity: event NewPriorityRequest(address sender, uint64 serialId, uint8 txType, bytes pubData, uint256 expirationBlock)
func (_Zkbas *ZkbasFilterer) WatchNewPriorityRequest(opts *bind.WatchOpts, sink chan<- *ZkbasNewPriorityRequest) (event.Subscription, error) {

	logs, sub, err := _Zkbas.contract.WatchLogs(opts, "NewPriorityRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkbasNewPriorityRequest)
				if err := _Zkbas.contract.UnpackLog(event, "NewPriorityRequest", log); err != nil {
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
func (_Zkbas *ZkbasFilterer) ParseNewPriorityRequest(log types.Log) (*ZkbasNewPriorityRequest, error) {
	event := new(ZkbasNewPriorityRequest)
	if err := _Zkbas.contract.UnpackLog(event, "NewPriorityRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkbasNoticePeriodChangeIterator is returned from FilterNoticePeriodChange and is used to iterate over the raw logs and unpacked data for NoticePeriodChange events raised by the Zkbas contract.
type ZkbasNoticePeriodChangeIterator struct {
	Event *ZkbasNoticePeriodChange // Event containing the contract specifics and raw log

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
func (it *ZkbasNoticePeriodChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkbasNoticePeriodChange)
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
		it.Event = new(ZkbasNoticePeriodChange)
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
func (it *ZkbasNoticePeriodChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkbasNoticePeriodChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkbasNoticePeriodChange represents a NoticePeriodChange event raised by the Zkbas contract.
type ZkbasNoticePeriodChange struct {
	NewNoticePeriod *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNoticePeriodChange is a free log retrieval operation binding the contract event 0xf2b18f8abbd8a0d0c1fb8245146eedf5304887b12f6395b548ca238e054a1483.
//
// Solidity: event NoticePeriodChange(uint256 newNoticePeriod)
func (_Zkbas *ZkbasFilterer) FilterNoticePeriodChange(opts *bind.FilterOpts) (*ZkbasNoticePeriodChangeIterator, error) {

	logs, sub, err := _Zkbas.contract.FilterLogs(opts, "NoticePeriodChange")
	if err != nil {
		return nil, err
	}
	return &ZkbasNoticePeriodChangeIterator{contract: _Zkbas.contract, event: "NoticePeriodChange", logs: logs, sub: sub}, nil
}

// WatchNoticePeriodChange is a free log subscription operation binding the contract event 0xf2b18f8abbd8a0d0c1fb8245146eedf5304887b12f6395b548ca238e054a1483.
//
// Solidity: event NoticePeriodChange(uint256 newNoticePeriod)
func (_Zkbas *ZkbasFilterer) WatchNoticePeriodChange(opts *bind.WatchOpts, sink chan<- *ZkbasNoticePeriodChange) (event.Subscription, error) {

	logs, sub, err := _Zkbas.contract.WatchLogs(opts, "NoticePeriodChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkbasNoticePeriodChange)
				if err := _Zkbas.contract.UnpackLog(event, "NoticePeriodChange", log); err != nil {
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
func (_Zkbas *ZkbasFilterer) ParseNoticePeriodChange(log types.Log) (*ZkbasNoticePeriodChange, error) {
	event := new(ZkbasNoticePeriodChange)
	if err := _Zkbas.contract.UnpackLog(event, "NoticePeriodChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkbasRegisterZNSIterator is returned from FilterRegisterZNS and is used to iterate over the raw logs and unpacked data for RegisterZNS events raised by the Zkbas contract.
type ZkbasRegisterZNSIterator struct {
	Event *ZkbasRegisterZNS // Event containing the contract specifics and raw log

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
func (it *ZkbasRegisterZNSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkbasRegisterZNS)
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
		it.Event = new(ZkbasRegisterZNS)
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
func (it *ZkbasRegisterZNSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkbasRegisterZNSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkbasRegisterZNS represents a RegisterZNS event raised by the Zkbas contract.
type ZkbasRegisterZNS struct {
	Name         string
	NameHash     [32]byte
	Owner        common.Address
	ZkbasPubKeyX [32]byte
	ZkbasPubKeyY [32]byte
	AccountIndex uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRegisterZNS is a free log retrieval operation binding the contract event 0xebe5b00e9e9052d45c93d72544961460aa0fa30e0142c07d7dd81025cf9dab6c.
//
// Solidity: event RegisterZNS(string name, bytes32 nameHash, address owner, bytes32 zkbasPubKeyX, bytes32 zkbasPubKeyY, uint32 accountIndex)
func (_Zkbas *ZkbasFilterer) FilterRegisterZNS(opts *bind.FilterOpts) (*ZkbasRegisterZNSIterator, error) {

	logs, sub, err := _Zkbas.contract.FilterLogs(opts, "RegisterZNS")
	if err != nil {
		return nil, err
	}
	return &ZkbasRegisterZNSIterator{contract: _Zkbas.contract, event: "RegisterZNS", logs: logs, sub: sub}, nil
}

// WatchRegisterZNS is a free log subscription operation binding the contract event 0xebe5b00e9e9052d45c93d72544961460aa0fa30e0142c07d7dd81025cf9dab6c.
//
// Solidity: event RegisterZNS(string name, bytes32 nameHash, address owner, bytes32 zkbasPubKeyX, bytes32 zkbasPubKeyY, uint32 accountIndex)
func (_Zkbas *ZkbasFilterer) WatchRegisterZNS(opts *bind.WatchOpts, sink chan<- *ZkbasRegisterZNS) (event.Subscription, error) {

	logs, sub, err := _Zkbas.contract.WatchLogs(opts, "RegisterZNS")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkbasRegisterZNS)
				if err := _Zkbas.contract.UnpackLog(event, "RegisterZNS", log); err != nil {
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
// Solidity: event RegisterZNS(string name, bytes32 nameHash, address owner, bytes32 zkbasPubKeyX, bytes32 zkbasPubKeyY, uint32 accountIndex)
func (_Zkbas *ZkbasFilterer) ParseRegisterZNS(log types.Log) (*ZkbasRegisterZNS, error) {
	event := new(ZkbasRegisterZNS)
	if err := _Zkbas.contract.UnpackLog(event, "RegisterZNS", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkbasUpdateTokenPairIterator is returned from FilterUpdateTokenPair and is used to iterate over the raw logs and unpacked data for UpdateTokenPair events raised by the Zkbas contract.
type ZkbasUpdateTokenPairIterator struct {
	Event *ZkbasUpdateTokenPair // Event containing the contract specifics and raw log

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
func (it *ZkbasUpdateTokenPairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkbasUpdateTokenPair)
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
		it.Event = new(ZkbasUpdateTokenPair)
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
func (it *ZkbasUpdateTokenPairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkbasUpdateTokenPairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkbasUpdateTokenPair represents a UpdateTokenPair event raised by the Zkbas contract.
type ZkbasUpdateTokenPair struct {
	PairIndex            uint16
	FeeRate              uint16
	TreasuryAccountIndex uint32
	TreasuryRate         uint16
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateTokenPair is a free log retrieval operation binding the contract event 0x8727712867f019c156b0c25f308498ea1524b635902129b0acd3e9ff2c96f1c8.
//
// Solidity: event UpdateTokenPair(uint16 pairIndex, uint16 feeRate, uint32 treasuryAccountIndex, uint16 treasuryRate)
func (_Zkbas *ZkbasFilterer) FilterUpdateTokenPair(opts *bind.FilterOpts) (*ZkbasUpdateTokenPairIterator, error) {

	logs, sub, err := _Zkbas.contract.FilterLogs(opts, "UpdateTokenPair")
	if err != nil {
		return nil, err
	}
	return &ZkbasUpdateTokenPairIterator{contract: _Zkbas.contract, event: "UpdateTokenPair", logs: logs, sub: sub}, nil
}

// WatchUpdateTokenPair is a free log subscription operation binding the contract event 0x8727712867f019c156b0c25f308498ea1524b635902129b0acd3e9ff2c96f1c8.
//
// Solidity: event UpdateTokenPair(uint16 pairIndex, uint16 feeRate, uint32 treasuryAccountIndex, uint16 treasuryRate)
func (_Zkbas *ZkbasFilterer) WatchUpdateTokenPair(opts *bind.WatchOpts, sink chan<- *ZkbasUpdateTokenPair) (event.Subscription, error) {

	logs, sub, err := _Zkbas.contract.WatchLogs(opts, "UpdateTokenPair")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkbasUpdateTokenPair)
				if err := _Zkbas.contract.UnpackLog(event, "UpdateTokenPair", log); err != nil {
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
func (_Zkbas *ZkbasFilterer) ParseUpdateTokenPair(log types.Log) (*ZkbasUpdateTokenPair, error) {
	event := new(ZkbasUpdateTokenPair)
	if err := _Zkbas.contract.UnpackLog(event, "UpdateTokenPair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkbasWithdrawNftIterator is returned from FilterWithdrawNft and is used to iterate over the raw logs and unpacked data for WithdrawNft events raised by the Zkbas contract.
type ZkbasWithdrawNftIterator struct {
	Event *ZkbasWithdrawNft // Event containing the contract specifics and raw log

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
func (it *ZkbasWithdrawNftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkbasWithdrawNft)
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
		it.Event = new(ZkbasWithdrawNft)
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
func (it *ZkbasWithdrawNftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkbasWithdrawNftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkbasWithdrawNft represents a WithdrawNft event raised by the Zkbas contract.
type ZkbasWithdrawNft struct {
	AccountIndex uint32
	NftL1Address common.Address
	ToAddress    common.Address
	NftL1TokenId *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawNft is a free log retrieval operation binding the contract event 0x001c1fcac2c66b39f6b0a825948e9f8892b2d4a17edc60aaf61ca474e08402ec.
//
// Solidity: event WithdrawNft(uint32 accountIndex, address nftL1Address, address toAddress, uint256 nftL1TokenId)
func (_Zkbas *ZkbasFilterer) FilterWithdrawNft(opts *bind.FilterOpts) (*ZkbasWithdrawNftIterator, error) {

	logs, sub, err := _Zkbas.contract.FilterLogs(opts, "WithdrawNft")
	if err != nil {
		return nil, err
	}
	return &ZkbasWithdrawNftIterator{contract: _Zkbas.contract, event: "WithdrawNft", logs: logs, sub: sub}, nil
}

// WatchWithdrawNft is a free log subscription operation binding the contract event 0x001c1fcac2c66b39f6b0a825948e9f8892b2d4a17edc60aaf61ca474e08402ec.
//
// Solidity: event WithdrawNft(uint32 accountIndex, address nftL1Address, address toAddress, uint256 nftL1TokenId)
func (_Zkbas *ZkbasFilterer) WatchWithdrawNft(opts *bind.WatchOpts, sink chan<- *ZkbasWithdrawNft) (event.Subscription, error) {

	logs, sub, err := _Zkbas.contract.WatchLogs(opts, "WithdrawNft")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkbasWithdrawNft)
				if err := _Zkbas.contract.UnpackLog(event, "WithdrawNft", log); err != nil {
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
func (_Zkbas *ZkbasFilterer) ParseWithdrawNft(log types.Log) (*ZkbasWithdrawNft, error) {
	event := new(ZkbasWithdrawNft)
	if err := _Zkbas.contract.UnpackLog(event, "WithdrawNft", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkbasWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the Zkbas contract.
type ZkbasWithdrawalIterator struct {
	Event *ZkbasWithdrawal // Event containing the contract specifics and raw log

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
func (it *ZkbasWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkbasWithdrawal)
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
		it.Event = new(ZkbasWithdrawal)
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
func (it *ZkbasWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkbasWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkbasWithdrawal represents a Withdrawal event raised by the Zkbas contract.
type ZkbasWithdrawal struct {
	AssetId uint16
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0xf4bf32c167ee6e782944cd1db8174729b46adcd3bc732e282cc4a80793933154.
//
// Solidity: event Withdrawal(uint16 assetId, uint128 amount)
func (_Zkbas *ZkbasFilterer) FilterWithdrawal(opts *bind.FilterOpts) (*ZkbasWithdrawalIterator, error) {

	logs, sub, err := _Zkbas.contract.FilterLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return &ZkbasWithdrawalIterator{contract: _Zkbas.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0xf4bf32c167ee6e782944cd1db8174729b46adcd3bc732e282cc4a80793933154.
//
// Solidity: event Withdrawal(uint16 assetId, uint128 amount)
func (_Zkbas *ZkbasFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *ZkbasWithdrawal) (event.Subscription, error) {

	logs, sub, err := _Zkbas.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkbasWithdrawal)
				if err := _Zkbas.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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
func (_Zkbas *ZkbasFilterer) ParseWithdrawal(log types.Log) (*ZkbasWithdrawal, error) {
	event := new(ZkbasWithdrawal)
	if err := _Zkbas.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkbasWithdrawalNFTPendingIterator is returned from FilterWithdrawalNFTPending and is used to iterate over the raw logs and unpacked data for WithdrawalNFTPending events raised by the Zkbas contract.
type ZkbasWithdrawalNFTPendingIterator struct {
	Event *ZkbasWithdrawalNFTPending // Event containing the contract specifics and raw log

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
func (it *ZkbasWithdrawalNFTPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkbasWithdrawalNFTPending)
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
		it.Event = new(ZkbasWithdrawalNFTPending)
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
func (it *ZkbasWithdrawalNFTPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkbasWithdrawalNFTPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkbasWithdrawalNFTPending represents a WithdrawalNFTPending event raised by the Zkbas contract.
type ZkbasWithdrawalNFTPending struct {
	NftIndex *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalNFTPending is a free log retrieval operation binding the contract event 0x71aea045e572384e5b53812f175c12adb074001c9c13d379730d15188d3729bc.
//
// Solidity: event WithdrawalNFTPending(uint40 indexed nftIndex)
func (_Zkbas *ZkbasFilterer) FilterWithdrawalNFTPending(opts *bind.FilterOpts, nftIndex []*big.Int) (*ZkbasWithdrawalNFTPendingIterator, error) {

	var nftIndexRule []interface{}
	for _, nftIndexItem := range nftIndex {
		nftIndexRule = append(nftIndexRule, nftIndexItem)
	}

	logs, sub, err := _Zkbas.contract.FilterLogs(opts, "WithdrawalNFTPending", nftIndexRule)
	if err != nil {
		return nil, err
	}
	return &ZkbasWithdrawalNFTPendingIterator{contract: _Zkbas.contract, event: "WithdrawalNFTPending", logs: logs, sub: sub}, nil
}

// WatchWithdrawalNFTPending is a free log subscription operation binding the contract event 0x71aea045e572384e5b53812f175c12adb074001c9c13d379730d15188d3729bc.
//
// Solidity: event WithdrawalNFTPending(uint40 indexed nftIndex)
func (_Zkbas *ZkbasFilterer) WatchWithdrawalNFTPending(opts *bind.WatchOpts, sink chan<- *ZkbasWithdrawalNFTPending, nftIndex []*big.Int) (event.Subscription, error) {

	var nftIndexRule []interface{}
	for _, nftIndexItem := range nftIndex {
		nftIndexRule = append(nftIndexRule, nftIndexItem)
	}

	logs, sub, err := _Zkbas.contract.WatchLogs(opts, "WithdrawalNFTPending", nftIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkbasWithdrawalNFTPending)
				if err := _Zkbas.contract.UnpackLog(event, "WithdrawalNFTPending", log); err != nil {
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
func (_Zkbas *ZkbasFilterer) ParseWithdrawalNFTPending(log types.Log) (*ZkbasWithdrawalNFTPending, error) {
	event := new(ZkbasWithdrawalNFTPending)
	if err := _Zkbas.contract.UnpackLog(event, "WithdrawalNFTPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
