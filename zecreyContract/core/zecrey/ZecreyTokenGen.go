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

// ZecreyTokenABI is the input ABI used to generate the binding from.
const ZecreyTokenABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ZecreyTokenBin is the compiled bytecode used for deploying new contracts.
var ZecreyTokenBin = "0x6080604052348015620000125760006000fd5b506040516200169a3803806200169a83398181016040526020811015620000395760006000fd5b81019080805190602001909291905050505b6040518060400160405280600681526020017f5a656372657900000000000000000000000000000000000000000000000000008152602001506040518060400160405280600381526020017f52455900000000000000000000000000000000000000000000000000000000008152602001505b8160036000509080519060200190620000d9929190620003c7565b508060046000509080519060200190620000f5929190620003c7565b506012600560006101000a81548160ff021916908360ff1602179055505b50506200012733826200012f60201b60201c565b5b5062000489565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151515620001d8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45524332303a206d696e7420746f20746865207a65726f20616464726573730081526020015060200191505060405180910390fd5b620001ec600083836200032b60201b60201c565b6200020b816002600050546200033160201b62000a6e1790919060201c565b60026000508190909055506200027581600060005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050546200033160201b62000a6e1790919060201c565b600060005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35b5050565b5b505050565b600060008284019050838110151515620003b6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081526020015060200191505060405180910390fd5b80915050620003c156505b92915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282620003ff576000855562000450565b82601f106200041a57805160ff191683800117855562000450565b8280016001018555821562000450579182015b828111156200044f57825182600050909055916020019190600101906200042d565b5b5090506200045f919062000463565b5090565b62000469565b8082111562000485576000818150600090555060010162000469565b5090565b61120180620004996000396000f3fe60806040523480156100115760006000fd5b50600436106100ae5760003560e01c80633950935111610072578063395093511461026157806370a08231146102c657806395d89b411461031f578063a457c2d7146103a3578063a9059cbb14610408578063dd62ed3e1461046d576100ae565b806306fdde03146100b4578063095ea7b31461013857806318160ddd1461019d57806323b872dd146101bb578063313ce56714610240576100ae565b60006000fd5b6100bc6104e6565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100fd5780820151818401525b6020810190506100e1565b50505050905090810190601f16801561012a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101856004803603604081101561014f5760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610590565b60405180821515815260200191505060405180910390f35b6101a56105bf565b6040518082815260200191505060405180910390f35b610228600480360360608110156101d25760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506105d1565b60405180821515815260200191505060405180910390f35b6102486106d0565b604051808260ff16815260200191505060405180910390f35b6102ae600480360360408110156102785760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506106ec565b60405180821515815260200191505060405180910390f35b610309600480360360208110156102dd5760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506107bf565b6040518082815260200191505060405180910390f35b610327610813565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103685780820151818401525b60208101905061034c565b50505050905090810190601f1680156103955780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103f0600480360360408110156103ba5760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506108bd565b60405180821515815260200191505060405180910390f35b6104556004803603604081101561041f5760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506109aa565b60405180821515815260200191505060405180910390f35b6104d0600480360360408110156104845760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506109d9565b6040518082815260200191505060405180910390f35b606060036000508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105815780601f1061055657610100808354040283529160200191610581565b820191906000526020600020905b81548152906001019060200180831161056457829003601f168201915b5050505050905061058d565b90565b60006105b06105a3610b0263ffffffff16565b8484610b0f63ffffffff16565b600190506105b9565b92915050565b600060026000505490506105ce565b90565b60006105e4848484610d1663ffffffff16565b6106c0846105f6610b0263ffffffff16565b6106b58560405180606001604052806028815260200161113660289139600160005060008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000506000610668610b0263ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050546110029092919063ffffffff16565b610b0f63ffffffff16565b600190506106c9565b9392505050565b6000600560009054906101000a900460ff1690506106e9565b90565b60006107b06106ff610b0263ffffffff16565b846107a58560016000506000610719610b0263ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005054610a6e90919063ffffffff16565b610b0f63ffffffff16565b600190506107b9565b92915050565b6000600060005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005054905061080e565b919050565b606060046000508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108ae5780601f10610883576101008083540402835291602001916108ae565b820191906000526020600020905b81548152906001019060200180831161089157829003601f168201915b505050505090506108ba565b90565b600061099b6108d0610b0263ffffffff16565b84610990856040518060600160405280602581526020016111a76025913960016000506000610903610b0263ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050546110029092919063ffffffff16565b610b0f63ffffffff16565b600190506109a4565b92915050565b60006109ca6109bd610b0263ffffffff16565b8484610d1663ffffffff16565b600190506109d3565b92915050565b6000600160005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050549050610a68565b92915050565b600060008284019050838110151515610af2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081526020015060200191505060405180910390fd5b80915050610afc56505b92915050565b6000339050610b0c565b90565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614151515610b97576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806111836024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151515610c1f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806110ee6022913960400191505060405180910390fd5b80600160005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a35b505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614151515610d9e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602581526020018061115e6025913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151515610e26576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001806110cb6023913960400191505060405180910390fd5b610e378383836110c463ffffffff16565b610ea98160405180606001604052806026815260200161111060269139600060005060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050546110029092919063ffffffff16565b600060005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050819090905550610f4c81600060005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005054610a6e90919063ffffffff16565b600060005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35b505050565b600083831115829015156110b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156110775780820151818401525b60208101905061105b565b50505050905090810190601f1680156110a45780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5082840390506110bd565b9392505050565b5b50505056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa2646970667358221220732f5b8e11cacefc24278532429f7f4f55a49c878e3d0f901db2f75666983fca64736f6c63430007060033"

// DeployZecreyToken deploys a new Ethereum contract, binding an instance of ZecreyToken to it.
func DeployZecreyToken(auth *bind.TransactOpts, backend bind.ContractBackend, _initialSupply *big.Int) (common.Address, *types.Transaction, *ZecreyToken, error) {
	parsed, err := abi.JSON(strings.NewReader(ZecreyTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZecreyTokenBin), backend, _initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZecreyToken{ZecreyTokenCaller: ZecreyTokenCaller{contract: contract}, ZecreyTokenTransactor: ZecreyTokenTransactor{contract: contract}, ZecreyTokenFilterer: ZecreyTokenFilterer{contract: contract}}, nil
}

// ZecreyToken is an auto generated Go binding around an Ethereum contract.
type ZecreyToken struct {
	ZecreyTokenCaller     // Read-only binding to the contract
	ZecreyTokenTransactor // Write-only binding to the contract
	ZecreyTokenFilterer   // Log filterer for contract events
}

// ZecreyTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZecreyTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZecreyTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZecreyTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZecreyTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZecreyTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZecreyTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZecreyTokenSession struct {
	Contract     *ZecreyToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZecreyTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZecreyTokenCallerSession struct {
	Contract *ZecreyTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ZecreyTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZecreyTokenTransactorSession struct {
	Contract     *ZecreyTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ZecreyTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZecreyTokenRaw struct {
	Contract *ZecreyToken // Generic contract binding to access the raw methods on
}

// ZecreyTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZecreyTokenCallerRaw struct {
	Contract *ZecreyTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ZecreyTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZecreyTokenTransactorRaw struct {
	Contract *ZecreyTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZecreyToken creates a new instance of ZecreyToken, bound to a specific deployed contract.
func NewZecreyToken(address common.Address, backend bind.ContractBackend) (*ZecreyToken, error) {
	contract, err := bindZecreyToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZecreyToken{ZecreyTokenCaller: ZecreyTokenCaller{contract: contract}, ZecreyTokenTransactor: ZecreyTokenTransactor{contract: contract}, ZecreyTokenFilterer: ZecreyTokenFilterer{contract: contract}}, nil
}

// NewZecreyTokenCaller creates a new read-only instance of ZecreyToken, bound to a specific deployed contract.
func NewZecreyTokenCaller(address common.Address, caller bind.ContractCaller) (*ZecreyTokenCaller, error) {
	contract, err := bindZecreyToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZecreyTokenCaller{contract: contract}, nil
}

// NewZecreyTokenTransactor creates a new write-only instance of ZecreyToken, bound to a specific deployed contract.
func NewZecreyTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ZecreyTokenTransactor, error) {
	contract, err := bindZecreyToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZecreyTokenTransactor{contract: contract}, nil
}

// NewZecreyTokenFilterer creates a new log filterer instance of ZecreyToken, bound to a specific deployed contract.
func NewZecreyTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ZecreyTokenFilterer, error) {
	contract, err := bindZecreyToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZecreyTokenFilterer{contract: contract}, nil
}

// bindZecreyToken binds a generic wrapper to an already deployed contract.
func bindZecreyToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZecreyTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZecreyToken *ZecreyTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZecreyToken.Contract.ZecreyTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZecreyToken *ZecreyTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZecreyToken.Contract.ZecreyTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZecreyToken *ZecreyTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZecreyToken.Contract.ZecreyTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZecreyToken *ZecreyTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZecreyToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZecreyToken *ZecreyTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZecreyToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZecreyToken *ZecreyTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZecreyToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZecreyToken *ZecreyTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZecreyToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZecreyToken *ZecreyTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ZecreyToken.Contract.Allowance(&_ZecreyToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZecreyToken *ZecreyTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ZecreyToken.Contract.Allowance(&_ZecreyToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZecreyToken *ZecreyTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZecreyToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZecreyToken *ZecreyTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ZecreyToken.Contract.BalanceOf(&_ZecreyToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZecreyToken *ZecreyTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ZecreyToken.Contract.BalanceOf(&_ZecreyToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZecreyToken *ZecreyTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZecreyToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZecreyToken *ZecreyTokenSession) Decimals() (uint8, error) {
	return _ZecreyToken.Contract.Decimals(&_ZecreyToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZecreyToken *ZecreyTokenCallerSession) Decimals() (uint8, error) {
	return _ZecreyToken.Contract.Decimals(&_ZecreyToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZecreyToken *ZecreyTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZecreyToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZecreyToken *ZecreyTokenSession) Name() (string, error) {
	return _ZecreyToken.Contract.Name(&_ZecreyToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZecreyToken *ZecreyTokenCallerSession) Name() (string, error) {
	return _ZecreyToken.Contract.Name(&_ZecreyToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZecreyToken *ZecreyTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZecreyToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZecreyToken *ZecreyTokenSession) Symbol() (string, error) {
	return _ZecreyToken.Contract.Symbol(&_ZecreyToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZecreyToken *ZecreyTokenCallerSession) Symbol() (string, error) {
	return _ZecreyToken.Contract.Symbol(&_ZecreyToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZecreyToken *ZecreyTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZecreyToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZecreyToken *ZecreyTokenSession) TotalSupply() (*big.Int, error) {
	return _ZecreyToken.Contract.TotalSupply(&_ZecreyToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZecreyToken *ZecreyTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ZecreyToken.Contract.TotalSupply(&_ZecreyToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ZecreyToken *ZecreyTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZecreyToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ZecreyToken *ZecreyTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZecreyToken.Contract.Approve(&_ZecreyToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ZecreyToken *ZecreyTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZecreyToken.Contract.Approve(&_ZecreyToken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ZecreyToken *ZecreyTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ZecreyToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ZecreyToken *ZecreyTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ZecreyToken.Contract.DecreaseAllowance(&_ZecreyToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ZecreyToken *ZecreyTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ZecreyToken.Contract.DecreaseAllowance(&_ZecreyToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ZecreyToken *ZecreyTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ZecreyToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ZecreyToken *ZecreyTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ZecreyToken.Contract.IncreaseAllowance(&_ZecreyToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ZecreyToken *ZecreyTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ZecreyToken.Contract.IncreaseAllowance(&_ZecreyToken.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ZecreyToken *ZecreyTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZecreyToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ZecreyToken *ZecreyTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZecreyToken.Contract.Transfer(&_ZecreyToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ZecreyToken *ZecreyTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZecreyToken.Contract.Transfer(&_ZecreyToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ZecreyToken *ZecreyTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZecreyToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ZecreyToken *ZecreyTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZecreyToken.Contract.TransferFrom(&_ZecreyToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ZecreyToken *ZecreyTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZecreyToken.Contract.TransferFrom(&_ZecreyToken.TransactOpts, sender, recipient, amount)
}

// ZecreyTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ZecreyToken contract.
type ZecreyTokenApprovalIterator struct {
	Event *ZecreyTokenApproval // Event containing the contract specifics and raw log

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
func (it *ZecreyTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreyTokenApproval)
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
		it.Event = new(ZecreyTokenApproval)
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
func (it *ZecreyTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreyTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreyTokenApproval represents a Approval event raised by the ZecreyToken contract.
type ZecreyTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZecreyToken *ZecreyTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ZecreyTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ZecreyToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ZecreyTokenApprovalIterator{contract: _ZecreyToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZecreyToken *ZecreyTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ZecreyTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ZecreyToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreyTokenApproval)
				if err := _ZecreyToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZecreyToken *ZecreyTokenFilterer) ParseApproval(log types.Log) (*ZecreyTokenApproval, error) {
	event := new(ZecreyTokenApproval)
	if err := _ZecreyToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZecreyTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ZecreyToken contract.
type ZecreyTokenTransferIterator struct {
	Event *ZecreyTokenTransfer // Event containing the contract specifics and raw log

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
func (it *ZecreyTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZecreyTokenTransfer)
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
		it.Event = new(ZecreyTokenTransfer)
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
func (it *ZecreyTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZecreyTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZecreyTokenTransfer represents a Transfer event raised by the ZecreyToken contract.
type ZecreyTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZecreyToken *ZecreyTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ZecreyTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZecreyToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ZecreyTokenTransferIterator{contract: _ZecreyToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZecreyToken *ZecreyTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ZecreyTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZecreyToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZecreyTokenTransfer)
				if err := _ZecreyToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZecreyToken *ZecreyTokenFilterer) ParseTransfer(log types.Log) (*ZecreyTokenTransfer, error) {
	event := new(ZecreyTokenTransfer)
	if err := _ZecreyToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
