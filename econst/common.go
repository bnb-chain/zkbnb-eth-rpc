package econst

import "time"

const (
	// test network
	InfuraRinkebyNetwork = "https://rinkeby.infura.io/v3/3ea7b84f02e34f36874d9e26a104b6e9"
	// max try times
	MaxTryTimes = 10
	// time interval for obtaining transaction status
	TryTimeInterval = time.Second * 2

	// super account
	SuperAddress = "0xFEf01c3494e9fbad65A5D12B3852CA87361Bc9B2"
	SuperSk      = "999435f70f6b448a3f751f879357c2870069cac8b9583a0780bfa36fe9835531"

	// regexp
	AddressRegexp = "^0x[0-9a-fA-F]{40}$"
	HashRegexp    = "^0x[0-9a-fA-F]{64}$"
	SkRegexp      = "[0-9a-fA-F]{64}$"
	SkRegexp0x    = "^0x[0-9a-fA-F]{64}$"

	// chain
	SuggestGasLimit         = uint64(21000) // wei
	SuggestContractGasLimit = uint64(300000)

	// Token address
	EVATokenAddress = "0x59e39e65fc4e35f7b1a21d51962109fa9eed47df"

	// Transaction status
	TransactionSuccess = uint64(1)
	TransactionFail    = uint64(0)
)
