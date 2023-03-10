// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ZkBNB

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
	_ = abi.ConvertType
)

// ExodusVerifierExitData is an auto generated low-level Go binding around an user-defined struct.
type ExodusVerifierExitData struct {
	AssetId                  uint32
	AccountId                uint32
	Amount                   *big.Int
	OfferCanceledOrFinalized *big.Int
	AccountNameHash          [32]byte
	PubKeyX                  [32]byte
	PubKeyY                  [32]byte
	Nonce                    *big.Int
	CollectionNonce          *big.Int
}

// ExodusVerifierExitNftData is an auto generated low-level Go binding around an user-defined struct.
type ExodusVerifierExitNftData struct {
	NftIndex            uint64
	CreatorAccountIndex *big.Int
	NftContentHash      [32]byte
	CreatorTreasuryRate *big.Int
	CollectionId        *big.Int
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

// ZkBNBCommitBlockInfo is an auto generated low-level Go binding around an user-defined struct.
type ZkBNBCommitBlockInfo struct {
	NewStateRoot      [32]byte
	PublicData        []byte
	Timestamp         *big.Int
	OnchainOperations []ZkBNBOnchainOperationData
	BlockNumber       uint32
	BlockSize         uint16
}

// ZkBNBOnchainOperationData is an auto generated low-level Go binding around an user-defined struct.
type ZkBNBOnchainOperationData struct {
	EthWitness       []byte
	PublicDataOffset uint32
}

// ZkBNBVerifyAndExecuteBlockInfo is an auto generated low-level Go binding around an user-defined struct.
type ZkBNBVerifyAndExecuteBlockInfo struct {
	BlockHeader              StorageStoredBlockInfo
	PendingOnchainOpsPubData [][]byte
}

// ZkBNBMetaData contains all meta data concerning the ZkBNB contract.
var ZkBNBMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"BlockCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"BlockVerification\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"totalBlocksVerified\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"totalBlocksCommitted\",\"type\":\"uint32\"}],\"name\":\"BlocksRevert\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"assetId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"zkbnbBlockNumber\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accountName\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"assetId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"DepositCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nftContentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"creatorTreasuryRate\",\"type\":\"uint16\"}],\"name\":\"DepositNft\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DesertMode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"zkbnbBlockId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tokenId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"FullExitCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serialId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumTxTypes.TxType\",\"name\":\"txType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expirationBlock\",\"type\":\"uint256\"}],\"name\":\"NewPriorityRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNoticePeriod\",\"type\":\"uint256\"}],\"name\":\"NoticePeriodChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"zkbnbPubKeyX\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"zkbnbPubKeyY\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"}],\"name\":\"RegisterZNS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftL1Address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftL1TokenId\",\"type\":\"uint256\"}],\"name\":\"WithdrawNft\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"assetId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"nftIndex\",\"type\":\"uint40\"}],\"name\":\"WithdrawalNFTPending\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_ACCOUNT_INDEX\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_AMOUNT_OF_REGISTERED_ASSETS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_DEPOSIT_AMOUNT\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FUNGIBLE_ASSET_ID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_NFT_INDEX\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SPECIAL_ACCOUNT_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SPECIAL_ACCOUNT_ID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWAL_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateDesertMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_n\",\"type\":\"uint64\"},{\"internalType\":\"bytes[]\",\"name\":\"_depositsPubData\",\"type\":\"bytes[]\"}],\"name\":\"cancelOutstandingDepositsForExodusMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"blockSize\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"priorityOperations\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"pendingOnchainOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structStorage.StoredBlockInfo\",\"name\":\"_lastCommittedBlockData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"ethWitness\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"publicDataOffset\",\"type\":\"uint32\"}],\"internalType\":\"structZkBNB.OnchainOperationData[]\",\"name\":\"onchainOperations\",\"type\":\"tuple[]\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"blockSize\",\"type\":\"uint16\"}],\"internalType\":\"structZkBNB.CommitBlockInfo[]\",\"name\":\"_newBlocksData\",\"type\":\"tuple[]\"}],\"name\":\"commitBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint104\",\"name\":\"_amount\",\"type\":\"uint104\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"depositBEP20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"depositBNB\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nftL1Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nftL1TokenId\",\"type\":\"uint256\"}],\"name\":\"depositNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"desertMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstPriorityRequestId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_assetAddr\",\"type\":\"address\"}],\"name\":\"getPendingBalance\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initializationParameters\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_nftRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"assetId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"accountId\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offerCanceledOrFinalized\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"accountNameHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"pubKeyX\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"pubKeyY\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collectionNonce\",\"type\":\"uint256\"}],\"internalType\":\"structExodusVerifier.ExitData\",\"name\":\"_exitData\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[16]\",\"name\":\"_assetMerkleProof\",\"type\":\"bytes32[16]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"_accountMerkleProof\",\"type\":\"bytes32[32]\"}],\"name\":\"performDesert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ownerAccountIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_accountRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nftIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"creatorAccountIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"nftContentHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"creatorTreasuryRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"}],\"internalType\":\"structExodusVerifier.ExitNftData[]\",\"name\":\"_exitNfts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32[40][]\",\"name\":\"_nftMerkleProofs\",\"type\":\"bytes32[40][]\"}],\"name\":\"performDesertNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_accountIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"requestFullExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_accountIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_creatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_nftIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_nftContentType\",\"type\":\"uint8\"}],\"name\":\"requestFullExitNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"blockSize\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"priorityOperations\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"pendingOnchainOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structStorage.StoredBlockInfo[]\",\"name\":\"_blocksToRevert\",\"type\":\"tuple[]\"}],\"name\":\"revertBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"storedBlockHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBlocksCommitted\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBlocksVerified\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalOpenPriorityRequests\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_maxAmount\",\"type\":\"uint128\"}],\"name\":\"transferERC20\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"withdrawnAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"upgradeParameters\",\"type\":\"bytes\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"blockSize\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"priorityOperations\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"pendingOnchainOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structStorage.StoredBlockInfo\",\"name\":\"blockHeader\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"pendingOnchainOpsPubData\",\"type\":\"bytes[]\"}],\"internalType\":\"structZkBNB.VerifyAndExecuteBlockInfo[]\",\"name\":\"_blocks\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_proofs\",\"type\":\"uint256[]\"}],\"name\":\"verifyAndExecuteBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_amount\",\"type\":\"uint128\"}],\"name\":\"withdrawPendingBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_nftIndex\",\"type\":\"uint40\"}],\"name\":\"withdrawPendingNFTBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
	parsed, err := ZkBNBMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) pure returns(bytes4)
func (_ZkBNB *ZkBNBCaller) OnERC721Received(opts *bind.CallOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	var out []interface{}
	err := _ZkBNB.contract.Call(opts, &out, "onERC721Received", operator, from, tokenId, data)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) pure returns(bytes4)
func (_ZkBNB *ZkBNBSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	return _ZkBNB.Contract.OnERC721Received(&_ZkBNB.CallOpts, operator, from, tokenId, data)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) pure returns(bytes4)
func (_ZkBNB *ZkBNBCallerSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	return _ZkBNB.Contract.OnERC721Received(&_ZkBNB.CallOpts, operator, from, tokenId, data)
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

// CancelOutstandingDepositsForExodusMode is a paid mutator transaction binding the contract method 0x7efcfe85.
//
// Solidity: function cancelOutstandingDepositsForExodusMode(uint64 _n, bytes[] _depositsPubData) returns()
func (_ZkBNB *ZkBNBTransactor) CancelOutstandingDepositsForExodusMode(opts *bind.TransactOpts, _n uint64, _depositsPubData [][]byte) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "cancelOutstandingDepositsForExodusMode", _n, _depositsPubData)
}

// CancelOutstandingDepositsForExodusMode is a paid mutator transaction binding the contract method 0x7efcfe85.
//
// Solidity: function cancelOutstandingDepositsForExodusMode(uint64 _n, bytes[] _depositsPubData) returns()
func (_ZkBNB *ZkBNBSession) CancelOutstandingDepositsForExodusMode(_n uint64, _depositsPubData [][]byte) (*types.Transaction, error) {
	return _ZkBNB.Contract.CancelOutstandingDepositsForExodusMode(&_ZkBNB.TransactOpts, _n, _depositsPubData)
}

// CancelOutstandingDepositsForExodusMode is a paid mutator transaction binding the contract method 0x7efcfe85.
//
// Solidity: function cancelOutstandingDepositsForExodusMode(uint64 _n, bytes[] _depositsPubData) returns()
func (_ZkBNB *ZkBNBTransactorSession) CancelOutstandingDepositsForExodusMode(_n uint64, _depositsPubData [][]byte) (*types.Transaction, error) {
	return _ZkBNB.Contract.CancelOutstandingDepositsForExodusMode(&_ZkBNB.TransactOpts, _n, _depositsPubData)
}

// CommitBlocks is a paid mutator transaction binding the contract method 0x23c4d62d.
//
// Solidity: function commitBlocks((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32) _lastCommittedBlockData, (bytes32,bytes,uint256,(bytes,uint32)[],uint32,uint16)[] _newBlocksData) returns()
func (_ZkBNB *ZkBNBTransactor) CommitBlocks(opts *bind.TransactOpts, _lastCommittedBlockData StorageStoredBlockInfo, _newBlocksData []ZkBNBCommitBlockInfo) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "commitBlocks", _lastCommittedBlockData, _newBlocksData)
}

// CommitBlocks is a paid mutator transaction binding the contract method 0x23c4d62d.
//
// Solidity: function commitBlocks((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32) _lastCommittedBlockData, (bytes32,bytes,uint256,(bytes,uint32)[],uint32,uint16)[] _newBlocksData) returns()
func (_ZkBNB *ZkBNBSession) CommitBlocks(_lastCommittedBlockData StorageStoredBlockInfo, _newBlocksData []ZkBNBCommitBlockInfo) (*types.Transaction, error) {
	return _ZkBNB.Contract.CommitBlocks(&_ZkBNB.TransactOpts, _lastCommittedBlockData, _newBlocksData)
}

// CommitBlocks is a paid mutator transaction binding the contract method 0x23c4d62d.
//
// Solidity: function commitBlocks((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32) _lastCommittedBlockData, (bytes32,bytes,uint256,(bytes,uint32)[],uint32,uint16)[] _newBlocksData) returns()
func (_ZkBNB *ZkBNBTransactorSession) CommitBlocks(_lastCommittedBlockData StorageStoredBlockInfo, _newBlocksData []ZkBNBCommitBlockInfo) (*types.Transaction, error) {
	return _ZkBNB.Contract.CommitBlocks(&_ZkBNB.TransactOpts, _lastCommittedBlockData, _newBlocksData)
}

// DepositBEP20 is a paid mutator transaction binding the contract method 0x10ff3764.
//
// Solidity: function depositBEP20(address _token, uint104 _amount, address _to) returns()
func (_ZkBNB *ZkBNBTransactor) DepositBEP20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "depositBEP20", _token, _amount, _to)
}

// DepositBEP20 is a paid mutator transaction binding the contract method 0x10ff3764.
//
// Solidity: function depositBEP20(address _token, uint104 _amount, address _to) returns()
func (_ZkBNB *ZkBNBSession) DepositBEP20(_token common.Address, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _ZkBNB.Contract.DepositBEP20(&_ZkBNB.TransactOpts, _token, _amount, _to)
}

// DepositBEP20 is a paid mutator transaction binding the contract method 0x10ff3764.
//
// Solidity: function depositBEP20(address _token, uint104 _amount, address _to) returns()
func (_ZkBNB *ZkBNBTransactorSession) DepositBEP20(_token common.Address, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _ZkBNB.Contract.DepositBEP20(&_ZkBNB.TransactOpts, _token, _amount, _to)
}

// DepositBNB is a paid mutator transaction binding the contract method 0x684a5843.
//
// Solidity: function depositBNB(address _to) payable returns()
func (_ZkBNB *ZkBNBTransactor) DepositBNB(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "depositBNB", _to)
}

// DepositBNB is a paid mutator transaction binding the contract method 0x684a5843.
//
// Solidity: function depositBNB(address _to) payable returns()
func (_ZkBNB *ZkBNBSession) DepositBNB(_to common.Address) (*types.Transaction, error) {
	return _ZkBNB.Contract.DepositBNB(&_ZkBNB.TransactOpts, _to)
}

// DepositBNB is a paid mutator transaction binding the contract method 0x684a5843.
//
// Solidity: function depositBNB(address _to) payable returns()
func (_ZkBNB *ZkBNBTransactorSession) DepositBNB(_to common.Address) (*types.Transaction, error) {
	return _ZkBNB.Contract.DepositBNB(&_ZkBNB.TransactOpts, _to)
}

// DepositNft is a paid mutator transaction binding the contract method 0x82a5b1aa.
//
// Solidity: function depositNft(address _to, address _nftL1Address, uint256 _nftL1TokenId) returns()
func (_ZkBNB *ZkBNBTransactor) DepositNft(opts *bind.TransactOpts, _to common.Address, _nftL1Address common.Address, _nftL1TokenId *big.Int) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "depositNft", _to, _nftL1Address, _nftL1TokenId)
}

// DepositNft is a paid mutator transaction binding the contract method 0x82a5b1aa.
//
// Solidity: function depositNft(address _to, address _nftL1Address, uint256 _nftL1TokenId) returns()
func (_ZkBNB *ZkBNBSession) DepositNft(_to common.Address, _nftL1Address common.Address, _nftL1TokenId *big.Int) (*types.Transaction, error) {
	return _ZkBNB.Contract.DepositNft(&_ZkBNB.TransactOpts, _to, _nftL1Address, _nftL1TokenId)
}

// DepositNft is a paid mutator transaction binding the contract method 0x82a5b1aa.
//
// Solidity: function depositNft(address _to, address _nftL1Address, uint256 _nftL1TokenId) returns()
func (_ZkBNB *ZkBNBTransactorSession) DepositNft(_to common.Address, _nftL1Address common.Address, _nftL1TokenId *big.Int) (*types.Transaction, error) {
	return _ZkBNB.Contract.DepositNft(&_ZkBNB.TransactOpts, _to, _nftL1Address, _nftL1TokenId)
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

// PerformDesert is a paid mutator transaction binding the contract method 0x7dc3b65b.
//
// Solidity: function performDesert(bytes32 _nftRoot, (uint32,uint32,uint256,uint256,bytes32,bytes32,bytes32,uint256,uint256) _exitData, bytes32[16] _assetMerkleProof, bytes32[32] _accountMerkleProof) returns()
func (_ZkBNB *ZkBNBTransactor) PerformDesert(opts *bind.TransactOpts, _nftRoot [32]byte, _exitData ExodusVerifierExitData, _assetMerkleProof [16][32]byte, _accountMerkleProof [32][32]byte) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "performDesert", _nftRoot, _exitData, _assetMerkleProof, _accountMerkleProof)
}

// PerformDesert is a paid mutator transaction binding the contract method 0x7dc3b65b.
//
// Solidity: function performDesert(bytes32 _nftRoot, (uint32,uint32,uint256,uint256,bytes32,bytes32,bytes32,uint256,uint256) _exitData, bytes32[16] _assetMerkleProof, bytes32[32] _accountMerkleProof) returns()
func (_ZkBNB *ZkBNBSession) PerformDesert(_nftRoot [32]byte, _exitData ExodusVerifierExitData, _assetMerkleProof [16][32]byte, _accountMerkleProof [32][32]byte) (*types.Transaction, error) {
	return _ZkBNB.Contract.PerformDesert(&_ZkBNB.TransactOpts, _nftRoot, _exitData, _assetMerkleProof, _accountMerkleProof)
}

// PerformDesert is a paid mutator transaction binding the contract method 0x7dc3b65b.
//
// Solidity: function performDesert(bytes32 _nftRoot, (uint32,uint32,uint256,uint256,bytes32,bytes32,bytes32,uint256,uint256) _exitData, bytes32[16] _assetMerkleProof, bytes32[32] _accountMerkleProof) returns()
func (_ZkBNB *ZkBNBTransactorSession) PerformDesert(_nftRoot [32]byte, _exitData ExodusVerifierExitData, _assetMerkleProof [16][32]byte, _accountMerkleProof [32][32]byte) (*types.Transaction, error) {
	return _ZkBNB.Contract.PerformDesert(&_ZkBNB.TransactOpts, _nftRoot, _exitData, _assetMerkleProof, _accountMerkleProof)
}

// PerformDesertNft is a paid mutator transaction binding the contract method 0xb944c17f.
//
// Solidity: function performDesertNft(uint256 _ownerAccountIndex, bytes32 _accountRoot, (uint64,uint256,bytes32,uint256,uint256)[] _exitNfts, bytes32[40][] _nftMerkleProofs) returns()
func (_ZkBNB *ZkBNBTransactor) PerformDesertNft(opts *bind.TransactOpts, _ownerAccountIndex *big.Int, _accountRoot [32]byte, _exitNfts []ExodusVerifierExitNftData, _nftMerkleProofs [][40][32]byte) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "performDesertNft", _ownerAccountIndex, _accountRoot, _exitNfts, _nftMerkleProofs)
}

// PerformDesertNft is a paid mutator transaction binding the contract method 0xb944c17f.
//
// Solidity: function performDesertNft(uint256 _ownerAccountIndex, bytes32 _accountRoot, (uint64,uint256,bytes32,uint256,uint256)[] _exitNfts, bytes32[40][] _nftMerkleProofs) returns()
func (_ZkBNB *ZkBNBSession) PerformDesertNft(_ownerAccountIndex *big.Int, _accountRoot [32]byte, _exitNfts []ExodusVerifierExitNftData, _nftMerkleProofs [][40][32]byte) (*types.Transaction, error) {
	return _ZkBNB.Contract.PerformDesertNft(&_ZkBNB.TransactOpts, _ownerAccountIndex, _accountRoot, _exitNfts, _nftMerkleProofs)
}

// PerformDesertNft is a paid mutator transaction binding the contract method 0xb944c17f.
//
// Solidity: function performDesertNft(uint256 _ownerAccountIndex, bytes32 _accountRoot, (uint64,uint256,bytes32,uint256,uint256)[] _exitNfts, bytes32[40][] _nftMerkleProofs) returns()
func (_ZkBNB *ZkBNBTransactorSession) PerformDesertNft(_ownerAccountIndex *big.Int, _accountRoot [32]byte, _exitNfts []ExodusVerifierExitNftData, _nftMerkleProofs [][40][32]byte) (*types.Transaction, error) {
	return _ZkBNB.Contract.PerformDesertNft(&_ZkBNB.TransactOpts, _ownerAccountIndex, _accountRoot, _exitNfts, _nftMerkleProofs)
}

// RequestFullExit is a paid mutator transaction binding the contract method 0xab9b2adf.
//
// Solidity: function requestFullExit(uint32 _accountIndex, address _asset) returns()
func (_ZkBNB *ZkBNBTransactor) RequestFullExit(opts *bind.TransactOpts, _accountIndex uint32, _asset common.Address) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "requestFullExit", _accountIndex, _asset)
}

// RequestFullExit is a paid mutator transaction binding the contract method 0xab9b2adf.
//
// Solidity: function requestFullExit(uint32 _accountIndex, address _asset) returns()
func (_ZkBNB *ZkBNBSession) RequestFullExit(_accountIndex uint32, _asset common.Address) (*types.Transaction, error) {
	return _ZkBNB.Contract.RequestFullExit(&_ZkBNB.TransactOpts, _accountIndex, _asset)
}

// RequestFullExit is a paid mutator transaction binding the contract method 0xab9b2adf.
//
// Solidity: function requestFullExit(uint32 _accountIndex, address _asset) returns()
func (_ZkBNB *ZkBNBTransactorSession) RequestFullExit(_accountIndex uint32, _asset common.Address) (*types.Transaction, error) {
	return _ZkBNB.Contract.RequestFullExit(&_ZkBNB.TransactOpts, _accountIndex, _asset)
}

// RequestFullExitNft is a paid mutator transaction binding the contract method 0x9f680531.
//
// Solidity: function requestFullExitNft(uint32 _accountIndex, address _creatorAddress, uint32 _nftIndex, uint8 _nftContentType) returns()
func (_ZkBNB *ZkBNBTransactor) RequestFullExitNft(opts *bind.TransactOpts, _accountIndex uint32, _creatorAddress common.Address, _nftIndex uint32, _nftContentType uint8) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "requestFullExitNft", _accountIndex, _creatorAddress, _nftIndex, _nftContentType)
}

// RequestFullExitNft is a paid mutator transaction binding the contract method 0x9f680531.
//
// Solidity: function requestFullExitNft(uint32 _accountIndex, address _creatorAddress, uint32 _nftIndex, uint8 _nftContentType) returns()
func (_ZkBNB *ZkBNBSession) RequestFullExitNft(_accountIndex uint32, _creatorAddress common.Address, _nftIndex uint32, _nftContentType uint8) (*types.Transaction, error) {
	return _ZkBNB.Contract.RequestFullExitNft(&_ZkBNB.TransactOpts, _accountIndex, _creatorAddress, _nftIndex, _nftContentType)
}

// RequestFullExitNft is a paid mutator transaction binding the contract method 0x9f680531.
//
// Solidity: function requestFullExitNft(uint32 _accountIndex, address _creatorAddress, uint32 _nftIndex, uint8 _nftContentType) returns()
func (_ZkBNB *ZkBNBTransactorSession) RequestFullExitNft(_accountIndex uint32, _creatorAddress common.Address, _nftIndex uint32, _nftContentType uint8) (*types.Transaction, error) {
	return _ZkBNB.Contract.RequestFullExitNft(&_ZkBNB.TransactOpts, _accountIndex, _creatorAddress, _nftIndex, _nftContentType)
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

// VerifyAndExecuteBlocks is a paid mutator transaction binding the contract method 0xf9ea2e71.
//
// Solidity: function verifyAndExecuteBlocks(((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32),bytes[])[] _blocks, uint256[] _proofs) returns()
func (_ZkBNB *ZkBNBTransactor) VerifyAndExecuteBlocks(opts *bind.TransactOpts, _blocks []ZkBNBVerifyAndExecuteBlockInfo, _proofs []*big.Int) (*types.Transaction, error) {
	return _ZkBNB.contract.Transact(opts, "verifyAndExecuteBlocks", _blocks, _proofs)
}

// VerifyAndExecuteBlocks is a paid mutator transaction binding the contract method 0xf9ea2e71.
//
// Solidity: function verifyAndExecuteBlocks(((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32),bytes[])[] _blocks, uint256[] _proofs) returns()
func (_ZkBNB *ZkBNBSession) VerifyAndExecuteBlocks(_blocks []ZkBNBVerifyAndExecuteBlockInfo, _proofs []*big.Int) (*types.Transaction, error) {
	return _ZkBNB.Contract.VerifyAndExecuteBlocks(&_ZkBNB.TransactOpts, _blocks, _proofs)
}

// VerifyAndExecuteBlocks is a paid mutator transaction binding the contract method 0xf9ea2e71.
//
// Solidity: function verifyAndExecuteBlocks(((uint16,uint32,uint64,bytes32,uint256,bytes32,bytes32),bytes[])[] _blocks, uint256[] _proofs) returns()
func (_ZkBNB *ZkBNBTransactorSession) VerifyAndExecuteBlocks(_blocks []ZkBNBVerifyAndExecuteBlockInfo, _proofs []*big.Int) (*types.Transaction, error) {
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
	AssetId uint16
	To      common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x19bfffca67bba360efb5a2e90ef97006a550aa25f2a4aa62b83434dcb9cae63b.
//
// Solidity: event Deposit(uint16 assetId, address to, uint128 amount)
func (_ZkBNB *ZkBNBFilterer) FilterDeposit(opts *bind.FilterOpts) (*ZkBNBDepositIterator, error) {

	logs, sub, err := _ZkBNB.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &ZkBNBDepositIterator{contract: _ZkBNB.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x19bfffca67bba360efb5a2e90ef97006a550aa25f2a4aa62b83434dcb9cae63b.
//
// Solidity: event Deposit(uint16 assetId, address to, uint128 amount)
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

// ParseDeposit is a log parse operation binding the contract event 0x19bfffca67bba360efb5a2e90ef97006a550aa25f2a4aa62b83434dcb9cae63b.
//
// Solidity: event Deposit(uint16 assetId, address to, uint128 amount)
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
	ZkbnbBlockNumber uint32
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
	To                  common.Address
	NftContentHash      [32]byte
	TokenAddress        common.Address
	NftTokenId          *big.Int
	CreatorTreasuryRate uint16
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDepositNft is a free log retrieval operation binding the contract event 0xb56226c9bf0f6973433206e903ba8b6c0ba30634b6b71fdc5c52518f9541aee2.
//
// Solidity: event DepositNft(address to, bytes32 nftContentHash, address tokenAddress, uint256 nftTokenId, uint16 creatorTreasuryRate)
func (_ZkBNB *ZkBNBFilterer) FilterDepositNft(opts *bind.FilterOpts) (*ZkBNBDepositNftIterator, error) {

	logs, sub, err := _ZkBNB.contract.FilterLogs(opts, "DepositNft")
	if err != nil {
		return nil, err
	}
	return &ZkBNBDepositNftIterator{contract: _ZkBNB.contract, event: "DepositNft", logs: logs, sub: sub}, nil
}

// WatchDepositNft is a free log subscription operation binding the contract event 0xb56226c9bf0f6973433206e903ba8b6c0ba30634b6b71fdc5c52518f9541aee2.
//
// Solidity: event DepositNft(address to, bytes32 nftContentHash, address tokenAddress, uint256 nftTokenId, uint16 creatorTreasuryRate)
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

// ParseDepositNft is a log parse operation binding the contract event 0xb56226c9bf0f6973433206e903ba8b6c0ba30634b6b71fdc5c52518f9541aee2.
//
// Solidity: event DepositNft(address to, bytes32 nftContentHash, address tokenAddress, uint256 nftTokenId, uint16 creatorTreasuryRate)
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
	ZkbnbBlockId uint32
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

// ZkBNBInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ZkBNB contract.
type ZkBNBInitializedIterator struct {
	Event *ZkBNBInitialized // Event containing the contract specifics and raw log

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
func (it *ZkBNBInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkBNBInitialized)
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
		it.Event = new(ZkBNBInitialized)
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
func (it *ZkBNBInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkBNBInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkBNBInitialized represents a Initialized event raised by the ZkBNB contract.
type ZkBNBInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ZkBNB *ZkBNBFilterer) FilterInitialized(opts *bind.FilterOpts) (*ZkBNBInitializedIterator, error) {

	logs, sub, err := _ZkBNB.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ZkBNBInitializedIterator{contract: _ZkBNB.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ZkBNB *ZkBNBFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ZkBNBInitialized) (event.Subscription, error) {

	logs, sub, err := _ZkBNB.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkBNBInitialized)
				if err := _ZkBNB.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ZkBNB *ZkBNBFilterer) ParseInitialized(log types.Log) (*ZkBNBInitialized, error) {
	event := new(ZkBNBInitialized)
	if err := _ZkBNB.contract.UnpackLog(event, "Initialized", log); err != nil {
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
	ZkbnbPubKeyX [32]byte
	ZkbnbPubKeyY [32]byte
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
