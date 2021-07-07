// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zecrey

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

// GovernanceABI is the input ABI used to generate the binding from.
const GovernanceABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"name\":\"AssetUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"providerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"name\":\"RollupProviderUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"}],\"name\":\"addAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"assetIds\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"assetId\",\"type\":\"uint32\"}],\"name\":\"getSupportedAsset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAssetEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAssetExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"assetId\",\"type\":\"uint32\"}],\"name\":\"isValidAsset\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"}],\"name\":\"isValidOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rollupProviders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"supportedAssets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"assetId\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"updateAssetStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"updateRollupProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GovernanceBin is the compiled bytecode used for deploying new contracts.
var GovernanceBin = "0x608060405260016201000560006101000a81548163ffffffff021916908363ffffffff1602179055503480156100355760006000fd5b505b5b600061004861020360201b60201c565b905080600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505b6000600090508060046000506000620100008110151561010457fe5b90900160005b6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600260005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600360005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550505b610210565b600033905061020d565b90565b6116118061021f6000396000f3fe60806040523480156100115760006000fd5b50600436106100ef5760003560e01c8063715018a61161008d578063d8e8dbc711610067578063d8e8dbc7146103bf578063e2d53d7c1461041e578063e53cea3314610479578063f2fde38b146104d8576100ef565b8063715018a6146103285780638da5cb5b14610332578063c68dbb3714610366576100ef565b806350921dc8116100c957806350921dc8146101d65780635ad1def3146102275780635b7df2b91461028257806367eb04e6146102cd576100ef565b8063114cb92f146100f5578063298410e514610150578063417a1c7914610195576100ef565b60006000fd5b6101386004803603602081101561010c5760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061051d565b60405180821515815260200191505060405180910390f35b610193600480360360208110156101675760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610542565b005b6101d4600480360360408110156101ac5760006000fd5b81019080803563ffffffff16906020019092919080351515906020019092919050505061098b565b005b610225600480360360408110156101ed5760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803515159060200190929190505050610c88565b005b61026a6004803603602081101561023e5760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610df8565b60405180821515815260200191505060405180910390f35b6102b5600480360360208110156102995760006000fd5b81019080803563ffffffff169060200190929190505050610e1d565b60405180821515815260200191505060405180910390f35b610310600480360360208110156102e45760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610fe9565b60405180821515815260200191505060405180910390f35b61033061100e565b005b61033a611191565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6103936004803603602081101561037d5760006000fd5b81019080803590602001909291905050506111c0565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610402600480360360208110156103d65760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611200565b604051808263ffffffff16815260200191505060405180910390f35b610461600480360360208110156104355760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061122a565b60405180821515815260200191505060405180910390f35b6104ac600480360360208110156104905760006000fd5b81019080803563ffffffff169060200190929190505050611288565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61051b600480360360208110156104ef5760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506112d5565b005b600260005060205280600052604060002060009150909054906101000a900460ff1681565b6105506114df63ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff1661057461119163ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff16141515610602576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281526020015060200191505060405180910390fd5b610611816114ec63ffffffff16565b1515610668576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806115946022913960400191505060405180910390fd5b600360005060008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151515610730576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f617373657420616c72656164792065786973747300000000000000000000000081526020015060200191505060405180910390fd5b6001600260005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600360005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508060046000506201000560009054906101000a900463ffffffff1663ffffffff16620100008110151561081557fe5b90900160005b6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506201000560009054906101000a900463ffffffff166201000460005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff16021790555060016201000560008282829054906101000a900463ffffffff160192506101000a81548163ffffffff021916908363ffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff1661094c60016201000560009054906101000a900463ffffffff1663ffffffff1661150690919063ffffffff16565b7fce6315d17c79350e737a11e641e32024ae26a2500d55b7e62cc61268811c7fbb600160405180821515815260200191505060405180910390a35b5b50565b6109996114df63ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff166109bd61119163ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff16141515610a4b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281526020015060200191505060405180910390fd5b6201000560009054906101000a900463ffffffff1663ffffffff168263ffffffff16101515610ae5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f696e76616c69642061737365742069640000000000000000000000000000000081526020015060200191505060405180910390fd5b6003600050600060046000508463ffffffff166201000081101515610b0657fe5b90900160005b9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515610bea576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f6173736574206e6f74206578697374730000000000000000000000000000000081526020015060200191505060405180910390fd5b806002600050600060046000508563ffffffff166201000081101515610c0c57fe5b90900160005b9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b5b5050565b610c966114df63ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff16610cba61119163ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff16141515610d48576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281526020015060200191505060405180910390fd5b80600160005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff167f46359ce9dbb6c7f9a375b44072210287916d3de725fc8927a8e762047e4a84248260405180821515815260200191505060405180910390a25b5b5050565b600160005060205280600052604060002060009150909054906101000a900460ff1681565b60006201000560009054906101000a900463ffffffff1663ffffffff168263ffffffff16101515610eb9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f696e76616c69642061737365744964000000000000000000000000000000000081526020015060200191505060405180910390fd5b6003600050600060046000508463ffffffff166201000081101515610eda57fe5b90900160005b9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff168015610fdd57506002600050600060046000508463ffffffff166201000081101515610f6f57fe5b90900160005b9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff165b9050610fe4565b919050565b600360005060205280600052604060002060009150909054906101000a900460ff1681565b61101c6114df63ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff1661104061119163ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff161415156110ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281526020015060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5b565b6000600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506111bd565b90565b6004600050816201000081106111d557600080fd5b90900160005b9150909054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6201000460005060205280600052604060002060009150909054906101000a900463ffffffff1681565b6000600160005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050611283565b919050565b600060046000508263ffffffff1662010000811015156112a457fe5b90900160005b9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506112d0565b919050565b6112e36114df63ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff1661130761119163ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff16141515611395576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281526020015060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561141d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806115b66026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5b50565b60003390506114e9565b90565b60006000823b90506000811191505061150156505b919050565b6000828211151515611583576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081526020015060200191505060405180910390fd5b818303905061158d565b9291505056fe616464726573732073686f756c6420626520636f6e747261637420616464726573734f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373a2646970667358221220ecdd93383e59b1c7cbb84059fd05266545b677a6142f42b075a8b8bde0aee46a64736f6c63430007060033"

// DeployGovernance deploys a new Ethereum contract, binding an instance of Governance to it.
func DeployGovernance(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Governance, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GovernanceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Governance{GovernanceCaller: GovernanceCaller{contract: contract}, GovernanceTransactor: GovernanceTransactor{contract: contract}, GovernanceFilterer: GovernanceFilterer{contract: contract}}, nil
}

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

// AssetIds is a free data retrieval call binding the contract method 0xd8e8dbc7.
//
// Solidity: function assetIds(address ) view returns(uint32)
func (_Governance *GovernanceCaller) AssetIds(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "assetIds", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AssetIds is a free data retrieval call binding the contract method 0xd8e8dbc7.
//
// Solidity: function assetIds(address ) view returns(uint32)
func (_Governance *GovernanceSession) AssetIds(arg0 common.Address) (uint32, error) {
	return _Governance.Contract.AssetIds(&_Governance.CallOpts, arg0)
}

// AssetIds is a free data retrieval call binding the contract method 0xd8e8dbc7.
//
// Solidity: function assetIds(address ) view returns(uint32)
func (_Governance *GovernanceCallerSession) AssetIds(arg0 common.Address) (uint32, error) {
	return _Governance.Contract.AssetIds(&_Governance.CallOpts, arg0)
}

// GetSupportedAsset is a free data retrieval call binding the contract method 0xe53cea33.
//
// Solidity: function getSupportedAsset(uint32 assetId) view returns(address)
func (_Governance *GovernanceCaller) GetSupportedAsset(opts *bind.CallOpts, assetId uint32) (common.Address, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getSupportedAsset", assetId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSupportedAsset is a free data retrieval call binding the contract method 0xe53cea33.
//
// Solidity: function getSupportedAsset(uint32 assetId) view returns(address)
func (_Governance *GovernanceSession) GetSupportedAsset(assetId uint32) (common.Address, error) {
	return _Governance.Contract.GetSupportedAsset(&_Governance.CallOpts, assetId)
}

// GetSupportedAsset is a free data retrieval call binding the contract method 0xe53cea33.
//
// Solidity: function getSupportedAsset(uint32 assetId) view returns(address)
func (_Governance *GovernanceCallerSession) GetSupportedAsset(assetId uint32) (common.Address, error) {
	return _Governance.Contract.GetSupportedAsset(&_Governance.CallOpts, assetId)
}

// IsAssetEnabled is a free data retrieval call binding the contract method 0x114cb92f.
//
// Solidity: function isAssetEnabled(address ) view returns(bool)
func (_Governance *GovernanceCaller) IsAssetEnabled(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "isAssetEnabled", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAssetEnabled is a free data retrieval call binding the contract method 0x114cb92f.
//
// Solidity: function isAssetEnabled(address ) view returns(bool)
func (_Governance *GovernanceSession) IsAssetEnabled(arg0 common.Address) (bool, error) {
	return _Governance.Contract.IsAssetEnabled(&_Governance.CallOpts, arg0)
}

// IsAssetEnabled is a free data retrieval call binding the contract method 0x114cb92f.
//
// Solidity: function isAssetEnabled(address ) view returns(bool)
func (_Governance *GovernanceCallerSession) IsAssetEnabled(arg0 common.Address) (bool, error) {
	return _Governance.Contract.IsAssetEnabled(&_Governance.CallOpts, arg0)
}

// IsAssetExists is a free data retrieval call binding the contract method 0x67eb04e6.
//
// Solidity: function isAssetExists(address ) view returns(bool)
func (_Governance *GovernanceCaller) IsAssetExists(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "isAssetExists", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAssetExists is a free data retrieval call binding the contract method 0x67eb04e6.
//
// Solidity: function isAssetExists(address ) view returns(bool)
func (_Governance *GovernanceSession) IsAssetExists(arg0 common.Address) (bool, error) {
	return _Governance.Contract.IsAssetExists(&_Governance.CallOpts, arg0)
}

// IsAssetExists is a free data retrieval call binding the contract method 0x67eb04e6.
//
// Solidity: function isAssetExists(address ) view returns(bool)
func (_Governance *GovernanceCallerSession) IsAssetExists(arg0 common.Address) (bool, error) {
	return _Governance.Contract.IsAssetExists(&_Governance.CallOpts, arg0)
}

// IsValidAsset is a free data retrieval call binding the contract method 0x5b7df2b9.
//
// Solidity: function isValidAsset(uint32 assetId) view returns(bool)
func (_Governance *GovernanceCaller) IsValidAsset(opts *bind.CallOpts, assetId uint32) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "isValidAsset", assetId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidAsset is a free data retrieval call binding the contract method 0x5b7df2b9.
//
// Solidity: function isValidAsset(uint32 assetId) view returns(bool)
func (_Governance *GovernanceSession) IsValidAsset(assetId uint32) (bool, error) {
	return _Governance.Contract.IsValidAsset(&_Governance.CallOpts, assetId)
}

// IsValidAsset is a free data retrieval call binding the contract method 0x5b7df2b9.
//
// Solidity: function isValidAsset(uint32 assetId) view returns(bool)
func (_Governance *GovernanceCallerSession) IsValidAsset(assetId uint32) (bool, error) {
	return _Governance.Contract.IsValidAsset(&_Governance.CallOpts, assetId)
}

// IsValidOperator is a free data retrieval call binding the contract method 0xe2d53d7c.
//
// Solidity: function isValidOperator(address providerAddr) view returns(bool)
func (_Governance *GovernanceCaller) IsValidOperator(opts *bind.CallOpts, providerAddr common.Address) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "isValidOperator", providerAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidOperator is a free data retrieval call binding the contract method 0xe2d53d7c.
//
// Solidity: function isValidOperator(address providerAddr) view returns(bool)
func (_Governance *GovernanceSession) IsValidOperator(providerAddr common.Address) (bool, error) {
	return _Governance.Contract.IsValidOperator(&_Governance.CallOpts, providerAddr)
}

// IsValidOperator is a free data retrieval call binding the contract method 0xe2d53d7c.
//
// Solidity: function isValidOperator(address providerAddr) view returns(bool)
func (_Governance *GovernanceCallerSession) IsValidOperator(providerAddr common.Address) (bool, error) {
	return _Governance.Contract.IsValidOperator(&_Governance.CallOpts, providerAddr)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Governance *GovernanceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Governance *GovernanceSession) Owner() (common.Address, error) {
	return _Governance.Contract.Owner(&_Governance.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Governance *GovernanceCallerSession) Owner() (common.Address, error) {
	return _Governance.Contract.Owner(&_Governance.CallOpts)
}

// RollupProviders is a free data retrieval call binding the contract method 0x5ad1def3.
//
// Solidity: function rollupProviders(address ) view returns(bool)
func (_Governance *GovernanceCaller) RollupProviders(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "rollupProviders", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RollupProviders is a free data retrieval call binding the contract method 0x5ad1def3.
//
// Solidity: function rollupProviders(address ) view returns(bool)
func (_Governance *GovernanceSession) RollupProviders(arg0 common.Address) (bool, error) {
	return _Governance.Contract.RollupProviders(&_Governance.CallOpts, arg0)
}

// RollupProviders is a free data retrieval call binding the contract method 0x5ad1def3.
//
// Solidity: function rollupProviders(address ) view returns(bool)
func (_Governance *GovernanceCallerSession) RollupProviders(arg0 common.Address) (bool, error) {
	return _Governance.Contract.RollupProviders(&_Governance.CallOpts, arg0)
}

// SupportedAssets is a free data retrieval call binding the contract method 0xc68dbb37.
//
// Solidity: function supportedAssets(uint256 ) view returns(address)
func (_Governance *GovernanceCaller) SupportedAssets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "supportedAssets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SupportedAssets is a free data retrieval call binding the contract method 0xc68dbb37.
//
// Solidity: function supportedAssets(uint256 ) view returns(address)
func (_Governance *GovernanceSession) SupportedAssets(arg0 *big.Int) (common.Address, error) {
	return _Governance.Contract.SupportedAssets(&_Governance.CallOpts, arg0)
}

// SupportedAssets is a free data retrieval call binding the contract method 0xc68dbb37.
//
// Solidity: function supportedAssets(uint256 ) view returns(address)
func (_Governance *GovernanceCallerSession) SupportedAssets(arg0 *big.Int) (common.Address, error) {
	return _Governance.Contract.SupportedAssets(&_Governance.CallOpts, arg0)
}

// AddAsset is a paid mutator transaction binding the contract method 0x298410e5.
//
// Solidity: function addAsset(address assetAddress) returns()
func (_Governance *GovernanceTransactor) AddAsset(opts *bind.TransactOpts, assetAddress common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "addAsset", assetAddress)
}

// AddAsset is a paid mutator transaction binding the contract method 0x298410e5.
//
// Solidity: function addAsset(address assetAddress) returns()
func (_Governance *GovernanceSession) AddAsset(assetAddress common.Address) (*types.Transaction, error) {
	return _Governance.Contract.AddAsset(&_Governance.TransactOpts, assetAddress)
}

// AddAsset is a paid mutator transaction binding the contract method 0x298410e5.
//
// Solidity: function addAsset(address assetAddress) returns()
func (_Governance *GovernanceTransactorSession) AddAsset(assetAddress common.Address) (*types.Transaction, error) {
	return _Governance.Contract.AddAsset(&_Governance.TransactOpts, assetAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Governance *GovernanceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Governance *GovernanceSession) RenounceOwnership() (*types.Transaction, error) {
	return _Governance.Contract.RenounceOwnership(&_Governance.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Governance *GovernanceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Governance.Contract.RenounceOwnership(&_Governance.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Governance *GovernanceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Governance *GovernanceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Governance.Contract.TransferOwnership(&_Governance.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Governance *GovernanceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Governance.Contract.TransferOwnership(&_Governance.TransactOpts, newOwner)
}

// UpdateAssetStatus is a paid mutator transaction binding the contract method 0x417a1c79.
//
// Solidity: function updateAssetStatus(uint32 assetId, bool status) returns()
func (_Governance *GovernanceTransactor) UpdateAssetStatus(opts *bind.TransactOpts, assetId uint32, status bool) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "updateAssetStatus", assetId, status)
}

// UpdateAssetStatus is a paid mutator transaction binding the contract method 0x417a1c79.
//
// Solidity: function updateAssetStatus(uint32 assetId, bool status) returns()
func (_Governance *GovernanceSession) UpdateAssetStatus(assetId uint32, status bool) (*types.Transaction, error) {
	return _Governance.Contract.UpdateAssetStatus(&_Governance.TransactOpts, assetId, status)
}

// UpdateAssetStatus is a paid mutator transaction binding the contract method 0x417a1c79.
//
// Solidity: function updateAssetStatus(uint32 assetId, bool status) returns()
func (_Governance *GovernanceTransactorSession) UpdateAssetStatus(assetId uint32, status bool) (*types.Transaction, error) {
	return _Governance.Contract.UpdateAssetStatus(&_Governance.TransactOpts, assetId, status)
}

// UpdateRollupProvider is a paid mutator transaction binding the contract method 0x50921dc8.
//
// Solidity: function updateRollupProvider(address providerAddr, bool status) returns()
func (_Governance *GovernanceTransactor) UpdateRollupProvider(opts *bind.TransactOpts, providerAddr common.Address, status bool) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "updateRollupProvider", providerAddr, status)
}

// UpdateRollupProvider is a paid mutator transaction binding the contract method 0x50921dc8.
//
// Solidity: function updateRollupProvider(address providerAddr, bool status) returns()
func (_Governance *GovernanceSession) UpdateRollupProvider(providerAddr common.Address, status bool) (*types.Transaction, error) {
	return _Governance.Contract.UpdateRollupProvider(&_Governance.TransactOpts, providerAddr, status)
}

// UpdateRollupProvider is a paid mutator transaction binding the contract method 0x50921dc8.
//
// Solidity: function updateRollupProvider(address providerAddr, bool status) returns()
func (_Governance *GovernanceTransactorSession) UpdateRollupProvider(providerAddr common.Address, status bool) (*types.Transaction, error) {
	return _Governance.Contract.UpdateRollupProvider(&_Governance.TransactOpts, providerAddr, status)
}

// GovernanceAssetUpdatedIterator is returned from FilterAssetUpdated and is used to iterate over the raw logs and unpacked data for AssetUpdated events raised by the Governance contract.
type GovernanceAssetUpdatedIterator struct {
	Event *GovernanceAssetUpdated // Event containing the contract specifics and raw log

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
func (it *GovernanceAssetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceAssetUpdated)
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
		it.Event = new(GovernanceAssetUpdated)
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
func (it *GovernanceAssetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceAssetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceAssetUpdated represents a AssetUpdated event raised by the Governance contract.
type GovernanceAssetUpdated struct {
	AssetId      *big.Int
	AssetAddress common.Address
	IsEnabled    bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAssetUpdated is a free log retrieval operation binding the contract event 0xce6315d17c79350e737a11e641e32024ae26a2500d55b7e62cc61268811c7fbb.
//
// Solidity: event AssetUpdated(uint256 indexed assetId, address indexed assetAddress, bool isEnabled)
func (_Governance *GovernanceFilterer) FilterAssetUpdated(opts *bind.FilterOpts, assetId []*big.Int, assetAddress []common.Address) (*GovernanceAssetUpdatedIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "AssetUpdated", assetIdRule, assetAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceAssetUpdatedIterator{contract: _Governance.contract, event: "AssetUpdated", logs: logs, sub: sub}, nil
}

// WatchAssetUpdated is a free log subscription operation binding the contract event 0xce6315d17c79350e737a11e641e32024ae26a2500d55b7e62cc61268811c7fbb.
//
// Solidity: event AssetUpdated(uint256 indexed assetId, address indexed assetAddress, bool isEnabled)
func (_Governance *GovernanceFilterer) WatchAssetUpdated(opts *bind.WatchOpts, sink chan<- *GovernanceAssetUpdated, assetId []*big.Int, assetAddress []common.Address) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "AssetUpdated", assetIdRule, assetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceAssetUpdated)
				if err := _Governance.contract.UnpackLog(event, "AssetUpdated", log); err != nil {
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

// ParseAssetUpdated is a log parse operation binding the contract event 0xce6315d17c79350e737a11e641e32024ae26a2500d55b7e62cc61268811c7fbb.
//
// Solidity: event AssetUpdated(uint256 indexed assetId, address indexed assetAddress, bool isEnabled)
func (_Governance *GovernanceFilterer) ParseAssetUpdated(log types.Log) (*GovernanceAssetUpdated, error) {
	event := new(GovernanceAssetUpdated)
	if err := _Governance.contract.UnpackLog(event, "AssetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Governance contract.
type GovernanceOwnershipTransferredIterator struct {
	Event *GovernanceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GovernanceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceOwnershipTransferred)
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
		it.Event = new(GovernanceOwnershipTransferred)
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
func (it *GovernanceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceOwnershipTransferred represents a OwnershipTransferred event raised by the Governance contract.
type GovernanceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Governance *GovernanceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GovernanceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceOwnershipTransferredIterator{contract: _Governance.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Governance *GovernanceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GovernanceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceOwnershipTransferred)
				if err := _Governance.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Governance *GovernanceFilterer) ParseOwnershipTransferred(log types.Log) (*GovernanceOwnershipTransferred, error) {
	event := new(GovernanceOwnershipTransferred)
	if err := _Governance.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceRollupProviderUpdatedIterator is returned from FilterRollupProviderUpdated and is used to iterate over the raw logs and unpacked data for RollupProviderUpdated events raised by the Governance contract.
type GovernanceRollupProviderUpdatedIterator struct {
	Event *GovernanceRollupProviderUpdated // Event containing the contract specifics and raw log

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
func (it *GovernanceRollupProviderUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceRollupProviderUpdated)
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
		it.Event = new(GovernanceRollupProviderUpdated)
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
func (it *GovernanceRollupProviderUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceRollupProviderUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceRollupProviderUpdated represents a RollupProviderUpdated event raised by the Governance contract.
type GovernanceRollupProviderUpdated struct {
	ProviderAddress common.Address
	Valid           bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRollupProviderUpdated is a free log retrieval operation binding the contract event 0x46359ce9dbb6c7f9a375b44072210287916d3de725fc8927a8e762047e4a8424.
//
// Solidity: event RollupProviderUpdated(address indexed providerAddress, bool valid)
func (_Governance *GovernanceFilterer) FilterRollupProviderUpdated(opts *bind.FilterOpts, providerAddress []common.Address) (*GovernanceRollupProviderUpdatedIterator, error) {

	var providerAddressRule []interface{}
	for _, providerAddressItem := range providerAddress {
		providerAddressRule = append(providerAddressRule, providerAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "RollupProviderUpdated", providerAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceRollupProviderUpdatedIterator{contract: _Governance.contract, event: "RollupProviderUpdated", logs: logs, sub: sub}, nil
}

// WatchRollupProviderUpdated is a free log subscription operation binding the contract event 0x46359ce9dbb6c7f9a375b44072210287916d3de725fc8927a8e762047e4a8424.
//
// Solidity: event RollupProviderUpdated(address indexed providerAddress, bool valid)
func (_Governance *GovernanceFilterer) WatchRollupProviderUpdated(opts *bind.WatchOpts, sink chan<- *GovernanceRollupProviderUpdated, providerAddress []common.Address) (event.Subscription, error) {

	var providerAddressRule []interface{}
	for _, providerAddressItem := range providerAddress {
		providerAddressRule = append(providerAddressRule, providerAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "RollupProviderUpdated", providerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceRollupProviderUpdated)
				if err := _Governance.contract.UnpackLog(event, "RollupProviderUpdated", log); err != nil {
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

// ParseRollupProviderUpdated is a log parse operation binding the contract event 0x46359ce9dbb6c7f9a375b44072210287916d3de725fc8927a8e762047e4a8424.
//
// Solidity: event RollupProviderUpdated(address indexed providerAddress, bool valid)
func (_Governance *GovernanceFilterer) ParseRollupProviderUpdated(log types.Log) (*GovernanceRollupProviderUpdated, error) {
	event := new(GovernanceRollupProviderUpdated)
	if err := _Governance.contract.UnpackLog(event, "RollupProviderUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
