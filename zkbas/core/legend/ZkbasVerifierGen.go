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

// VerifierMetaData contains all meta data concerning the Verifier contract.
var VerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ScalarField\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"upgradeParameters\",\"type\":\"bytes\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"in_proof\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"proof_inputs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"num_proofs\",\"type\":\"uint256\"}],\"name\":\"verifyBatchProofs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"in_proof\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"proof_inputs\",\"type\":\"uint256[]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifierMetaData.ABI instead.
var VerifierABI = VerifierMetaData.ABI

// Verifier is an auto generated Go binding around an Ethereum contract.
type Verifier struct {
	VerifierCaller     // Read-only binding to the contract
	VerifierTransactor // Write-only binding to the contract
	VerifierFilterer   // Log filterer for contract events
}

// VerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifierSession struct {
	Contract     *Verifier         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifierCallerSession struct {
	Contract *VerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// VerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifierTransactorSession struct {
	Contract     *VerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifierRaw struct {
	Contract *Verifier // Generic contract binding to access the raw methods on
}

// VerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifierCallerRaw struct {
	Contract *VerifierCaller // Generic read-only contract binding to access the raw methods on
}

// VerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifierTransactorRaw struct {
	Contract *VerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifier creates a new instance of Verifier, bound to a specific deployed contract.
func NewVerifier(address common.Address, backend bind.ContractBackend) (*Verifier, error) {
	contract, err := bindVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verifier{VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

// NewVerifierCaller creates a new read-only instance of Verifier, bound to a specific deployed contract.
func NewVerifierCaller(address common.Address, caller bind.ContractCaller) (*VerifierCaller, error) {
	contract, err := bindVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierCaller{contract: contract}, nil
}

// NewVerifierTransactor creates a new write-only instance of Verifier, bound to a specific deployed contract.
func NewVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierTransactor, error) {
	contract, err := bindVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierTransactor{contract: contract}, nil
}

// NewVerifierFilterer creates a new log filterer instance of Verifier, bound to a specific deployed contract.
func NewVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierFilterer, error) {
	contract, err := bindVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierFilterer{contract: contract}, nil
}

// bindVerifier binds a generic wrapper to an already deployed contract.
func bindVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifier *VerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.VerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifier *VerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifier *VerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifier *VerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifier *VerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifier *VerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transact(opts, method, params...)
}

// ScalarField is a free data retrieval call binding the contract method 0xdaad1e63.
//
// Solidity: function ScalarField() pure returns(uint256)
func (_Verifier *VerifierCaller) ScalarField(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "ScalarField")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScalarField is a free data retrieval call binding the contract method 0xdaad1e63.
//
// Solidity: function ScalarField() pure returns(uint256)
func (_Verifier *VerifierSession) ScalarField() (*big.Int, error) {
	return _Verifier.Contract.ScalarField(&_Verifier.CallOpts)
}

// ScalarField is a free data retrieval call binding the contract method 0xdaad1e63.
//
// Solidity: function ScalarField() pure returns(uint256)
func (_Verifier *VerifierCallerSession) ScalarField() (*big.Int, error) {
	return _Verifier.Contract.ScalarField(&_Verifier.CallOpts)
}

// VerifyBatchProofs is a free data retrieval call binding the contract method 0xc50e8263.
//
// Solidity: function verifyBatchProofs(uint256[] in_proof, uint256[] proof_inputs, uint256 num_proofs) view returns(bool success)
func (_Verifier *VerifierCaller) VerifyBatchProofs(opts *bind.CallOpts, in_proof []*big.Int, proof_inputs []*big.Int, num_proofs *big.Int) (bool, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "verifyBatchProofs", in_proof, proof_inputs, num_proofs)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyBatchProofs is a free data retrieval call binding the contract method 0xc50e8263.
//
// Solidity: function verifyBatchProofs(uint256[] in_proof, uint256[] proof_inputs, uint256 num_proofs) view returns(bool success)
func (_Verifier *VerifierSession) VerifyBatchProofs(in_proof []*big.Int, proof_inputs []*big.Int, num_proofs *big.Int) (bool, error) {
	return _Verifier.Contract.VerifyBatchProofs(&_Verifier.CallOpts, in_proof, proof_inputs, num_proofs)
}

// VerifyBatchProofs is a free data retrieval call binding the contract method 0xc50e8263.
//
// Solidity: function verifyBatchProofs(uint256[] in_proof, uint256[] proof_inputs, uint256 num_proofs) view returns(bool success)
func (_Verifier *VerifierCallerSession) VerifyBatchProofs(in_proof []*big.Int, proof_inputs []*big.Int, num_proofs *big.Int) (bool, error) {
	return _Verifier.Contract.VerifyBatchProofs(&_Verifier.CallOpts, in_proof, proof_inputs, num_proofs)
}

// VerifyProof is a free data retrieval call binding the contract method 0x721ea4ac.
//
// Solidity: function verifyProof(uint256[] in_proof, uint256[] proof_inputs) view returns(bool)
func (_Verifier *VerifierCaller) VerifyProof(opts *bind.CallOpts, in_proof []*big.Int, proof_inputs []*big.Int) (bool, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "verifyProof", in_proof, proof_inputs)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x721ea4ac.
//
// Solidity: function verifyProof(uint256[] in_proof, uint256[] proof_inputs) view returns(bool)
func (_Verifier *VerifierSession) VerifyProof(in_proof []*big.Int, proof_inputs []*big.Int) (bool, error) {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, in_proof, proof_inputs)
}

// VerifyProof is a free data retrieval call binding the contract method 0x721ea4ac.
//
// Solidity: function verifyProof(uint256[] in_proof, uint256[] proof_inputs) view returns(bool)
func (_Verifier *VerifierCallerSession) VerifyProof(in_proof []*big.Int, proof_inputs []*big.Int) (bool, error) {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, in_proof, proof_inputs)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes ) returns()
func (_Verifier *VerifierTransactor) Initialize(opts *bind.TransactOpts, arg0 []byte) (*types.Transaction, error) {
	return _Verifier.contract.Transact(opts, "initialize", arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes ) returns()
func (_Verifier *VerifierSession) Initialize(arg0 []byte) (*types.Transaction, error) {
	return _Verifier.Contract.Initialize(&_Verifier.TransactOpts, arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes ) returns()
func (_Verifier *VerifierTransactorSession) Initialize(arg0 []byte) (*types.Transaction, error) {
	return _Verifier.Contract.Initialize(&_Verifier.TransactOpts, arg0)
}

// Upgrade is a paid mutator transaction binding the contract method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_Verifier *VerifierTransactor) Upgrade(opts *bind.TransactOpts, upgradeParameters []byte) (*types.Transaction, error) {
	return _Verifier.contract.Transact(opts, "upgrade", upgradeParameters)
}

// Upgrade is a paid mutator transaction binding the contract method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_Verifier *VerifierSession) Upgrade(upgradeParameters []byte) (*types.Transaction, error) {
	return _Verifier.Contract.Upgrade(&_Verifier.TransactOpts, upgradeParameters)
}

// Upgrade is a paid mutator transaction binding the contract method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_Verifier *VerifierTransactorSession) Upgrade(upgradeParameters []byte) (*types.Transaction, error) {
	return _Verifier.Contract.Upgrade(&_Verifier.TransactOpts, upgradeParameters)
}
