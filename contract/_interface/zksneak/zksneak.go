// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zksneak

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

// ZksneakABI is the input ABI used to generate the binding from.
const ZksneakABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"addr\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oldRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"RootUpdateForSwap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oldRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"RootUpdateForTransfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pk\",\"type\":\"bytes32\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"verifyAtomicSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"verifyTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ZksneakBin is the compiled bytecode used for deploying new contracts.
var ZksneakBin = "0x60806040523480156100115760006000fd5b50610017565b610566806100266000396000f3fe6080604052600436106100435760003560e01c806365f8872b14610049578063b214faa514610073578063e7192eb41461008f578063f3fef3a3146100b957610043565b60006000fd5b3480156100565760006000fd5b50610071600480360381019061006c919061031f565b6100e3565b005b61008d6004803603810190610088919061031f565b610135565b005b34801561009c5760006000fd5b506100b760048036038101906100b2919061031f565b6101b7565b005b3480156100c65760006000fd5b506100e160048036038101906100dc91906102e0565b610209565b005b806000600050819090600019169055507f2307bf71351a24961c73b390c1ab365eaa7312d1edc3841cd990a688fa4d882a600060005054826040516101299291906103ec565b60405180910390a15b50565b60003411151561017a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161017190610461565b60405180910390fd5b7f98e783c3864bbf744a057ef605a2a61701c3b62b5ed68b3745b99094497daf1f81346040516101ab929190610416565b60405180910390a15b50565b806000600050819090600019169055507f4d6456b216ce6c1e8c608e3195c9d15123d88fc2098a1634f085d64b2b1941eb600060005054826040516101fd9291906103ec565b60405180910390a15b50565b804711151561024d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024490610440565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610294573d600060003e3d6000fd5b505b50505661052f565b6000813590506102ad816104de565b5b92915050565b6000813590506102c3816104f9565b5b92915050565b6000813590506102d981610514565b5b92915050565b60006000604083850312156102f55760006000fd5b60006103038582860161029e565b9250506020610314858286016102ca565b9150505b9250929050565b6000602082840312156103325760006000fd5b6000610340848285016102b4565b9150505b92915050565b610353816104a7565b82525b5050565b6000610367600e83610482565b91507f746f6f206d756368206d6f6e657900000000000000000000000000000000000060008301526020820190505b919050565b60006103a8601d83610482565b91507f796f752073686f756c64206465706f73697420736f6d65206d6f6e657900000060008301526020820190505b919050565b6103e5816104d3565b82525b5050565b6000604082019050610401600083018561034a565b61040e602083018461034a565b5b9392505050565b600060408201905061042b600083018561034a565b61043860208301846103dc565b5b9392505050565b600060208201905081810360008301526104598161035a565b90505b919050565b6000602082019050818103600083015261047a8161039b565b90505b919050565b60008282526020820190505b92915050565b600061049f826104b2565b90505b919050565b60008190505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b6104e781610494565b811415156104f55760006000fd5b5b50565b610502816104a7565b811415156105105760006000fd5b5b50565b61051d816104d3565b8114151561052b5760006000fd5b5b50565bfea2646970667358221220f063c5f5497ae8c20a5d494ff26de4bcc51f013b6daff363cc9919164cc180ed64736f6c63430007060033"

// DeployZksneak deploys a new Ethereum contract, binding an instance of Zksneak to it.
func DeployZksneak(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Zksneak, error) {
	parsed, err := abi.JSON(strings.NewReader(ZksneakABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZksneakBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Zksneak{ZksneakCaller: ZksneakCaller{contract: contract}, ZksneakTransactor: ZksneakTransactor{contract: contract}, ZksneakFilterer: ZksneakFilterer{contract: contract}}, nil
}

// Zksneak is an auto generated Go binding around an Ethereum contract.
type Zksneak struct {
	ZksneakCaller     // Read-only binding to the contract
	ZksneakTransactor // Write-only binding to the contract
	ZksneakFilterer   // Log filterer for contract events
}

// ZksneakCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZksneakCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZksneakTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZksneakTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZksneakFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZksneakFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZksneakSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZksneakSession struct {
	Contract     *Zksneak          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZksneakCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZksneakCallerSession struct {
	Contract *ZksneakCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ZksneakTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZksneakTransactorSession struct {
	Contract     *ZksneakTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ZksneakRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZksneakRaw struct {
	Contract *Zksneak // Generic contract binding to access the raw methods on
}

// ZksneakCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZksneakCallerRaw struct {
	Contract *ZksneakCaller // Generic read-only contract binding to access the raw methods on
}

// ZksneakTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZksneakTransactorRaw struct {
	Contract *ZksneakTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZksneak creates a new instance of Zksneak, bound to a specific deployed contract.
func NewZksneak(address common.Address, backend bind.ContractBackend) (*Zksneak, error) {
	contract, err := bindZksneak(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Zksneak{ZksneakCaller: ZksneakCaller{contract: contract}, ZksneakTransactor: ZksneakTransactor{contract: contract}, ZksneakFilterer: ZksneakFilterer{contract: contract}}, nil
}

// NewZksneakCaller creates a new read-only instance of Zksneak, bound to a specific deployed contract.
func NewZksneakCaller(address common.Address, caller bind.ContractCaller) (*ZksneakCaller, error) {
	contract, err := bindZksneak(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZksneakCaller{contract: contract}, nil
}

// NewZksneakTransactor creates a new write-only instance of Zksneak, bound to a specific deployed contract.
func NewZksneakTransactor(address common.Address, transactor bind.ContractTransactor) (*ZksneakTransactor, error) {
	contract, err := bindZksneak(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZksneakTransactor{contract: contract}, nil
}

// NewZksneakFilterer creates a new log filterer instance of Zksneak, bound to a specific deployed contract.
func NewZksneakFilterer(address common.Address, filterer bind.ContractFilterer) (*ZksneakFilterer, error) {
	contract, err := bindZksneak(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZksneakFilterer{contract: contract}, nil
}

// bindZksneak binds a generic wrapper to an already deployed contract.
func bindZksneak(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZksneakABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zksneak *ZksneakRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Zksneak.Contract.ZksneakCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zksneak *ZksneakRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zksneak.Contract.ZksneakTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zksneak *ZksneakRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zksneak.Contract.ZksneakTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zksneak *ZksneakCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Zksneak.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zksneak *ZksneakTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zksneak.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zksneak *ZksneakTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zksneak.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xb214faa5.
//
// Solidity: function deposit(bytes32 pk) payable returns()
func (_Zksneak *ZksneakTransactor) Deposit(opts *bind.TransactOpts, pk [32]byte) (*types.Transaction, error) {
	return _Zksneak.contract.Transact(opts, "deposit", pk)
}

// Deposit is a paid mutator transaction binding the contract method 0xb214faa5.
//
// Solidity: function deposit(bytes32 pk) payable returns()
func (_Zksneak *ZksneakSession) Deposit(pk [32]byte) (*types.Transaction, error) {
	return _Zksneak.Contract.Deposit(&_Zksneak.TransactOpts, pk)
}

// Deposit is a paid mutator transaction binding the contract method 0xb214faa5.
//
// Solidity: function deposit(bytes32 pk) payable returns()
func (_Zksneak *ZksneakTransactorSession) Deposit(pk [32]byte) (*types.Transaction, error) {
	return _Zksneak.Contract.Deposit(&_Zksneak.TransactOpts, pk)
}

// VerifyAtomicSwap is a paid mutator transaction binding the contract method 0x65f8872b.
//
// Solidity: function verifyAtomicSwap(bytes32 newRoot) returns()
func (_Zksneak *ZksneakTransactor) VerifyAtomicSwap(opts *bind.TransactOpts, newRoot [32]byte) (*types.Transaction, error) {
	return _Zksneak.contract.Transact(opts, "verifyAtomicSwap", newRoot)
}

// VerifyAtomicSwap is a paid mutator transaction binding the contract method 0x65f8872b.
//
// Solidity: function verifyAtomicSwap(bytes32 newRoot) returns()
func (_Zksneak *ZksneakSession) VerifyAtomicSwap(newRoot [32]byte) (*types.Transaction, error) {
	return _Zksneak.Contract.VerifyAtomicSwap(&_Zksneak.TransactOpts, newRoot)
}

// VerifyAtomicSwap is a paid mutator transaction binding the contract method 0x65f8872b.
//
// Solidity: function verifyAtomicSwap(bytes32 newRoot) returns()
func (_Zksneak *ZksneakTransactorSession) VerifyAtomicSwap(newRoot [32]byte) (*types.Transaction, error) {
	return _Zksneak.Contract.VerifyAtomicSwap(&_Zksneak.TransactOpts, newRoot)
}

// VerifyTransfer is a paid mutator transaction binding the contract method 0xe7192eb4.
//
// Solidity: function verifyTransfer(bytes32 newRoot) returns()
func (_Zksneak *ZksneakTransactor) VerifyTransfer(opts *bind.TransactOpts, newRoot [32]byte) (*types.Transaction, error) {
	return _Zksneak.contract.Transact(opts, "verifyTransfer", newRoot)
}

// VerifyTransfer is a paid mutator transaction binding the contract method 0xe7192eb4.
//
// Solidity: function verifyTransfer(bytes32 newRoot) returns()
func (_Zksneak *ZksneakSession) VerifyTransfer(newRoot [32]byte) (*types.Transaction, error) {
	return _Zksneak.Contract.VerifyTransfer(&_Zksneak.TransactOpts, newRoot)
}

// VerifyTransfer is a paid mutator transaction binding the contract method 0xe7192eb4.
//
// Solidity: function verifyTransfer(bytes32 newRoot) returns()
func (_Zksneak *ZksneakTransactorSession) VerifyTransfer(newRoot [32]byte) (*types.Transaction, error) {
	return _Zksneak.Contract.VerifyTransfer(&_Zksneak.TransactOpts, newRoot)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _addr, uint256 amount) returns()
func (_Zksneak *ZksneakTransactor) Withdraw(opts *bind.TransactOpts, _addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Zksneak.contract.Transact(opts, "withdraw", _addr, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _addr, uint256 amount) returns()
func (_Zksneak *ZksneakSession) Withdraw(_addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Zksneak.Contract.Withdraw(&_Zksneak.TransactOpts, _addr, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _addr, uint256 amount) returns()
func (_Zksneak *ZksneakTransactorSession) Withdraw(_addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Zksneak.Contract.Withdraw(&_Zksneak.TransactOpts, _addr, amount)
}

// ZksneakDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Zksneak contract.
type ZksneakDepositIterator struct {
	Event *ZksneakDeposit // Event containing the contract specifics and raw log

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
func (it *ZksneakDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZksneakDeposit)
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
		it.Event = new(ZksneakDeposit)
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
func (it *ZksneakDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZksneakDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZksneakDeposit represents a Deposit event raised by the Zksneak contract.
type ZksneakDeposit struct {
	Addr   [32]byte
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x98e783c3864bbf744a057ef605a2a61701c3b62b5ed68b3745b99094497daf1f.
//
// Solidity: event Deposit(bytes32 addr, uint256 amount)
func (_Zksneak *ZksneakFilterer) FilterDeposit(opts *bind.FilterOpts) (*ZksneakDepositIterator, error) {

	logs, sub, err := _Zksneak.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &ZksneakDepositIterator{contract: _Zksneak.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x98e783c3864bbf744a057ef605a2a61701c3b62b5ed68b3745b99094497daf1f.
//
// Solidity: event Deposit(bytes32 addr, uint256 amount)
func (_Zksneak *ZksneakFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ZksneakDeposit) (event.Subscription, error) {

	logs, sub, err := _Zksneak.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZksneakDeposit)
				if err := _Zksneak.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x98e783c3864bbf744a057ef605a2a61701c3b62b5ed68b3745b99094497daf1f.
//
// Solidity: event Deposit(bytes32 addr, uint256 amount)
func (_Zksneak *ZksneakFilterer) ParseDeposit(log types.Log) (*ZksneakDeposit, error) {
	event := new(ZksneakDeposit)
	if err := _Zksneak.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ZksneakRootUpdateForSwapIterator is returned from FilterRootUpdateForSwap and is used to iterate over the raw logs and unpacked data for RootUpdateForSwap events raised by the Zksneak contract.
type ZksneakRootUpdateForSwapIterator struct {
	Event *ZksneakRootUpdateForSwap // Event containing the contract specifics and raw log

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
func (it *ZksneakRootUpdateForSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZksneakRootUpdateForSwap)
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
		it.Event = new(ZksneakRootUpdateForSwap)
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
func (it *ZksneakRootUpdateForSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZksneakRootUpdateForSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZksneakRootUpdateForSwap represents a RootUpdateForSwap event raised by the Zksneak contract.
type ZksneakRootUpdateForSwap struct {
	OldRoot [32]byte
	NewRoot [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRootUpdateForSwap is a free log retrieval operation binding the contract event 0x2307bf71351a24961c73b390c1ab365eaa7312d1edc3841cd990a688fa4d882a.
//
// Solidity: event RootUpdateForSwap(bytes32 oldRoot, bytes32 newRoot)
func (_Zksneak *ZksneakFilterer) FilterRootUpdateForSwap(opts *bind.FilterOpts) (*ZksneakRootUpdateForSwapIterator, error) {

	logs, sub, err := _Zksneak.contract.FilterLogs(opts, "RootUpdateForSwap")
	if err != nil {
		return nil, err
	}
	return &ZksneakRootUpdateForSwapIterator{contract: _Zksneak.contract, event: "RootUpdateForSwap", logs: logs, sub: sub}, nil
}

// WatchRootUpdateForSwap is a free log subscription operation binding the contract event 0x2307bf71351a24961c73b390c1ab365eaa7312d1edc3841cd990a688fa4d882a.
//
// Solidity: event RootUpdateForSwap(bytes32 oldRoot, bytes32 newRoot)
func (_Zksneak *ZksneakFilterer) WatchRootUpdateForSwap(opts *bind.WatchOpts, sink chan<- *ZksneakRootUpdateForSwap) (event.Subscription, error) {

	logs, sub, err := _Zksneak.contract.WatchLogs(opts, "RootUpdateForSwap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZksneakRootUpdateForSwap)
				if err := _Zksneak.contract.UnpackLog(event, "RootUpdateForSwap", log); err != nil {
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

// ParseRootUpdateForSwap is a log parse operation binding the contract event 0x2307bf71351a24961c73b390c1ab365eaa7312d1edc3841cd990a688fa4d882a.
//
// Solidity: event RootUpdateForSwap(bytes32 oldRoot, bytes32 newRoot)
func (_Zksneak *ZksneakFilterer) ParseRootUpdateForSwap(log types.Log) (*ZksneakRootUpdateForSwap, error) {
	event := new(ZksneakRootUpdateForSwap)
	if err := _Zksneak.contract.UnpackLog(event, "RootUpdateForSwap", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ZksneakRootUpdateForTransferIterator is returned from FilterRootUpdateForTransfer and is used to iterate over the raw logs and unpacked data for RootUpdateForTransfer events raised by the Zksneak contract.
type ZksneakRootUpdateForTransferIterator struct {
	Event *ZksneakRootUpdateForTransfer // Event containing the contract specifics and raw log

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
func (it *ZksneakRootUpdateForTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZksneakRootUpdateForTransfer)
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
		it.Event = new(ZksneakRootUpdateForTransfer)
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
func (it *ZksneakRootUpdateForTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZksneakRootUpdateForTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZksneakRootUpdateForTransfer represents a RootUpdateForTransfer event raised by the Zksneak contract.
type ZksneakRootUpdateForTransfer struct {
	OldRoot [32]byte
	NewRoot [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRootUpdateForTransfer is a free log retrieval operation binding the contract event 0x4d6456b216ce6c1e8c608e3195c9d15123d88fc2098a1634f085d64b2b1941eb.
//
// Solidity: event RootUpdateForTransfer(bytes32 oldRoot, bytes32 newRoot)
func (_Zksneak *ZksneakFilterer) FilterRootUpdateForTransfer(opts *bind.FilterOpts) (*ZksneakRootUpdateForTransferIterator, error) {

	logs, sub, err := _Zksneak.contract.FilterLogs(opts, "RootUpdateForTransfer")
	if err != nil {
		return nil, err
	}
	return &ZksneakRootUpdateForTransferIterator{contract: _Zksneak.contract, event: "RootUpdateForTransfer", logs: logs, sub: sub}, nil
}

// WatchRootUpdateForTransfer is a free log subscription operation binding the contract event 0x4d6456b216ce6c1e8c608e3195c9d15123d88fc2098a1634f085d64b2b1941eb.
//
// Solidity: event RootUpdateForTransfer(bytes32 oldRoot, bytes32 newRoot)
func (_Zksneak *ZksneakFilterer) WatchRootUpdateForTransfer(opts *bind.WatchOpts, sink chan<- *ZksneakRootUpdateForTransfer) (event.Subscription, error) {

	logs, sub, err := _Zksneak.contract.WatchLogs(opts, "RootUpdateForTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZksneakRootUpdateForTransfer)
				if err := _Zksneak.contract.UnpackLog(event, "RootUpdateForTransfer", log); err != nil {
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

// ParseRootUpdateForTransfer is a log parse operation binding the contract event 0x4d6456b216ce6c1e8c608e3195c9d15123d88fc2098a1634f085d64b2b1941eb.
//
// Solidity: event RootUpdateForTransfer(bytes32 oldRoot, bytes32 newRoot)
func (_Zksneak *ZksneakFilterer) ParseRootUpdateForTransfer(log types.Log) (*ZksneakRootUpdateForTransfer, error) {
	event := new(ZksneakRootUpdateForTransfer)
	if err := _Zksneak.contract.UnpackLog(event, "RootUpdateForTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
