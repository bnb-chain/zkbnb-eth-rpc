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

// GovernanceABI is the input ABI used to generate the binding from.
const GovernanceABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"AssetPausedUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"assetId\",\"type\":\"uint16\"}],\"name\":\"NewAsset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAssetGovernance\",\"name\":\"newAssetGovernance\",\"type\":\"address\"}],\"name\":\"NewAssetGovernance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newGovernor\",\"type\":\"address\"}],\"name\":\"NewGovernor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"ValidatorStatusUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_ACCOUNT_INDEX\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_AMOUNT_OF_REGISTERED_ASSETS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_DEPOSIT_AMOUNT\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FUNGIBLE_ASSET_ID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECURITY_COUNCIL_MEMBERS_NUMBER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SPECIAL_ACCOUNT_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SPECIAL_ACCOUNT_ID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_NOTICE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWAL_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"addAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"assetAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetGovernance\",\"outputs\":[{\"internalType\":\"contractAssetGovernance\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"assetsList\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAssetGovernance\",\"name\":\"_newAssetGovernance\",\"type\":\"address\"}],\"name\":\"changeAssetGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGovernor\",\"type\":\"address\"}],\"name\":\"changeGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initializationParameters\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAddressExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkGovernor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"pausedAssets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"requireActiveValidator\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"requireGovernor\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_assetAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_assetPaused\",\"type\":\"bool\"}],\"name\":\"setAssetPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"upgradeParameters\",\"type\":\"bytes\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_assetAddr\",\"type\":\"address\"}],\"name\":\"validateAssetAddress\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// GovernanceBin is the compiled bytecode used for deploying new contracts.
var GovernanceBin = "0x608060405234801561001057600080fd5b50610c18806100206000396000f3fe608060405234801561001057600080fd5b50600436106101a95760003560e01c80634a51a71f116100f9578063d87e374811610097578063f39349ef11610071578063f39349ef14610517578063f5e7d6fd1461051f578063f5f84ed414610527578063fa52c7d81461054d576101a9565b8063d87e3748146104aa578063dbfc2967146104d0578063e4c0aaf4146104f1576101a9565b80637ea399c1116100d35780637ea399c1146104505780639bd7760914610474578063c701f9551461049a578063cc375fb7146104a2576101a9565b80634a51a71f146103e35780634b18bd0f146103fd5780634c34a98214610423576101a9565b806331d8687b11610166578063437545f911610140578063437545f91461033d578063437da02f1461033d578063439fab91146103455780634623c91d146103b5576101a9565b806331d8687b146102c1578063321e182b146102f65780634242d5b31461031c576101a9565b806301e1d114146101ae5780630d360b7f146101cd5780631e763ee3146101d55780632520ce5a146101fb578063253946451461022b578063298410e51461029b575b600080fd5b6101b6610573565b6040805161ffff9092168252519081900360200190f35b6101b6610584565b6101b6600480360360208110156101eb57600080fd5b50356001600160a01b031661058a565b6102296004803603604081101561021157600080fd5b506001600160a01b03813516906020013515156105a0565b005b6102296004803603602081101561024157600080fd5b81019060208101813564010000000081111561025c57600080fd5b82018360208201111561026e57600080fd5b8035906020019184600183028401116401000000008311171561029057600080fd5b5090925090506106af565b610229600480360360208110156102b157600080fd5b50356001600160a01b03166106b3565b6102e2600480360360208110156102d757600080fd5b503561ffff1661082f565b604080519115158252519081900360200190f35b6102e26004803603602081101561030c57600080fd5b50356001600160a01b0316610844565b610324610859565b6040805163ffffffff9092168252519081900360200190f35b61032461085e565b6102296004803603602081101561035b57600080fd5b81019060208101813564010000000081111561037657600080fd5b82018360208201111561038857600080fd5b803590602001918460018302840111640100000000831117156103aa57600080fd5b509092509050610866565b610229600480360360408110156103cb57600080fd5b506001600160a01b038135169060200135151561089f565b6103eb61092e565b60408051918252519081900360200190f35b6102296004803603602081101561041357600080fd5b50356001600160a01b0316610933565b61042b610997565b604080516fffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6104586109a8565b604080516001600160a01b039092168252519081900360200190f35b6101b66004803603602081101561048a57600080fd5b50356001600160a01b03166109b3565b6103eb610a5b565b6103eb610a62565b610229600480360360208110156104c057600080fd5b50356001600160a01b0316610a69565b610458600480360360208110156104e657600080fd5b503561ffff16610adb565b6102296004803603602081101561050757600080fd5b50356001600160a01b0316610af6565b610458610b68565b610458610b77565b6102296004803603602081101561053d57600080fd5b50356001600160a01b0316610b86565b6102e26004803603602081101561056357600080fd5b50356001600160a01b0316610bcd565b600054600160a01b900461ffff1681565b61ffff81565b60036020526000908152604090205461ffff1681565b6105a933610b86565b6000306001600160a01b0316639bd77609846040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156105f857600080fd5b505afa15801561060c573d6000803e3d6000fd5b505050506040513d602081101561062257600080fd5b505161ffff811660009081526002602052604090205490915060ff161515821515146106aa5761ffff8116600090815260026020908152604091829020805460ff1916851515908117909155825190815291516001600160a01b038616927ff7ca5545623b85829d3abcb6729eb021070bbe93cba64f5c4de6b17b805b7d0292908290030190a25b505050565b5050565b6006546001600160a01b031633146106f7576040805162461bcd60e51b8152602060048201526002602482015261314560f01b604482015290519081900360640190fd5b6001600160a01b03811660009081526003602052604090205461ffff161561074b576040805162461bcd60e51b8152602060048201526002602482015261316560f01b604482015290519081900360640190fd5b60005461ffff600160a01b909104811610610792576040805162461bcd60e51b815260206004820152600260248201526118b360f11b604482015290519081900360640190fd5b60008054600161ffff600160a01b8084048216929092018116820261ffff60a01b1990931692909217808455041680825260046020908152604080842080546001600160a01b0387166001600160a01b031990911681179091558085526003909252808420805461ffff19168417905551919283927f990c18cf226b253448a6a051b5cadf51a61f4ca56703374cdd8bf9df04b2f9019190a35050565b60026020526000908152604090205460ff1681565b60056020526000908152604090205460ff1681565b600081565b63ffffffff81565b60008282602081101561087857600080fd5b506000805491356001600160a01b03166001600160a01b0319909216919091179055505050565b6108a833610b86565b6001600160a01b03821660009081526001602052604090205460ff161515811515146106af576001600160a01b038216600081815260016020908152604091829020805460ff1916851515908117909155825190815291517f065b77b53864e46fda3d8986acb51696223d6dde7ced42441eb150bae6d481369281900390910190a25050565b600381565b6001600160a01b03811660009081526001602052604090205460ff16610994576040805162461bcd60e51b815260206004820152601160248201527034b73b30b634b2103b30b634b230ba37b960791b604482015290519081900360640190fd5b50565b6cffffffffffffffffffffffffff81565b6001600160a01b0381565b6001600160a01b03811660009081526003602052604081205461ffff1680610a07576040805162461bcd60e51b8152602060048201526002602482015261316960f01b604482015290519081900360640190fd5b61ffff811660009081526002602052604090205460ff1615610a55576040805162461bcd60e51b8152602060048201526002602482015261326960f01b604482015290519081900360640190fd5b92915050565b620186a081565b6224ea0081565b610a7233610b86565b6006546001600160a01b0382811691161461099457600680546001600160a01b0383166001600160a01b0319909116811790915560408051918252517fc7708091594580aae77e08cb1e2e3b9d2bd0cab0d2d95797248dd95f02d7299e9181900360200190a150565b6004602052600090815260409020546001600160a01b031681565b610aff33610b86565b6000546001600160a01b0382811691161461099457600080546001600160a01b0383166001600160a01b0319909116811790915560408051918252517f5425363a03f182281120f5919107c49c7a1a623acc1cbc6df468b6f0c11fcf8c9181900360200190a150565b6000546001600160a01b031681565b6006546001600160a01b031681565b6000546001600160a01b03828116911614610994576040805162461bcd60e51b8152602060048201526002602482015261316760f01b604482015290519081900360640190fd5b60016020526000908152604090205460ff168156fea2646970667358221220bb4641da870d05b1551de19e59342231b862f7dbad93d6c235c1f08abe1dfbdc64736f6c63430007060033"

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
// Solidity: event AssetPausedUpdate(address indexed token, bool paused)
func (_Governance *GovernanceFilterer) FilterAssetPausedUpdate(opts *bind.FilterOpts, token []common.Address) (*GovernanceAssetPausedUpdateIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "AssetPausedUpdate", tokenRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceAssetPausedUpdateIterator{contract: _Governance.contract, event: "AssetPausedUpdate", logs: logs, sub: sub}, nil
}

// WatchAssetPausedUpdate is a free log subscription operation binding the contract event 0xf7ca5545623b85829d3abcb6729eb021070bbe93cba64f5c4de6b17b805b7d02.
//
// Solidity: event AssetPausedUpdate(address indexed token, bool paused)
func (_Governance *GovernanceFilterer) WatchAssetPausedUpdate(opts *bind.WatchOpts, sink chan<- *GovernanceAssetPausedUpdate, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "AssetPausedUpdate", tokenRule)
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
// Solidity: event AssetPausedUpdate(address indexed token, bool paused)
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
// Solidity: event NewAsset(address indexed assetAddress, uint16 indexed assetId)
func (_Governance *GovernanceFilterer) FilterNewAsset(opts *bind.FilterOpts, assetAddress []common.Address, assetId []uint16) (*GovernanceNewAssetIterator, error) {

	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "NewAsset", assetAddressRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceNewAssetIterator{contract: _Governance.contract, event: "NewAsset", logs: logs, sub: sub}, nil
}

// WatchNewAsset is a free log subscription operation binding the contract event 0x990c18cf226b253448a6a051b5cadf51a61f4ca56703374cdd8bf9df04b2f901.
//
// Solidity: event NewAsset(address indexed assetAddress, uint16 indexed assetId)
func (_Governance *GovernanceFilterer) WatchNewAsset(opts *bind.WatchOpts, sink chan<- *GovernanceNewAsset, assetAddress []common.Address, assetId []uint16) (event.Subscription, error) {

	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "NewAsset", assetAddressRule, assetIdRule)
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
// Solidity: event NewAsset(address indexed assetAddress, uint16 indexed assetId)
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
// Solidity: event ValidatorStatusUpdate(address indexed validatorAddress, bool isActive)
func (_Governance *GovernanceFilterer) FilterValidatorStatusUpdate(opts *bind.FilterOpts, validatorAddress []common.Address) (*GovernanceValidatorStatusUpdateIterator, error) {

	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ValidatorStatusUpdate", validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceValidatorStatusUpdateIterator{contract: _Governance.contract, event: "ValidatorStatusUpdate", logs: logs, sub: sub}, nil
}

// WatchValidatorStatusUpdate is a free log subscription operation binding the contract event 0x065b77b53864e46fda3d8986acb51696223d6dde7ced42441eb150bae6d48136.
//
// Solidity: event ValidatorStatusUpdate(address indexed validatorAddress, bool isActive)
func (_Governance *GovernanceFilterer) WatchValidatorStatusUpdate(opts *bind.WatchOpts, sink chan<- *GovernanceValidatorStatusUpdate, validatorAddress []common.Address) (event.Subscription, error) {

	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ValidatorStatusUpdate", validatorAddressRule)
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
// Solidity: event ValidatorStatusUpdate(address indexed validatorAddress, bool isActive)
func (_Governance *GovernanceFilterer) ParseValidatorStatusUpdate(log types.Log) (*GovernanceValidatorStatusUpdate, error) {
	event := new(GovernanceValidatorStatusUpdate)
	if err := _Governance.contract.UnpackLog(event, "ValidatorStatusUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
