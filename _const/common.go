package _const

import "time"

const (
	// test network
	InfuraRinkebyNetwork = "wss://rinkeby.infura.io/ws/v3/787bc04c3f044d77a538d519ef26b53e"
	RopstenNetwork       = "wss://ropsten.infura.io/ws/v3/787bc04c3f044d77a538d519ef26b53e"
	LocalNetwork         = "http://localhost:8545"
	// max try times
	MaxTryTimes = 10
	// time interval for obtaining transaction status
	TryTimeInterval = time.Second * 2

	// rinkeby super account
	RinkebySuperAddress = "0xFEf01c3494e9fbad65A5D12B3852CA87361Bc9B2"
	RinkebySuperSk      = "999435f70f6b448a3f751f879357c2870069cac8b9583a0780bfa36fe9835531"

	// local network super account
	LocalSuperAddress = "0x89D37ea8a0f102D90C424141F897A6a764A291AF"
	LocalSuperSk      = "ad6ad08487f7d8c96450f71dba9d8a4dc7e0924bdd62eda59962685577db1068"

	// test address
	ToAddress = "0xbE008826B14b8E65Dee9104fB67D47a09Cdbbd0E"

	// regexp
	AddressRegexp = "^0x[0-9a-fA-F]{40}$"
	HashRegexp    = "^0x[0-9a-fA-F]{64}$"
	SkRegexp      = "[0-9a-fA-F]{64}$"
	SkRegexp0x    = "^0x[0-9a-fA-F]{64}$"

	// chain
	SuggestGasLimit         = uint64(21000) // wei
	SuggestFunctionGasLimit = uint64(80000)
	SuggestContractGasLimit = uint64(300000)
	SuggestHighGasLimit     = uint64(1000000)

	// Token address
	EVATokenAddress = "0x59e39e65fc4e35f7b1a21d51962109fa9eed47df"
	ZKSNeakAddress  = "0x95F45C568f146c409F92321B1123505Cbe84719B"

	// Transaction status
	TransactionSuccess = uint64(1)
	TransactionFail    = uint64(0)
)
