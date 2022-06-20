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

// GovernanceMetaData contains all meta data concerning the Governance contract.
var GovernanceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"AssetPausedUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"assetId\",\"type\":\"uint16\"}],\"name\":\"NewAsset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAssetGovernance\",\"name\":\"newAssetGovernance\",\"type\":\"address\"}],\"name\":\"NewAssetGovernance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newGovernor\",\"type\":\"address\"}],\"name\":\"NewGovernor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"ValidatorStatusUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_ACCOUNT_INDEX\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_AMOUNT_OF_REGISTERED_ASSETS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_DEPOSIT_AMOUNT\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FUNGIBLE_ASSET_ID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_NFT_INDEX\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECURITY_COUNCIL_MEMBERS_NUMBER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHORTEST_UPGRADE_NOTICE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SPECIAL_ACCOUNT_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SPECIAL_ACCOUNT_ID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TX_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_NOTICE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWAL_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"addAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"assetAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetGovernance\",\"outputs\":[{\"internalType\":\"contractAssetGovernance\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"assetsList\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAssetGovernance\",\"name\":\"_newAssetGovernance\",\"type\":\"address\"}],\"name\":\"changeAssetGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGovernor\",\"type\":\"address\"}],\"name\":\"changeGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initializationParameters\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAddressExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkGovernor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"pausedAssets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"requireActiveValidator\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"requireGovernor\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_assetAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_assetPaused\",\"type\":\"bool\"}],\"name\":\"setAssetPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"upgradeParameters\",\"type\":\"bytes\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_assetAddr\",\"type\":\"address\"}],\"name\":\"validateAssetAddress\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"validateAssetTokenLister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GovernanceABI is the input ABI used to generate the binding from.
// Deprecated: Use GovernanceMetaData.ABI instead.
var GovernanceABI = GovernanceMetaData.ABI

// Governance is an auto generated Go binding around an Ethereum contract.
type Governance struct {
	GovernanceCaller     // Read-only binding to the contract
	GovernanceTransactor // Write-only binding to the contract
	GovernanceFilterer   // Log filterer for contract events
}

// GovernanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernanceSession struct {
	Contract     *Governance       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernanceCallerSession struct {
	Contract *GovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernanceTransactorSession struct {
	Contract     *GovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GovernanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernanceRaw struct {
	Contract *Governance // Generic contract binding to access the raw methods on
}

// GovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernanceCallerRaw struct {
	Contract *GovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// GovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernanceTransactorRaw struct {
	Contract *GovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernance creates a new instance of Governance, bound to a specific deployed contract.
func NewGovernance(address common.Address, backend bind.ContractBackend) (*Governance, error) {
	contract, err := bindGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Governance{GovernanceCaller: GovernanceCaller{contract: contract}, GovernanceTransactor: GovernanceTransactor{contract: contract}, GovernanceFilterer: GovernanceFilterer{contract: contract}}, nil
}

// NewGovernanceCaller creates a new read-only instance of Governance, bound to a specific deployed contract.
func NewGovernanceCaller(address common.Address, caller bind.ContractCaller) (*GovernanceCaller, error) {
	contract, err := bindGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceCaller{contract: contract}, nil
}

// NewGovernanceTransactor creates a new write-only instance of Governance, bound to a specific deployed contract.
func NewGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernanceTransactor, error) {
	contract, err := bindGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceTransactor{contract: contract}, nil
}

// NewGovernanceFilterer creates a new log filterer instance of Governance, bound to a specific deployed contract.
func NewGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernanceFilterer, error) {
	contract, err := bindGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernanceFilterer{contract: contract}, nil
}

// bindGovernance binds a generic wrapper to an already deployed contract.
func bindGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governance *GovernanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Governance.Contract.GovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governance *GovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.Contract.GovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governance *GovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governance.Contract.GovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governance *GovernanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Governance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governance *GovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governance *GovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governance.Contract.contract.Transact(opts, method, params...)
}

// MAXACCOUNTINDEX is a free data retrieval call binding the contract method 0x437545f9.
//
// Solidity: function MAX_ACCOUNT_INDEX() view returns(uint32)
func (_Governance *GovernanceCaller) MAXACCOUNTINDEX(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "MAX_ACCOUNT_INDEX")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXACCOUNTINDEX is a free data retrieval call binding the contract method 0x437545f9.
//
// Solidity: function MAX_ACCOUNT_INDEX() view returns(uint32)
func (_Governance *GovernanceSession) MAXACCOUNTINDEX() (uint32, error) {
	return _Governance.Contract.MAXACCOUNTINDEX(&_Governance.CallOpts)
}

// MAXACCOUNTINDEX is a free data retrieval call binding the contract method 0x437545f9.
//
// Solidity: function MAX_ACCOUNT_INDEX() view returns(uint32)
func (_Governance *GovernanceCallerSession) MAXACCOUNTINDEX() (uint32, error) {
	return _Governance.Contract.MAXACCOUNTINDEX(&_Governance.CallOpts)
}

// MAXAMOUNTOFREGISTEREDASSETS is a free data retrieval call binding the contract method 0x0d360b7f.
//
// Solidity: function MAX_AMOUNT_OF_REGISTERED_ASSETS() view returns(uint16)
func (_Governance *GovernanceCaller) MAXAMOUNTOFREGISTEREDASSETS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "MAX_AMOUNT_OF_REGISTERED_ASSETS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXAMOUNTOFREGISTEREDASSETS is a free data retrieval call binding the contract method 0x0d360b7f.
//
// Solidity: function MAX_AMOUNT_OF_REGISTERED_ASSETS() view returns(uint16)
func (_Governance *GovernanceSession) MAXAMOUNTOFREGISTEREDASSETS() (uint16, error) {
	return _Governance.Contract.MAXAMOUNTOFREGISTEREDASSETS(&_Governance.CallOpts)
}

// MAXAMOUNTOFREGISTEREDASSETS is a free data retrieval call binding the contract method 0x0d360b7f.
//
// Solidity: function MAX_AMOUNT_OF_REGISTERED_ASSETS() view returns(uint16)
func (_Governance *GovernanceCallerSession) MAXAMOUNTOFREGISTEREDASSETS() (uint16, error) {
	return _Governance.Contract.MAXAMOUNTOFREGISTEREDASSETS(&_Governance.CallOpts)
}

// MAXDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x4c34a982.
//
// Solidity: function MAX_DEPOSIT_AMOUNT() view returns(uint128)
func (_Governance *GovernanceCaller) MAXDEPOSITAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "MAX_DEPOSIT_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x4c34a982.
//
// Solidity: function MAX_DEPOSIT_AMOUNT() view returns(uint128)
func (_Governance *GovernanceSession) MAXDEPOSITAMOUNT() (*big.Int, error) {
	return _Governance.Contract.MAXDEPOSITAMOUNT(&_Governance.CallOpts)
}

// MAXDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x4c34a982.
//
// Solidity: function MAX_DEPOSIT_AMOUNT() view returns(uint128)
func (_Governance *GovernanceCallerSession) MAXDEPOSITAMOUNT() (*big.Int, error) {
	return _Governance.Contract.MAXDEPOSITAMOUNT(&_Governance.CallOpts)
}

// MAXFUNGIBLEASSETID is a free data retrieval call binding the contract method 0x437da02f.
//
// Solidity: function MAX_FUNGIBLE_ASSET_ID() view returns(uint32)
func (_Governance *GovernanceCaller) MAXFUNGIBLEASSETID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "MAX_FUNGIBLE_ASSET_ID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXFUNGIBLEASSETID is a free data retrieval call binding the contract method 0x437da02f.
//
// Solidity: function MAX_FUNGIBLE_ASSET_ID() view returns(uint32)
func (_Governance *GovernanceSession) MAXFUNGIBLEASSETID() (uint32, error) {
	return _Governance.Contract.MAXFUNGIBLEASSETID(&_Governance.CallOpts)
}

// MAXFUNGIBLEASSETID is a free data retrieval call binding the contract method 0x437da02f.
//
// Solidity: function MAX_FUNGIBLE_ASSET_ID() view returns(uint32)
func (_Governance *GovernanceCallerSession) MAXFUNGIBLEASSETID() (uint32, error) {
	return _Governance.Contract.MAXFUNGIBLEASSETID(&_Governance.CallOpts)
}

// MAXNFTINDEX is a free data retrieval call binding the contract method 0x14791ad2.
//
// Solidity: function MAX_NFT_INDEX() view returns(uint40)
func (_Governance *GovernanceCaller) MAXNFTINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "MAX_NFT_INDEX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXNFTINDEX is a free data retrieval call binding the contract method 0x14791ad2.
//
// Solidity: function MAX_NFT_INDEX() view returns(uint40)
func (_Governance *GovernanceSession) MAXNFTINDEX() (*big.Int, error) {
	return _Governance.Contract.MAXNFTINDEX(&_Governance.CallOpts)
}

// MAXNFTINDEX is a free data retrieval call binding the contract method 0x14791ad2.
//
// Solidity: function MAX_NFT_INDEX() view returns(uint40)
func (_Governance *GovernanceCallerSession) MAXNFTINDEX() (*big.Int, error) {
	return _Governance.Contract.MAXNFTINDEX(&_Governance.CallOpts)
}

// SECURITYCOUNCILMEMBERSNUMBER is a free data retrieval call binding the contract method 0x4a51a71f.
//
// Solidity: function SECURITY_COUNCIL_MEMBERS_NUMBER() view returns(uint256)
func (_Governance *GovernanceCaller) SECURITYCOUNCILMEMBERSNUMBER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "SECURITY_COUNCIL_MEMBERS_NUMBER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECURITYCOUNCILMEMBERSNUMBER is a free data retrieval call binding the contract method 0x4a51a71f.
//
// Solidity: function SECURITY_COUNCIL_MEMBERS_NUMBER() view returns(uint256)
func (_Governance *GovernanceSession) SECURITYCOUNCILMEMBERSNUMBER() (*big.Int, error) {
	return _Governance.Contract.SECURITYCOUNCILMEMBERSNUMBER(&_Governance.CallOpts)
}

// SECURITYCOUNCILMEMBERSNUMBER is a free data retrieval call binding the contract method 0x4a51a71f.
//
// Solidity: function SECURITY_COUNCIL_MEMBERS_NUMBER() view returns(uint256)
func (_Governance *GovernanceCallerSession) SECURITYCOUNCILMEMBERSNUMBER() (*big.Int, error) {
	return _Governance.Contract.SECURITYCOUNCILMEMBERSNUMBER(&_Governance.CallOpts)
}

// SHORTESTUPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0x85053581.
//
// Solidity: function SHORTEST_UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_Governance *GovernanceCaller) SHORTESTUPGRADENOTICEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "SHORTEST_UPGRADE_NOTICE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SHORTESTUPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0x85053581.
//
// Solidity: function SHORTEST_UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_Governance *GovernanceSession) SHORTESTUPGRADENOTICEPERIOD() (*big.Int, error) {
	return _Governance.Contract.SHORTESTUPGRADENOTICEPERIOD(&_Governance.CallOpts)
}

// SHORTESTUPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0x85053581.
//
// Solidity: function SHORTEST_UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_Governance *GovernanceCallerSession) SHORTESTUPGRADENOTICEPERIOD() (*big.Int, error) {
	return _Governance.Contract.SHORTESTUPGRADENOTICEPERIOD(&_Governance.CallOpts)
}

// SPECIALACCOUNTADDRESS is a free data retrieval call binding the contract method 0x7ea399c1.
//
// Solidity: function SPECIAL_ACCOUNT_ADDRESS() view returns(address)
func (_Governance *GovernanceCaller) SPECIALACCOUNTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "SPECIAL_ACCOUNT_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SPECIALACCOUNTADDRESS is a free data retrieval call binding the contract method 0x7ea399c1.
//
// Solidity: function SPECIAL_ACCOUNT_ADDRESS() view returns(address)
func (_Governance *GovernanceSession) SPECIALACCOUNTADDRESS() (common.Address, error) {
	return _Governance.Contract.SPECIALACCOUNTADDRESS(&_Governance.CallOpts)
}

// SPECIALACCOUNTADDRESS is a free data retrieval call binding the contract method 0x7ea399c1.
//
// Solidity: function SPECIAL_ACCOUNT_ADDRESS() view returns(address)
func (_Governance *GovernanceCallerSession) SPECIALACCOUNTADDRESS() (common.Address, error) {
	return _Governance.Contract.SPECIALACCOUNTADDRESS(&_Governance.CallOpts)
}

// SPECIALACCOUNTID is a free data retrieval call binding the contract method 0x4242d5b3.
//
// Solidity: function SPECIAL_ACCOUNT_ID() view returns(uint32)
func (_Governance *GovernanceCaller) SPECIALACCOUNTID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "SPECIAL_ACCOUNT_ID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SPECIALACCOUNTID is a free data retrieval call binding the contract method 0x4242d5b3.
//
// Solidity: function SPECIAL_ACCOUNT_ID() view returns(uint32)
func (_Governance *GovernanceSession) SPECIALACCOUNTID() (uint32, error) {
	return _Governance.Contract.SPECIALACCOUNTID(&_Governance.CallOpts)
}

// SPECIALACCOUNTID is a free data retrieval call binding the contract method 0x4242d5b3.
//
// Solidity: function SPECIAL_ACCOUNT_ID() view returns(uint32)
func (_Governance *GovernanceCallerSession) SPECIALACCOUNTID() (uint32, error) {
	return _Governance.Contract.SPECIALACCOUNTID(&_Governance.CallOpts)
}

// TXSIZE is a free data retrieval call binding the contract method 0xe6e3c012.
//
// Solidity: function TX_SIZE() view returns(uint256)
func (_Governance *GovernanceCaller) TXSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "TX_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TXSIZE is a free data retrieval call binding the contract method 0xe6e3c012.
//
// Solidity: function TX_SIZE() view returns(uint256)
func (_Governance *GovernanceSession) TXSIZE() (*big.Int, error) {
	return _Governance.Contract.TXSIZE(&_Governance.CallOpts)
}

// TXSIZE is a free data retrieval call binding the contract method 0xe6e3c012.
//
// Solidity: function TX_SIZE() view returns(uint256)
func (_Governance *GovernanceCallerSession) TXSIZE() (*big.Int, error) {
	return _Governance.Contract.TXSIZE(&_Governance.CallOpts)
}

// UPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0xcc375fb7.
//
// Solidity: function UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_Governance *GovernanceCaller) UPGRADENOTICEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "UPGRADE_NOTICE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0xcc375fb7.
//
// Solidity: function UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_Governance *GovernanceSession) UPGRADENOTICEPERIOD() (*big.Int, error) {
	return _Governance.Contract.UPGRADENOTICEPERIOD(&_Governance.CallOpts)
}

// UPGRADENOTICEPERIOD is a free data retrieval call binding the contract method 0xcc375fb7.
//
// Solidity: function UPGRADE_NOTICE_PERIOD() view returns(uint256)
func (_Governance *GovernanceCallerSession) UPGRADENOTICEPERIOD() (*big.Int, error) {
	return _Governance.Contract.UPGRADENOTICEPERIOD(&_Governance.CallOpts)
}

// WITHDRAWALGASLIMIT is a free data retrieval call binding the contract method 0xc701f955.
//
// Solidity: function WITHDRAWAL_GAS_LIMIT() view returns(uint256)
func (_Governance *GovernanceCaller) WITHDRAWALGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "WITHDRAWAL_GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WITHDRAWALGASLIMIT is a free data retrieval call binding the contract method 0xc701f955.
//
// Solidity: function WITHDRAWAL_GAS_LIMIT() view returns(uint256)
func (_Governance *GovernanceSession) WITHDRAWALGASLIMIT() (*big.Int, error) {
	return _Governance.Contract.WITHDRAWALGASLIMIT(&_Governance.CallOpts)
}

// WITHDRAWALGASLIMIT is a free data retrieval call binding the contract method 0xc701f955.
//
// Solidity: function WITHDRAWAL_GAS_LIMIT() view returns(uint256)
func (_Governance *GovernanceCallerSession) WITHDRAWALGASLIMIT() (*big.Int, error) {
	return _Governance.Contract.WITHDRAWALGASLIMIT(&_Governance.CallOpts)
}

// AssetAddresses is a free data retrieval call binding the contract method 0xdbfc2967.
//
// Solidity: function assetAddresses(uint16 ) view returns(address)
func (_Governance *GovernanceCaller) AssetAddresses(opts *bind.CallOpts, arg0 uint16) (common.Address, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "assetAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetAddresses is a free data retrieval call binding the contract method 0xdbfc2967.
//
// Solidity: function assetAddresses(uint16 ) view returns(address)
func (_Governance *GovernanceSession) AssetAddresses(arg0 uint16) (common.Address, error) {
	return _Governance.Contract.AssetAddresses(&_Governance.CallOpts, arg0)
}

// AssetAddresses is a free data retrieval call binding the contract method 0xdbfc2967.
//
// Solidity: function assetAddresses(uint16 ) view returns(address)
func (_Governance *GovernanceCallerSession) AssetAddresses(arg0 uint16) (common.Address, error) {
	return _Governance.Contract.AssetAddresses(&_Governance.CallOpts, arg0)
}

// AssetGovernance is a free data retrieval call binding the contract method 0xf5e7d6fd.
//
// Solidity: function assetGovernance() view returns(address)
func (_Governance *GovernanceCaller) AssetGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "assetGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetGovernance is a free data retrieval call binding the contract method 0xf5e7d6fd.
//
// Solidity: function assetGovernance() view returns(address)
func (_Governance *GovernanceSession) AssetGovernance() (common.Address, error) {
	return _Governance.Contract.AssetGovernance(&_Governance.CallOpts)
}

// AssetGovernance is a free data retrieval call binding the contract method 0xf5e7d6fd.
//
// Solidity: function assetGovernance() view returns(address)
func (_Governance *GovernanceCallerSession) AssetGovernance() (common.Address, error) {
	return _Governance.Contract.AssetGovernance(&_Governance.CallOpts)
}

// AssetsList is a free data retrieval call binding the contract method 0x1e763ee3.
//
// Solidity: function assetsList(address ) view returns(uint16)
func (_Governance *GovernanceCaller) AssetsList(opts *bind.CallOpts, arg0 common.Address) (uint16, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "assetsList", arg0)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// AssetsList is a free data retrieval call binding the contract method 0x1e763ee3.
//
// Solidity: function assetsList(address ) view returns(uint16)
func (_Governance *GovernanceSession) AssetsList(arg0 common.Address) (uint16, error) {
	return _Governance.Contract.AssetsList(&_Governance.CallOpts, arg0)
}

// AssetsList is a free data retrieval call binding the contract method 0x1e763ee3.
//
// Solidity: function assetsList(address ) view returns(uint16)
func (_Governance *GovernanceCallerSession) AssetsList(arg0 common.Address) (uint16, error) {
	return _Governance.Contract.AssetsList(&_Governance.CallOpts, arg0)
}

// IsAddressExists is a free data retrieval call binding the contract method 0x321e182b.
//
// Solidity: function isAddressExists(address ) view returns(bool)
func (_Governance *GovernanceCaller) IsAddressExists(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "isAddressExists", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAddressExists is a free data retrieval call binding the contract method 0x321e182b.
//
// Solidity: function isAddressExists(address ) view returns(bool)
func (_Governance *GovernanceSession) IsAddressExists(arg0 common.Address) (bool, error) {
	return _Governance.Contract.IsAddressExists(&_Governance.CallOpts, arg0)
}

// IsAddressExists is a free data retrieval call binding the contract method 0x321e182b.
//
// Solidity: function isAddressExists(address ) view returns(bool)
func (_Governance *GovernanceCallerSession) IsAddressExists(arg0 common.Address) (bool, error) {
	return _Governance.Contract.IsAddressExists(&_Governance.CallOpts, arg0)
}

// NetworkGovernor is a free data retrieval call binding the contract method 0xf39349ef.
//
// Solidity: function networkGovernor() view returns(address)
func (_Governance *GovernanceCaller) NetworkGovernor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "networkGovernor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NetworkGovernor is a free data retrieval call binding the contract method 0xf39349ef.
//
// Solidity: function networkGovernor() view returns(address)
func (_Governance *GovernanceSession) NetworkGovernor() (common.Address, error) {
	return _Governance.Contract.NetworkGovernor(&_Governance.CallOpts)
}

// NetworkGovernor is a free data retrieval call binding the contract method 0xf39349ef.
//
// Solidity: function networkGovernor() view returns(address)
func (_Governance *GovernanceCallerSession) NetworkGovernor() (common.Address, error) {
	return _Governance.Contract.NetworkGovernor(&_Governance.CallOpts)
}

// PausedAssets is a free data retrieval call binding the contract method 0x31d8687b.
//
// Solidity: function pausedAssets(uint16 ) view returns(bool)
func (_Governance *GovernanceCaller) PausedAssets(opts *bind.CallOpts, arg0 uint16) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "pausedAssets", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PausedAssets is a free data retrieval call binding the contract method 0x31d8687b.
//
// Solidity: function pausedAssets(uint16 ) view returns(bool)
func (_Governance *GovernanceSession) PausedAssets(arg0 uint16) (bool, error) {
	return _Governance.Contract.PausedAssets(&_Governance.CallOpts, arg0)
}

// PausedAssets is a free data retrieval call binding the contract method 0x31d8687b.
//
// Solidity: function pausedAssets(uint16 ) view returns(bool)
func (_Governance *GovernanceCallerSession) PausedAssets(arg0 uint16) (bool, error) {
	return _Governance.Contract.PausedAssets(&_Governance.CallOpts, arg0)
}

// RequireActiveValidator is a free data retrieval call binding the contract method 0x4b18bd0f.
//
// Solidity: function requireActiveValidator(address _address) view returns()
func (_Governance *GovernanceCaller) RequireActiveValidator(opts *bind.CallOpts, _address common.Address) error {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "requireActiveValidator", _address)

	if err != nil {
		return err
	}

	return err

}

// RequireActiveValidator is a free data retrieval call binding the contract method 0x4b18bd0f.
//
// Solidity: function requireActiveValidator(address _address) view returns()
func (_Governance *GovernanceSession) RequireActiveValidator(_address common.Address) error {
	return _Governance.Contract.RequireActiveValidator(&_Governance.CallOpts, _address)
}

// RequireActiveValidator is a free data retrieval call binding the contract method 0x4b18bd0f.
//
// Solidity: function requireActiveValidator(address _address) view returns()
func (_Governance *GovernanceCallerSession) RequireActiveValidator(_address common.Address) error {
	return _Governance.Contract.RequireActiveValidator(&_Governance.CallOpts, _address)
}

// RequireGovernor is a free data retrieval call binding the contract method 0xf5f84ed4.
//
// Solidity: function requireGovernor(address _address) view returns()
func (_Governance *GovernanceCaller) RequireGovernor(opts *bind.CallOpts, _address common.Address) error {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "requireGovernor", _address)

	if err != nil {
		return err
	}

	return err

}

// RequireGovernor is a free data retrieval call binding the contract method 0xf5f84ed4.
//
// Solidity: function requireGovernor(address _address) view returns()
func (_Governance *GovernanceSession) RequireGovernor(_address common.Address) error {
	return _Governance.Contract.RequireGovernor(&_Governance.CallOpts, _address)
}

// RequireGovernor is a free data retrieval call binding the contract method 0xf5f84ed4.
//
// Solidity: function requireGovernor(address _address) view returns()
func (_Governance *GovernanceCallerSession) RequireGovernor(_address common.Address) error {
	return _Governance.Contract.RequireGovernor(&_Governance.CallOpts, _address)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint16)
func (_Governance *GovernanceCaller) TotalAssets(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint16)
func (_Governance *GovernanceSession) TotalAssets() (uint16, error) {
	return _Governance.Contract.TotalAssets(&_Governance.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint16)
func (_Governance *GovernanceCallerSession) TotalAssets() (uint16, error) {
	return _Governance.Contract.TotalAssets(&_Governance.CallOpts)
}

// ValidateAssetAddress is a free data retrieval call binding the contract method 0x9bd77609.
//
// Solidity: function validateAssetAddress(address _assetAddr) view returns(uint16)
func (_Governance *GovernanceCaller) ValidateAssetAddress(opts *bind.CallOpts, _assetAddr common.Address) (uint16, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "validateAssetAddress", _assetAddr)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ValidateAssetAddress is a free data retrieval call binding the contract method 0x9bd77609.
//
// Solidity: function validateAssetAddress(address _assetAddr) view returns(uint16)
func (_Governance *GovernanceSession) ValidateAssetAddress(_assetAddr common.Address) (uint16, error) {
	return _Governance.Contract.ValidateAssetAddress(&_Governance.CallOpts, _assetAddr)
}

// ValidateAssetAddress is a free data retrieval call binding the contract method 0x9bd77609.
//
// Solidity: function validateAssetAddress(address _assetAddr) view returns(uint16)
func (_Governance *GovernanceCallerSession) ValidateAssetAddress(_assetAddr common.Address) (uint16, error) {
	return _Governance.Contract.ValidateAssetAddress(&_Governance.CallOpts, _assetAddr)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(bool)
func (_Governance *GovernanceCaller) Validators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "validators", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(bool)
func (_Governance *GovernanceSession) Validators(arg0 common.Address) (bool, error) {
	return _Governance.Contract.Validators(&_Governance.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(bool)
func (_Governance *GovernanceCallerSession) Validators(arg0 common.Address) (bool, error) {
	return _Governance.Contract.Validators(&_Governance.CallOpts, arg0)
}

// AddAsset is a paid mutator transaction binding the contract method 0x298410e5.
//
// Solidity: function addAsset(address _asset) returns()
func (_Governance *GovernanceTransactor) AddAsset(opts *bind.TransactOpts, _asset common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "addAsset", _asset)
}

// AddAsset is a paid mutator transaction binding the contract method 0x298410e5.
//
// Solidity: function addAsset(address _asset) returns()
func (_Governance *GovernanceSession) AddAsset(_asset common.Address) (*types.Transaction, error) {
	return _Governance.Contract.AddAsset(&_Governance.TransactOpts, _asset)
}

// AddAsset is a paid mutator transaction binding the contract method 0x298410e5.
//
// Solidity: function addAsset(address _asset) returns()
func (_Governance *GovernanceTransactorSession) AddAsset(_asset common.Address) (*types.Transaction, error) {
	return _Governance.Contract.AddAsset(&_Governance.TransactOpts, _asset)
}

// ChangeAssetGovernance is a paid mutator transaction binding the contract method 0xd87e3748.
//
// Solidity: function changeAssetGovernance(address _newAssetGovernance) returns()
func (_Governance *GovernanceTransactor) ChangeAssetGovernance(opts *bind.TransactOpts, _newAssetGovernance common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "changeAssetGovernance", _newAssetGovernance)
}

// ChangeAssetGovernance is a paid mutator transaction binding the contract method 0xd87e3748.
//
// Solidity: function changeAssetGovernance(address _newAssetGovernance) returns()
func (_Governance *GovernanceSession) ChangeAssetGovernance(_newAssetGovernance common.Address) (*types.Transaction, error) {
	return _Governance.Contract.ChangeAssetGovernance(&_Governance.TransactOpts, _newAssetGovernance)
}

// ChangeAssetGovernance is a paid mutator transaction binding the contract method 0xd87e3748.
//
// Solidity: function changeAssetGovernance(address _newAssetGovernance) returns()
func (_Governance *GovernanceTransactorSession) ChangeAssetGovernance(_newAssetGovernance common.Address) (*types.Transaction, error) {
	return _Governance.Contract.ChangeAssetGovernance(&_Governance.TransactOpts, _newAssetGovernance)
}

// ChangeGovernor is a paid mutator transaction binding the contract method 0xe4c0aaf4.
//
// Solidity: function changeGovernor(address _newGovernor) returns()
func (_Governance *GovernanceTransactor) ChangeGovernor(opts *bind.TransactOpts, _newGovernor common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "changeGovernor", _newGovernor)
}

// ChangeGovernor is a paid mutator transaction binding the contract method 0xe4c0aaf4.
//
// Solidity: function changeGovernor(address _newGovernor) returns()
func (_Governance *GovernanceSession) ChangeGovernor(_newGovernor common.Address) (*types.Transaction, error) {
	return _Governance.Contract.ChangeGovernor(&_Governance.TransactOpts, _newGovernor)
}

// ChangeGovernor is a paid mutator transaction binding the contract method 0xe4c0aaf4.
//
// Solidity: function changeGovernor(address _newGovernor) returns()
func (_Governance *GovernanceTransactorSession) ChangeGovernor(_newGovernor common.Address) (*types.Transaction, error) {
	return _Governance.Contract.ChangeGovernor(&_Governance.TransactOpts, _newGovernor)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializationParameters) returns()
func (_Governance *GovernanceTransactor) Initialize(opts *bind.TransactOpts, initializationParameters []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "initialize", initializationParameters)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializationParameters) returns()
func (_Governance *GovernanceSession) Initialize(initializationParameters []byte) (*types.Transaction, error) {
	return _Governance.Contract.Initialize(&_Governance.TransactOpts, initializationParameters)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializationParameters) returns()
func (_Governance *GovernanceTransactorSession) Initialize(initializationParameters []byte) (*types.Transaction, error) {
	return _Governance.Contract.Initialize(&_Governance.TransactOpts, initializationParameters)
}

// SetAssetPaused is a paid mutator transaction binding the contract method 0x2520ce5a.
//
// Solidity: function setAssetPaused(address _assetAddress, bool _assetPaused) returns()
func (_Governance *GovernanceTransactor) SetAssetPaused(opts *bind.TransactOpts, _assetAddress common.Address, _assetPaused bool) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setAssetPaused", _assetAddress, _assetPaused)
}

// SetAssetPaused is a paid mutator transaction binding the contract method 0x2520ce5a.
//
// Solidity: function setAssetPaused(address _assetAddress, bool _assetPaused) returns()
func (_Governance *GovernanceSession) SetAssetPaused(_assetAddress common.Address, _assetPaused bool) (*types.Transaction, error) {
	return _Governance.Contract.SetAssetPaused(&_Governance.TransactOpts, _assetAddress, _assetPaused)
}

// SetAssetPaused is a paid mutator transaction binding the contract method 0x2520ce5a.
//
// Solidity: function setAssetPaused(address _assetAddress, bool _assetPaused) returns()
func (_Governance *GovernanceTransactorSession) SetAssetPaused(_assetAddress common.Address, _assetPaused bool) (*types.Transaction, error) {
	return _Governance.Contract.SetAssetPaused(&_Governance.TransactOpts, _assetAddress, _assetPaused)
}

// SetValidator is a paid mutator transaction binding the contract method 0x4623c91d.
//
// Solidity: function setValidator(address _validator, bool _active) returns()
func (_Governance *GovernanceTransactor) SetValidator(opts *bind.TransactOpts, _validator common.Address, _active bool) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setValidator", _validator, _active)
}

// SetValidator is a paid mutator transaction binding the contract method 0x4623c91d.
//
// Solidity: function setValidator(address _validator, bool _active) returns()
func (_Governance *GovernanceSession) SetValidator(_validator common.Address, _active bool) (*types.Transaction, error) {
	return _Governance.Contract.SetValidator(&_Governance.TransactOpts, _validator, _active)
}

// SetValidator is a paid mutator transaction binding the contract method 0x4623c91d.
//
// Solidity: function setValidator(address _validator, bool _active) returns()
func (_Governance *GovernanceTransactorSession) SetValidator(_validator common.Address, _active bool) (*types.Transaction, error) {
	return _Governance.Contract.SetValidator(&_Governance.TransactOpts, _validator, _active)
}

// Upgrade is a paid mutator transaction binding the contract method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_Governance *GovernanceTransactor) Upgrade(opts *bind.TransactOpts, upgradeParameters []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "upgrade", upgradeParameters)
}

// Upgrade is a paid mutator transaction binding the contract method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_Governance *GovernanceSession) Upgrade(upgradeParameters []byte) (*types.Transaction, error) {
	return _Governance.Contract.Upgrade(&_Governance.TransactOpts, upgradeParameters)
}

// Upgrade is a paid mutator transaction binding the contract method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_Governance *GovernanceTransactorSession) Upgrade(upgradeParameters []byte) (*types.Transaction, error) {
	return _Governance.Contract.Upgrade(&_Governance.TransactOpts, upgradeParameters)
}

// ValidateAssetTokenLister is a paid mutator transaction binding the contract method 0xc2001b38.
//
// Solidity: function validateAssetTokenLister(address _address) returns()
func (_Governance *GovernanceTransactor) ValidateAssetTokenLister(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "validateAssetTokenLister", _address)
}

// ValidateAssetTokenLister is a paid mutator transaction binding the contract method 0xc2001b38.
//
// Solidity: function validateAssetTokenLister(address _address) returns()
func (_Governance *GovernanceSession) ValidateAssetTokenLister(_address common.Address) (*types.Transaction, error) {
	return _Governance.Contract.ValidateAssetTokenLister(&_Governance.TransactOpts, _address)
}

// ValidateAssetTokenLister is a paid mutator transaction binding the contract method 0xc2001b38.
//
// Solidity: function validateAssetTokenLister(address _address) returns()
func (_Governance *GovernanceTransactorSession) ValidateAssetTokenLister(_address common.Address) (*types.Transaction, error) {
	return _Governance.Contract.ValidateAssetTokenLister(&_Governance.TransactOpts, _address)
}

// GovernanceAssetPausedUpdateIterator is returned from FilterAssetPausedUpdate and is used to iterate over the raw logs and unpacked data for AssetPausedUpdate events raised by the Governance contract.
type GovernanceAssetPausedUpdateIterator struct {
	Event *GovernanceAssetPausedUpdate // Event containing the contract specifics and raw log

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
func (it *GovernanceAssetPausedUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceAssetPausedUpdate)
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
		it.Event = new(GovernanceAssetPausedUpdate)
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
func (it *GovernanceAssetPausedUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceAssetPausedUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceAssetPausedUpdate represents a AssetPausedUpdate event raised by the Governance contract.
type GovernanceAssetPausedUpdate struct {
	Token  common.Address
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAssetPausedUpdate is a free log retrieval operation binding the contract event 0xf7ca5545623b85829d3abcb6729eb021070bbe93cba64f5c4de6b17b805b7d02.
//
// Solidity: event AssetPausedUpdate(address token, bool paused)
func (_Governance *GovernanceFilterer) FilterAssetPausedUpdate(opts *bind.FilterOpts) (*GovernanceAssetPausedUpdateIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "AssetPausedUpdate")
	if err != nil {
		return nil, err
	}
	return &GovernanceAssetPausedUpdateIterator{contract: _Governance.contract, event: "AssetPausedUpdate", logs: logs, sub: sub}, nil
}

// WatchAssetPausedUpdate is a free log subscription operation binding the contract event 0xf7ca5545623b85829d3abcb6729eb021070bbe93cba64f5c4de6b17b805b7d02.
//
// Solidity: event AssetPausedUpdate(address token, bool paused)
func (_Governance *GovernanceFilterer) WatchAssetPausedUpdate(opts *bind.WatchOpts, sink chan<- *GovernanceAssetPausedUpdate) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "AssetPausedUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceAssetPausedUpdate)
				if err := _Governance.contract.UnpackLog(event, "AssetPausedUpdate", log); err != nil {
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

// ParseAssetPausedUpdate is a log parse operation binding the contract event 0xf7ca5545623b85829d3abcb6729eb021070bbe93cba64f5c4de6b17b805b7d02.
//
// Solidity: event AssetPausedUpdate(address token, bool paused)
func (_Governance *GovernanceFilterer) ParseAssetPausedUpdate(log types.Log) (*GovernanceAssetPausedUpdate, error) {
	event := new(GovernanceAssetPausedUpdate)
	if err := _Governance.contract.UnpackLog(event, "AssetPausedUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceNewAssetIterator is returned from FilterNewAsset and is used to iterate over the raw logs and unpacked data for NewAsset events raised by the Governance contract.
type GovernanceNewAssetIterator struct {
	Event *GovernanceNewAsset // Event containing the contract specifics and raw log

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
func (it *GovernanceNewAssetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceNewAsset)
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
		it.Event = new(GovernanceNewAsset)
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
func (it *GovernanceNewAssetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceNewAssetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceNewAsset represents a NewAsset event raised by the Governance contract.
type GovernanceNewAsset struct {
	AssetAddress common.Address
	AssetId      uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewAsset is a free log retrieval operation binding the contract event 0x990c18cf226b253448a6a051b5cadf51a61f4ca56703374cdd8bf9df04b2f901.
//
// Solidity: event NewAsset(address assetAddress, uint16 assetId)
func (_Governance *GovernanceFilterer) FilterNewAsset(opts *bind.FilterOpts) (*GovernanceNewAssetIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "NewAsset")
	if err != nil {
		return nil, err
	}
	return &GovernanceNewAssetIterator{contract: _Governance.contract, event: "NewAsset", logs: logs, sub: sub}, nil
}

// WatchNewAsset is a free log subscription operation binding the contract event 0x990c18cf226b253448a6a051b5cadf51a61f4ca56703374cdd8bf9df04b2f901.
//
// Solidity: event NewAsset(address assetAddress, uint16 assetId)
func (_Governance *GovernanceFilterer) WatchNewAsset(opts *bind.WatchOpts, sink chan<- *GovernanceNewAsset) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "NewAsset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceNewAsset)
				if err := _Governance.contract.UnpackLog(event, "NewAsset", log); err != nil {
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

// ParseNewAsset is a log parse operation binding the contract event 0x990c18cf226b253448a6a051b5cadf51a61f4ca56703374cdd8bf9df04b2f901.
//
// Solidity: event NewAsset(address assetAddress, uint16 assetId)
func (_Governance *GovernanceFilterer) ParseNewAsset(log types.Log) (*GovernanceNewAsset, error) {
	event := new(GovernanceNewAsset)
	if err := _Governance.contract.UnpackLog(event, "NewAsset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceNewAssetGovernanceIterator is returned from FilterNewAssetGovernance and is used to iterate over the raw logs and unpacked data for NewAssetGovernance events raised by the Governance contract.
type GovernanceNewAssetGovernanceIterator struct {
	Event *GovernanceNewAssetGovernance // Event containing the contract specifics and raw log

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
func (it *GovernanceNewAssetGovernanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceNewAssetGovernance)
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
		it.Event = new(GovernanceNewAssetGovernance)
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
func (it *GovernanceNewAssetGovernanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceNewAssetGovernanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceNewAssetGovernance represents a NewAssetGovernance event raised by the Governance contract.
type GovernanceNewAssetGovernance struct {
	NewAssetGovernance common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewAssetGovernance is a free log retrieval operation binding the contract event 0xc7708091594580aae77e08cb1e2e3b9d2bd0cab0d2d95797248dd95f02d7299e.
//
// Solidity: event NewAssetGovernance(address newAssetGovernance)
func (_Governance *GovernanceFilterer) FilterNewAssetGovernance(opts *bind.FilterOpts) (*GovernanceNewAssetGovernanceIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "NewAssetGovernance")
	if err != nil {
		return nil, err
	}
	return &GovernanceNewAssetGovernanceIterator{contract: _Governance.contract, event: "NewAssetGovernance", logs: logs, sub: sub}, nil
}

// WatchNewAssetGovernance is a free log subscription operation binding the contract event 0xc7708091594580aae77e08cb1e2e3b9d2bd0cab0d2d95797248dd95f02d7299e.
//
// Solidity: event NewAssetGovernance(address newAssetGovernance)
func (_Governance *GovernanceFilterer) WatchNewAssetGovernance(opts *bind.WatchOpts, sink chan<- *GovernanceNewAssetGovernance) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "NewAssetGovernance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceNewAssetGovernance)
				if err := _Governance.contract.UnpackLog(event, "NewAssetGovernance", log); err != nil {
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

// ParseNewAssetGovernance is a log parse operation binding the contract event 0xc7708091594580aae77e08cb1e2e3b9d2bd0cab0d2d95797248dd95f02d7299e.
//
// Solidity: event NewAssetGovernance(address newAssetGovernance)
func (_Governance *GovernanceFilterer) ParseNewAssetGovernance(log types.Log) (*GovernanceNewAssetGovernance, error) {
	event := new(GovernanceNewAssetGovernance)
	if err := _Governance.contract.UnpackLog(event, "NewAssetGovernance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceNewGovernorIterator is returned from FilterNewGovernor and is used to iterate over the raw logs and unpacked data for NewGovernor events raised by the Governance contract.
type GovernanceNewGovernorIterator struct {
	Event *GovernanceNewGovernor // Event containing the contract specifics and raw log

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
func (it *GovernanceNewGovernorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceNewGovernor)
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
		it.Event = new(GovernanceNewGovernor)
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
func (it *GovernanceNewGovernorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceNewGovernorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceNewGovernor represents a NewGovernor event raised by the Governance contract.
type GovernanceNewGovernor struct {
	NewGovernor common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewGovernor is a free log retrieval operation binding the contract event 0x5425363a03f182281120f5919107c49c7a1a623acc1cbc6df468b6f0c11fcf8c.
//
// Solidity: event NewGovernor(address newGovernor)
func (_Governance *GovernanceFilterer) FilterNewGovernor(opts *bind.FilterOpts) (*GovernanceNewGovernorIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "NewGovernor")
	if err != nil {
		return nil, err
	}
	return &GovernanceNewGovernorIterator{contract: _Governance.contract, event: "NewGovernor", logs: logs, sub: sub}, nil
}

// WatchNewGovernor is a free log subscription operation binding the contract event 0x5425363a03f182281120f5919107c49c7a1a623acc1cbc6df468b6f0c11fcf8c.
//
// Solidity: event NewGovernor(address newGovernor)
func (_Governance *GovernanceFilterer) WatchNewGovernor(opts *bind.WatchOpts, sink chan<- *GovernanceNewGovernor) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "NewGovernor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceNewGovernor)
				if err := _Governance.contract.UnpackLog(event, "NewGovernor", log); err != nil {
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

// ParseNewGovernor is a log parse operation binding the contract event 0x5425363a03f182281120f5919107c49c7a1a623acc1cbc6df468b6f0c11fcf8c.
//
// Solidity: event NewGovernor(address newGovernor)
func (_Governance *GovernanceFilterer) ParseNewGovernor(log types.Log) (*GovernanceNewGovernor, error) {
	event := new(GovernanceNewGovernor)
	if err := _Governance.contract.UnpackLog(event, "NewGovernor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceValidatorStatusUpdateIterator is returned from FilterValidatorStatusUpdate and is used to iterate over the raw logs and unpacked data for ValidatorStatusUpdate events raised by the Governance contract.
type GovernanceValidatorStatusUpdateIterator struct {
	Event *GovernanceValidatorStatusUpdate // Event containing the contract specifics and raw log

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
func (it *GovernanceValidatorStatusUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceValidatorStatusUpdate)
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
		it.Event = new(GovernanceValidatorStatusUpdate)
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
func (it *GovernanceValidatorStatusUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceValidatorStatusUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceValidatorStatusUpdate represents a ValidatorStatusUpdate event raised by the Governance contract.
type GovernanceValidatorStatusUpdate struct {
	ValidatorAddress common.Address
	IsActive         bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterValidatorStatusUpdate is a free log retrieval operation binding the contract event 0x065b77b53864e46fda3d8986acb51696223d6dde7ced42441eb150bae6d48136.
//
// Solidity: event ValidatorStatusUpdate(address validatorAddress, bool isActive)
func (_Governance *GovernanceFilterer) FilterValidatorStatusUpdate(opts *bind.FilterOpts) (*GovernanceValidatorStatusUpdateIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ValidatorStatusUpdate")
	if err != nil {
		return nil, err
	}
	return &GovernanceValidatorStatusUpdateIterator{contract: _Governance.contract, event: "ValidatorStatusUpdate", logs: logs, sub: sub}, nil
}

// WatchValidatorStatusUpdate is a free log subscription operation binding the contract event 0x065b77b53864e46fda3d8986acb51696223d6dde7ced42441eb150bae6d48136.
//
// Solidity: event ValidatorStatusUpdate(address validatorAddress, bool isActive)
func (_Governance *GovernanceFilterer) WatchValidatorStatusUpdate(opts *bind.WatchOpts, sink chan<- *GovernanceValidatorStatusUpdate) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ValidatorStatusUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceValidatorStatusUpdate)
				if err := _Governance.contract.UnpackLog(event, "ValidatorStatusUpdate", log); err != nil {
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

// ParseValidatorStatusUpdate is a log parse operation binding the contract event 0x065b77b53864e46fda3d8986acb51696223d6dde7ced42441eb150bae6d48136.
//
// Solidity: event ValidatorStatusUpdate(address validatorAddress, bool isActive)
func (_Governance *GovernanceFilterer) ParseValidatorStatusUpdate(log types.Log) (*GovernanceValidatorStatusUpdate, error) {
	event := new(GovernanceValidatorStatusUpdate)
	if err := _Governance.contract.UnpackLog(event, "ValidatorStatusUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
