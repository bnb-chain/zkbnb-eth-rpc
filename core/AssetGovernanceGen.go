// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package core

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

// AssetGovernanceMetaData contains all meta data concerning the AssetGovernance contract.
var AssetGovernanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_listingFeeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_listingFee\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_listingCap\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_treasuryAccountIndex\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newListingCap\",\"type\":\"uint16\"}],\"name\":\"ListingCapUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"newListingFeeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newListingFee\",\"type\":\"uint256\"}],\"name\":\"ListingFeeTokenUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newListingFee\",\"type\":\"uint256\"}],\"name\":\"ListingFeeUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenLister\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"TokenListerUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"TreasuryUpdate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_assetAddress\",\"type\":\"address\"}],\"name\":\"addAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"contractGovernance\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listingCap\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listingFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listingFeeToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_listerAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"setLister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_newListingCap\",\"type\":\"uint16\"}],\"name\":\"setListingCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newListingFee\",\"type\":\"uint256\"}],\"name\":\"setListingFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_newListingFeeAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newListingFee\",\"type\":\"uint256\"}],\"name\":\"setListingFeeAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newTreasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenLister\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AssetGovernanceABI is the input ABI used to generate the binding from.
// Deprecated: Use AssetGovernanceMetaData.ABI instead.
var AssetGovernanceABI = AssetGovernanceMetaData.ABI

// AssetGovernance is an auto generated Go binding around an Ethereum contract.
type AssetGovernance struct {
	AssetGovernanceCaller     // Read-only binding to the contract
	AssetGovernanceTransactor // Write-only binding to the contract
	AssetGovernanceFilterer   // Log filterer for contract events
}

// AssetGovernanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssetGovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetGovernanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetGovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetGovernanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssetGovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetGovernanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetGovernanceSession struct {
	Contract     *AssetGovernance  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetGovernanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetGovernanceCallerSession struct {
	Contract *AssetGovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AssetGovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetGovernanceTransactorSession struct {
	Contract     *AssetGovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AssetGovernanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssetGovernanceRaw struct {
	Contract *AssetGovernance // Generic contract binding to access the raw methods on
}

// AssetGovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetGovernanceCallerRaw struct {
	Contract *AssetGovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// AssetGovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetGovernanceTransactorRaw struct {
	Contract *AssetGovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssetGovernance creates a new instance of AssetGovernance, bound to a specific deployed contract.
func NewAssetGovernance(address common.Address, backend bind.ContractBackend) (*AssetGovernance, error) {
	contract, err := bindAssetGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssetGovernance{AssetGovernanceCaller: AssetGovernanceCaller{contract: contract}, AssetGovernanceTransactor: AssetGovernanceTransactor{contract: contract}, AssetGovernanceFilterer: AssetGovernanceFilterer{contract: contract}}, nil
}

// NewAssetGovernanceCaller creates a new read-only instance of AssetGovernance, bound to a specific deployed contract.
func NewAssetGovernanceCaller(address common.Address, caller bind.ContractCaller) (*AssetGovernanceCaller, error) {
	contract, err := bindAssetGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetGovernanceCaller{contract: contract}, nil
}

// NewAssetGovernanceTransactor creates a new write-only instance of AssetGovernance, bound to a specific deployed contract.
func NewAssetGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetGovernanceTransactor, error) {
	contract, err := bindAssetGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetGovernanceTransactor{contract: contract}, nil
}

// NewAssetGovernanceFilterer creates a new log filterer instance of AssetGovernance, bound to a specific deployed contract.
func NewAssetGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*AssetGovernanceFilterer, error) {
	contract, err := bindAssetGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetGovernanceFilterer{contract: contract}, nil
}

// bindAssetGovernance binds a generic wrapper to an already deployed contract.
func bindAssetGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetGovernanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetGovernance *AssetGovernanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetGovernance.Contract.AssetGovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetGovernance *AssetGovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetGovernance.Contract.AssetGovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetGovernance *AssetGovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetGovernance.Contract.AssetGovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetGovernance *AssetGovernanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetGovernance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetGovernance *AssetGovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetGovernance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetGovernance *AssetGovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetGovernance.Contract.contract.Transact(opts, method, params...)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_AssetGovernance *AssetGovernanceCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssetGovernance.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_AssetGovernance *AssetGovernanceSession) Governance() (common.Address, error) {
	return _AssetGovernance.Contract.Governance(&_AssetGovernance.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_AssetGovernance *AssetGovernanceCallerSession) Governance() (common.Address, error) {
	return _AssetGovernance.Contract.Governance(&_AssetGovernance.CallOpts)
}

// ListingCap is a free data retrieval call binding the contract method 0x9b453945.
//
// Solidity: function listingCap() view returns(uint16)
func (_AssetGovernance *AssetGovernanceCaller) ListingCap(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _AssetGovernance.contract.Call(opts, &out, "listingCap")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ListingCap is a free data retrieval call binding the contract method 0x9b453945.
//
// Solidity: function listingCap() view returns(uint16)
func (_AssetGovernance *AssetGovernanceSession) ListingCap() (uint16, error) {
	return _AssetGovernance.Contract.ListingCap(&_AssetGovernance.CallOpts)
}

// ListingCap is a free data retrieval call binding the contract method 0x9b453945.
//
// Solidity: function listingCap() view returns(uint16)
func (_AssetGovernance *AssetGovernanceCallerSession) ListingCap() (uint16, error) {
	return _AssetGovernance.Contract.ListingCap(&_AssetGovernance.CallOpts)
}

// ListingFee is a free data retrieval call binding the contract method 0x6a1b7ecc.
//
// Solidity: function listingFee() view returns(uint256)
func (_AssetGovernance *AssetGovernanceCaller) ListingFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssetGovernance.contract.Call(opts, &out, "listingFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ListingFee is a free data retrieval call binding the contract method 0x6a1b7ecc.
//
// Solidity: function listingFee() view returns(uint256)
func (_AssetGovernance *AssetGovernanceSession) ListingFee() (*big.Int, error) {
	return _AssetGovernance.Contract.ListingFee(&_AssetGovernance.CallOpts)
}

// ListingFee is a free data retrieval call binding the contract method 0x6a1b7ecc.
//
// Solidity: function listingFee() view returns(uint256)
func (_AssetGovernance *AssetGovernanceCallerSession) ListingFee() (*big.Int, error) {
	return _AssetGovernance.Contract.ListingFee(&_AssetGovernance.CallOpts)
}

// ListingFeeToken is a free data retrieval call binding the contract method 0xa1a18c38.
//
// Solidity: function listingFeeToken() view returns(address)
func (_AssetGovernance *AssetGovernanceCaller) ListingFeeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssetGovernance.contract.Call(opts, &out, "listingFeeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ListingFeeToken is a free data retrieval call binding the contract method 0xa1a18c38.
//
// Solidity: function listingFeeToken() view returns(address)
func (_AssetGovernance *AssetGovernanceSession) ListingFeeToken() (common.Address, error) {
	return _AssetGovernance.Contract.ListingFeeToken(&_AssetGovernance.CallOpts)
}

// ListingFeeToken is a free data retrieval call binding the contract method 0xa1a18c38.
//
// Solidity: function listingFeeToken() view returns(address)
func (_AssetGovernance *AssetGovernanceCallerSession) ListingFeeToken() (common.Address, error) {
	return _AssetGovernance.Contract.ListingFeeToken(&_AssetGovernance.CallOpts)
}

// TokenLister is a free data retrieval call binding the contract method 0xaf3d414b.
//
// Solidity: function tokenLister(address ) view returns(bool)
func (_AssetGovernance *AssetGovernanceCaller) TokenLister(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AssetGovernance.contract.Call(opts, &out, "tokenLister", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenLister is a free data retrieval call binding the contract method 0xaf3d414b.
//
// Solidity: function tokenLister(address ) view returns(bool)
func (_AssetGovernance *AssetGovernanceSession) TokenLister(arg0 common.Address) (bool, error) {
	return _AssetGovernance.Contract.TokenLister(&_AssetGovernance.CallOpts, arg0)
}

// TokenLister is a free data retrieval call binding the contract method 0xaf3d414b.
//
// Solidity: function tokenLister(address ) view returns(bool)
func (_AssetGovernance *AssetGovernanceCallerSession) TokenLister(arg0 common.Address) (bool, error) {
	return _AssetGovernance.Contract.TokenLister(&_AssetGovernance.CallOpts, arg0)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_AssetGovernance *AssetGovernanceCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssetGovernance.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_AssetGovernance *AssetGovernanceSession) Treasury() (common.Address, error) {
	return _AssetGovernance.Contract.Treasury(&_AssetGovernance.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_AssetGovernance *AssetGovernanceCallerSession) Treasury() (common.Address, error) {
	return _AssetGovernance.Contract.Treasury(&_AssetGovernance.CallOpts)
}

// AddAsset is a paid mutator transaction binding the contract method 0x298410e5.
//
// Solidity: function addAsset(address _assetAddress) returns()
func (_AssetGovernance *AssetGovernanceTransactor) AddAsset(opts *bind.TransactOpts, _assetAddress common.Address) (*types.Transaction, error) {
	return _AssetGovernance.contract.Transact(opts, "addAsset", _assetAddress)
}

// AddAsset is a paid mutator transaction binding the contract method 0x298410e5.
//
// Solidity: function addAsset(address _assetAddress) returns()
func (_AssetGovernance *AssetGovernanceSession) AddAsset(_assetAddress common.Address) (*types.Transaction, error) {
	return _AssetGovernance.Contract.AddAsset(&_AssetGovernance.TransactOpts, _assetAddress)
}

// AddAsset is a paid mutator transaction binding the contract method 0x298410e5.
//
// Solidity: function addAsset(address _assetAddress) returns()
func (_AssetGovernance *AssetGovernanceTransactorSession) AddAsset(_assetAddress common.Address) (*types.Transaction, error) {
	return _AssetGovernance.Contract.AddAsset(&_AssetGovernance.TransactOpts, _assetAddress)
}

// SetLister is a paid mutator transaction binding the contract method 0x719c150c.
//
// Solidity: function setLister(address _listerAddress, bool _active) returns()
func (_AssetGovernance *AssetGovernanceTransactor) SetLister(opts *bind.TransactOpts, _listerAddress common.Address, _active bool) (*types.Transaction, error) {
	return _AssetGovernance.contract.Transact(opts, "setLister", _listerAddress, _active)
}

// SetLister is a paid mutator transaction binding the contract method 0x719c150c.
//
// Solidity: function setLister(address _listerAddress, bool _active) returns()
func (_AssetGovernance *AssetGovernanceSession) SetLister(_listerAddress common.Address, _active bool) (*types.Transaction, error) {
	return _AssetGovernance.Contract.SetLister(&_AssetGovernance.TransactOpts, _listerAddress, _active)
}

// SetLister is a paid mutator transaction binding the contract method 0x719c150c.
//
// Solidity: function setLister(address _listerAddress, bool _active) returns()
func (_AssetGovernance *AssetGovernanceTransactorSession) SetLister(_listerAddress common.Address, _active bool) (*types.Transaction, error) {
	return _AssetGovernance.Contract.SetLister(&_AssetGovernance.TransactOpts, _listerAddress, _active)
}

// SetListingCap is a paid mutator transaction binding the contract method 0xb42a9d80.
//
// Solidity: function setListingCap(uint16 _newListingCap) returns()
func (_AssetGovernance *AssetGovernanceTransactor) SetListingCap(opts *bind.TransactOpts, _newListingCap uint16) (*types.Transaction, error) {
	return _AssetGovernance.contract.Transact(opts, "setListingCap", _newListingCap)
}

// SetListingCap is a paid mutator transaction binding the contract method 0xb42a9d80.
//
// Solidity: function setListingCap(uint16 _newListingCap) returns()
func (_AssetGovernance *AssetGovernanceSession) SetListingCap(_newListingCap uint16) (*types.Transaction, error) {
	return _AssetGovernance.Contract.SetListingCap(&_AssetGovernance.TransactOpts, _newListingCap)
}

// SetListingCap is a paid mutator transaction binding the contract method 0xb42a9d80.
//
// Solidity: function setListingCap(uint16 _newListingCap) returns()
func (_AssetGovernance *AssetGovernanceTransactorSession) SetListingCap(_newListingCap uint16) (*types.Transaction, error) {
	return _AssetGovernance.Contract.SetListingCap(&_AssetGovernance.TransactOpts, _newListingCap)
}

// SetListingFee is a paid mutator transaction binding the contract method 0x131dbd09.
//
// Solidity: function setListingFee(uint256 _newListingFee) returns()
func (_AssetGovernance *AssetGovernanceTransactor) SetListingFee(opts *bind.TransactOpts, _newListingFee *big.Int) (*types.Transaction, error) {
	return _AssetGovernance.contract.Transact(opts, "setListingFee", _newListingFee)
}

// SetListingFee is a paid mutator transaction binding the contract method 0x131dbd09.
//
// Solidity: function setListingFee(uint256 _newListingFee) returns()
func (_AssetGovernance *AssetGovernanceSession) SetListingFee(_newListingFee *big.Int) (*types.Transaction, error) {
	return _AssetGovernance.Contract.SetListingFee(&_AssetGovernance.TransactOpts, _newListingFee)
}

// SetListingFee is a paid mutator transaction binding the contract method 0x131dbd09.
//
// Solidity: function setListingFee(uint256 _newListingFee) returns()
func (_AssetGovernance *AssetGovernanceTransactorSession) SetListingFee(_newListingFee *big.Int) (*types.Transaction, error) {
	return _AssetGovernance.Contract.SetListingFee(&_AssetGovernance.TransactOpts, _newListingFee)
}

// SetListingFeeAsset is a paid mutator transaction binding the contract method 0x4ddf6237.
//
// Solidity: function setListingFeeAsset(address _newListingFeeAsset, uint256 _newListingFee) returns()
func (_AssetGovernance *AssetGovernanceTransactor) SetListingFeeAsset(opts *bind.TransactOpts, _newListingFeeAsset common.Address, _newListingFee *big.Int) (*types.Transaction, error) {
	return _AssetGovernance.contract.Transact(opts, "setListingFeeAsset", _newListingFeeAsset, _newListingFee)
}

// SetListingFeeAsset is a paid mutator transaction binding the contract method 0x4ddf6237.
//
// Solidity: function setListingFeeAsset(address _newListingFeeAsset, uint256 _newListingFee) returns()
func (_AssetGovernance *AssetGovernanceSession) SetListingFeeAsset(_newListingFeeAsset common.Address, _newListingFee *big.Int) (*types.Transaction, error) {
	return _AssetGovernance.Contract.SetListingFeeAsset(&_AssetGovernance.TransactOpts, _newListingFeeAsset, _newListingFee)
}

// SetListingFeeAsset is a paid mutator transaction binding the contract method 0x4ddf6237.
//
// Solidity: function setListingFeeAsset(address _newListingFeeAsset, uint256 _newListingFee) returns()
func (_AssetGovernance *AssetGovernanceTransactorSession) SetListingFeeAsset(_newListingFeeAsset common.Address, _newListingFee *big.Int) (*types.Transaction, error) {
	return _AssetGovernance.Contract.SetListingFeeAsset(&_AssetGovernance.TransactOpts, _newListingFeeAsset, _newListingFee)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _newTreasury) returns()
func (_AssetGovernance *AssetGovernanceTransactor) SetTreasury(opts *bind.TransactOpts, _newTreasury common.Address) (*types.Transaction, error) {
	return _AssetGovernance.contract.Transact(opts, "setTreasury", _newTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _newTreasury) returns()
func (_AssetGovernance *AssetGovernanceSession) SetTreasury(_newTreasury common.Address) (*types.Transaction, error) {
	return _AssetGovernance.Contract.SetTreasury(&_AssetGovernance.TransactOpts, _newTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _newTreasury) returns()
func (_AssetGovernance *AssetGovernanceTransactorSession) SetTreasury(_newTreasury common.Address) (*types.Transaction, error) {
	return _AssetGovernance.Contract.SetTreasury(&_AssetGovernance.TransactOpts, _newTreasury)
}

// AssetGovernanceListingCapUpdateIterator is returned from FilterListingCapUpdate and is used to iterate over the raw logs and unpacked data for ListingCapUpdate events raised by the AssetGovernance contract.
type AssetGovernanceListingCapUpdateIterator struct {
	Event *AssetGovernanceListingCapUpdate // Event containing the contract specifics and raw log

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
func (it *AssetGovernanceListingCapUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetGovernanceListingCapUpdate)
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
		it.Event = new(AssetGovernanceListingCapUpdate)
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
func (it *AssetGovernanceListingCapUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetGovernanceListingCapUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetGovernanceListingCapUpdate represents a ListingCapUpdate event raised by the AssetGovernance contract.
type AssetGovernanceListingCapUpdate struct {
	NewListingCap uint16
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterListingCapUpdate is a free log retrieval operation binding the contract event 0xd07b88b359bb918e12349b2779862f7c6954810991cfcf664e36bd2168f17b0d.
//
// Solidity: event ListingCapUpdate(uint16 newListingCap)
func (_AssetGovernance *AssetGovernanceFilterer) FilterListingCapUpdate(opts *bind.FilterOpts) (*AssetGovernanceListingCapUpdateIterator, error) {

	logs, sub, err := _AssetGovernance.contract.FilterLogs(opts, "ListingCapUpdate")
	if err != nil {
		return nil, err
	}
	return &AssetGovernanceListingCapUpdateIterator{contract: _AssetGovernance.contract, event: "ListingCapUpdate", logs: logs, sub: sub}, nil
}

// WatchListingCapUpdate is a free log subscription operation binding the contract event 0xd07b88b359bb918e12349b2779862f7c6954810991cfcf664e36bd2168f17b0d.
//
// Solidity: event ListingCapUpdate(uint16 newListingCap)
func (_AssetGovernance *AssetGovernanceFilterer) WatchListingCapUpdate(opts *bind.WatchOpts, sink chan<- *AssetGovernanceListingCapUpdate) (event.Subscription, error) {

	logs, sub, err := _AssetGovernance.contract.WatchLogs(opts, "ListingCapUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetGovernanceListingCapUpdate)
				if err := _AssetGovernance.contract.UnpackLog(event, "ListingCapUpdate", log); err != nil {
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

// ParseListingCapUpdate is a log parse operation binding the contract event 0xd07b88b359bb918e12349b2779862f7c6954810991cfcf664e36bd2168f17b0d.
//
// Solidity: event ListingCapUpdate(uint16 newListingCap)
func (_AssetGovernance *AssetGovernanceFilterer) ParseListingCapUpdate(log types.Log) (*AssetGovernanceListingCapUpdate, error) {
	event := new(AssetGovernanceListingCapUpdate)
	if err := _AssetGovernance.contract.UnpackLog(event, "ListingCapUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetGovernanceListingFeeTokenUpdateIterator is returned from FilterListingFeeTokenUpdate and is used to iterate over the raw logs and unpacked data for ListingFeeTokenUpdate events raised by the AssetGovernance contract.
type AssetGovernanceListingFeeTokenUpdateIterator struct {
	Event *AssetGovernanceListingFeeTokenUpdate // Event containing the contract specifics and raw log

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
func (it *AssetGovernanceListingFeeTokenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetGovernanceListingFeeTokenUpdate)
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
		it.Event = new(AssetGovernanceListingFeeTokenUpdate)
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
func (it *AssetGovernanceListingFeeTokenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetGovernanceListingFeeTokenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetGovernanceListingFeeTokenUpdate represents a ListingFeeTokenUpdate event raised by the AssetGovernance contract.
type AssetGovernanceListingFeeTokenUpdate struct {
	NewListingFeeToken common.Address
	NewListingFee      *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterListingFeeTokenUpdate is a free log retrieval operation binding the contract event 0x06c002e270943576c6f430ffa27b7c21b4f12f8398079a843e82978a718fe139.
//
// Solidity: event ListingFeeTokenUpdate(address indexed newListingFeeToken, uint256 newListingFee)
func (_AssetGovernance *AssetGovernanceFilterer) FilterListingFeeTokenUpdate(opts *bind.FilterOpts, newListingFeeToken []common.Address) (*AssetGovernanceListingFeeTokenUpdateIterator, error) {

	var newListingFeeTokenRule []interface{}
	for _, newListingFeeTokenItem := range newListingFeeToken {
		newListingFeeTokenRule = append(newListingFeeTokenRule, newListingFeeTokenItem)
	}

	logs, sub, err := _AssetGovernance.contract.FilterLogs(opts, "ListingFeeTokenUpdate", newListingFeeTokenRule)
	if err != nil {
		return nil, err
	}
	return &AssetGovernanceListingFeeTokenUpdateIterator{contract: _AssetGovernance.contract, event: "ListingFeeTokenUpdate", logs: logs, sub: sub}, nil
}

// WatchListingFeeTokenUpdate is a free log subscription operation binding the contract event 0x06c002e270943576c6f430ffa27b7c21b4f12f8398079a843e82978a718fe139.
//
// Solidity: event ListingFeeTokenUpdate(address indexed newListingFeeToken, uint256 newListingFee)
func (_AssetGovernance *AssetGovernanceFilterer) WatchListingFeeTokenUpdate(opts *bind.WatchOpts, sink chan<- *AssetGovernanceListingFeeTokenUpdate, newListingFeeToken []common.Address) (event.Subscription, error) {

	var newListingFeeTokenRule []interface{}
	for _, newListingFeeTokenItem := range newListingFeeToken {
		newListingFeeTokenRule = append(newListingFeeTokenRule, newListingFeeTokenItem)
	}

	logs, sub, err := _AssetGovernance.contract.WatchLogs(opts, "ListingFeeTokenUpdate", newListingFeeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetGovernanceListingFeeTokenUpdate)
				if err := _AssetGovernance.contract.UnpackLog(event, "ListingFeeTokenUpdate", log); err != nil {
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

// ParseListingFeeTokenUpdate is a log parse operation binding the contract event 0x06c002e270943576c6f430ffa27b7c21b4f12f8398079a843e82978a718fe139.
//
// Solidity: event ListingFeeTokenUpdate(address indexed newListingFeeToken, uint256 newListingFee)
func (_AssetGovernance *AssetGovernanceFilterer) ParseListingFeeTokenUpdate(log types.Log) (*AssetGovernanceListingFeeTokenUpdate, error) {
	event := new(AssetGovernanceListingFeeTokenUpdate)
	if err := _AssetGovernance.contract.UnpackLog(event, "ListingFeeTokenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetGovernanceListingFeeUpdateIterator is returned from FilterListingFeeUpdate and is used to iterate over the raw logs and unpacked data for ListingFeeUpdate events raised by the AssetGovernance contract.
type AssetGovernanceListingFeeUpdateIterator struct {
	Event *AssetGovernanceListingFeeUpdate // Event containing the contract specifics and raw log

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
func (it *AssetGovernanceListingFeeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetGovernanceListingFeeUpdate)
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
		it.Event = new(AssetGovernanceListingFeeUpdate)
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
func (it *AssetGovernanceListingFeeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetGovernanceListingFeeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetGovernanceListingFeeUpdate represents a ListingFeeUpdate event raised by the AssetGovernance contract.
type AssetGovernanceListingFeeUpdate struct {
	NewListingFee *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterListingFeeUpdate is a free log retrieval operation binding the contract event 0xa0951033a9fa04df3877d5fdba9b43d00127caf3d6ff546d0bc633f122844ee6.
//
// Solidity: event ListingFeeUpdate(uint256 newListingFee)
func (_AssetGovernance *AssetGovernanceFilterer) FilterListingFeeUpdate(opts *bind.FilterOpts) (*AssetGovernanceListingFeeUpdateIterator, error) {

	logs, sub, err := _AssetGovernance.contract.FilterLogs(opts, "ListingFeeUpdate")
	if err != nil {
		return nil, err
	}
	return &AssetGovernanceListingFeeUpdateIterator{contract: _AssetGovernance.contract, event: "ListingFeeUpdate", logs: logs, sub: sub}, nil
}

// WatchListingFeeUpdate is a free log subscription operation binding the contract event 0xa0951033a9fa04df3877d5fdba9b43d00127caf3d6ff546d0bc633f122844ee6.
//
// Solidity: event ListingFeeUpdate(uint256 newListingFee)
func (_AssetGovernance *AssetGovernanceFilterer) WatchListingFeeUpdate(opts *bind.WatchOpts, sink chan<- *AssetGovernanceListingFeeUpdate) (event.Subscription, error) {

	logs, sub, err := _AssetGovernance.contract.WatchLogs(opts, "ListingFeeUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetGovernanceListingFeeUpdate)
				if err := _AssetGovernance.contract.UnpackLog(event, "ListingFeeUpdate", log); err != nil {
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

// ParseListingFeeUpdate is a log parse operation binding the contract event 0xa0951033a9fa04df3877d5fdba9b43d00127caf3d6ff546d0bc633f122844ee6.
//
// Solidity: event ListingFeeUpdate(uint256 newListingFee)
func (_AssetGovernance *AssetGovernanceFilterer) ParseListingFeeUpdate(log types.Log) (*AssetGovernanceListingFeeUpdate, error) {
	event := new(AssetGovernanceListingFeeUpdate)
	if err := _AssetGovernance.contract.UnpackLog(event, "ListingFeeUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetGovernanceTokenListerUpdateIterator is returned from FilterTokenListerUpdate and is used to iterate over the raw logs and unpacked data for TokenListerUpdate events raised by the AssetGovernance contract.
type AssetGovernanceTokenListerUpdateIterator struct {
	Event *AssetGovernanceTokenListerUpdate // Event containing the contract specifics and raw log

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
func (it *AssetGovernanceTokenListerUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetGovernanceTokenListerUpdate)
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
		it.Event = new(AssetGovernanceTokenListerUpdate)
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
func (it *AssetGovernanceTokenListerUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetGovernanceTokenListerUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetGovernanceTokenListerUpdate represents a TokenListerUpdate event raised by the AssetGovernance contract.
type AssetGovernanceTokenListerUpdate struct {
	TokenLister common.Address
	IsActive    bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenListerUpdate is a free log retrieval operation binding the contract event 0x947516bb7f07e641084b9d05943fb1c89e05da3d4d17f707ab4efd6ab58f0910.
//
// Solidity: event TokenListerUpdate(address indexed tokenLister, bool isActive)
func (_AssetGovernance *AssetGovernanceFilterer) FilterTokenListerUpdate(opts *bind.FilterOpts, tokenLister []common.Address) (*AssetGovernanceTokenListerUpdateIterator, error) {

	var tokenListerRule []interface{}
	for _, tokenListerItem := range tokenLister {
		tokenListerRule = append(tokenListerRule, tokenListerItem)
	}

	logs, sub, err := _AssetGovernance.contract.FilterLogs(opts, "TokenListerUpdate", tokenListerRule)
	if err != nil {
		return nil, err
	}
	return &AssetGovernanceTokenListerUpdateIterator{contract: _AssetGovernance.contract, event: "TokenListerUpdate", logs: logs, sub: sub}, nil
}

// WatchTokenListerUpdate is a free log subscription operation binding the contract event 0x947516bb7f07e641084b9d05943fb1c89e05da3d4d17f707ab4efd6ab58f0910.
//
// Solidity: event TokenListerUpdate(address indexed tokenLister, bool isActive)
func (_AssetGovernance *AssetGovernanceFilterer) WatchTokenListerUpdate(opts *bind.WatchOpts, sink chan<- *AssetGovernanceTokenListerUpdate, tokenLister []common.Address) (event.Subscription, error) {

	var tokenListerRule []interface{}
	for _, tokenListerItem := range tokenLister {
		tokenListerRule = append(tokenListerRule, tokenListerItem)
	}

	logs, sub, err := _AssetGovernance.contract.WatchLogs(opts, "TokenListerUpdate", tokenListerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetGovernanceTokenListerUpdate)
				if err := _AssetGovernance.contract.UnpackLog(event, "TokenListerUpdate", log); err != nil {
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

// ParseTokenListerUpdate is a log parse operation binding the contract event 0x947516bb7f07e641084b9d05943fb1c89e05da3d4d17f707ab4efd6ab58f0910.
//
// Solidity: event TokenListerUpdate(address indexed tokenLister, bool isActive)
func (_AssetGovernance *AssetGovernanceFilterer) ParseTokenListerUpdate(log types.Log) (*AssetGovernanceTokenListerUpdate, error) {
	event := new(AssetGovernanceTokenListerUpdate)
	if err := _AssetGovernance.contract.UnpackLog(event, "TokenListerUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetGovernanceTreasuryUpdateIterator is returned from FilterTreasuryUpdate and is used to iterate over the raw logs and unpacked data for TreasuryUpdate events raised by the AssetGovernance contract.
type AssetGovernanceTreasuryUpdateIterator struct {
	Event *AssetGovernanceTreasuryUpdate // Event containing the contract specifics and raw log

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
func (it *AssetGovernanceTreasuryUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetGovernanceTreasuryUpdate)
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
		it.Event = new(AssetGovernanceTreasuryUpdate)
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
func (it *AssetGovernanceTreasuryUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetGovernanceTreasuryUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetGovernanceTreasuryUpdate represents a TreasuryUpdate event raised by the AssetGovernance contract.
type AssetGovernanceTreasuryUpdate struct {
	NewTreasury common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTreasuryUpdate is a free log retrieval operation binding the contract event 0x8a3509a4057c89a5993a4a3140c2ebf7e829d325d8998eaa6c48adcff98b2cef.
//
// Solidity: event TreasuryUpdate(address newTreasury)
func (_AssetGovernance *AssetGovernanceFilterer) FilterTreasuryUpdate(opts *bind.FilterOpts) (*AssetGovernanceTreasuryUpdateIterator, error) {

	logs, sub, err := _AssetGovernance.contract.FilterLogs(opts, "TreasuryUpdate")
	if err != nil {
		return nil, err
	}
	return &AssetGovernanceTreasuryUpdateIterator{contract: _AssetGovernance.contract, event: "TreasuryUpdate", logs: logs, sub: sub}, nil
}

// WatchTreasuryUpdate is a free log subscription operation binding the contract event 0x8a3509a4057c89a5993a4a3140c2ebf7e829d325d8998eaa6c48adcff98b2cef.
//
// Solidity: event TreasuryUpdate(address newTreasury)
func (_AssetGovernance *AssetGovernanceFilterer) WatchTreasuryUpdate(opts *bind.WatchOpts, sink chan<- *AssetGovernanceTreasuryUpdate) (event.Subscription, error) {

	logs, sub, err := _AssetGovernance.contract.WatchLogs(opts, "TreasuryUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetGovernanceTreasuryUpdate)
				if err := _AssetGovernance.contract.UnpackLog(event, "TreasuryUpdate", log); err != nil {
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

// ParseTreasuryUpdate is a log parse operation binding the contract event 0x8a3509a4057c89a5993a4a3140c2ebf7e829d325d8998eaa6c48adcff98b2cef.
//
// Solidity: event TreasuryUpdate(address newTreasury)
func (_AssetGovernance *AssetGovernanceFilterer) ParseTreasuryUpdate(log types.Log) (*AssetGovernanceTreasuryUpdate, error) {
	event := new(AssetGovernanceTreasuryUpdate)
	if err := _AssetGovernance.contract.UnpackLog(event, "TreasuryUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
