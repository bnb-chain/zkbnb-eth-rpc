package _const

import "time"

const (
	// test network
	InfuraRinkebyNetwork = "https://rinkeby.infura.io/v3/787bc04c3f044d77a538d519ef26b53e"
	RopstenNetwork       = "https://ropsten.infura.io/v3/787bc04c3f044d77a538d519ef26b53e"
	LocalNetwork         = "http://localhost:8545"
	// max try times
	MaxTryTimes = 10
	// time interval for obtaining transaction status
	TryTimeInterval = time.Second * 2

	// super account
	SuperAddress = "0xFEf01c3494e9fbad65A5D12B3852CA87361Bc9B2"
	SuperSk      = "999435f70f6b448a3f751f879357c2870069cac8b9583a0780bfa36fe9835531"

	// test address
	ToAddress = "0xE9b15a2D396B349ABF60e53ec66Bcf9af262D449"

	// regexp
	AddressRegexp = "^0x[0-9a-fA-F]{40}$"
	HashRegexp    = "^0x[0-9a-fA-F]{64}$"
	SkRegexp      = "[0-9a-fA-F]{64}$"
	SkRegexp0x    = "^0x[0-9a-fA-F]{64}$"

	// chain
	SuggestGasLimit         = uint64(21000) // wei
	SuggestFunctionGasLimit = uint64(80000)
	SuggestContractGasLimit = uint64(300000)

	// Token address
	EVATokenAddress = "0x59e39e65fc4e35f7b1a21d51962109fa9eed47df"

	// Transaction status
	TransactionSuccess = uint64(1)
	TransactionFail    = uint64(0)
)
