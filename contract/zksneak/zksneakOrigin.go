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
const ZksneakABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"addr\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oldRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"NewBlock\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"accountRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pk\",\"type\":\"bytes32\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"verifyBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ZksneakBin is the compiled bytecode used for deploying new contracts.
var ZksneakBin = "0x60806040523480156100115760006000fd5b50610017565b61053b806100266000396000f3fe6080604052600436106100435760003560e01c8063a7ccec2f14610049578063b214faa514610075578063f3fef3a314610091578063f4fbb459146100bb57610043565b60006000fd5b3480156100565760006000fd5b5061005f6100e5565b60405161006c91906103a5565b60405180910390f35b61008f600480360381019061008a91906102d8565b6100ee565b005b34801561009e5760006000fd5b506100b960048036038101906100b49190610299565b610170565b005b3480156100c85760006000fd5b506100e360048036038101906100de91906102d8565b610201565b005b60006000505481565b600034111515610133576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161012a90610436565b60405180910390fd5b7f98e783c3864bbf744a057ef605a2a61701c3b62b5ed68b3745b99094497daf1f81346040516101649291906103eb565b60405180910390a15b50565b80471115156101b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ab90610415565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156101fb573d600060003e3d6000fd5b505b5050565b806000600050819090600019169055507f01d8a02b7dd57c2d711c59ff367c5077e4cec95780b088a91e7b4fb5159a3d91600060005054826040516102479291906103c1565b60405180910390a15b5056610504565b600081359050610266816104b3565b5b92915050565b60008135905061027c816104ce565b5b92915050565b600081359050610292816104e9565b5b92915050565b60006000604083850312156102ae5760006000fd5b60006102bc85828601610257565b92505060206102cd85828601610283565b9150505b9250929050565b6000602082840312156102eb5760006000fd5b60006102f98482850161026d565b9150505b92915050565b61030c8161047c565b82525b5050565b6000610320600e83610457565b91507f746f6f206d756368206d6f6e657900000000000000000000000000000000000060008301526020820190505b919050565b6000610361601d83610457565b91507f796f752073686f756c64206465706f73697420736f6d65206d6f6e657900000060008301526020820190505b919050565b61039e816104a8565b82525b5050565b60006020820190506103ba6000830184610303565b5b92915050565b60006040820190506103d66000830185610303565b6103e36020830184610303565b5b9392505050565b60006040820190506104006000830185610303565b61040d6020830184610395565b5b9392505050565b6000602082019050818103600083015261042e81610313565b90505b919050565b6000602082019050818103600083015261044f81610354565b90505b919050565b60008282526020820190505b92915050565b600061047482610487565b90505b919050565b60008190505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b6104bc81610469565b811415156104ca5760006000fd5b5b50565b6104d78161047c565b811415156104e55760006000fd5b5b50565b6104f2816104a8565b811415156105005760006000fd5b5b50565bfea2646970667358221220de005f80a83c0b2e902ac2f13aa53ccb3fae7f42e76544aaab0073f583bfa87164736f6c63430007060033"

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
func (_Zksneak *ZksneakRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_Zksneak *ZksneakCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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

// AccountRoot is a free data retrieval call binding the contract method 0xa7ccec2f.
//
// Solidity: function accountRoot() view returns(bytes32)
func (_Zksneak *ZksneakCaller) AccountRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Zksneak.contract.Call(opts, &out, "accountRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AccountRoot is a free data retrieval call binding the contract method 0xa7ccec2f.
//
// Solidity: function accountRoot() view returns(bytes32)
func (_Zksneak *ZksneakSession) AccountRoot() ([32]byte, error) {
	return _Zksneak.Contract.AccountRoot(&_Zksneak.CallOpts)
}

// AccountRoot is a free data retrieval call binding the contract method 0xa7ccec2f.
//
// Solidity: function accountRoot() view returns(bytes32)
func (_Zksneak *ZksneakCallerSession) AccountRoot() ([32]byte, error) {
	return _Zksneak.Contract.AccountRoot(&_Zksneak.CallOpts)
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

// VerifyBlock is a paid mutator transaction binding the contract method 0xf4fbb459.
//
// Solidity: function verifyBlock(bytes32 newRoot) returns()
func (_Zksneak *ZksneakTransactor) VerifyBlock(opts *bind.TransactOpts, newRoot [32]byte) (*types.Transaction, error) {
	return _Zksneak.contract.Transact(opts, "verifyBlock", newRoot)
}

// VerifyBlock is a paid mutator transaction binding the contract method 0xf4fbb459.
//
// Solidity: function verifyBlock(bytes32 newRoot) returns()
func (_Zksneak *ZksneakSession) VerifyBlock(newRoot [32]byte) (*types.Transaction, error) {
	return _Zksneak.Contract.VerifyBlock(&_Zksneak.TransactOpts, newRoot)
}

// VerifyBlock is a paid mutator transaction binding the contract method 0xf4fbb459.
//
// Solidity: function verifyBlock(bytes32 newRoot) returns()
func (_Zksneak *ZksneakTransactorSession) VerifyBlock(newRoot [32]byte) (*types.Transaction, error) {
	return _Zksneak.Contract.VerifyBlock(&_Zksneak.TransactOpts, newRoot)
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
	event.Raw = log
	return event, nil
}

// ZksneakNewBlockIterator is returned from FilterNewBlock and is used to iterate over the raw logs and unpacked data for NewBlock events raised by the Zksneak contract.
type ZksneakNewBlockIterator struct {
	Event *ZksneakNewBlock // Event containing the contract specifics and raw log

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
func (it *ZksneakNewBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZksneakNewBlock)
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
		it.Event = new(ZksneakNewBlock)
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
func (it *ZksneakNewBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZksneakNewBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZksneakNewBlock represents a NewBlock event raised by the Zksneak contract.
type ZksneakNewBlock struct {
	OldRoot [32]byte
	NewRoot [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewBlock is a free log retrieval operation binding the contract event 0x01d8a02b7dd57c2d711c59ff367c5077e4cec95780b088a91e7b4fb5159a3d91.
//
// Solidity: event NewBlock(bytes32 oldRoot, bytes32 newRoot)
func (_Zksneak *ZksneakFilterer) FilterNewBlock(opts *bind.FilterOpts) (*ZksneakNewBlockIterator, error) {

	logs, sub, err := _Zksneak.contract.FilterLogs(opts, "NewBlock")
	if err != nil {
		return nil, err
	}
	return &ZksneakNewBlockIterator{contract: _Zksneak.contract, event: "NewBlock", logs: logs, sub: sub}, nil
}

// WatchNewBlock is a free log subscription operation binding the contract event 0x01d8a02b7dd57c2d711c59ff367c5077e4cec95780b088a91e7b4fb5159a3d91.
//
// Solidity: event NewBlock(bytes32 oldRoot, bytes32 newRoot)
func (_Zksneak *ZksneakFilterer) WatchNewBlock(opts *bind.WatchOpts, sink chan<- *ZksneakNewBlock) (event.Subscription, error) {

	logs, sub, err := _Zksneak.contract.WatchLogs(opts, "NewBlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZksneakNewBlock)
				if err := _Zksneak.contract.UnpackLog(event, "NewBlock", log); err != nil {
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

// ParseNewBlock is a log parse operation binding the contract event 0x01d8a02b7dd57c2d711c59ff367c5077e4cec95780b088a91e7b4fb5159a3d91.
//
// Solidity: event NewBlock(bytes32 oldRoot, bytes32 newRoot)
func (_Zksneak *ZksneakFilterer) ParseNewBlock(log types.Log) (*ZksneakNewBlock, error) {
	event := new(ZksneakNewBlock)
	if err := _Zksneak.contract.UnpackLog(event, "NewBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
