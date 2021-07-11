package zecrey

import (
	"Zecrey-eth-rpc/_rpc"
	"math/big"
)

var (
	InitialTokenSupply, _ = new(big.Int).SetString("100000000000000000000000000", 10)
	RinkebyChainId        = _rpc.RinkebyChainId
)
