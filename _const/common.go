package _const

import (
	"time"
)

const (
	InfuraRinkebyNetwork = "wss://rinkeby.infura.io/ws/v3/787bc04c3f044d77a538d519ef26b53e"
	LocalNetwork         = "http://localhost:8545"

	MaxTryTimes = 20

	TryTimeInterval = time.Second * 2

	RinkebySuperSk = "acbaa269bd7573ff12361be4b97201aef019776ea13384681d4e5ba6a88367d9"

	AddressRegexp = "^0x[0-9a-fA-F]{40}$"
	HashRegexp    = "^0x[0-9a-fA-F]{64}$"
	SkRegexp      = "[0-9a-fA-F]{64}$"
	SkRegexp0x    = "^0x[0-9a-fA-F]{64}$"

	SuggestGasLimit         = uint64(21000) // wei
	SuggestContractGasLimit = uint64(300000)

	TransactionSuccess = uint64(1)
	TransactionFail    = uint64(0)
)
