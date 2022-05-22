// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zecreyLegend

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StorageBlockHeader is an auto generated low-level Go binding around an user-defined struct.
type StorageBlockHeader struct {
	BlockNumber                  uint32
	PriorityOperations           uint64
	PendingOnchainOperationsHash [32]byte
	Timestamp                    *big.Int
	AccountRoot                  [32]byte
	Commitment                   [32]byte
}

// ZecreyLegendCommitBlockInfo is an auto generated low-level Go binding around an user-defined struct.
type ZecreyLegendCommitBlockInfo struct {
	NewAccountRoot    [32]byte
	PublicData        []byte
	Timestamp         *big.Int
	PublicDataOffsets []uint32
	BlockNumber       uint32
}

// ZecreyLegendVerifyBlockInfo is an auto generated low-level Go binding around an user-defined struct.
type ZecreyLegendVerifyBlockInfo struct {
	BlockHeader              StorageBlockHeader
	PendingOnchainOpsPubdata [][]byte
}

// ZecreyLegendABI is the input ABI used to generate the binding from.
const ZecreyLegendABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"BlockCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"BlockVerification\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"totalBlocksVerified\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"totalBlocksCommitted\",\"type\":\"uint32\"}],\"name\":\"BlocksRevert\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"pairIndex\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"asset0Id\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"asset1Id\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"feeRate\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"treasuryAccountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"treasuryRate\",\"type\":\"uint16\"}],\"name\":\"CreateTokenPair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"assetId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accountName\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"zecreyBlockNumber\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accountName\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"assetId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"DepositCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accountNameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"creatorTreasuryRate\",\"type\":\"uint16\"}],\"name\":\"DepositNft\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DesertMode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"zecreyBlockId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tokenId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"FullExitCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serialId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumTxTypes.TxType\",\"name\":\"txType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expirationBlock\",\"type\":\"uint256\"}],\"name\":\"NewPriorityRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNoticePeriod\",\"type\":\"uint256\"}],\"name\":\"NoticePeriodChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"zecreyPubKey\",\"type\":\"bytes32\"}],\"name\":\"RegisterZNS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"pairIndex\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"feeRate\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"treasuryAccountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"treasuryRate\",\"type\":\"uint16\"}],\"name\":\"UpdateTokenPair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accountNameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxyAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"name\":\"WithdrawNft\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"assetId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"assetId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"WithdrawalPending\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_ACCOUNT_INDEX\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_AMOUNT_OF_REGISTERED_ASSETS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_DEPOSIT_AMOUNT\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FUNGIBLE_ASSET_ID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECURITY_COUNCIL_MEMBERS_NUMBER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHORTEST_UPGRADE_NOTICE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SPECIAL_ACCOUNT_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SPECIAL_ACCOUNT_ID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_NOTICE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWAL_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_maxAmount\",\"type\":\"uint128\"}],\"name\":\"_transferERC20\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"withdrawnAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateDesertMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"priorityOperations\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"pendingOnchainOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"accountRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structStorage.BlockHeader\",\"name\":\"_lastCommittedBlockData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"newAccountRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint32[]\",\"name\":\"publicDataOffsets\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"internalType\":\"structZecreyLegend.CommitBlockInfo[]\",\"name\":\"_newBlocksData\",\"type\":\"tuple[]\"}],\"name\":\"commitBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cutUpgradeNoticePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint104\",\"name\":\"_amount\",\"type\":\"uint104\"},{\"internalType\":\"bytes32\",\"name\":\"_accountNameHash\",\"type\":\"bytes32\"}],\"name\":\"depositBEP20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_accountNameHash\",\"type\":\"bytes32\"}],\"name\":\"depositBNB\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_accountNameHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nftTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_creatorTreasuryRate\",\"type\":\"uint16\"}],\"name\":\"depositNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"desertMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstPriorityRequestId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"accountNameHash\",\"type\":\"bytes32\"}],\"name\":\"getAddressByAccountNameHash\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNoticePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_assetAddr\",\"type\":\"address\"}],\"name\":\"getPendingBalance\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initializationParameters\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isReadyForUpgrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_zecreyPubKey\",\"type\":\"bytes32\"}],\"name\":\"registerZNS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_accountNameHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"requestFullExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_accountNameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_nftIndex\",\"type\":\"uint32\"}],\"name\":\"requestFullExitNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"priorityOperations\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"pendingOnchainOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"accountRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structStorage.BlockHeader[]\",\"name\":\"_blocksToRevert\",\"type\":\"tuple[]\"}],\"name\":\"revertBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"storedBlockHeaderHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBlocksCommitted\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBlocksVerified\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalOpenPriorityRequests\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTokenPairs\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_pairIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_feeRate\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"_treasuryAccountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"_treasuryRate\",\"type\":\"uint16\"}],\"name\":\"updatePairRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"upgradeParameters\",\"type\":\"bytes\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeCanceled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeFinishes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeNoticePeriodStarted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradePreparationStarted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"priorityOperations\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"pendingOnchainOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"accountRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structStorage.BlockHeader\",\"name\":\"blockHeader\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"pendingOnchainOpsPubdata\",\"type\":\"bytes[]\"}],\"internalType\":\"structZecreyLegend.VerifyBlockInfo[]\",\"name\":\"_blocks\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_proofs\",\"type\":\"uint256[]\"}],\"name\":\"verifyBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_amount\",\"type\":\"uint128\"}],\"name\":\"withdrawPendingBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ZecreyLegend is an auto generated Go binding around an Ethereum contract.
type ZecreyLegend struct {
	ZecreyLegendCaller     // Read-only binding to the contract
	ZecreyLegendTransactor // Write-only binding to the contract
	ZecreyLegendFilterer   // Log filterer for contract events
}

// ZecreyLegendCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZecreyLegendCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZecreyLegendTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZecreyLegendTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZecreyLegendFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZecreyLegendFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZecreyLegendSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZecreyLegendSession struct {
	Contract     *ZecreyLegend     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZecreyLegendCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZecreyLegendCallerSession struct {
	Contract *ZecreyLegendCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ZecreyLegendTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZecreyLegendTransactorSession struct {
	Contract     *ZecreyLegendTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ZecreyLegendRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZecreyLegendRaw struct {
	Contract *ZecreyLegend // Generic contract binding to access the raw methods on
}

// ZecreyLegendCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZecreyLegendCallerRaw struct {
	Contract *ZecreyLegendCaller // Generic read-only contract binding to access the raw methods on
}

// ZecreyLegendTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZecreyLegendTransactorRaw struct {
	Contract *ZecreyLegendTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZecreyLegend creates a new instance of ZecreyLegend, bound to a specific deployed contract.
func NewZecreyLegend(address common.Address, backend bind.ContractBackend) (*ZecreyLegend, error) {
	contract, err := bindZecreyLegend(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZecreyLegend{ZecreyLegendCaller: ZecreyLegendCaller{contract: contract}, ZecreyLegendTransactor: ZecreyLegendTransactor{contract: contract}, ZecreyLegendFilterer: ZecreyLegendFilterer{contract: contract}}, nil
}

// NewZecreyLegendCaller creates a new read-only instance of ZecreyLegend, bound to a specific deployed contract.
func NewZecreyLegendCaller(address common.Address, caller bind.ContractCaller) (*ZecreyLegendCaller, error) {
	contract, err := bindZecreyLegend(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZecreyLegendCaller{contract: contract}, nil
}

// NewZecreyLegendTransactor creates a new write-only instance of ZecreyLegend, bound to a specific deployed contract.
func NewZecreyLegendTransactor(address common.Address, transactor bind.ContractTransactor) (*ZecreyLegendTransactor, error) {
	contract, err := bindZecreyLegend(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZecreyLegendTransactor{contract: contract}, nil
}

// NewZecreyLegendFilterer creates a new log filterer instance of ZecreyLegend, bound to a specific deployed contract.
func NewZecreyLegendFilterer(address common.Address, filterer bind.ContractFilterer) (*ZecreyLegendFilterer, error) {
	contract, err := bindZecreyLegend(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZecreyLegendFilterer{contract: contract}, nil
}

// bindZecreyLegend binds a generic wrapper to an already deployed contract.
func bindZecreyLegend(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZecreyLegendABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZecreyLegend *ZecreyLegendRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZecreyLegend.Contract.ZecreyLegendCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZecreyLegend *ZecreyLegendRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.ZecreyLegendTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZecreyLegend *ZecreyLegendRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.ZecreyLegendTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZecreyLegend *ZecreyLegendCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZecreyLegend.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZecreyLegend *ZecreyLegendTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZecreyLegend *ZecreyLegendTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.contract.Transact(opts, method, params...)
}

// MAXACCOUNTINDEX is a free data retrieval call binding the contract method 0x437545f9.
//
// Solidity: function MAX_ACCOUNT_INDEX() view returns(uint32)
func (_ZecreyLegend *ZecreyLegendCaller) MAXACCOUNTINDEX(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ZecreyLegend.contract.Call(opts, &out, "MAX_ACCOUNT_INDEX")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXACCOUNTINDEX is a free data retrieval call binding the contract method 0x437545f9.
//
// Solidity: function MAX_ACCOUNT_INDEX() view returns(uint32)
func (_ZecreyLegend *ZecreyLegendSession) MAXACCOUNTINDEX() (uint32, error) {
	return _ZecreyLegend.Contract.MAXACCOUNTINDEX(&_ZecreyLegend.CallOpts)
}

// MAXACCOUNTINDEX is a free data retrieval call binding the contract method 0x437545f9.
//
// Solidity: function MAX_ACCOUNT_INDEX() view returns(uint32)
func (_ZecreyLegend *ZecreyLegendCallerSession) MAXACCOUNTINDEX() (uint32, error) {
	return _ZecreyLegend.Contract.MAXACCOUNTINDEX(&_ZecreyLegend.CallOpts)
}

// MAXAMOUNTOFREGISTEREDASSETS is a free data retrieval call binding the contract method 0x0d360b7f.
//
// Solidity: function MAX_AMOUNT_OF_REGISTERED_ASSETS() view returns(uint16)
func (_ZecreyLegend *ZecreyLegendCaller) MAXAMOUNTOFREGISTEREDASSETS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ZecreyLegend.contract.Call(opts, &out, "MAX_AMOUNT_OF_REGISTERED_ASSETS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXAMOUNTOFREGISTEREDASSETS is a free data retrieval call binding the contract method 0x0d360b7f.
//
// Solidity: function MAX_AMOUNT_OF_REGISTERED_ASSETS() view returns(uint16)
func (_ZecreyLegend *ZecreyLegendSession) MAXAMOUNTOFREGISTEREDASSETS() (uint16, error) {
	return _ZecreyLegend.Contract.MAXAMOUNTOFREGISTEREDASSETS(&_ZecreyLegend.CallOpts)
}

// MAXAMOUNTOFREGISTEREDASSETS is a free data retrieval call binding the contract method 0x0d360b7f.
//
// Solidity: function MAX_AMOUNT_OF_REGISTERED_ASSETS() view returns(uint16)
func (_ZecreyLegend *ZecreyLegendCallerSession) MAXAMOUNTOFREGISTEREDASSETS() (uint16, error) {
	return _ZecreyLegend.Contract.MAXAMOUNTOFREGISTEREDASSETS(&_ZecreyLegend.CallOpts)
}

// MAXDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x4c34a982.
//
// Solidity: function MAX_DEPOSIT_AMOUNT() view returns(uint128)
func (_ZecreyLegend *ZecreyLegendCaller) MAXDEPOSITAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZecreyLegend.contract.Call(opts, &out, "MAX_DEPOSIT_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x4c34a982.
//
// Solidity: function MAX_DEPOSIT_AMOUNT() view returns(uint128)
func (_ZecreyLegend *ZecreyLegendSession) MAXDEPOSITAMOUNT() (*big.Int, error) {
	return _ZecreyLegend.Contract.MAXDEPOSITAMOUNT(&_ZecreyLegend.CallOpts)
}

// MAXDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x4c34a982.
//
// Solidity: function MAX_DEPOSIT_AMOUNT() view returns(uint128)
func (_ZecreyLegend *ZecreyLegendCallerSession) MAXDEPOSITAMOUNT() (*big.Int, error) {
	return _ZecreyLegend.Contract.MAXDEPOSITAMOUNT(&_ZecreyLegend.CallOpts)
}

// MAXFUNGIBLEASSETID is a free data retrieval call binding the contract method 0x437da02f.
//
// Solidity: function MAX_FUNGIBLE_ASSET_ID() view returns(uint32)
func (_ZecreyLegend *ZecreyLegendCaller) MAXFUNGIBLEASSETID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ZecreyLegend.contract.Call(opts, &out, "MAX_FUNGIBLE_ASSET_ID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXFUNGIBLEASSETID is a free data retrieval call binding the contract method 0x437da02f.
//
// Solidity: function MAX_FUNGIBLE_ASSET_ID() view returns(uint32)
func (_ZecreyLegend *ZecreyLegendSession) MAXFUNGIBLEASSETID() (uint32, error) {
	return _ZecreyLegend.Contract.MAXFUNGIBLEASSETID(&_ZecreyLegend.CallOpts)
}

// MAXFUNGIBLEASSETID is a free data retrieval call binding the contract method 0x437da02f.
//
// Solidity: function MAX_FUNGIBLE_ASSET_ID() view returns(uint32)
func (_ZecreyLegend *ZecreyLegendCallerSession) MAXFUNGIBLEASSETID() (uint32, error) {
	return _ZecreyLegend.Contract.MAXFUNGIBLEASSETID(&_ZecreyLegend.CallOpts)
}

// SECURITYCOUNCILMEMBERSNUMBER is a free data retrieval call binding the contract method 0x4a51a71f.
//
// Solidity: function SECURITY_COUNCIL_MEMBERS_NUMBER() view returns(uint256)
func (_ZecreyLegend *ZecreyLegendCaller) SECURITYCOUNCILMEMBERSNUMBER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZecreyLegend.contract.Call(opts, &out, "SECURITY_COUNCIL_MEMBERS_NUMBER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECURITYCOUNCILMEMBERSNUMBER is a free data retrieval call binding the contract method 0x4a51a71f.
//
// Solidity: function SECURITY_COUNCIL_MEMBERS_NUMBER() view returns(uint256)
func (_ZecreyLegend *ZecreyLegendSession) SECURITYCOUNCILMEMBERSNUMBER() (*big.Int, error) {
	return _ZecreyLegend.Contract.SECURITYCOUNCILMEMBERSNUMBER(&_ZecreyLegend.CallOpts)
}

// SECURITYCOUNCILMEMBERSNUMBER is a free data retrieval call binding the contract method 0x4a51a71f.
//
// Solidity: function SECURITY_COUNCIL_MEMBERS_NUMBER() view returns(uint256)
func (_ZecreyLegend *ZecreyLegendCallerSession) SECURITYCOUNCILMEMBERSNUMBER() (*big.Int, error) {
	return _ZecreyLegend.Contract.SECURITYCOUNCILMEMBERSNUMBER(&_ZecreyLegend.CallOpts)
}

// SHORTESTUPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0x85053581.
//
// Solidity: function SHORTEST_UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_ZecreyLegend *ZecreyLegendCaller) SHORTESTUPGRADENOTICEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZecreyLegend.contract.Call(opts, &out, "SHORTEST_UPGRADE_NOTICE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SHORTESTUPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0x85053581.
//
// Solidity: function SHORTEST_UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_ZecreyLegend *ZecreyLegendSession) SHORTESTUPGRADENOTICEPERIOD() (*big.Int, error) {
	return _ZecreyLegend.Contract.SHORTESTUPGRADENOTICEPERIOD(&_ZecreyLegend.CallOpts)
}

// SHORTESTUPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0x85053581.
//
// Solidity: function SHORTEST_UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_ZecreyLegend *ZecreyLegendCallerSession) SHORTESTUPGRADENOTICEPERIOD() (*big.Int, error) {
	return _ZecreyLegend.Contract.SHORTESTUPGRADENOTICEPERIOD(&_ZecreyLegend.CallOpts)
}

// SPECIALACCOUNTADDRESS is a free data retrieval call binding the contract method 0x7ea399c1.
//
// Solidity: function SPECIAL_ACCOUNT_ADDRESS() view returns(address)
func (_ZecreyLegend *ZecreyLegendCaller) SPECIALACCOUNTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZecreyLegend.contract.Call(opts, &out, "SPECIAL_ACCOUNT_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SPECIALACCOUNTADDRESS is a free data retrieval call binding the contract method 0x7ea399c1.
//
// Solidity: function SPECIAL_ACCOUNT_ADDRESS() view returns(address)
func (_ZecreyLegend *ZecreyLegendSession) SPECIALACCOUNTADDRESS() (common.Address, error) {
	return _ZecreyLegend.Contract.SPECIALACCOUNTADDRESS(&_ZecreyLegend.CallOpts)
}

// SPECIALACCOUNTADDRESS is a free data retrieval call binding the contract method 0x7ea399c1.
//
// Solidity: function SPECIAL_ACCOUNT_ADDRESS() view returns(address)
func (_ZecreyLegend *ZecreyLegendCallerSession) SPECIALACCOUNTADDRESS() (common.Address, error) {
	return _ZecreyLegend.Contract.SPECIALACCOUNTADDRESS(&_ZecreyLegend.CallOpts)
}

// SPECIALACCOUNTID is a free data retrieval call binding the contract method 0x4242d5b3.
//
// Solidity: function SPECIAL_ACCOUNT_ID() view returns(uint32)
func (_ZecreyLegend *ZecreyLegendCaller) SPECIALACCOUNTID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ZecreyLegend.contract.Call(opts, &out, "SPECIAL_ACCOUNT_ID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SPECIALACCOUNTID is a free data retrieval call binding the contract method 0x4242d5b3.
//
// Solidity: function SPECIAL_ACCOUNT_ID() view returns(uint32)
func (_ZecreyLegend *ZecreyLegendSession) SPECIALACCOUNTID() (uint32, error) {
	return _ZecreyLegend.Contract.SPECIALACCOUNTID(&_ZecreyLegend.CallOpts)
}

// SPECIALACCOUNTID is a free data retrieval call binding the contract method 0x4242d5b3.
//
// Solidity: function SPECIAL_ACCOUNT_ID() view returns(uint32)
func (_ZecreyLegend *ZecreyLegendCallerSession) SPECIALACCOUNTID() (uint32, error) {
	return _ZecreyLegend.Contract.SPECIALACCOUNTID(&_ZecreyLegend.CallOpts)
}

// UPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0xcc375fb7.
//
// Solidity: function UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_ZecreyLegend *ZecreyLegendCaller) UPGRADENOTICEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZecreyLegend.contract.Call(opts, &out, "UPGRADE_NOTICE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0xcc375fb7.
//
// Solidity: function UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_ZecreyLegend *ZecreyLegendSession) UPGRADENOTICEPERIOD() (*big.Int, error) {
	return _ZecreyLegend.Contract.UPGRADENOTICEPERIOD(&_ZecreyLegend.CallOpts)
}

// UPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0xcc375fb7.
//
// Solidity: function UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_ZecreyLegend *ZecreyLegendCallerSession) UPGRADENOTICEPERIOD() (*big.Int, error) {
	return _ZecreyLegend.Contract.UPGRADENOTICEPERIOD(&_ZecreyLegend.CallOpts)
}

// WITHDRAWALGASLIMIT is a free data retrieval call binding the contract method 0xc701f955.
//
// Solidity: function WITHDRAWAL_GAS_LIMIT() view returns(uint256)
func (_ZecreyLegend *ZecreyLegendCaller) WITHDRAWALGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZecreyLegend.contract.Call(opts, &out, "WITHDRAWAL_GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WITHDRAWALGASLIMIT is a free data retrieval call binding the contract method 0xc701f955.
//
// Solidity: function WITHDRAWAL_GAS_LIMIT() view returns(uint256)
func (_ZecreyLegend *ZecreyLegendSession) WITHDRAWALGASLIMIT() (*big.Int, error) {
	return _ZecreyLegend.Contract.WITHDRAWALGASLIMIT(&_ZecreyLegend.CallOpts)
}

// WITHDRAWALGASLIMIT is a free data retrieval call binding the contract method 0xc701f955.
//
// Solidity: function WITHDRAWAL_GAS_LIMIT() view returns(uint256)
func (_ZecreyLegend *ZecreyLegendCallerSession) WITHDRAWALGASLIMIT() (*big.Int, error) {
	return _ZecreyLegend.Contract.WITHDRAWALGASLIMIT(&_ZecreyLegend.CallOpts)
}

// AccountRoot is a free data retrieval call binding the contract method 0xa7ccec2f.
//
// Solidity: function accountRoot() view returns(bytes32)
func (_ZecreyLegend *ZecreyLegendCaller) AccountRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZecreyLegend.contract.Call(opts, &out, "accountRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AccountRoot is a free data retrieval call binding the contract method 0xa7ccec2f.
//
// Solidity: function accountRoot() view returns(bytes32)
func (_ZecreyLegend *ZecreyLegendSession) AccountRoot() ([32]byte, error) {
	return _ZecreyLegend.Contract.AccountRoot(&_ZecreyLegend.CallOpts)
}

// AccountRoot is a free data retrieval call binding the contract method 0xa7ccec2f.
//
// Solidity: function accountRoot() view returns(bytes32)
func (_ZecreyLegend *ZecreyLegendCallerSession) AccountRoot() ([32]byte, error) {
	return _ZecreyLegend.Contract.AccountRoot(&_ZecreyLegend.CallOpts)
}

// DesertMode is a free data retrieval call binding the contract method 0x02cfb563.
//
// Solidity: function desertMode() view returns(bool)
func (_ZecreyLegend *ZecreyLegendCaller) DesertMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZecreyLegend.contract.Call(opts, &out, "desertMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DesertMode is a free data retrieval call binding the contract method 0x02cfb563.
//
// Solidity: function desertMode() view returns(bool)
func (_ZecreyLegend *ZecreyLegendSession) DesertMode() (bool, error) {
	return _ZecreyLegend.Contract.DesertMode(&_ZecreyLegend.CallOpts)
}

// DesertMode is a free data retrieval call binding the contract method 0x02cfb563.
//
// Solidity: function desertMode() view returns(bool)
func (_ZecreyLegend *ZecreyLegendCallerSession) DesertMode() (bool, error) {
	return _ZecreyLegend.Contract.DesertMode(&_ZecreyLegend.CallOpts)
}

// FirstPriorityRequestId is a free data retrieval call binding the contract method 0x67708dae.
//
// Solidity: function firstPriorityRequestId() view returns(uint64)
func (_ZecreyLegend *ZecreyLegendCaller) FirstPriorityRequestId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ZecreyLegend.contract.Call(opts, &out, "firstPriorityRequestId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// FirstPriorityRequestId is a free data retrieval call binding the contract method 0x67708dae.
//
// Solidity: function firstPriorityRequestId() view returns(uint64)
func (_ZecreyLegend *ZecreyLegendSession) FirstPriorityRequestId() (uint64, error) {
	return _ZecreyLegend.Contract.FirstPriorityRequestId(&_ZecreyLegend.CallOpts)
}

// FirstPriorityRequestId is a free data retrieval call binding the contract method 0x67708dae.
//
// Solidity: function firstPriorityRequestId() view returns(uint64)
func (_ZecreyLegend *ZecreyLegendCallerSession) FirstPriorityRequestId() (uint64, error) {
	return _ZecreyLegend.Contract.FirstPriorityRequestId(&_ZecreyLegend.CallOpts)
}

// GetAddressByAccountNameHash is a free data retrieval call binding the contract method 0x0b8f1c0c.
//
// Solidity: function getAddressByAccountNameHash(bytes32 accountNameHash) view returns(address)
func (_ZecreyLegend *ZecreyLegendCaller) GetAddressByAccountNameHash(opts *bind.CallOpts, accountNameHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ZecreyLegend.contract.Call(opts, &out, "getAddressByAccountNameHash", accountNameHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressByAccountNameHash is a free data retrieval call binding the contract method 0x0b8f1c0c.
//
// Solidity: function getAddressByAccountNameHash(bytes32 accountNameHash) view returns(address)
func (_ZecreyLegend *ZecreyLegendSession) GetAddressByAccountNameHash(accountNameHash [32]byte) (common.Address, error) {
	return _ZecreyLegend.Contract.GetAddressByAccountNameHash(&_ZecreyLegend.CallOpts, accountNameHash)
}

// GetAddressByAccountNameHash is a free data retrieval call binding the contract method 0x0b8f1c0c.
//
// Solidity: function getAddressByAccountNameHash(bytes32 accountNameHash) view returns(address)
func (_ZecreyLegend *ZecreyLegendCallerSession) GetAddressByAccountNameHash(accountNameHash [32]byte) (common.Address, error) {
	return _ZecreyLegend.Contract.GetAddressByAccountNameHash(&_ZecreyLegend.CallOpts, accountNameHash)
}

// GetNoticePeriod is a free data retrieval call binding the contract method 0x2a3174f4.
//
// Solidity: function getNoticePeriod() pure returns(uint256)
func (_ZecreyLegend *ZecreyLegendCaller) GetNoticePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZecreyLegend.contract.Call(opts, &out, "getNoticePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNoticePeriod is a free data retrieval call binding the contract method 0x2a3174f4.
//
// Solidity: function getNoticePeriod() pure returns(uint256)
func (_ZecreyLegend *ZecreyLegendSession) GetNoticePeriod() (*big.Int, error) {
	return _ZecreyLegend.Contract.GetNoticePeriod(&_ZecreyLegend.CallOpts)
}

// GetNoticePeriod is a free data retrieval call binding the contract method 0x2a3174f4.
//
// Solidity: function getNoticePeriod() pure returns(uint256)
func (_ZecreyLegend *ZecreyLegendCallerSession) GetNoticePeriod() (*big.Int, error) {
	return _ZecreyLegend.Contract.GetNoticePeriod(&_ZecreyLegend.CallOpts)
}

// GetPendingBalance is a free data retrieval call binding the contract method 0x5aca41f6.
//
// Solidity: function getPendingBalance(address _address, address _assetAddr) view returns(uint128)
func (_ZecreyLegend *ZecreyLegendCaller) GetPendingBalance(opts *bind.CallOpts, _address common.Address, _assetAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZecreyLegend.contract.Call(opts, &out, "getPendingBalance", _address, _assetAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingBalance is a free data retrieval call binding the contract method 0x5aca41f6.
//
// Solidity: function getPendingBalance(address _address, address _assetAddr) view returns(uint128)
func (_ZecreyLegend *ZecreyLegendSession) GetPendingBalance(_address common.Address, _assetAddr common.Address) (*big.Int, error) {
	return _ZecreyLegend.Contract.GetPendingBalance(&_ZecreyLegend.CallOpts, _address, _assetAddr)
}

// GetPendingBalance is a free data retrieval call binding the contract method 0x5aca41f6.
//
// Solidity: function getPendingBalance(address _address, address _assetAddr) view returns(uint128)
func (_ZecreyLegend *ZecreyLegendCallerSession) GetPendingBalance(_address common.Address, _assetAddr common.Address) (*big.Int, error) {
	return _ZecreyLegend.Contract.GetPendingBalance(&_ZecreyLegend.CallOpts, _address, _assetAddr)
}

// IsReadyForUpgrade is a free data retrieval call binding the contract method 0x8773334c.
//
// Solidity: function isReadyForUpgrade() view returns(bool)
func (_ZecreyLegend *ZecreyLegendCaller) IsReadyForUpgrade(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZecreyLegend.contract.Call(opts, &out, "isReadyForUpgrade")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReadyForUpgrade is a free data retrieval call binding the contract method 0x8773334c.
//
// Solidity: function isReadyForUpgrade() view returns(bool)
func (_ZecreyLegend *ZecreyLegendSession) IsReadyForUpgrade() (bool, error) {
	return _ZecreyLegend.Contract.IsReadyForUpgrade(&_ZecreyLegend.CallOpts)
}

// IsReadyForUpgrade is a free data retrieval call binding the contract method 0x8773334c.
//
// Solidity: function isReadyForUpgrade() view returns(bool)
func (_ZecreyLegend *ZecreyLegendCallerSession) IsReadyForUpgrade() (bool, error) {
	return _ZecreyLegend.Contract.IsReadyForUpgrade(&_ZecreyLegend.CallOpts)
}

// StoredBlockHeaderHashes is a free data retrieval call binding the contract method 0x4167d002.
//
// Solidity: function storedBlockHeaderHashes(uint32 ) view returns(bytes32)
func (_ZecreyLegend *ZecreyLegendCaller) StoredBlockHeaderHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ZecreyLegend.contract.Call(opts, &out, "storedBlockHeaderHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StoredBlockHeaderHashes is a free data retrieval call binding the contract method 0x4167d002.
//
// Solidity: function storedBlockHeaderHashes(uint32 ) view returns(bytes32)
func (_ZecreyLegend *ZecreyLegendSession) StoredBlockHeaderHashes(arg0 uint32) ([32]byte, error) {
	return _ZecreyLegend.Contract.StoredBlockHeaderHashes(&_ZecreyLegend.CallOpts, arg0)
}

// StoredBlockHeaderHashes is a free data retrieval call binding the contract method 0x4167d002.
//
// Solidity: function storedBlockHeaderHashes(uint32 ) view returns(bytes32)
func (_ZecreyLegend *ZecreyLegendCallerSession) StoredBlockHeaderHashes(arg0 uint32) ([32]byte, error) {
	return _ZecreyLegend.Contract.StoredBlockHeaderHashes(&_ZecreyLegend.CallOpts, arg0)
}

// TotalBlocksCommitted is a free data retrieval call binding the contract method 0xfaf4d8cb.
//
// Solidity: function totalBlocksCommitted() view returns(uint32)
func (_ZecreyLegend *ZecreyLegendCaller) TotalBlocksCommitted(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ZecreyLegend.contract.Call(opts, &out, "totalBlocksCommitted")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TotalBlocksCommitted is a free data retrieval call binding the contract method 0xfaf4d8cb.
//
// Solidity: function totalBlocksCommitted() view returns(uint32)
func (_ZecreyLegend *ZecreyLegendSession) TotalBlocksCommitted() (uint32, error) {
	return _ZecreyLegend.Contract.TotalBlocksCommitted(&_ZecreyLegend.CallOpts)
}

// TotalBlocksCommitted is a free data retrieval call binding the contract method 0xfaf4d8cb.
//
// Solidity: function totalBlocksCommitted() view returns(uint32)
func (_ZecreyLegend *ZecreyLegendCallerSession) TotalBlocksCommitted() (uint32, error) {
	return _ZecreyLegend.Contract.TotalBlocksCommitted(&_ZecreyLegend.CallOpts)
}

// TotalBlocksVerified is a free data retrieval call binding the contract method 0x2d24006c.
//
// Solidity: function totalBlocksVerified() view returns(uint32)
func (_ZecreyLegend *ZecreyLegendCaller) TotalBlocksVerified(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ZecreyLegend.contract.Call(opts, &out, "totalBlocksVerified")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TotalBlocksVerified is a free data retrieval call binding the contract method 0x2d24006c.
//
// Solidity: function totalBlocksVerified() view returns(uint32)
func (_ZecreyLegend *ZecreyLegendSession) TotalBlocksVerified() (uint32, error) {
	return _ZecreyLegend.Contract.TotalBlocksVerified(&_ZecreyLegend.CallOpts)
}

// TotalBlocksVerified is a free data retrieval call binding the contract method 0x2d24006c.
//
// Solidity: function totalBlocksVerified() view returns(uint32)
func (_ZecreyLegend *ZecreyLegendCallerSession) TotalBlocksVerified() (uint32, error) {
	return _ZecreyLegend.Contract.TotalBlocksVerified(&_ZecreyLegend.CallOpts)
}

// TotalOpenPriorityRequests is a free data retrieval call binding the contract method 0xc57b22be.
//
// Solidity: function totalOpenPriorityRequests() view returns(uint64)
func (_ZecreyLegend *ZecreyLegendCaller) TotalOpenPriorityRequests(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ZecreyLegend.contract.Call(opts, &out, "totalOpenPriorityRequests")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalOpenPriorityRequests is a free data retrieval call binding the contract method 0xc57b22be.
//
// Solidity: function totalOpenPriorityRequests() view returns(uint64)
func (_ZecreyLegend *ZecreyLegendSession) TotalOpenPriorityRequests() (uint64, error) {
	return _ZecreyLegend.Contract.TotalOpenPriorityRequests(&_ZecreyLegend.CallOpts)
}

// TotalOpenPriorityRequests is a free data retrieval call binding the contract method 0xc57b22be.
//
// Solidity: function totalOpenPriorityRequests() view returns(uint64)
func (_ZecreyLegend *ZecreyLegendCallerSession) TotalOpenPriorityRequests() (uint64, error) {
	return _ZecreyLegend.Contract.TotalOpenPriorityRequests(&_ZecreyLegend.CallOpts)
}

// TotalTokenPairs is a free data retrieval call binding the contract method 0xd0ad718d.
//
// Solidity: function totalTokenPairs() view returns(uint16)
func (_ZecreyLegend *ZecreyLegendCaller) TotalTokenPairs(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ZecreyLegend.contract.Call(opts, &out, "totalTokenPairs")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// TotalTokenPairs is a free data retrieval call binding the contract method 0xd0ad718d.
//
// Solidity: function totalTokenPairs() view returns(uint16)
func (_ZecreyLegend *ZecreyLegendSession) TotalTokenPairs() (uint16, error) {
	return _ZecreyLegend.Contract.TotalTokenPairs(&_ZecreyLegend.CallOpts)
}

// TotalTokenPairs is a free data retrieval call binding the contract method 0xd0ad718d.
//
// Solidity: function totalTokenPairs() view returns(uint16)
func (_ZecreyLegend *ZecreyLegendCallerSession) TotalTokenPairs() (uint16, error) {
	return _ZecreyLegend.Contract.TotalTokenPairs(&_ZecreyLegend.CallOpts)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x8ee1a74e.
//
// Solidity: function _transferERC20(address _token, address _to, uint128 _amount, uint128 _maxAmount) returns(uint128 withdrawnAmount)
func (_ZecreyLegend *ZecreyLegendTransactor) TransferERC20(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _maxAmount *big.Int) (*types.Transaction, error) {
	return _ZecreyLegend.contract.Transact(opts, "_transferERC20", _token, _to, _amount, _maxAmount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x8ee1a74e.
//
// Solidity: function _transferERC20(address _token, address _to, uint128 _amount, uint128 _maxAmount) returns(uint128 withdrawnAmount)
func (_ZecreyLegend *ZecreyLegendSession) TransferERC20(_token common.Address, _to common.Address, _amount *big.Int, _maxAmount *big.Int) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.TransferERC20(&_ZecreyLegend.TransactOpts, _token, _to, _amount, _maxAmount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x8ee1a74e.
//
// Solidity: function _transferERC20(address _token, address _to, uint128 _amount, uint128 _maxAmount) returns(uint128 withdrawnAmount)
func (_ZecreyLegend *ZecreyLegendTransactorSession) TransferERC20(_token common.Address, _to common.Address, _amount *big.Int, _maxAmount *big.Int) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.TransferERC20(&_ZecreyLegend.TransactOpts, _token, _to, _amount, _maxAmount)
}

// ActivateDesertMode is a paid mutator transaction binding the contract method 0x22b22256.
//
// Solidity: function activateDesertMode() returns(bool)
func (_ZecreyLegend *ZecreyLegendTransactor) ActivateDesertMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZecreyLegend.contract.Transact(opts, "activateDesertMode")
}

// ActivateDesertMode is a paid mutator transaction binding the contract method 0x22b22256.
//
// Solidity: function activateDesertMode() returns(bool)
func (_ZecreyLegend *ZecreyLegendSession) ActivateDesertMode() (*types.Transaction, error) {
	return _ZecreyLegend.Contract.ActivateDesertMode(&_ZecreyLegend.TransactOpts)
}

// ActivateDesertMode is a paid mutator transaction binding the contract method 0x22b22256.
//
// Solidity: function activateDesertMode() returns(bool)
func (_ZecreyLegend *ZecreyLegendTransactorSession) ActivateDesertMode() (*types.Transaction, error) {
	return _ZecreyLegend.Contract.ActivateDesertMode(&_ZecreyLegend.TransactOpts)
}

// CommitBlocks is a paid mutator transaction binding the contract method 0xaad1820e.
//
// Solidity: function commitBlocks((uint32,uint64,bytes32,uint256,bytes32,bytes32) _lastCommittedBlockData, (bytes32,bytes,uint256,uint32[],uint32)[] _newBlocksData) returns()
func (_ZecreyLegend *ZecreyLegendTransactor) CommitBlocks(opts *bind.TransactOpts, _lastCommittedBlockData StorageBlockHeader, _newBlocksData []ZecreyLegendCommitBlockInfo) (*types.Transaction, error) {
	return _ZecreyLegend.contract.Transact(opts, "commitBlocks", _lastCommittedBlockData, _newBlocksData)
}

// CommitBlocks is a paid mutator transaction binding the contract method 0xaad1820e.
//
// Solidity: function commitBlocks((uint32,uint64,bytes32,uint256,bytes32,bytes32) _lastCommittedBlockData, (bytes32,bytes,uint256,uint32[],uint32)[] _newBlocksData) returns()
func (_ZecreyLegend *ZecreyLegendSession) CommitBlocks(_lastCommittedBlockData StorageBlockHeader, _newBlocksData []ZecreyLegendCommitBlockInfo) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.CommitBlocks(&_ZecreyLegend.TransactOpts, _lastCommittedBlockData, _newBlocksData)
}

// CommitBlocks is a paid mutator transaction binding the contract method 0xaad1820e.
//
// Solidity: function commitBlocks((uint32,uint64,bytes32,uint256,bytes32,bytes32) _lastCommittedBlockData, (bytes32,bytes,uint256,uint32[],uint32)[] _newBlocksData) returns()
func (_ZecreyLegend *ZecreyLegendTransactorSession) CommitBlocks(_lastCommittedBlockData StorageBlockHeader, _newBlocksData []ZecreyLegendCommitBlockInfo) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.CommitBlocks(&_ZecreyLegend.TransactOpts, _lastCommittedBlockData, _newBlocksData)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address _tokenA, address _tokenB) returns()
func (_ZecreyLegend *ZecreyLegendTransactor) CreatePair(opts *bind.TransactOpts, _tokenA common.Address, _tokenB common.Address) (*types.Transaction, error) {
	return _ZecreyLegend.contract.Transact(opts, "createPair", _tokenA, _tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address _tokenA, address _tokenB) returns()
func (_ZecreyLegend *ZecreyLegendSession) CreatePair(_tokenA common.Address, _tokenB common.Address) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.CreatePair(&_ZecreyLegend.TransactOpts, _tokenA, _tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address _tokenA, address _tokenB) returns()
func (_ZecreyLegend *ZecreyLegendTransactorSession) CreatePair(_tokenA common.Address, _tokenB common.Address) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.CreatePair(&_ZecreyLegend.TransactOpts, _tokenA, _tokenB)
}

// CutUpgradeNoticePeriod is a paid mutator transaction binding the contract method 0x3e71e1e7.
//
// Solidity: function cutUpgradeNoticePeriod() returns()
func (_ZecreyLegend *ZecreyLegendTransactor) CutUpgradeNoticePeriod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZecreyLegend.contract.Transact(opts, "cutUpgradeNoticePeriod")
}

// CutUpgradeNoticePeriod is a paid mutator transaction binding the contract method 0x3e71e1e7.
//
// Solidity: function cutUpgradeNoticePeriod() returns()
func (_ZecreyLegend *ZecreyLegendSession) CutUpgradeNoticePeriod() (*types.Transaction, error) {
	return _ZecreyLegend.Contract.CutUpgradeNoticePeriod(&_ZecreyLegend.TransactOpts)
}

// CutUpgradeNoticePeriod is a paid mutator transaction binding the contract method 0x3e71e1e7.
//
// Solidity: function cutUpgradeNoticePeriod() returns()
func (_ZecreyLegend *ZecreyLegendTransactorSession) CutUpgradeNoticePeriod() (*types.Transaction, error) {
	return _ZecreyLegend.Contract.CutUpgradeNoticePeriod(&_ZecreyLegend.TransactOpts)
}

// DepositBEP20 is a paid mutator transaction binding the contract method 0x03892c3c.
//
// Solidity: function depositBEP20(address _token, uint104 _amount, bytes32 _accountNameHash) returns()
func (_ZecreyLegend *ZecreyLegendTransactor) DepositBEP20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _accountNameHash [32]byte) (*types.Transaction, error) {
	return _ZecreyLegend.contract.Transact(opts, "depositBEP20", _token, _amount, _accountNameHash)
}

// DepositBEP20 is a paid mutator transaction binding the contract method 0x03892c3c.
//
// Solidity: function depositBEP20(address _token, uint104 _amount, bytes32 _accountNameHash) returns()
func (_ZecreyLegend *ZecreyLegendSession) DepositBEP20(_token common.Address, _amount *big.Int, _accountNameHash [32]byte) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.DepositBEP20(&_ZecreyLegend.TransactOpts, _token, _amount, _accountNameHash)
}

// DepositBEP20 is a paid mutator transaction binding the contract method 0x03892c3c.
//
// Solidity: function depositBEP20(address _token, uint104 _amount, bytes32 _accountNameHash) returns()
func (_ZecreyLegend *ZecreyLegendTransactorSession) DepositBEP20(_token common.Address, _amount *big.Int, _accountNameHash [32]byte) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.DepositBEP20(&_ZecreyLegend.TransactOpts, _token, _amount, _accountNameHash)
}

// DepositBNB is a paid mutator transaction binding the contract method 0x3d002926.
//
// Solidity: function depositBNB(bytes32 _accountNameHash) payable returns()
func (_ZecreyLegend *ZecreyLegendTransactor) DepositBNB(opts *bind.TransactOpts, _accountNameHash [32]byte) (*types.Transaction, error) {
	return _ZecreyLegend.contract.Transact(opts, "depositBNB", _accountNameHash)
}

// DepositBNB is a paid mutator transaction binding the contract method 0x3d002926.
//
// Solidity: function depositBNB(bytes32 _accountNameHash) payable returns()
func (_ZecreyLegend *ZecreyLegendSession) DepositBNB(_accountNameHash [32]byte) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.DepositBNB(&_ZecreyLegend.TransactOpts, _accountNameHash)
}

// DepositBNB is a paid mutator transaction binding the contract method 0x3d002926.
//
// Solidity: function depositBNB(bytes32 _accountNameHash) payable returns()
func (_ZecreyLegend *ZecreyLegendTransactorSession) DepositBNB(_accountNameHash [32]byte) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.DepositBNB(&_ZecreyLegend.TransactOpts, _accountNameHash)
}

// DepositNft is a paid mutator transaction binding the contract method 0x8900e4c5.
//
// Solidity: function depositNft(bytes32 _accountNameHash, address _tokenAddress, uint256 _nftTokenId, uint16 _creatorTreasuryRate) returns()
func (_ZecreyLegend *ZecreyLegendTransactor) DepositNft(opts *bind.TransactOpts, _accountNameHash [32]byte, _tokenAddress common.Address, _nftTokenId *big.Int, _creatorTreasuryRate uint16) (*types.Transaction, error) {
	return _ZecreyLegend.contract.Transact(opts, "depositNft", _accountNameHash, _tokenAddress, _nftTokenId, _creatorTreasuryRate)
}

// DepositNft is a paid mutator transaction binding the contract method 0x8900e4c5.
//
// Solidity: function depositNft(bytes32 _accountNameHash, address _tokenAddress, uint256 _nftTokenId, uint16 _creatorTreasuryRate) returns()
func (_ZecreyLegend *ZecreyLegendSession) DepositNft(_accountNameHash [32]byte, _tokenAddress common.Address, _nftTokenId *big.Int, _creatorTreasuryRate uint16) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.DepositNft(&_ZecreyLegend.TransactOpts, _accountNameHash, _tokenAddress, _nftTokenId, _creatorTreasuryRate)
}

// DepositNft is a paid mutator transaction binding the contract method 0x8900e4c5.
//
// Solidity: function depositNft(bytes32 _accountNameHash, address _tokenAddress, uint256 _nftTokenId, uint16 _creatorTreasuryRate) returns()
func (_ZecreyLegend *ZecreyLegendTransactorSession) DepositNft(_accountNameHash [32]byte, _tokenAddress common.Address, _nftTokenId *big.Int, _creatorTreasuryRate uint16) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.DepositNft(&_ZecreyLegend.TransactOpts, _accountNameHash, _tokenAddress, _nftTokenId, _creatorTreasuryRate)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializationParameters) returns()
func (_ZecreyLegend *ZecreyLegendTransactor) Initialize(opts *bind.TransactOpts, initializationParameters []byte) (*types.Transaction, error) {
	return _ZecreyLegend.contract.Transact(opts, "initialize", initializationParameters)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializationParameters) returns()
func (_ZecreyLegend *ZecreyLegendSession) Initialize(initializationParameters []byte) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.Initialize(&_ZecreyLegend.TransactOpts, initializationParameters)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializationParameters) returns()
func (_ZecreyLegend *ZecreyLegendTransactorSession) Initialize(initializationParameters []byte) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.Initialize(&_ZecreyLegend.TransactOpts, initializationParameters)
}

// RegisterZNS is a paid mutator transaction binding the contract method 0xd7757da1.
//
// Solidity: function registerZNS(string _name, address _owner, bytes32 _zecreyPubKey) returns()
func (_ZecreyLegend *ZecreyLegendTransactor) RegisterZNS(opts *bind.TransactOpts, _name string, _owner common.Address, _zecreyPubKey [32]byte) (*types.Transaction, error) {
	return _ZecreyLegend.contract.Transact(opts, "registerZNS", _name, _owner, _zecreyPubKey)
}

// RegisterZNS is a paid mutator transaction binding the contract method 0xd7757da1.
//
// Solidity: function registerZNS(string _name, address _owner, bytes32 _zecreyPubKey) returns()
func (_ZecreyLegend *ZecreyLegendSession) RegisterZNS(_name string, _owner common.Address, _zecreyPubKey [32]byte) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.RegisterZNS(&_ZecreyLegend.TransactOpts, _name, _owner, _zecreyPubKey)
}

// RegisterZNS is a paid mutator transaction binding the contract method 0xd7757da1.
//
// Solidity: function registerZNS(string _name, address _owner, bytes32 _zecreyPubKey) returns()
func (_ZecreyLegend *ZecreyLegendTransactorSession) RegisterZNS(_name string, _owner common.Address, _zecreyPubKey [32]byte) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.RegisterZNS(&_ZecreyLegend.TransactOpts, _name, _owner, _zecreyPubKey)
}

// RequestFullExit is a paid mutator transaction binding the contract method 0x5fa8fc89.
//
// Solidity: function requestFullExit(bytes32 _accountNameHash, address _asset) returns()
func (_ZecreyLegend *ZecreyLegendTransactor) RequestFullExit(opts *bind.TransactOpts, _accountNameHash [32]byte, _asset common.Address) (*types.Transaction, error) {
	return _ZecreyLegend.contract.Transact(opts, "requestFullExit", _accountNameHash, _asset)
}

// RequestFullExit is a paid mutator transaction binding the contract method 0x5fa8fc89.
//
// Solidity: function requestFullExit(bytes32 _accountNameHash, address _asset) returns()
func (_ZecreyLegend *ZecreyLegendSession) RequestFullExit(_accountNameHash [32]byte, _asset common.Address) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.RequestFullExit(&_ZecreyLegend.TransactOpts, _accountNameHash, _asset)
}

// RequestFullExit is a paid mutator transaction binding the contract method 0x5fa8fc89.
//
// Solidity: function requestFullExit(bytes32 _accountNameHash, address _asset) returns()
func (_ZecreyLegend *ZecreyLegendTransactorSession) RequestFullExit(_accountNameHash [32]byte, _asset common.Address) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.RequestFullExit(&_ZecreyLegend.TransactOpts, _accountNameHash, _asset)
}

// RequestFullExitNft is a paid mutator transaction binding the contract method 0xc5d9f7cb.
//
// Solidity: function requestFullExitNft(bytes32 _accountNameHash, uint32 _nftIndex) returns()
func (_ZecreyLegend *ZecreyLegendTransactor) RequestFullExitNft(opts *bind.TransactOpts, _accountNameHash [32]byte, _nftIndex uint32) (*types.Transaction, error) {
	return _ZecreyLegend.contract.Transact(opts, "requestFullExitNft", _accountNameHash, _nftIndex)
}

// RequestFullExitNft is a paid mutator transaction binding the contract method 0xc5d9f7cb.
//
// Solidity: function requestFullExitNft(bytes32 _accountNameHash, uint32 _nftIndex) returns()
func (_ZecreyLegend *ZecreyLegendSession) RequestFullExitNft(_accountNameHash [32]byte, _nftIndex uint32) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.RequestFullExitNft(&_ZecreyLegend.TransactOpts, _accountNameHash, _nftIndex)
}

// RequestFullExitNft is a paid mutator transaction binding the contract method 0xc5d9f7cb.
//
// Solidity: function requestFullExitNft(bytes32 _accountNameHash, uint32 _nftIndex) returns()
func (_ZecreyLegend *ZecreyLegendTransactorSession) RequestFullExitNft(_accountNameHash [32]byte, _nftIndex uint32) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.RequestFullExitNft(&_ZecreyLegend.TransactOpts, _accountNameHash, _nftIndex)
}

// RevertBlocks is a paid mutator transaction binding the contract method 0xb4a8498c.
//
// Solidity: function revertBlocks((uint32,uint64,bytes32,uint256,bytes32,bytes32)[] _blocksToRevert) returns()
func (_ZecreyLegend *ZecreyLegendTransactor) RevertBlocks(opts *bind.TransactOpts, _blocksToRevert []StorageBlockHeader) (*types.Transaction, error) {
	return _ZecreyLegend.contract.Transact(opts, "revertBlocks", _blocksToRevert)
}

// RevertBlocks is a paid mutator transaction binding the contract method 0xb4a8498c.
//
// Solidity: function revertBlocks((uint32,uint64,bytes32,uint256,bytes32,bytes32)[] _blocksToRevert) returns()
func (_ZecreyLegend *ZecreyLegendSession) RevertBlocks(_blocksToRevert []StorageBlockHeader) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.RevertBlocks(&_ZecreyLegend.TransactOpts, _blocksToRevert)
}

// RevertBlocks is a paid mutator transaction binding the contract method 0xb4a8498c.
//
// Solidity: function revertBlocks((uint32,uint64,bytes32,uint256,bytes32,bytes32)[] _blocksToRevert) returns()
func (_ZecreyLegend *ZecreyLegendTransactorSession) RevertBlocks(_blocksToRevert []StorageBlockHeader) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.RevertBlocks(&_ZecreyLegend.TransactOpts, _blocksToRevert)
}

// UpdatePairRate is a paid mutator transaction binding the contract method 0xd01b2308.
//
// Solidity: function updatePairRate(uint16 _pairIndex, uint16 _feeRate, uint32 _treasuryAccountIndex, uint16 _treasuryRate) returns()
func (_ZecreyLegend *ZecreyLegendTransactor) UpdatePairRate(opts *bind.TransactOpts, _pairIndex uint16, _feeRate uint16, _treasuryAccountIndex uint32, _treasuryRate uint16) (*types.Transaction, error) {
	return _ZecreyLegend.contract.Transact(opts, "updatePairRate", _pairIndex, _feeRate, _treasuryAccountIndex, _treasuryRate)
}

// UpdatePairRate is a paid mutator transaction binding the contract method 0xd01b2308.
//
// Solidity: function updatePairRate(uint16 _pairIndex, uint16 _feeRate, uint32 _treasuryAccountIndex, uint16 _treasuryRate) returns()
func (_ZecreyLegend *ZecreyLegendSession) UpdatePairRate(_pairIndex uint16, _feeRate uint16, _treasuryAccountIndex uint32, _treasuryRate uint16) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.UpdatePairRate(&_ZecreyLegend.TransactOpts, _pairIndex, _feeRate, _treasuryAccountIndex, _treasuryRate)
}

// UpdatePairRate is a paid mutator transaction binding the contract method 0xd01b2308.
//
// Solidity: function updatePairRate(uint16 _pairIndex, uint16 _feeRate, uint32 _treasuryAccountIndex, uint16 _treasuryRate) returns()
func (_ZecreyLegend *ZecreyLegendTransactorSession) UpdatePairRate(_pairIndex uint16, _feeRate uint16, _treasuryAccountIndex uint32, _treasuryRate uint16) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.UpdatePairRate(&_ZecreyLegend.TransactOpts, _pairIndex, _feeRate, _treasuryAccountIndex, _treasuryRate)
}

// Upgrade is a paid mutator transaction binding the contract method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_ZecreyLegend *ZecreyLegendTransactor) Upgrade(opts *bind.TransactOpts, upgradeParameters []byte) (*types.Transaction, error) {
	return _ZecreyLegend.contract.Transact(opts, "upgrade", upgradeParameters)
}

// Upgrade is a paid mutator transaction binding the contract method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_ZecreyLegend *ZecreyLegendSession) Upgrade(upgradeParameters []byte) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.Upgrade(&_ZecreyLegend.TransactOpts, upgradeParameters)
}

// Upgrade is a paid mutator transaction binding the contract method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_ZecreyLegend *ZecreyLegendTransactorSession) Upgrade(upgradeParameters []byte) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.Upgrade(&_ZecreyLegend.TransactOpts, upgradeParameters)
}

// UpgradeCanceled is a paid mutator transaction binding the contract method 0x871b8ff1.
//
// Solidity: function upgradeCanceled() returns()
func (_ZecreyLegend *ZecreyLegendTransactor) UpgradeCanceled(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZecreyLegend.contract.Transact(opts, "upgradeCanceled")
}

// UpgradeCanceled is a paid mutator transaction binding the contract method 0x871b8ff1.
//
// Solidity: function upgradeCanceled() returns()
func (_ZecreyLegend *ZecreyLegendSession) UpgradeCanceled() (*types.Transaction, error) {
	return _ZecreyLegend.Contract.UpgradeCanceled(&_ZecreyLegend.TransactOpts)
}

// UpgradeCanceled is a paid mutator transaction binding the contract method 0x871b8ff1.
//
// Solidity: function upgradeCanceled() returns()
func (_ZecreyLegend *ZecreyLegendTransactorSession) UpgradeCanceled() (*types.Transaction, error) {
	return _ZecreyLegend.Contract.UpgradeCanceled(&_ZecreyLegend.TransactOpts)
}

// UpgradeFinishes is a paid mutator transaction binding the contract method 0xb269b9ae.
//
// Solidity: function upgradeFinishes() returns()
func (_ZecreyLegend *ZecreyLegendTransactor) UpgradeFinishes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZecreyLegend.contract.Transact(opts, "upgradeFinishes")
}

// UpgradeFinishes is a paid mutator transaction binding the contract method 0xb269b9ae.
//
// Solidity: function upgradeFinishes() returns()
func (_ZecreyLegend *ZecreyLegendSession) UpgradeFinishes() (*types.Transaction, error) {
	return _ZecreyLegend.Contract.UpgradeFinishes(&_ZecreyLegend.TransactOpts)
}

// UpgradeFinishes is a paid mutator transaction binding the contract method 0xb269b9ae.
//
// Solidity: function upgradeFinishes() returns()
func (_ZecreyLegend *ZecreyLegendTransactorSession) UpgradeFinishes() (*types.Transaction, error) {
	return _ZecreyLegend.Contract.UpgradeFinishes(&_ZecreyLegend.TransactOpts)
}

// UpgradeNoticePeriodStarted is a paid mutator transaction binding the contract method 0x3b154b73.
//
// Solidity: function upgradeNoticePeriodStarted() returns()
func (_ZecreyLegend *ZecreyLegendTransactor) UpgradeNoticePeriodStarted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZecreyLegend.contract.Transact(opts, "upgradeNoticePeriodStarted")
}

// UpgradeNoticePeriodStarted is a paid mutator transaction binding the contract method 0x3b154b73.
//
// Solidity: function upgradeNoticePeriodStarted() returns()
func (_ZecreyLegend *ZecreyLegendSession) UpgradeNoticePeriodStarted() (*types.Transaction, error) {
	return _ZecreyLegend.Contract.UpgradeNoticePeriodStarted(&_ZecreyLegend.TransactOpts)
}

// UpgradeNoticePeriodStarted is a paid mutator transaction binding the contract method 0x3b154b73.
//
// Solidity: function upgradeNoticePeriodStarted() returns()
func (_ZecreyLegend *ZecreyLegendTransactorSession) UpgradeNoticePeriodStarted() (*types.Transaction, error) {
	return _ZecreyLegend.Contract.UpgradeNoticePeriodStarted(&_ZecreyLegend.TransactOpts)
}

// UpgradePreparationStarted is a paid mutator transaction binding the contract method 0x78b91e70.
//
// Solidity: function upgradePreparationStarted() returns()
func (_ZecreyLegend *ZecreyLegendTransactor) UpgradePreparationStarted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZecreyLegend.contract.Transact(opts, "upgradePreparationStarted")
}

// UpgradePreparationStarted is a paid mutator transaction binding the contract method 0x78b91e70.
//
// Solidity: function upgradePreparationStarted() returns()
func (_ZecreyLegend *ZecreyLegendSession) UpgradePreparationStarted() (*types.Transaction, error) {
	return _ZecreyLegend.Contract.UpgradePreparationStarted(&_ZecreyLegend.TransactOpts)
}

// UpgradePreparationStarted is a paid mutator transaction binding the contract method 0x78b91e70.
//
// Solidity: function upgradePreparationStarted() returns()
func (_ZecreyLegend *ZecreyLegendTransactorSession) UpgradePreparationStarted() (*types.Transaction, error) {
	return _ZecreyLegend.Contract.UpgradePreparationStarted(&_ZecreyLegend.TransactOpts)
}

// VerifyBlocks is a paid mutator transaction binding the contract method 0xb77a1ef5.
//
// Solidity: function verifyBlocks(((uint32,uint64,bytes32,uint256,bytes32,bytes32),bytes[])[] _blocks, uint256[] _proofs) returns()
func (_ZecreyLegend *ZecreyLegendTransactor) VerifyBlocks(opts *bind.TransactOpts, _blocks []ZecreyLegendVerifyBlockInfo, _proofs []*big.Int) (*types.Transaction, error) {
	return _ZecreyLegend.contract.Transact(opts, "verifyBlocks", _blocks, _proofs)
}

// VerifyBlocks is a paid mutator transaction binding the contract method 0xb77a1ef5.
//
// Solidity: function verifyBlocks(((uint32,uint64,bytes32,uint256,bytes32,bytes32),bytes[])[] _blocks, uint256[] _proofs) returns()
func (_ZecreyLegend *ZecreyLegendSession) VerifyBlocks(_blocks []ZecreyLegendVerifyBlockInfo, _proofs []*big.Int) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.VerifyBlocks(&_ZecreyLegend.TransactOpts, _blocks, _proofs)
}

// VerifyBlocks is a paid mutator transaction binding the contract method 0xb77a1ef5.
//
// Solidity: function verifyBlocks(((uint32,uint64,bytes32,uint256,bytes32,bytes32),bytes[])[] _blocks, uint256[] _proofs) returns()
func (_ZecreyLegend *ZecreyLegendTransactorSession) VerifyBlocks(_blocks []ZecreyLegendVerifyBlockInfo, _proofs []*big.Int) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.VerifyBlocks(&_ZecreyLegend.TransactOpts, _blocks, _proofs)
}

// WithdrawPendingBalance is a paid mutator transaction binding the contract method 0xd514da50.
//
// Solidity: function withdrawPendingBalance(address _owner, address _token, uint128 _amount) returns()
func (_ZecreyLegend *ZecreyLegendTransactor) WithdrawPendingBalance(opts *bind.TransactOpts, _owner common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ZecreyLegend.contract.Transact(opts, "withdrawPendingBalance", _owner, _token, _amount)
}

// WithdrawPendingBalance is a paid mutator transaction binding the contract method 0xd514da50.
//
// Solidity: function withdrawPendingBalance(address _owner, address _token, uint128 _amount) returns()
func (_ZecreyLegend *ZecreyLegendSession) WithdrawPendingBalance(_owner common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.WithdrawPendingBalance(&_ZecreyLegend.TransactOpts, _owner, _token, _amount)
}

// WithdrawPendingBalance is a paid mutator transaction binding the contract method 0xd514da50.
//
// Solidity: function withdrawPendingBalance(address _owner, address _token, uint128 _amount) returns()
func (_ZecreyLegend *ZecreyLegendTransactorSession) WithdrawPendingBalance(_owner common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ZecreyLegend.Contract.WithdrawPendingBalance(&_ZecreyLegend.TransactOpts, _owner, _token, _amount)
}

// ZecreyLegendBlockCommitIterator is returned from FilterBlockCommit and is used to iterate over the raw logs and unpacked data for BlockCommit events raised by the ZecreyLegend contract.
type ZecreyLegendBlockCommitIterator struct {
	Event *ZecreyLegendBlockCommit // Event containing the contract specifics and raw log

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
func (it *ZecreyLegendBlockCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreyLegendBlockCommit)
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
		it.Event = new(ZecreyLegendBlockCommit)
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
func (it *ZecreyLegendBlockCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreyLegendBlockCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreyLegendBlockCommit represents a BlockCommit event raised by the ZecreyLegend contract.
type ZecreyLegendBlockCommit struct {
	BlockNumber uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockCommit is a free log retrieval operation binding the contract event 0x81a92942d0f9c33b897a438384c9c3d88be397776138efa3ba1a4fc8b6268424.
//
// Solidity: event BlockCommit(uint32 blockNumber)
func (_ZecreyLegend *ZecreyLegendFilterer) FilterBlockCommit(opts *bind.FilterOpts) (*ZecreyLegendBlockCommitIterator, error) {

	logs, sub, err := _ZecreyLegend.contract.FilterLogs(opts, "BlockCommit")
	if err != nil {
		return nil, err
	}
	return &ZecreyLegendBlockCommitIterator{contract: _ZecreyLegend.contract, event: "BlockCommit", logs: logs, sub: sub}, nil
}

// WatchBlockCommit is a free log subscription operation binding the contract event 0x81a92942d0f9c33b897a438384c9c3d88be397776138efa3ba1a4fc8b6268424.
//
// Solidity: event BlockCommit(uint32 blockNumber)
func (_ZecreyLegend *ZecreyLegendFilterer) WatchBlockCommit(opts *bind.WatchOpts, sink chan<- *ZecreyLegendBlockCommit) (event.Subscription, error) {

	logs, sub, err := _ZecreyLegend.contract.WatchLogs(opts, "BlockCommit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreyLegendBlockCommit)
				if err := _ZecreyLegend.contract.UnpackLog(event, "BlockCommit", log); err != nil {
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
func (_ZecreyLegend *ZecreyLegendFilterer) ParseBlockCommit(log types.Log) (*ZecreyLegendBlockCommit, error) {
	event := new(ZecreyLegendBlockCommit)
	if err := _ZecreyLegend.contract.UnpackLog(event, "BlockCommit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreyLegendBlockVerificationIterator is returned from FilterBlockVerification and is used to iterate over the raw logs and unpacked data for BlockVerification events raised by the ZecreyLegend contract.
type ZecreyLegendBlockVerificationIterator struct {
	Event *ZecreyLegendBlockVerification // Event containing the contract specifics and raw log

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
func (it *ZecreyLegendBlockVerificationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreyLegendBlockVerification)
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
		it.Event = new(ZecreyLegendBlockVerification)
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
func (it *ZecreyLegendBlockVerificationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreyLegendBlockVerificationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreyLegendBlockVerification represents a BlockVerification event raised by the ZecreyLegend contract.
type ZecreyLegendBlockVerification struct {
	BlockNumber uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockVerification is a free log retrieval operation binding the contract event 0x0cdbd8bd7813095001c5fe7917bd69d834dc01db7c1dfcf52ca135bd20384413.
//
// Solidity: event BlockVerification(uint32 blockNumber)
func (_ZecreyLegend *ZecreyLegendFilterer) FilterBlockVerification(opts *bind.FilterOpts) (*ZecreyLegendBlockVerificationIterator, error) {

	logs, sub, err := _ZecreyLegend.contract.FilterLogs(opts, "BlockVerification")
	if err != nil {
		return nil, err
	}
	return &ZecreyLegendBlockVerificationIterator{contract: _ZecreyLegend.contract, event: "BlockVerification", logs: logs, sub: sub}, nil
}

// WatchBlockVerification is a free log subscription operation binding the contract event 0x0cdbd8bd7813095001c5fe7917bd69d834dc01db7c1dfcf52ca135bd20384413.
//
// Solidity: event BlockVerification(uint32 blockNumber)
func (_ZecreyLegend *ZecreyLegendFilterer) WatchBlockVerification(opts *bind.WatchOpts, sink chan<- *ZecreyLegendBlockVerification) (event.Subscription, error) {

	logs, sub, err := _ZecreyLegend.contract.WatchLogs(opts, "BlockVerification")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreyLegendBlockVerification)
				if err := _ZecreyLegend.contract.UnpackLog(event, "BlockVerification", log); err != nil {
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
func (_ZecreyLegend *ZecreyLegendFilterer) ParseBlockVerification(log types.Log) (*ZecreyLegendBlockVerification, error) {
	event := new(ZecreyLegendBlockVerification)
	if err := _ZecreyLegend.contract.UnpackLog(event, "BlockVerification", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreyLegendBlocksRevertIterator is returned from FilterBlocksRevert and is used to iterate over the raw logs and unpacked data for BlocksRevert events raised by the ZecreyLegend contract.
type ZecreyLegendBlocksRevertIterator struct {
	Event *ZecreyLegendBlocksRevert // Event containing the contract specifics and raw log

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
func (it *ZecreyLegendBlocksRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreyLegendBlocksRevert)
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
		it.Event = new(ZecreyLegendBlocksRevert)
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
func (it *ZecreyLegendBlocksRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreyLegendBlocksRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreyLegendBlocksRevert represents a BlocksRevert event raised by the ZecreyLegend contract.
type ZecreyLegendBlocksRevert struct {
	TotalBlocksVerified  uint32
	TotalBlocksCommitted uint32
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterBlocksRevert is a free log retrieval operation binding the contract event 0x6f3a8259cce1ea2680115053d21c971aa1764295a45850f520525f2bfdf3c9d3.
//
// Solidity: event BlocksRevert(uint32 totalBlocksVerified, uint32 totalBlocksCommitted)
func (_ZecreyLegend *ZecreyLegendFilterer) FilterBlocksRevert(opts *bind.FilterOpts) (*ZecreyLegendBlocksRevertIterator, error) {

	logs, sub, err := _ZecreyLegend.contract.FilterLogs(opts, "BlocksRevert")
	if err != nil {
		return nil, err
	}
	return &ZecreyLegendBlocksRevertIterator{contract: _ZecreyLegend.contract, event: "BlocksRevert", logs: logs, sub: sub}, nil
}

// WatchBlocksRevert is a free log subscription operation binding the contract event 0x6f3a8259cce1ea2680115053d21c971aa1764295a45850f520525f2bfdf3c9d3.
//
// Solidity: event BlocksRevert(uint32 totalBlocksVerified, uint32 totalBlocksCommitted)
func (_ZecreyLegend *ZecreyLegendFilterer) WatchBlocksRevert(opts *bind.WatchOpts, sink chan<- *ZecreyLegendBlocksRevert) (event.Subscription, error) {

	logs, sub, err := _ZecreyLegend.contract.WatchLogs(opts, "BlocksRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreyLegendBlocksRevert)
				if err := _ZecreyLegend.contract.UnpackLog(event, "BlocksRevert", log); err != nil {
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
func (_ZecreyLegend *ZecreyLegendFilterer) ParseBlocksRevert(log types.Log) (*ZecreyLegendBlocksRevert, error) {
	event := new(ZecreyLegendBlocksRevert)
	if err := _ZecreyLegend.contract.UnpackLog(event, "BlocksRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreyLegendCreateTokenPairIterator is returned from FilterCreateTokenPair and is used to iterate over the raw logs and unpacked data for CreateTokenPair events raised by the ZecreyLegend contract.
type ZecreyLegendCreateTokenPairIterator struct {
	Event *ZecreyLegendCreateTokenPair // Event containing the contract specifics and raw log

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
func (it *ZecreyLegendCreateTokenPairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreyLegendCreateTokenPair)
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
		it.Event = new(ZecreyLegendCreateTokenPair)
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
func (it *ZecreyLegendCreateTokenPairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreyLegendCreateTokenPairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreyLegendCreateTokenPair represents a CreateTokenPair event raised by the ZecreyLegend contract.
type ZecreyLegendCreateTokenPair struct {
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
func (_ZecreyLegend *ZecreyLegendFilterer) FilterCreateTokenPair(opts *bind.FilterOpts) (*ZecreyLegendCreateTokenPairIterator, error) {

	logs, sub, err := _ZecreyLegend.contract.FilterLogs(opts, "CreateTokenPair")
	if err != nil {
		return nil, err
	}
	return &ZecreyLegendCreateTokenPairIterator{contract: _ZecreyLegend.contract, event: "CreateTokenPair", logs: logs, sub: sub}, nil
}

// WatchCreateTokenPair is a free log subscription operation binding the contract event 0x82e21fd4ced46b510ee1bf5a34eea3dc34aa1f5460f53a25dd22828cde30e087.
//
// Solidity: event CreateTokenPair(uint16 pairIndex, uint16 asset0Id, uint16 asset1Id, uint16 feeRate, uint32 treasuryAccountIndex, uint16 treasuryRate)
func (_ZecreyLegend *ZecreyLegendFilterer) WatchCreateTokenPair(opts *bind.WatchOpts, sink chan<- *ZecreyLegendCreateTokenPair) (event.Subscription, error) {

	logs, sub, err := _ZecreyLegend.contract.WatchLogs(opts, "CreateTokenPair")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreyLegendCreateTokenPair)
				if err := _ZecreyLegend.contract.UnpackLog(event, "CreateTokenPair", log); err != nil {
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
func (_ZecreyLegend *ZecreyLegendFilterer) ParseCreateTokenPair(log types.Log) (*ZecreyLegendCreateTokenPair, error) {
	event := new(ZecreyLegendCreateTokenPair)
	if err := _ZecreyLegend.contract.UnpackLog(event, "CreateTokenPair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreyLegendDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ZecreyLegend contract.
type ZecreyLegendDepositIterator struct {
	Event *ZecreyLegendDeposit // Event containing the contract specifics and raw log

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
func (it *ZecreyLegendDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreyLegendDeposit)
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
		it.Event = new(ZecreyLegendDeposit)
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
func (it *ZecreyLegendDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreyLegendDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreyLegendDeposit represents a Deposit event raised by the ZecreyLegend contract.
type ZecreyLegendDeposit struct {
	AssetId     uint16
	AccountName [32]byte
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xaa46d46658f805449ac7eaaf11d3eaebb6f508930128d276dfefcc8dc02a13c7.
//
// Solidity: event Deposit(uint16 assetId, bytes32 accountName, uint128 amount)
func (_ZecreyLegend *ZecreyLegendFilterer) FilterDeposit(opts *bind.FilterOpts) (*ZecreyLegendDepositIterator, error) {

	logs, sub, err := _ZecreyLegend.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &ZecreyLegendDepositIterator{contract: _ZecreyLegend.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xaa46d46658f805449ac7eaaf11d3eaebb6f508930128d276dfefcc8dc02a13c7.
//
// Solidity: event Deposit(uint16 assetId, bytes32 accountName, uint128 amount)
func (_ZecreyLegend *ZecreyLegendFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ZecreyLegendDeposit) (event.Subscription, error) {

	logs, sub, err := _ZecreyLegend.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreyLegendDeposit)
				if err := _ZecreyLegend.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_ZecreyLegend *ZecreyLegendFilterer) ParseDeposit(log types.Log) (*ZecreyLegendDeposit, error) {
	event := new(ZecreyLegendDeposit)
	if err := _ZecreyLegend.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreyLegendDepositCommitIterator is returned from FilterDepositCommit and is used to iterate over the raw logs and unpacked data for DepositCommit events raised by the ZecreyLegend contract.
type ZecreyLegendDepositCommitIterator struct {
	Event *ZecreyLegendDepositCommit // Event containing the contract specifics and raw log

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
func (it *ZecreyLegendDepositCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreyLegendDepositCommit)
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
		it.Event = new(ZecreyLegendDepositCommit)
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
func (it *ZecreyLegendDepositCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreyLegendDepositCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreyLegendDepositCommit represents a DepositCommit event raised by the ZecreyLegend contract.
type ZecreyLegendDepositCommit struct {
	ZecreyBlockNumber uint32
	AccountIndex      uint32
	AccountName       [32]byte
	AssetId           uint16
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDepositCommit is a free log retrieval operation binding the contract event 0x9b3e77f05cee61edb7f3e0e85d71a29e89cb18cb4236a8b8cc3e4d66640d8f71.
//
// Solidity: event DepositCommit(uint32 indexed zecreyBlockNumber, uint32 indexed accountIndex, bytes32 accountName, uint16 indexed assetId, uint128 amount)
func (_ZecreyLegend *ZecreyLegendFilterer) FilterDepositCommit(opts *bind.FilterOpts, zecreyBlockNumber []uint32, accountIndex []uint32, assetId []uint16) (*ZecreyLegendDepositCommitIterator, error) {

	var zecreyBlockNumberRule []interface{}
	for _, zecreyBlockNumberItem := range zecreyBlockNumber {
		zecreyBlockNumberRule = append(zecreyBlockNumberRule, zecreyBlockNumberItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _ZecreyLegend.contract.FilterLogs(opts, "DepositCommit", zecreyBlockNumberRule, accountIndexRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &ZecreyLegendDepositCommitIterator{contract: _ZecreyLegend.contract, event: "DepositCommit", logs: logs, sub: sub}, nil
}

// WatchDepositCommit is a free log subscription operation binding the contract event 0x9b3e77f05cee61edb7f3e0e85d71a29e89cb18cb4236a8b8cc3e4d66640d8f71.
//
// Solidity: event DepositCommit(uint32 indexed zecreyBlockNumber, uint32 indexed accountIndex, bytes32 accountName, uint16 indexed assetId, uint128 amount)
func (_ZecreyLegend *ZecreyLegendFilterer) WatchDepositCommit(opts *bind.WatchOpts, sink chan<- *ZecreyLegendDepositCommit, zecreyBlockNumber []uint32, accountIndex []uint32, assetId []uint16) (event.Subscription, error) {

	var zecreyBlockNumberRule []interface{}
	for _, zecreyBlockNumberItem := range zecreyBlockNumber {
		zecreyBlockNumberRule = append(zecreyBlockNumberRule, zecreyBlockNumberItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _ZecreyLegend.contract.WatchLogs(opts, "DepositCommit", zecreyBlockNumberRule, accountIndexRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreyLegendDepositCommit)
				if err := _ZecreyLegend.contract.UnpackLog(event, "DepositCommit", log); err != nil {
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
// Solidity: event DepositCommit(uint32 indexed zecreyBlockNumber, uint32 indexed accountIndex, bytes32 accountName, uint16 indexed assetId, uint128 amount)
func (_ZecreyLegend *ZecreyLegendFilterer) ParseDepositCommit(log types.Log) (*ZecreyLegendDepositCommit, error) {
	event := new(ZecreyLegendDepositCommit)
	if err := _ZecreyLegend.contract.UnpackLog(event, "DepositCommit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreyLegendDepositNftIterator is returned from FilterDepositNft and is used to iterate over the raw logs and unpacked data for DepositNft events raised by the ZecreyLegend contract.
type ZecreyLegendDepositNftIterator struct {
	Event *ZecreyLegendDepositNft // Event containing the contract specifics and raw log

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
func (it *ZecreyLegendDepositNftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreyLegendDepositNft)
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
		it.Event = new(ZecreyLegendDepositNft)
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
func (it *ZecreyLegendDepositNftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreyLegendDepositNftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreyLegendDepositNft represents a DepositNft event raised by the ZecreyLegend contract.
type ZecreyLegendDepositNft struct {
	AccountNameHash     [32]byte
	TokenAddress        common.Address
	NftTokenId          *big.Int
	CreatorTreasuryRate uint16
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDepositNft is a free log retrieval operation binding the contract event 0x5e7ed95c5de0831db760fa0f4bddf5b3752edc7b22814e7cfaceddf7fb9f9ed6.
//
// Solidity: event DepositNft(bytes32 accountNameHash, address tokenAddress, uint256 nftTokenId, uint16 creatorTreasuryRate)
func (_ZecreyLegend *ZecreyLegendFilterer) FilterDepositNft(opts *bind.FilterOpts) (*ZecreyLegendDepositNftIterator, error) {

	logs, sub, err := _ZecreyLegend.contract.FilterLogs(opts, "DepositNft")
	if err != nil {
		return nil, err
	}
	return &ZecreyLegendDepositNftIterator{contract: _ZecreyLegend.contract, event: "DepositNft", logs: logs, sub: sub}, nil
}

// WatchDepositNft is a free log subscription operation binding the contract event 0x5e7ed95c5de0831db760fa0f4bddf5b3752edc7b22814e7cfaceddf7fb9f9ed6.
//
// Solidity: event DepositNft(bytes32 accountNameHash, address tokenAddress, uint256 nftTokenId, uint16 creatorTreasuryRate)
func (_ZecreyLegend *ZecreyLegendFilterer) WatchDepositNft(opts *bind.WatchOpts, sink chan<- *ZecreyLegendDepositNft) (event.Subscription, error) {

	logs, sub, err := _ZecreyLegend.contract.WatchLogs(opts, "DepositNft")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreyLegendDepositNft)
				if err := _ZecreyLegend.contract.UnpackLog(event, "DepositNft", log); err != nil {
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

// ParseDepositNft is a log parse operation binding the contract event 0x5e7ed95c5de0831db760fa0f4bddf5b3752edc7b22814e7cfaceddf7fb9f9ed6.
//
// Solidity: event DepositNft(bytes32 accountNameHash, address tokenAddress, uint256 nftTokenId, uint16 creatorTreasuryRate)
func (_ZecreyLegend *ZecreyLegendFilterer) ParseDepositNft(log types.Log) (*ZecreyLegendDepositNft, error) {
	event := new(ZecreyLegendDepositNft)
	if err := _ZecreyLegend.contract.UnpackLog(event, "DepositNft", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreyLegendDesertModeIterator is returned from FilterDesertMode and is used to iterate over the raw logs and unpacked data for DesertMode events raised by the ZecreyLegend contract.
type ZecreyLegendDesertModeIterator struct {
	Event *ZecreyLegendDesertMode // Event containing the contract specifics and raw log

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
func (it *ZecreyLegendDesertModeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreyLegendDesertMode)
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
		it.Event = new(ZecreyLegendDesertMode)
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
func (it *ZecreyLegendDesertModeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreyLegendDesertModeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreyLegendDesertMode represents a DesertMode event raised by the ZecreyLegend contract.
type ZecreyLegendDesertMode struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDesertMode is a free log retrieval operation binding the contract event 0x9f7e400a81dddbf1c18b1c37f82aa303d166295ca4b577eb2a7c23d4b704ba89.
//
// Solidity: event DesertMode()
func (_ZecreyLegend *ZecreyLegendFilterer) FilterDesertMode(opts *bind.FilterOpts) (*ZecreyLegendDesertModeIterator, error) {

	logs, sub, err := _ZecreyLegend.contract.FilterLogs(opts, "DesertMode")
	if err != nil {
		return nil, err
	}
	return &ZecreyLegendDesertModeIterator{contract: _ZecreyLegend.contract, event: "DesertMode", logs: logs, sub: sub}, nil
}

// WatchDesertMode is a free log subscription operation binding the contract event 0x9f7e400a81dddbf1c18b1c37f82aa303d166295ca4b577eb2a7c23d4b704ba89.
//
// Solidity: event DesertMode()
func (_ZecreyLegend *ZecreyLegendFilterer) WatchDesertMode(opts *bind.WatchOpts, sink chan<- *ZecreyLegendDesertMode) (event.Subscription, error) {

	logs, sub, err := _ZecreyLegend.contract.WatchLogs(opts, "DesertMode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreyLegendDesertMode)
				if err := _ZecreyLegend.contract.UnpackLog(event, "DesertMode", log); err != nil {
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
func (_ZecreyLegend *ZecreyLegendFilterer) ParseDesertMode(log types.Log) (*ZecreyLegendDesertMode, error) {
	event := new(ZecreyLegendDesertMode)
	if err := _ZecreyLegend.contract.UnpackLog(event, "DesertMode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreyLegendFullExitCommitIterator is returned from FilterFullExitCommit and is used to iterate over the raw logs and unpacked data for FullExitCommit events raised by the ZecreyLegend contract.
type ZecreyLegendFullExitCommitIterator struct {
	Event *ZecreyLegendFullExitCommit // Event containing the contract specifics and raw log

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
func (it *ZecreyLegendFullExitCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreyLegendFullExitCommit)
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
		it.Event = new(ZecreyLegendFullExitCommit)
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
func (it *ZecreyLegendFullExitCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreyLegendFullExitCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreyLegendFullExitCommit represents a FullExitCommit event raised by the ZecreyLegend contract.
type ZecreyLegendFullExitCommit struct {
	ZecreyBlockId uint32
	AccountId     uint32
	Owner         common.Address
	TokenId       uint16
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFullExitCommit is a free log retrieval operation binding the contract event 0x66fc63d751ecbefca61d4e2e7c534e4f29c61aed8ece23ed635277a7ea6f9bc4.
//
// Solidity: event FullExitCommit(uint32 indexed zecreyBlockId, uint32 indexed accountId, address owner, uint16 indexed tokenId, uint128 amount)
func (_ZecreyLegend *ZecreyLegendFilterer) FilterFullExitCommit(opts *bind.FilterOpts, zecreyBlockId []uint32, accountId []uint32, tokenId []uint16) (*ZecreyLegendFullExitCommitIterator, error) {

	var zecreyBlockIdRule []interface{}
	for _, zecreyBlockIdItem := range zecreyBlockId {
		zecreyBlockIdRule = append(zecreyBlockIdRule, zecreyBlockIdItem)
	}
	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ZecreyLegend.contract.FilterLogs(opts, "FullExitCommit", zecreyBlockIdRule, accountIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ZecreyLegendFullExitCommitIterator{contract: _ZecreyLegend.contract, event: "FullExitCommit", logs: logs, sub: sub}, nil
}

// WatchFullExitCommit is a free log subscription operation binding the contract event 0x66fc63d751ecbefca61d4e2e7c534e4f29c61aed8ece23ed635277a7ea6f9bc4.
//
// Solidity: event FullExitCommit(uint32 indexed zecreyBlockId, uint32 indexed accountId, address owner, uint16 indexed tokenId, uint128 amount)
func (_ZecreyLegend *ZecreyLegendFilterer) WatchFullExitCommit(opts *bind.WatchOpts, sink chan<- *ZecreyLegendFullExitCommit, zecreyBlockId []uint32, accountId []uint32, tokenId []uint16) (event.Subscription, error) {

	var zecreyBlockIdRule []interface{}
	for _, zecreyBlockIdItem := range zecreyBlockId {
		zecreyBlockIdRule = append(zecreyBlockIdRule, zecreyBlockIdItem)
	}
	var accountIdRule []interface{}
	for _, accountIdItem := range accountId {
		accountIdRule = append(accountIdRule, accountIdItem)
	}

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ZecreyLegend.contract.WatchLogs(opts, "FullExitCommit", zecreyBlockIdRule, accountIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreyLegendFullExitCommit)
				if err := _ZecreyLegend.contract.UnpackLog(event, "FullExitCommit", log); err != nil {
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
// Solidity: event FullExitCommit(uint32 indexed zecreyBlockId, uint32 indexed accountId, address owner, uint16 indexed tokenId, uint128 amount)
func (_ZecreyLegend *ZecreyLegendFilterer) ParseFullExitCommit(log types.Log) (*ZecreyLegendFullExitCommit, error) {
	event := new(ZecreyLegendFullExitCommit)
	if err := _ZecreyLegend.contract.UnpackLog(event, "FullExitCommit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreyLegendNewPriorityRequestIterator is returned from FilterNewPriorityRequest and is used to iterate over the raw logs and unpacked data for NewPriorityRequest events raised by the ZecreyLegend contract.
type ZecreyLegendNewPriorityRequestIterator struct {
	Event *ZecreyLegendNewPriorityRequest // Event containing the contract specifics and raw log

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
func (it *ZecreyLegendNewPriorityRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreyLegendNewPriorityRequest)
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
		it.Event = new(ZecreyLegendNewPriorityRequest)
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
func (it *ZecreyLegendNewPriorityRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreyLegendNewPriorityRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreyLegendNewPriorityRequest represents a NewPriorityRequest event raised by the ZecreyLegend contract.
type ZecreyLegendNewPriorityRequest struct {
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
func (_ZecreyLegend *ZecreyLegendFilterer) FilterNewPriorityRequest(opts *bind.FilterOpts) (*ZecreyLegendNewPriorityRequestIterator, error) {

	logs, sub, err := _ZecreyLegend.contract.FilterLogs(opts, "NewPriorityRequest")
	if err != nil {
		return nil, err
	}
	return &ZecreyLegendNewPriorityRequestIterator{contract: _ZecreyLegend.contract, event: "NewPriorityRequest", logs: logs, sub: sub}, nil
}

// WatchNewPriorityRequest is a free log subscription operation binding the contract event 0xd0943372c08b438a88d4b39d77216901079eda9ca59d45349841c099083b6830.
//
// Solidity: event NewPriorityRequest(address sender, uint64 serialId, uint8 txType, bytes pubData, uint256 expirationBlock)
func (_ZecreyLegend *ZecreyLegendFilterer) WatchNewPriorityRequest(opts *bind.WatchOpts, sink chan<- *ZecreyLegendNewPriorityRequest) (event.Subscription, error) {

	logs, sub, err := _ZecreyLegend.contract.WatchLogs(opts, "NewPriorityRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreyLegendNewPriorityRequest)
				if err := _ZecreyLegend.contract.UnpackLog(event, "NewPriorityRequest", log); err != nil {
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
func (_ZecreyLegend *ZecreyLegendFilterer) ParseNewPriorityRequest(log types.Log) (*ZecreyLegendNewPriorityRequest, error) {
	event := new(ZecreyLegendNewPriorityRequest)
	if err := _ZecreyLegend.contract.UnpackLog(event, "NewPriorityRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreyLegendNoticePeriodChangeIterator is returned from FilterNoticePeriodChange and is used to iterate over the raw logs and unpacked data for NoticePeriodChange events raised by the ZecreyLegend contract.
type ZecreyLegendNoticePeriodChangeIterator struct {
	Event *ZecreyLegendNoticePeriodChange // Event containing the contract specifics and raw log

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
func (it *ZecreyLegendNoticePeriodChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreyLegendNoticePeriodChange)
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
		it.Event = new(ZecreyLegendNoticePeriodChange)
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
func (it *ZecreyLegendNoticePeriodChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreyLegendNoticePeriodChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreyLegendNoticePeriodChange represents a NoticePeriodChange event raised by the ZecreyLegend contract.
type ZecreyLegendNoticePeriodChange struct {
	NewNoticePeriod *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNoticePeriodChange is a free log retrieval operation binding the contract event 0xf2b18f8abbd8a0d0c1fb8245146eedf5304887b12f6395b548ca238e054a1483.
//
// Solidity: event NoticePeriodChange(uint256 newNoticePeriod)
func (_ZecreyLegend *ZecreyLegendFilterer) FilterNoticePeriodChange(opts *bind.FilterOpts) (*ZecreyLegendNoticePeriodChangeIterator, error) {

	logs, sub, err := _ZecreyLegend.contract.FilterLogs(opts, "NoticePeriodChange")
	if err != nil {
		return nil, err
	}
	return &ZecreyLegendNoticePeriodChangeIterator{contract: _ZecreyLegend.contract, event: "NoticePeriodChange", logs: logs, sub: sub}, nil
}

// WatchNoticePeriodChange is a free log subscription operation binding the contract event 0xf2b18f8abbd8a0d0c1fb8245146eedf5304887b12f6395b548ca238e054a1483.
//
// Solidity: event NoticePeriodChange(uint256 newNoticePeriod)
func (_ZecreyLegend *ZecreyLegendFilterer) WatchNoticePeriodChange(opts *bind.WatchOpts, sink chan<- *ZecreyLegendNoticePeriodChange) (event.Subscription, error) {

	logs, sub, err := _ZecreyLegend.contract.WatchLogs(opts, "NoticePeriodChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreyLegendNoticePeriodChange)
				if err := _ZecreyLegend.contract.UnpackLog(event, "NoticePeriodChange", log); err != nil {
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
func (_ZecreyLegend *ZecreyLegendFilterer) ParseNoticePeriodChange(log types.Log) (*ZecreyLegendNoticePeriodChange, error) {
	event := new(ZecreyLegendNoticePeriodChange)
	if err := _ZecreyLegend.contract.UnpackLog(event, "NoticePeriodChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreyLegendRegisterZNSIterator is returned from FilterRegisterZNS and is used to iterate over the raw logs and unpacked data for RegisterZNS events raised by the ZecreyLegend contract.
type ZecreyLegendRegisterZNSIterator struct {
	Event *ZecreyLegendRegisterZNS // Event containing the contract specifics and raw log

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
func (it *ZecreyLegendRegisterZNSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreyLegendRegisterZNS)
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
		it.Event = new(ZecreyLegendRegisterZNS)
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
func (it *ZecreyLegendRegisterZNSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreyLegendRegisterZNSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreyLegendRegisterZNS represents a RegisterZNS event raised by the ZecreyLegend contract.
type ZecreyLegendRegisterZNS struct {
	Name         string
	NameHash     [32]byte
	Owner        common.Address
	ZecreyPubKey [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRegisterZNS is a free log retrieval operation binding the contract event 0x8f3bddba5ebbeea69f799226d7c87c68ca4c0e45fe9131cf1c43fa29239be932.
//
// Solidity: event RegisterZNS(string name, bytes32 nameHash, address owner, bytes32 zecreyPubKey)
func (_ZecreyLegend *ZecreyLegendFilterer) FilterRegisterZNS(opts *bind.FilterOpts) (*ZecreyLegendRegisterZNSIterator, error) {

	logs, sub, err := _ZecreyLegend.contract.FilterLogs(opts, "RegisterZNS")
	if err != nil {
		return nil, err
	}
	return &ZecreyLegendRegisterZNSIterator{contract: _ZecreyLegend.contract, event: "RegisterZNS", logs: logs, sub: sub}, nil
}

// WatchRegisterZNS is a free log subscription operation binding the contract event 0x8f3bddba5ebbeea69f799226d7c87c68ca4c0e45fe9131cf1c43fa29239be932.
//
// Solidity: event RegisterZNS(string name, bytes32 nameHash, address owner, bytes32 zecreyPubKey)
func (_ZecreyLegend *ZecreyLegendFilterer) WatchRegisterZNS(opts *bind.WatchOpts, sink chan<- *ZecreyLegendRegisterZNS) (event.Subscription, error) {

	logs, sub, err := _ZecreyLegend.contract.WatchLogs(opts, "RegisterZNS")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreyLegendRegisterZNS)
				if err := _ZecreyLegend.contract.UnpackLog(event, "RegisterZNS", log); err != nil {
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

// ParseRegisterZNS is a log parse operation binding the contract event 0x8f3bddba5ebbeea69f799226d7c87c68ca4c0e45fe9131cf1c43fa29239be932.
//
// Solidity: event RegisterZNS(string name, bytes32 nameHash, address owner, bytes32 zecreyPubKey)
func (_ZecreyLegend *ZecreyLegendFilterer) ParseRegisterZNS(log types.Log) (*ZecreyLegendRegisterZNS, error) {
	event := new(ZecreyLegendRegisterZNS)
	if err := _ZecreyLegend.contract.UnpackLog(event, "RegisterZNS", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreyLegendUpdateTokenPairIterator is returned from FilterUpdateTokenPair and is used to iterate over the raw logs and unpacked data for UpdateTokenPair events raised by the ZecreyLegend contract.
type ZecreyLegendUpdateTokenPairIterator struct {
	Event *ZecreyLegendUpdateTokenPair // Event containing the contract specifics and raw log

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
func (it *ZecreyLegendUpdateTokenPairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreyLegendUpdateTokenPair)
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
		it.Event = new(ZecreyLegendUpdateTokenPair)
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
func (it *ZecreyLegendUpdateTokenPairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreyLegendUpdateTokenPairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreyLegendUpdateTokenPair represents a UpdateTokenPair event raised by the ZecreyLegend contract.
type ZecreyLegendUpdateTokenPair struct {
	PairIndex            uint16
	FeeRate              uint16
	TreasuryAccountIndex uint32
	TreasuryRate         uint16
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateTokenPair is a free log retrieval operation binding the contract event 0x8727712867f019c156b0c25f308498ea1524b635902129b0acd3e9ff2c96f1c8.
//
// Solidity: event UpdateTokenPair(uint16 pairIndex, uint16 feeRate, uint32 treasuryAccountIndex, uint16 treasuryRate)
func (_ZecreyLegend *ZecreyLegendFilterer) FilterUpdateTokenPair(opts *bind.FilterOpts) (*ZecreyLegendUpdateTokenPairIterator, error) {

	logs, sub, err := _ZecreyLegend.contract.FilterLogs(opts, "UpdateTokenPair")
	if err != nil {
		return nil, err
	}
	return &ZecreyLegendUpdateTokenPairIterator{contract: _ZecreyLegend.contract, event: "UpdateTokenPair", logs: logs, sub: sub}, nil
}

// WatchUpdateTokenPair is a free log subscription operation binding the contract event 0x8727712867f019c156b0c25f308498ea1524b635902129b0acd3e9ff2c96f1c8.
//
// Solidity: event UpdateTokenPair(uint16 pairIndex, uint16 feeRate, uint32 treasuryAccountIndex, uint16 treasuryRate)
func (_ZecreyLegend *ZecreyLegendFilterer) WatchUpdateTokenPair(opts *bind.WatchOpts, sink chan<- *ZecreyLegendUpdateTokenPair) (event.Subscription, error) {

	logs, sub, err := _ZecreyLegend.contract.WatchLogs(opts, "UpdateTokenPair")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreyLegendUpdateTokenPair)
				if err := _ZecreyLegend.contract.UnpackLog(event, "UpdateTokenPair", log); err != nil {
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
func (_ZecreyLegend *ZecreyLegendFilterer) ParseUpdateTokenPair(log types.Log) (*ZecreyLegendUpdateTokenPair, error) {
	event := new(ZecreyLegendUpdateTokenPair)
	if err := _ZecreyLegend.contract.UnpackLog(event, "UpdateTokenPair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreyLegendWithdrawNftIterator is returned from FilterWithdrawNft and is used to iterate over the raw logs and unpacked data for WithdrawNft events raised by the ZecreyLegend contract.
type ZecreyLegendWithdrawNftIterator struct {
	Event *ZecreyLegendWithdrawNft // Event containing the contract specifics and raw log

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
func (it *ZecreyLegendWithdrawNftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreyLegendWithdrawNft)
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
		it.Event = new(ZecreyLegendWithdrawNft)
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
func (it *ZecreyLegendWithdrawNftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreyLegendWithdrawNftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreyLegendWithdrawNft represents a WithdrawNft event raised by the ZecreyLegend contract.
type ZecreyLegendWithdrawNft struct {
	AccountNameHash [32]byte
	TokenAddress    common.Address
	ToAddress       common.Address
	ProxyAddress    common.Address
	NftTokenId      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawNft is a free log retrieval operation binding the contract event 0x49f385f20711f01d7f0a1465c0fccbfeff7c753dda812cb3d7c9b6252f6cd543.
//
// Solidity: event WithdrawNft(bytes32 accountNameHash, address tokenAddress, address toAddress, address proxyAddress, uint256 nftTokenId)
func (_ZecreyLegend *ZecreyLegendFilterer) FilterWithdrawNft(opts *bind.FilterOpts) (*ZecreyLegendWithdrawNftIterator, error) {

	logs, sub, err := _ZecreyLegend.contract.FilterLogs(opts, "WithdrawNft")
	if err != nil {
		return nil, err
	}
	return &ZecreyLegendWithdrawNftIterator{contract: _ZecreyLegend.contract, event: "WithdrawNft", logs: logs, sub: sub}, nil
}

// WatchWithdrawNft is a free log subscription operation binding the contract event 0x49f385f20711f01d7f0a1465c0fccbfeff7c753dda812cb3d7c9b6252f6cd543.
//
// Solidity: event WithdrawNft(bytes32 accountNameHash, address tokenAddress, address toAddress, address proxyAddress, uint256 nftTokenId)
func (_ZecreyLegend *ZecreyLegendFilterer) WatchWithdrawNft(opts *bind.WatchOpts, sink chan<- *ZecreyLegendWithdrawNft) (event.Subscription, error) {

	logs, sub, err := _ZecreyLegend.contract.WatchLogs(opts, "WithdrawNft")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreyLegendWithdrawNft)
				if err := _ZecreyLegend.contract.UnpackLog(event, "WithdrawNft", log); err != nil {
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

// ParseWithdrawNft is a log parse operation binding the contract event 0x49f385f20711f01d7f0a1465c0fccbfeff7c753dda812cb3d7c9b6252f6cd543.
//
// Solidity: event WithdrawNft(bytes32 accountNameHash, address tokenAddress, address toAddress, address proxyAddress, uint256 nftTokenId)
func (_ZecreyLegend *ZecreyLegendFilterer) ParseWithdrawNft(log types.Log) (*ZecreyLegendWithdrawNft, error) {
	event := new(ZecreyLegendWithdrawNft)
	if err := _ZecreyLegend.contract.UnpackLog(event, "WithdrawNft", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreyLegendWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the ZecreyLegend contract.
type ZecreyLegendWithdrawalIterator struct {
	Event *ZecreyLegendWithdrawal // Event containing the contract specifics and raw log

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
func (it *ZecreyLegendWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreyLegendWithdrawal)
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
		it.Event = new(ZecreyLegendWithdrawal)
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
func (it *ZecreyLegendWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreyLegendWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreyLegendWithdrawal represents a Withdrawal event raised by the ZecreyLegend contract.
type ZecreyLegendWithdrawal struct {
	AssetId uint16
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0xf4bf32c167ee6e782944cd1db8174729b46adcd3bc732e282cc4a80793933154.
//
// Solidity: event Withdrawal(uint16 assetId, uint128 amount)
func (_ZecreyLegend *ZecreyLegendFilterer) FilterWithdrawal(opts *bind.FilterOpts) (*ZecreyLegendWithdrawalIterator, error) {

	logs, sub, err := _ZecreyLegend.contract.FilterLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return &ZecreyLegendWithdrawalIterator{contract: _ZecreyLegend.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0xf4bf32c167ee6e782944cd1db8174729b46adcd3bc732e282cc4a80793933154.
//
// Solidity: event Withdrawal(uint16 assetId, uint128 amount)
func (_ZecreyLegend *ZecreyLegendFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *ZecreyLegendWithdrawal) (event.Subscription, error) {

	logs, sub, err := _ZecreyLegend.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreyLegendWithdrawal)
				if err := _ZecreyLegend.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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
func (_ZecreyLegend *ZecreyLegendFilterer) ParseWithdrawal(log types.Log) (*ZecreyLegendWithdrawal, error) {
	event := new(ZecreyLegendWithdrawal)
	if err := _ZecreyLegend.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreyLegendWithdrawalPendingIterator is returned from FilterWithdrawalPending and is used to iterate over the raw logs and unpacked data for WithdrawalPending events raised by the ZecreyLegend contract.
type ZecreyLegendWithdrawalPendingIterator struct {
	Event *ZecreyLegendWithdrawalPending // Event containing the contract specifics and raw log

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
func (it *ZecreyLegendWithdrawalPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreyLegendWithdrawalPending)
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
		it.Event = new(ZecreyLegendWithdrawalPending)
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
func (it *ZecreyLegendWithdrawalPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreyLegendWithdrawalPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreyLegendWithdrawalPending represents a WithdrawalPending event raised by the ZecreyLegend contract.
type ZecreyLegendWithdrawalPending struct {
	AssetId uint16
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalPending is a free log retrieval operation binding the contract event 0xd19cf67bbb6c320849f41b650b1179fb06a3f104451c75109c3b006a385c1688.
//
// Solidity: event WithdrawalPending(uint16 assetId, uint128 amount)
func (_ZecreyLegend *ZecreyLegendFilterer) FilterWithdrawalPending(opts *bind.FilterOpts) (*ZecreyLegendWithdrawalPendingIterator, error) {

	logs, sub, err := _ZecreyLegend.contract.FilterLogs(opts, "WithdrawalPending")
	if err != nil {
		return nil, err
	}
	return &ZecreyLegendWithdrawalPendingIterator{contract: _ZecreyLegend.contract, event: "WithdrawalPending", logs: logs, sub: sub}, nil
}

// WatchWithdrawalPending is a free log subscription operation binding the contract event 0xd19cf67bbb6c320849f41b650b1179fb06a3f104451c75109c3b006a385c1688.
//
// Solidity: event WithdrawalPending(uint16 assetId, uint128 amount)
func (_ZecreyLegend *ZecreyLegendFilterer) WatchWithdrawalPending(opts *bind.WatchOpts, sink chan<- *ZecreyLegendWithdrawalPending) (event.Subscription, error) {

	logs, sub, err := _ZecreyLegend.contract.WatchLogs(opts, "WithdrawalPending")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreyLegendWithdrawalPending)
				if err := _ZecreyLegend.contract.UnpackLog(event, "WithdrawalPending", log); err != nil {
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

// ParseWithdrawalPending is a log parse operation binding the contract event 0xd19cf67bbb6c320849f41b650b1179fb06a3f104451c75109c3b006a385c1688.
//
// Solidity: event WithdrawalPending(uint16 assetId, uint128 amount)
func (_ZecreyLegend *ZecreyLegendFilterer) ParseWithdrawalPending(log types.Log) (*ZecreyLegendWithdrawalPending, error) {
	event := new(ZecreyLegendWithdrawalPending)
	if err := _ZecreyLegend.contract.UnpackLog(event, "WithdrawalPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
