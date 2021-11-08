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

// VerifierABI is the input ABI used to generate the binding from.
const VerifierABI = "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"input\",\"type\":\"uint256[2]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VerifierBin is the compiled bytecode used for deploying new contracts.
var VerifierBin = "0x608060405234801561001057600080fd5b50610f5a806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063f5c9d69e14610030575b600080fd5b610129600480360361014081101561004757600080fd5b6040805180820182529183019291818301918390600290839083908082843760009201829052506040805180820190915293969594608081019493509150600290835b828210156100c8576040805180820182529080840286019060029083908390808284376000920191909152505050815260019091019060200161008a565b50506040805180820182529396959481810194935091506002908390839080828437600092019190915250506040805180820182529295949381810193925090600290839083908082843760009201919091525091945061013d9350505050565b604080519115158252519081900360200190f35b6000610147610dac565b6040805180820182528751815260208089015181830152908352815160808101835287515181840190815288518301516060830152815282518084018452888301805151825251830151818401528183015283820152815180830183528651815286820151918101919091529082015260006101c1610645565b604080518082019091526000808252602082015283515191925090600080516020610f058339815191521161023d576040805162461bcd60e51b815260206004820152601760248201527f76657269666965722d61582d6774652d7072696d652d71000000000000000000604482015290519081900360640190fd5b825160200151600080516020610f05833981519152116102a4576040805162461bcd60e51b815260206004820152601760248201527f76657269666965722d61592d6774652d7072696d652d71000000000000000000604482015290519081900360640190fd5b60208301515151600080516020610f058339815191521161030c576040805162461bcd60e51b815260206004820152601860248201527f76657269666965722d6258302d6774652d7072696d652d710000000000000000604482015290519081900360640190fd5b602083810151015151600080516020610f0583398151915211610376576040805162461bcd60e51b815260206004820152601860248201527f76657269666965722d6259302d6774652d7072696d652d710000000000000000604482015290519081900360640190fd5b602083810151510151600080516020610f05833981519152116103e0576040805162461bcd60e51b815260206004820152601860248201527f76657269666965722d6258312d6774652d7072696d652d710000000000000000604482015290519081900360640190fd5b6020838101518101510151600080516020610f058339815191521161044c576040805162461bcd60e51b815260206004820152601860248201527f76657269666965722d6259312d6774652d7072696d652d710000000000000000604482015290519081900360640190fd5b604083015151600080516020610f05833981519152116104b3576040805162461bcd60e51b815260206004820152601760248201527f76657269666965722d63582d6774652d7072696d652d71000000000000000000604482015290519081900360640190fd5b600080516020610f058339815191528360400151602001511061051d576040805162461bcd60e51b815260206004820152601760248201527f76657269666965722d63592d6774652d7072696d652d71000000000000000000604482015290519081900360640190fd5b60005b60028110156105f1577f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000186826002811061055657fe5b6020020151106105ad576040805162461bcd60e51b815260206004820152601f60248201527f76657269666965722d6774652d736e61726b2d7363616c61722d6669656c6400604482015290519081900360640190fd5b6105e7826105e2856080015184600101600381106105c757fe5b60200201518985600281106105d857fe5b60200201516109bc565b610a51565b9150600101610520565b50608082015151610603908290610a51565b90506106396106158460000151610ae2565b84602001518460000151856020015185876040015189604001518960600151610b65565b98975050505050505050565b61064d610dde565b6040805180820182527f0c7509c153ce226f510b4f599eaf17a5ae1a59c95d40a240b370ab38ca75370481527f1aa5ff7f04e3b8132032ced67a5013b5cc2b508c75e58e6fc28aa364596191bc6020808301919091529083528151608080820184527f022bd68e4b74a5a21cd3978d57a7f5d41443e32ad47eba3acf504d6f352b295c8285019081527f26e1f310e560cd871efcfda76ee3e23eee1f1fd08cde0cd323080a08b765195f606080850191909152908352845180860186527f2b74c16cec250634d82672a28e161a80fd796bfa6c554ac72e9b0696d7925a0c81527f0cda2e7448fd8312f189ac1b17dd60cc7752782215659f0604ad30308377c567818601528385015285840192909252835180820185527f1510119c86df39d9a77da91550581aab38f8828a1eabce6285b380cbf99857038186019081527ec77eb3c390b84ad18b8ba5c2af50fb4792af36b1c2fd3dbdc6ee5b5ec51506828501528152845180860186527f1f97654a2ad8ec22053a824fe6e796b8c2ffbc224f835b79bee2b591fe73e6f581527f139665663c7eab792cceb88de58ba7f3a83171beebe888daf54c8875ce53464b818601528185015285850152835180820185527f0efd410606bc9137ee5f38e08f6ac2ec2b3cc35b37a0c63598891d0156f096658186019081527f2b456c7aed96a7ed9f2349ff15bbb4012b65215663d2173cb732b5d018722498828501528152845180860186527e92a4155a9ed4a07b53f6320b93637c36e954f8970e92e2fa06075750c42a3281527f1f03264a4f869e2108d5bbb392f88186657d02751ff79a31bac99ec4a9de1153818601528185015291850191909152825180840184527f0f9bd37b3a0136a23f716a8f8d9ccb179ffe7cd6fb247405696207bfce026ff581527f12ae39135bf9e01fb9bd89af169a65ba465eff07b396e27332a1345c72733aee81840152908401805191909152825180840184527f228d861a32ecca69d116a73b9862f8bc69f8c7b6d718b0aa422c6c70e54c857581527f242ec2d5211a3a3a3435178c300a100f049c839b77a97807aa2747338926e88c818401528151830152825180840184527f1b69b846c07bdb7a0502f49742171e7f55b5c4d62f566fe51c0756489b27f77581527f2fa40c58a1d862c05b8efdcedf80609f3099f925cad35a91ae730e938c1562f292810192909252519091015290565b6109c4610e25565b6109cc610e3f565b835181526020808501519082015260408101839052600060608360808460076107d05a03fa90508080156109ff57610a01565bfe5b5080610a49576040805162461bcd60e51b81526020600482015260126024820152711c185a5c9a5b99cb5b5d5b0b59985a5b195960721b604482015290519081900360640190fd5b505092915050565b610a59610e25565b610a61610e5d565b8351815260208085015181830152835160408301528301516060808301919091526000908360c08460066107d05a03fa90508080156109ff575080610a49576040805162461bcd60e51b81526020600482015260126024820152711c185a5c9a5b99cb5859190b59985a5b195960721b604482015290519081900360640190fd5b610aea610e25565b8151158015610afb57506020820151155b15610b1a57506040805180820190915260008082526020820152610b60565b604051806040016040528083600001518152602001600080516020610f05833981519152846020015181610b4a57fe5b06600080516020610f0583398151915203905290505b919050565b60408051608080820183528a825260208083018a90528284018890526060808401879052845192830185528b83528282018a9052828501889052820185905283516018808252610320820190955260009491859190839082016103008036833701905050905060005b6004811015610d255760068102858260048110610be757fe5b6020020151518351849083908110610bfb57fe5b602002602001018181525050858260048110610c1357fe5b602002015160200151838260010181518110610c2b57fe5b602002602001018181525050848260048110610c4357fe5b602002015151518351849060028401908110610c5b57fe5b602002602001018181525050848260048110610c7357fe5b60200201515160016020020151838260030181518110610c8f57fe5b602002602001018181525050848260048110610ca757fe5b602002015160200151600060028110610cbc57fe5b6020020151838260040181518110610cd057fe5b602002602001018181525050848260048110610ce857fe5b602002015160200151600160028110610cfd57fe5b6020020151838260050181518110610d1157fe5b602090810291909101015250600101610bce565b50610d2e610e7b565b6000602082602086026020860160086107d05a03fa90508080156109ff575080610d97576040805162461bcd60e51b81526020600482015260156024820152741c185a5c9a5b99cb5bdc18dbd9194b59985a5b1959605a1b604482015290519081900360640190fd5b505115159d9c50505050505050505050505050565b6040518060600160405280610dbf610e25565b8152602001610dcc610e99565b8152602001610dd9610e25565b905290565b6040518060a00160405280610df1610e25565b8152602001610dfe610e99565b8152602001610e0b610e99565b8152602001610e18610e99565b8152602001610dd9610eb9565b604051806040016040528060008152602001600081525090565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060400160405280610eac610ee6565b8152602001610dd9610ee6565b60405180606001604052806003905b610ed0610e25565b815260200190600190039081610ec85790505090565b6040518060400160405280600290602082028036833750919291505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a264697066735822122069f548af800ddcec6c015f2471b0ffc975101a0daf79674830fc0e3387335b8e64736f6c63430007060033"

// DeployVerifier deploys a new Ethereum contract, binding an instance of Verifier to it.
func DeployVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Verifier, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifierABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Verifier{VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

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

// VerifyProof is a free data retrieval call binding the contract method 0xf5c9d69e.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[2] input) view returns(bool r)
func (_Verifier *VerifierCaller) VerifyProof(opts *bind.CallOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [2]*big.Int) (bool, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "verifyProof", a, b, c, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0xf5c9d69e.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[2] input) view returns(bool r)
func (_Verifier *VerifierSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [2]*big.Int) (bool, error) {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, a, b, c, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xf5c9d69e.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[2] input) view returns(bool r)
func (_Verifier *VerifierCallerSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [2]*big.Int) (bool, error) {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, a, b, c, input)
}
