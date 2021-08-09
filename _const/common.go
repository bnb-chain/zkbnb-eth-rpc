package _const

import "time"

const (
	// test network
	InfuraRinkebyNetwork = "wss://rinkeby.infura.io/ws/v3/787bc04c3f044d77a538d519ef26b53e"
	RopstenNetwork       = "wss://ropsten.infura.io/ws/v3/787bc04c3f044d77a538d519ef26b53e"
	LocalNetwork         = "http://localhost:8545"
	AuroraTestNetwork    = "https://testnet.aurora.dev"
	AuroraMainNetwork    = "https://mainnet.aurora.dev"
	ArbitrumTestNetwork  = "https://rinkeby.arbitrum.io/rpc"
	PolyTestNetwork      = "https://matic-mumbai.chainstacklabs.com"

	// max try times
	MaxTryTimes = 10
	// time interval for obtaining transaction status
	TryTimeInterval = time.Second * 2

	// rinkeby super account
	RinkebySuperAddress = "0xE9b15a2D396B349ABF60e53ec66Bcf9af262D449"
	RinkebySuperSk      = "acbaa269bd7573ff12361be4b97201aef019776ea13384681d4e5ba6a88367d9"

	// local network super account
	LocalSuperAddress  = "0x89D37ea8a0f102D90C424141F897A6a764A291AF"
	LocalSuperSk       = "ad6ad08487f7d8c96450f71dba9d8a4dc7e0924bdd62eda59962685577db1068"
	LocalSuperSkBigInt = "78438847730043804073117709805401039734901665551211393154078937917603918319720"

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
	SuggestHighGasLimit     = uint64(5000000)

	// Token address
	EVATokenAddress = "0x59e39e65fc4e35f7b1a21d51962109fa9eed47df"
	ZKSneakAddress  = "0x95F45C568f146c409F92321B1123505Cbe84719B"

	// Token address
	ZecreyTokenAddress      = "0x1874D6B0B85fF4AD526163a4178064aEaB55D9DD"
	ZecreyAddress           = "0x9c460ccF59E01fFc66a48378Fb3957BEbf2123a2"
	ZecreyGovernanceAddress = "0xa54f5Eb49cd28088D91aD201387Bfc42e1711F69"

	// Transaction status
	TransactionSuccess = uint64(1)
	TransactionFail    = uint64(0)
)
