// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package basic

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

// AssetGovernanceABI is the input ABI used to generate the binding from.
const AssetGovernanceABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenLister\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"TokenListerUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"contractGovernance\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initializationParameters\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listingCap\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_assetId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_assetAddress\",\"type\":\"address\"}],\"name\":\"setAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_listerAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"setLister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenLister\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"upgradeParameters\",\"type\":\"bytes\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AssetGovernanceBin is the compiled bytecode used for deploying new contracts.
var AssetGovernanceBin = "0x608060405234801561001057600080fd5b5061001961001e565b610081565b7f8e94fed44239eb2314ab7a406345e6c5a8f0ccedf3b600de3d004e672c33abf480546001909155801561007e576040805162461bcd60e51b815260206004820152600260248201526118a160f11b604482015290519081900360640190fd5b50565b611093806100906000396000f3fe608060405234801561001057600080fd5b50600436106101585760003560e01c80635f4fea26116100c3578063d87e37481161007c578063d87e3748146104d1578063dbfc2967146104f7578063e4c0aaf414610534578063f39349ef1461055a578063f5e7d6fd14610562578063f5f84ed41461056a57610158565b80635f4fea26146103e55780636c8244871461040b5780637a3459d4146104315780639bd7760914610457578063a28e00321461047d578063ac9b5671146104a357610158565b806326eb0f4a1161011557806326eb0f4a1461029657806328b58e1a146102c457806331d8687b146102fe578063321e182b1461031f578063439fab91146103455780635bab15ca146103b557610158565b806301e1d1141461015d5780631afcac961461017c5780631c021956146101ac5780631e763ee3146101d25780632520ce5a146101f85780632539464514610226575b600080fd5b610165610590565b6040805161ffff9092168252519081900360200190f35b6101aa6004803603604081101561019257600080fd5b506001600160a01b03813516906020013515156105a1565b005b6101aa600480360360208110156101c257600080fd5b50356001600160a01b0316610631565b610165600480360360208110156101e857600080fd5b50356001600160a01b0316610693565b6101aa6004803603604081101561020e57600080fd5b506001600160a01b03813516906020013515156106a9565b6101aa6004803603602081101561023c57600080fd5b81019060208101813564010000000081111561025757600080fd5b82018360208201111561026957600080fd5b8035906020019184600183028401116401000000008311171561028b57600080fd5b50909250905061062d565b6101aa600480360360408110156102ac57600080fd5b506001600160a01b03813516906020013515156107b8565b6102ea600480360360208110156102da57600080fd5b50356001600160a01b0316610847565b604080519115158252519081900360200190f35b6102ea6004803603602081101561031457600080fd5b503561ffff1661085c565b6102ea6004803603602081101561033557600080fd5b50356001600160a01b0316610871565b6101aa6004803603602081101561035b57600080fd5b81019060208101813564010000000081111561037657600080fd5b82018360208201111561038857600080fd5b803590602001918460018302840111640100000000831117156103aa57600080fd5b509092509050610886565b6101aa600480360360408110156103cb57600080fd5b50803561ffff1690602001356001600160a01b03166108de565b6102ea600480360360208110156103fb57600080fd5b50356001600160a01b0316610c01565b6102ea6004803603602081101561042157600080fd5b50356001600160a01b0316610c16565b6101aa6004803603602081101561044757600080fd5b50356001600160a01b0316610c2b565b6101656004803603602081101561046d57600080fd5b50356001600160a01b0316610c8b565b6101aa6004803603602081101561049357600080fd5b50356001600160a01b0316610da9565b6101aa600480360360408110156104b957600080fd5b506001600160a01b0381351690602001351515610e0a565b6101aa600480360360208110156104e757600080fd5b50356001600160a01b0316610e99565b6105186004803603602081101561050d57600080fd5b503561ffff16610f0b565b604080516001600160a01b039092168252519081900360200190f35b6101aa6004803603602081101561054a57600080fd5b50356001600160a01b0316610f26565b610518610f98565b610518610fa7565b6101aa6004803603602081101561058057600080fd5b50356001600160a01b0316610fb6565b600154600160a01b900461ffff1681565b6105aa33610fb6565b6001600160a01b03821660009081526002602052604090205460ff1615158115151461062d576001600160a01b038216600081815260026020908152604091829020805460ff1916851515908117909155825190815291517fd158a06928d89556421a4e2eabd98c7ea10e974c07e399e3afc449d19eb391b29281900390910190a25b5050565b6001600160a01b03811660009081526004602052604090205460ff16610690576040805162461bcd60e51b815260206004820152600f60248201526e34b73b30b634b21036b7b734ba37b960891b604482015290519081900360640190fd5b50565b60066020526000908152604090205461ffff1681565b6106b233610fb6565b6000306001600160a01b0316639bd77609846040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561070157600080fd5b505afa158015610715573d6000803e3d6000fd5b505050506040513d602081101561072b57600080fd5b505161ffff811660009081526005602052604090205490915060ff161515821515146107b35761ffff8116600090815260056020908152604091829020805460ff1916851515908117909155825190815291516001600160a01b038616927ff7ca5545623b85829d3abcb6729eb021070bbe93cba64f5c4de6b17b805b7d0292908290030190a25b505050565b6107c133610fb6565b6001600160a01b03821660009081526004602052604090205460ff1615158115151461062d576001600160a01b038216600081815260046020908152604091829020805460ff1916851515908117909155825190815291517f6fbe3fec04b6d2f79e569f1a46ec4d1621e751112483b5ce9d81832f19fee5c19281900390910190a25050565b60046020526000908152604090205460ff1681565b60056020526000908152604090205460ff1681565b60086020526000908152604090205460ff1681565b61088e610ffd565b600080838360408110156108a157600080fd5b506001805482356001600160a01b039081166001600160a01b03199283161790925560008054602090940135909216921691909117905550505050565b6009546001600160a01b03163314610922576040805162461bcd60e51b8152602060048201526002602482015261314560f01b604482015290519081900360640190fd5b60008054906101000a90046001600160a01b03166001600160a01b0316637af87d196040518163ffffffff1660e01b815260040160206040518083038186803b15801561096e57600080fd5b505afa158015610982573d6000803e3d6000fd5b505050506040513d602081101561099857600080fd5b505161ffff838116911614156109da576040805162461bcd60e51b8152602060048201526002602482015261316760f01b604482015290519081900360640190fd5b61ffff82166000908152600760205260409020546001600160a01b031615610a2e576040805162461bcd60e51b815260206004820152600260248201526118b360f11b604482015290519081900360640190fd5b6001600160a01b03811660009081526008602052604090205460ff1615610a81576040805162461bcd60e51b8152602060048201526002602482015261316560f01b604482015290519081900360640190fd5b60008054906101000a90046001600160a01b03166001600160a01b0316637af87d196040518163ffffffff1660e01b815260040160206040518083038186803b158015610acd57600080fd5b505afa158015610ae1573d6000803e3d6000fd5b505050506040513d6020811015610af757600080fd5b50516001600160a01b03821660009081526006602052604090205461ffff90811691161415610b52576040805162461bcd60e51b8152602060048201526002602482015261316760f01b604482015290519081900360640190fd5b6001805461ffff60a01b198116600160a01b9182900461ffff908116840181169092021782556001600160a01b0383166000818152600660209081526040808320805461ffff191695891695861790558483526007825280832080546001600160a01b031916851790558383526008909152808220805460ff19169095179094559251919290917f990c18cf226b253448a6a051b5cadf51a61f4ca56703374cdd8bf9df04b2f9019190a35050565b60026020526000908152604090205460ff1681565b60036020526000908152604090205460ff1681565b6001600160a01b03811660009081526003602052604090205460ff16610690576040805162461bcd60e51b815260206004820152601060248201526f34b73b30b634b2103b32b934b334b2b960811b604482015290519081900360640190fd5b6001600160a01b0380821660009081526006602090815260408083205483548251637af87d1960e01b81529251949561ffff90921694911692637af87d199260048082019391829003018186803b158015610ce557600080fd5b505afa158015610cf9573d6000803e3d6000fd5b505050506040513d6020811015610d0f57600080fd5b505161ffff82811691161415610d51576040805162461bcd60e51b8152602060048201526002602482015261316960f01b604482015290519081900360640190fd5b6001600160a01b03831660009081526008602052604090205460ff16610da3576040805162461bcd60e51b8152602060048201526002602482015261326960f01b604482015290519081900360640190fd5b92915050565b6001600160a01b03811660009081526002602052604090205460ff16610690576040805162461bcd60e51b815260206004820152601160248201527034b73b30b634b21031b7b6b6b4ba3a32b960791b604482015290519081900360640190fd5b610e1333610fb6565b6001600160a01b03821660009081526003602052604090205460ff1615158115151461062d576001600160a01b038216600081815260036020908152604091829020805460ff1916851515908117909155825190815291517ff09ef994a79ba4638c91b1d9595c89e208c3a9b9403769b68670c678782f82d09281900390910190a25050565b610ea233610fb6565b6009546001600160a01b0382811691161461069057600980546001600160a01b0383166001600160a01b0319909116811790915560408051918252517fc7708091594580aae77e08cb1e2e3b9d2bd0cab0d2d95797248dd95f02d7299e9181900360200190a150565b6007602052600090815260409020546001600160a01b031681565b610f2f33610fb6565b6001546001600160a01b0382811691161461069057600180546001600160a01b0383166001600160a01b0319909116811790915560408051918252517f5425363a03f182281120f5919107c49c7a1a623acc1cbc6df468b6f0c11fcf8c9181900360200190a150565b6001546001600160a01b031681565b6009546001600160a01b031681565b6001546001600160a01b03828116911614610690576040805162461bcd60e51b8152602060048201526002602482015261316760f01b604482015290519081900360640190fd5b7f8e94fed44239eb2314ab7a406345e6c5a8f0ccedf3b600de3d004e672c33abf4805460019091558015610690576040805162461bcd60e51b815260206004820152600260248201526118a160f11b604482015290519081900360640190fdfea26469706673582212208296ed07e151b8f63b8a3a49e798b48f7b2570b7fa1972307ce6369dbc8c7d5664736f6c63430007060033"

// DeployAssetGovernance deploys a new Ethereum zecrey, binding an instance of AssetGovernance to it.
func DeployAssetGovernance(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AssetGovernance, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetGovernanceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AssetGovernanceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AssetGovernance{AssetGovernanceCaller: AssetGovernanceCaller{contract: contract}, AssetGovernanceTransactor: AssetGovernanceTransactor{contract: contract}, AssetGovernanceFilterer: AssetGovernanceFilterer{contract: contract}}, nil
}

// AssetGovernance is an auto generated Go binding around an Ethereum zecrey.
type AssetGovernance struct {
	AssetGovernanceCaller     // Read-only binding to the zecrey
	AssetGovernanceTransactor // Write-only binding to the zecrey
	AssetGovernanceFilterer   // Log filterer for zecrey events
}

// AssetGovernanceCaller is an auto generated read-only Go binding around an Ethereum zecrey.
type AssetGovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetGovernanceTransactor is an auto generated write-only Go binding around an Ethereum zecrey.
type AssetGovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetGovernanceFilterer is an auto generated log filtering Go binding around an Ethereum zecrey events.
type AssetGovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetGovernanceSession is an auto generated Go binding around an Ethereum zecrey,
// with pre-set call and transact options.
type AssetGovernanceSession struct {
	Contract     *AssetGovernance  // Generic zecrey binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetGovernanceCallerSession is an auto generated read-only Go binding around an Ethereum zecrey,
// with pre-set call options.
type AssetGovernanceCallerSession struct {
	Contract *AssetGovernanceCaller // Generic zecrey caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AssetGovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum zecrey,
// with pre-set transact options.
type AssetGovernanceTransactorSession struct {
	Contract     *AssetGovernanceTransactor // Generic zecrey transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AssetGovernanceRaw is an auto generated low-level Go binding around an Ethereum zecrey.
type AssetGovernanceRaw struct {
	Contract *AssetGovernance // Generic zecrey binding to access the raw methods on
}

// AssetGovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum zecrey.
type AssetGovernanceCallerRaw struct {
	Contract *AssetGovernanceCaller // Generic read-only zecrey binding to access the raw methods on
}

// AssetGovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum zecrey.
type AssetGovernanceTransactorRaw struct {
	Contract *AssetGovernanceTransactor // Generic write-only zecrey binding to access the raw methods on
}

// NewAssetGovernance creates a new instance of AssetGovernance, bound to a specific deployed zecrey.
func NewAssetGovernance(address common.Address, backend bind.ContractBackend) (*AssetGovernance, error) {
	contract, err := bindAssetGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssetGovernance{AssetGovernanceCaller: AssetGovernanceCaller{contract: contract}, AssetGovernanceTransactor: AssetGovernanceTransactor{contract: contract}, AssetGovernanceFilterer: AssetGovernanceFilterer{contract: contract}}, nil
}

// NewAssetGovernanceCaller creates a new read-only instance of AssetGovernance, bound to a specific deployed zecrey.
func NewAssetGovernanceCaller(address common.Address, caller bind.ContractCaller) (*AssetGovernanceCaller, error) {
	contract, err := bindAssetGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetGovernanceCaller{contract: contract}, nil
}

// NewAssetGovernanceTransactor creates a new write-only instance of AssetGovernance, bound to a specific deployed zecrey.
func NewAssetGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetGovernanceTransactor, error) {
	contract, err := bindAssetGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetGovernanceTransactor{contract: contract}, nil
}

// NewAssetGovernanceFilterer creates a new log filterer instance of AssetGovernance, bound to a specific deployed zecrey.
func NewAssetGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*AssetGovernanceFilterer, error) {
	contract, err := bindAssetGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetGovernanceFilterer{contract: contract}, nil
}

// bindAssetGovernance binds a generic wrapper to an already deployed zecrey.
func bindAssetGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetGovernanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) zecrey method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetGovernance *AssetGovernanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetGovernance.Contract.AssetGovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the zecrey, calling
// its default method if one is available.
func (_AssetGovernance *AssetGovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetGovernance.Contract.AssetGovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) zecrey method with params as input values.
func (_AssetGovernance *AssetGovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetGovernance.Contract.AssetGovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) zecrey method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetGovernance *AssetGovernanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetGovernance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the zecrey, calling
// its default method if one is available.
func (_AssetGovernance *AssetGovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetGovernance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) zecrey method with params as input values.
func (_AssetGovernance *AssetGovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetGovernance.Contract.contract.Transact(opts, method, params...)
}

// Governance is a free data retrieval call binding the zecrey method 0x5aa6e675.
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

// Governance is a free data retrieval call binding the zecrey method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_AssetGovernance *AssetGovernanceSession) Governance() (common.Address, error) {
	return _AssetGovernance.Contract.Governance(&_AssetGovernance.CallOpts)
}

// Governance is a free data retrieval call binding the zecrey method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_AssetGovernance *AssetGovernanceCallerSession) Governance() (common.Address, error) {
	return _AssetGovernance.Contract.Governance(&_AssetGovernance.CallOpts)
}

// ListingCap is a free data retrieval call binding the zecrey method 0x9b453945.
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

// ListingCap is a free data retrieval call binding the zecrey method 0x9b453945.
//
// Solidity: function listingCap() view returns(uint16)
func (_AssetGovernance *AssetGovernanceSession) ListingCap() (uint16, error) {
	return _AssetGovernance.Contract.ListingCap(&_AssetGovernance.CallOpts)
}

// ListingCap is a free data retrieval call binding the zecrey method 0x9b453945.
//
// Solidity: function listingCap() view returns(uint16)
func (_AssetGovernance *AssetGovernanceCallerSession) ListingCap() (uint16, error) {
	return _AssetGovernance.Contract.ListingCap(&_AssetGovernance.CallOpts)
}

// TokenLister is a free data retrieval call binding the zecrey method 0xaf3d414b.
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

// TokenLister is a free data retrieval call binding the zecrey method 0xaf3d414b.
//
// Solidity: function tokenLister(address ) view returns(bool)
func (_AssetGovernance *AssetGovernanceSession) TokenLister(arg0 common.Address) (bool, error) {
	return _AssetGovernance.Contract.TokenLister(&_AssetGovernance.CallOpts, arg0)
}

// TokenLister is a free data retrieval call binding the zecrey method 0xaf3d414b.
//
// Solidity: function tokenLister(address ) view returns(bool)
func (_AssetGovernance *AssetGovernanceCallerSession) TokenLister(arg0 common.Address) (bool, error) {
	return _AssetGovernance.Contract.TokenLister(&_AssetGovernance.CallOpts, arg0)
}

// Treasury is a free data retrieval call binding the zecrey method 0x61d027b3.
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

// Treasury is a free data retrieval call binding the zecrey method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_AssetGovernance *AssetGovernanceSession) Treasury() (common.Address, error) {
	return _AssetGovernance.Contract.Treasury(&_AssetGovernance.CallOpts)
}

// Treasury is a free data retrieval call binding the zecrey method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_AssetGovernance *AssetGovernanceCallerSession) Treasury() (common.Address, error) {
	return _AssetGovernance.Contract.Treasury(&_AssetGovernance.CallOpts)
}

// Initialize is a paid mutator transaction binding the zecrey method 0x439fab91.
//
// Solidity: function initialize(bytes initializationParameters) returns()
func (_AssetGovernance *AssetGovernanceTransactor) Initialize(opts *bind.TransactOpts, initializationParameters []byte) (*types.Transaction, error) {
	return _AssetGovernance.contract.Transact(opts, "initialize", initializationParameters)
}

// Initialize is a paid mutator transaction binding the zecrey method 0x439fab91.
//
// Solidity: function initialize(bytes initializationParameters) returns()
func (_AssetGovernance *AssetGovernanceSession) Initialize(initializationParameters []byte) (*types.Transaction, error) {
	return _AssetGovernance.Contract.Initialize(&_AssetGovernance.TransactOpts, initializationParameters)
}

// Initialize is a paid mutator transaction binding the zecrey method 0x439fab91.
//
// Solidity: function initialize(bytes initializationParameters) returns()
func (_AssetGovernance *AssetGovernanceTransactorSession) Initialize(initializationParameters []byte) (*types.Transaction, error) {
	return _AssetGovernance.Contract.Initialize(&_AssetGovernance.TransactOpts, initializationParameters)
}

// SetAsset is a paid mutator transaction binding the zecrey method 0x5bab15ca.
//
// Solidity: function setAsset(uint16 _assetId, address _assetAddress) returns()
func (_AssetGovernance *AssetGovernanceTransactor) SetAsset(opts *bind.TransactOpts, _assetId uint16, _assetAddress common.Address) (*types.Transaction, error) {
	return _AssetGovernance.contract.Transact(opts, "setAsset", _assetId, _assetAddress)
}

// SetAsset is a paid mutator transaction binding the zecrey method 0x5bab15ca.
//
// Solidity: function setAsset(uint16 _assetId, address _assetAddress) returns()
func (_AssetGovernance *AssetGovernanceSession) SetAsset(_assetId uint16, _assetAddress common.Address) (*types.Transaction, error) {
	return _AssetGovernance.Contract.SetAsset(&_AssetGovernance.TransactOpts, _assetId, _assetAddress)
}

// SetAsset is a paid mutator transaction binding the zecrey method 0x5bab15ca.
//
// Solidity: function setAsset(uint16 _assetId, address _assetAddress) returns()
func (_AssetGovernance *AssetGovernanceTransactorSession) SetAsset(_assetId uint16, _assetAddress common.Address) (*types.Transaction, error) {
	return _AssetGovernance.Contract.SetAsset(&_AssetGovernance.TransactOpts, _assetId, _assetAddress)
}

// SetLister is a paid mutator transaction binding the zecrey method 0x719c150c.
//
// Solidity: function setLister(address _listerAddress, bool _active) returns()
func (_AssetGovernance *AssetGovernanceTransactor) SetLister(opts *bind.TransactOpts, _listerAddress common.Address, _active bool) (*types.Transaction, error) {
	return _AssetGovernance.contract.Transact(opts, "setLister", _listerAddress, _active)
}

// SetLister is a paid mutator transaction binding the zecrey method 0x719c150c.
//
// Solidity: function setLister(address _listerAddress, bool _active) returns()
func (_AssetGovernance *AssetGovernanceSession) SetLister(_listerAddress common.Address, _active bool) (*types.Transaction, error) {
	return _AssetGovernance.Contract.SetLister(&_AssetGovernance.TransactOpts, _listerAddress, _active)
}

// SetLister is a paid mutator transaction binding the zecrey method 0x719c150c.
//
// Solidity: function setLister(address _listerAddress, bool _active) returns()
func (_AssetGovernance *AssetGovernanceTransactorSession) SetLister(_listerAddress common.Address, _active bool) (*types.Transaction, error) {
	return _AssetGovernance.Contract.SetLister(&_AssetGovernance.TransactOpts, _listerAddress, _active)
}

// Upgrade is a paid mutator transaction binding the zecrey method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_AssetGovernance *AssetGovernanceTransactor) Upgrade(opts *bind.TransactOpts, upgradeParameters []byte) (*types.Transaction, error) {
	return _AssetGovernance.contract.Transact(opts, "upgrade", upgradeParameters)
}

// Upgrade is a paid mutator transaction binding the zecrey method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_AssetGovernance *AssetGovernanceSession) Upgrade(upgradeParameters []byte) (*types.Transaction, error) {
	return _AssetGovernance.Contract.Upgrade(&_AssetGovernance.TransactOpts, upgradeParameters)
}

// Upgrade is a paid mutator transaction binding the zecrey method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_AssetGovernance *AssetGovernanceTransactorSession) Upgrade(upgradeParameters []byte) (*types.Transaction, error) {
	return _AssetGovernance.Contract.Upgrade(&_AssetGovernance.TransactOpts, upgradeParameters)
}

// AssetGovernanceTokenListerUpdateIterator is returned from FilterTokenListerUpdate and is used to iterate over the raw logs and unpacked data for TokenListerUpdate events raised by the AssetGovernance zecrey.
type AssetGovernanceTokenListerUpdateIterator struct {
	Event *AssetGovernanceTokenListerUpdate // Event containing the zecrey specifics and raw log

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

// AssetGovernanceTokenListerUpdate represents a TokenListerUpdate event raised by the AssetGovernance zecrey.
type AssetGovernanceTokenListerUpdate struct {
	TokenLister common.Address
	IsActive    bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenListerUpdate is a free log retrieval operation binding the zecrey event 0x947516bb7f07e641084b9d05943fb1c89e05da3d4d17f707ab4efd6ab58f0910.
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

// WatchTokenListerUpdate is a free log subscription operation binding the zecrey event 0x947516bb7f07e641084b9d05943fb1c89e05da3d4d17f707ab4efd6ab58f0910.
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

// ParseTokenListerUpdate is a log parse operation binding the zecrey event 0x947516bb7f07e641084b9d05943fb1c89e05da3d4d17f707ab4efd6ab58f0910.
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
