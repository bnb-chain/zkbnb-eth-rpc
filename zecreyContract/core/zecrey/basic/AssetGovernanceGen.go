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
const AssetGovernanceABI = "[{\"inputs\":[{\"internalType\":\"contractGovernance\",\"name\":\"_governance\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_listingFeeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_listingFee\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_listingCap\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newListingCap\",\"type\":\"uint16\"}],\"name\":\"ListingCapUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"newListingFeeToken\",\"type\":\"address\"}],\"name\":\"ListingFeeTokenUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newListingFee\",\"type\":\"uint256\"}],\"name\":\"ListingFeeUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenLister\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"TokenListerUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"TreasuryUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"contractGovernance\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listingCap\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listingFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listingFeeToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_assetId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_assetAddress\",\"type\":\"address\"}],\"name\":\"setAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_listerAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"setLister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_newListingCap\",\"type\":\"uint16\"}],\"name\":\"setListingCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newListingFee\",\"type\":\"uint256\"}],\"name\":\"setListingFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_newListingFeeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newListingFee\",\"type\":\"uint256\"}],\"name\":\"setListingFeeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newTreasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenLister\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AssetGovernanceBin is the compiled bytecode used for deploying new contracts.
var AssetGovernanceBin = "0x608060405234801561001057600080fd5b50604051610ffc380380610ffc8339818101604052606081101561003357600080fd5b50805160208201516040909201516000805460ff191660ff9093169290921762ffff00191661010061ffff948516021764ffff000000191663010000009390911692909202919091178155610f6e90819061008e90396000f3fe608060405234801561001057600080fd5b50600436106101a95760003560e01c80636c824487116100f9578063d021f32a11610097578063e4c0aaf411610071578063e4c0aaf4146105fa578063f39349ef14610620578063f5e7d6fd14610628578063f5f84ed414610630576101a9565b8063d021f32a14610571578063d87e374814610597578063dbfc2967146105bd576101a9565b80639ac2a011116100d35780639ac2a011146104d15780639bd77609146104f7578063a28e00321461051d578063ac9b567114610543576101a9565b80636c824487146104645780637a3459d41461048a578063953f9371146104b0576101a9565b8063253946451161016657806331d8687b1161014057806331d8687b1461037d578063439fab911461039e5780635bab15ca1461040e5780635f4fea261461043e576101a9565b806325394645146102a557806326eb0f4a1461031557806328b58e1a14610343576101a9565b806301e1d114146101ae5780631afcac96146101cd5780631c021956146101fd5780631e1bff3f146102235780631e763ee3146102515780632520ce5a14610277575b600080fd5b6101b6610656565b6040805161ffff9092168252519081900360200190f35b6101fb600480360360408110156101e357600080fd5b506001600160a01b0381351690602001351515610667565b005b6101fb6004803603602081101561021357600080fd5b50356001600160a01b03166106f7565b6101fb6004803603604081101561023957600080fd5b506001600160a01b0381351690602001351515610759565b6101b66004803603602081101561026757600080fd5b50356001600160a01b03166107e8565b6101fb6004803603604081101561028d57600080fd5b506001600160a01b03813516906020013515156107fe565b6101fb600480360360208110156102bb57600080fd5b8101906020810181356401000000008111156102d657600080fd5b8201836020820111156102e857600080fd5b8035906020019184600183028401116401000000008311171561030a57600080fd5b5090925090506106f3565b6101fb6004803603604081101561032b57600080fd5b506001600160a01b038135169060200135151561090d565b6103696004803603602081101561035957600080fd5b50356001600160a01b031661099c565b604080519115158252519081900360200190f35b6103696004803603602081101561039357600080fd5b503561ffff166109b1565b6101fb600480360360208110156103b457600080fd5b8101906020810181356401000000008111156103cf57600080fd5b8201836020820111156103e157600080fd5b8035906020019184600183028401116401000000008311171561040357600080fd5b5090925090506109c6565b6101fb6004803603604081101561042457600080fd5b50803561ffff1690602001356001600160a01b0316610a08565b6103696004803603602081101561045457600080fd5b50356001600160a01b0316610b43565b6103696004803603602081101561047a57600080fd5b50356001600160a01b0316610b58565b6101fb600480360360208110156104a057600080fd5b50356001600160a01b0316610b6d565b610369600480360360208110156104c657600080fd5b503561ffff16610bcd565b610369600480360360208110156104e757600080fd5b50356001600160a01b0316610be4565b6101b66004803603602081101561050d57600080fd5b50356001600160a01b0316610bf9565b6101fb6004803603602081101561053357600080fd5b50356001600160a01b0316610c61565b6101fb6004803603604081101561055957600080fd5b506001600160a01b0381351690602001351515610cc2565b6101fb6004803603602081101561058757600080fd5b50356001600160a01b0316610d51565b6101fb600480360360208110156105ad57600080fd5b50356001600160a01b0316610db1565b6105de600480360360208110156105d357600080fd5b503561ffff16610e23565b604080516001600160a01b039092168252519081900360200190f35b6101fb6004803603602081101561061057600080fd5b50356001600160a01b0316610e3e565b6105de610ec5565b6105de610edb565b6101fb6004803603602081101561064657600080fd5b50356001600160a01b0316610eea565b600054600160c81b900461ffff1681565b61067033610eea565b6001600160a01b03821660009081526001602052604090205460ff161515811515146106f3576001600160a01b038216600081815260016020908152604091829020805460ff1916851515908117909155825190815291517fd158a06928d89556421a4e2eabd98c7ea10e974c07e399e3afc449d19eb391b29281900390910190a25b5050565b6001600160a01b03811660009081526004602052604090205460ff16610756576040805162461bcd60e51b815260206004820152600f60248201526e34b73b30b634b21036b7b734ba37b960891b604482015290519081900360640190fd5b50565b61076233610eea565b6001600160a01b03821660009081526003602052604090205460ff161515811515146106f3576001600160a01b038216600081815260036020908152604091829020805460ff1916851515908117909155825190815291517f9a01f9639f4fd4c2eea316eabcc61b3674b410002afdb96bf605cf288d37f9da9281900390910190a25050565b60066020526000908152604090205461ffff1681565b61080733610eea565b6000306001600160a01b0316639bd77609846040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561085657600080fd5b505afa15801561086a573d6000803e3d6000fd5b505050506040513d602081101561088057600080fd5b505161ffff811660009081526005602052604090205490915060ff161515821515146109085761ffff8116600090815260056020908152604091829020805460ff1916851515908117909155825190815291516001600160a01b038616927ff7ca5545623b85829d3abcb6729eb021070bbe93cba64f5c4de6b17b805b7d0292908290030190a25b505050565b61091633610eea565b6001600160a01b03821660009081526004602052604090205460ff161515811515146106f3576001600160a01b038216600081815260046020908152604091829020805460ff1916851515908117909155825190815291517f6fbe3fec04b6d2f79e569f1a46ec4d1621e751112483b5ce9d81832f19fee5c19281900390910190a25050565b60046020526000908152604090205460ff1681565b60056020526000908152604090205460ff1681565b6000828260208110156109d857600080fd5b506000805465010000000000600160c81b03191691356001600160a01b0316600160281b02919091179055505050565b6008546001600160a01b03163314610a4c576040805162461bcd60e51b8152602060048201526002602482015261314560f01b604482015290519081900360640190fd5b600080546001600160a01b0383168252600660205260409091205461ffff908116610100909204161415610aac576040805162461bcd60e51b8152602060048201526002602482015261316560f01b604482015290519081900360640190fd5b6000805461ffff60c81b198116600160c81b9182900461ffff90811660010181169092021782556001600160a01b038316808352600660209081526040808520805461ffff19169488169485179055838552600790915280842080546001600160a01b0319168317905551919290917f990c18cf226b253448a6a051b5cadf51a61f4ca56703374cdd8bf9df04b2f9019190a35050565b60016020526000908152604090205460ff1681565b60026020526000908152604090205460ff1681565b6001600160a01b03811660009081526002602052604090205460ff16610756576040805162461bcd60e51b815260206004820152601060248201526f34b73b30b634b2103b32b934b334b2b960811b604482015290519081900360640190fd5b600054600160c81b900461ffff9081169116111590565b60036020526000908152604090205460ff1681565b6001600160a01b038116600090815260066020526040812054815461ffff9182169161010090910416811415610c5b576040805162461bcd60e51b8152602060048201526002602482015261316960f01b604482015290519081900360640190fd5b92915050565b6001600160a01b03811660009081526001602052604090205460ff16610756576040805162461bcd60e51b815260206004820152601160248201527034b73b30b634b21031b7b6b6b4ba3a32b960791b604482015290519081900360640190fd5b610ccb33610eea565b6001600160a01b03821660009081526002602052604090205460ff161515811515146106f3576001600160a01b038216600081815260026020908152604091829020805460ff1916851515908117909155825190815291517ff09ef994a79ba4638c91b1d9595c89e208c3a9b9403769b68670c678782f82d09281900390910190a25050565b6001600160a01b03811660009081526003602052604090205460ff16610756576040805162461bcd60e51b815260206004820152601060248201526f34b73b30b634b21032bc32b1baba37b960811b604482015290519081900360640190fd5b610dba33610eea565b6008546001600160a01b0382811691161461075657600880546001600160a01b0383166001600160a01b0319909116811790915560408051918252517fc7708091594580aae77e08cb1e2e3b9d2bd0cab0d2d95797248dd95f02d7299e9181900360200190a150565b6007602052600090815260409020546001600160a01b031681565b610e4733610eea565b6000546001600160a01b03828116600160281b909204161461075657600080546001600160a01b038316600160281b810265010000000000600160c81b03199092169190911790915560408051918252517f5425363a03f182281120f5919107c49c7a1a623acc1cbc6df468b6f0c11fcf8c9181900360200190a150565b600054600160281b90046001600160a01b031681565b6008546001600160a01b031681565b6000546001600160a01b03828116600160281b9092041614610756576040805162461bcd60e51b8152602060048201526002602482015261316760f01b604482015290519081900360640190fdfea2646970667358221220ff968465ea595be0acda19e3412da63769d2579044357b73b822dbc567c7b19264736f6c63430007060033"

// DeployAssetGovernance deploys a new Ethereum contract, binding an instance of AssetGovernance to it.
func DeployAssetGovernance(auth *bind.TransactOpts, backend bind.ContractBackend, _governance common.Address, _listingFeeToken common.Address, _listingFee *big.Int, _listingCap uint16, _treasury common.Address) (common.Address, *types.Transaction, *AssetGovernance, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetGovernanceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AssetGovernanceBin), backend, _governance, _listingFeeToken, _listingFee, _listingCap, _treasury)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AssetGovernance{AssetGovernanceCaller: AssetGovernanceCaller{contract: contract}, AssetGovernanceTransactor: AssetGovernanceTransactor{contract: contract}, AssetGovernanceFilterer: AssetGovernanceFilterer{contract: contract}}, nil
}

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

// SetAsset is a paid mutator transaction binding the contract method 0x5bab15ca.
//
// Solidity: function setAsset(uint16 _assetId, address _assetAddress) returns()
func (_AssetGovernance *AssetGovernanceTransactor) SetAsset(opts *bind.TransactOpts, _assetId uint16, _assetAddress common.Address) (*types.Transaction, error) {
	return _AssetGovernance.contract.Transact(opts, "setAsset", _assetId, _assetAddress)
}

// SetAsset is a paid mutator transaction binding the contract method 0x5bab15ca.
//
// Solidity: function setAsset(uint16 _assetId, address _assetAddress) returns()
func (_AssetGovernance *AssetGovernanceSession) SetAsset(_assetId uint16, _assetAddress common.Address) (*types.Transaction, error) {
	return _AssetGovernance.Contract.SetAsset(&_AssetGovernance.TransactOpts, _assetId, _assetAddress)
}

// SetAsset is a paid mutator transaction binding the contract method 0x5bab15ca.
//
// Solidity: function setAsset(uint16 _assetId, address _assetAddress) returns()
func (_AssetGovernance *AssetGovernanceTransactorSession) SetAsset(_assetId uint16, _assetAddress common.Address) (*types.Transaction, error) {
	return _AssetGovernance.Contract.SetAsset(&_AssetGovernance.TransactOpts, _assetId, _assetAddress)
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

// SetListingFeeToken is a paid mutator transaction binding the contract method 0x70de40b0.
//
// Solidity: function setListingFeeToken(address _newListingFeeToken, uint256 _newListingFee) returns()
func (_AssetGovernance *AssetGovernanceTransactor) SetListingFeeToken(opts *bind.TransactOpts, _newListingFeeToken common.Address, _newListingFee *big.Int) (*types.Transaction, error) {
	return _AssetGovernance.contract.Transact(opts, "setListingFeeToken", _newListingFeeToken, _newListingFee)
}

// SetListingFeeToken is a paid mutator transaction binding the contract method 0x70de40b0.
//
// Solidity: function setListingFeeToken(address _newListingFeeToken, uint256 _newListingFee) returns()
func (_AssetGovernance *AssetGovernanceSession) SetListingFeeToken(_newListingFeeToken common.Address, _newListingFee *big.Int) (*types.Transaction, error) {
	return _AssetGovernance.Contract.SetListingFeeToken(&_AssetGovernance.TransactOpts, _newListingFeeToken, _newListingFee)
}

// SetListingFeeToken is a paid mutator transaction binding the contract method 0x70de40b0.
//
// Solidity: function setListingFeeToken(address _newListingFeeToken, uint256 _newListingFee) returns()
func (_AssetGovernance *AssetGovernanceTransactorSession) SetListingFeeToken(_newListingFeeToken common.Address, _newListingFee *big.Int) (*types.Transaction, error) {
	return _AssetGovernance.Contract.SetListingFeeToken(&_AssetGovernance.TransactOpts, _newListingFeeToken, _newListingFee)
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
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterListingFeeTokenUpdate is a free log retrieval operation binding the contract event 0x3246cb0d7c19b6340b39fa43b91e3b663fb1eb1a2f5eb747c8f09646f6300555.
//
// Solidity: event ListingFeeTokenUpdate(address indexed newListingFeeToken)
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

// WatchListingFeeTokenUpdate is a free log subscription operation binding the contract event 0x3246cb0d7c19b6340b39fa43b91e3b663fb1eb1a2f5eb747c8f09646f6300555.
//
// Solidity: event ListingFeeTokenUpdate(address indexed newListingFeeToken)
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

// ParseListingFeeTokenUpdate is a log parse operation binding the contract event 0x3246cb0d7c19b6340b39fa43b91e3b663fb1eb1a2f5eb747c8f09646f6300555.
//
// Solidity: event ListingFeeTokenUpdate(address indexed newListingFeeToken)
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
