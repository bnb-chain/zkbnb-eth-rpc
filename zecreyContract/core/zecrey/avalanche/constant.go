package avalanche

import (
	"github.com/zecrey-labs/zecrey-eth-rpc/_const"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
	"math/big"
)

const (
	NetworkEndPoint = "http://localhost:8545"
	SuperAddress    = "0x89D37ea8a0f102D90C424141F897A6a764A291AF"
	SuperSk         = "ad6ad08487f7d8c96450f71dba9d8a4dc7e0924bdd62eda59962685577db1068"
)

var (
	ChainId    = big.NewInt(4)
	Cli, _     = _rpc.NewClient(NetworkEndPoint)
	AuthCli, _ = _rpc.NewAuthClient(Cli, _const.LocalSuperSk, ChainId)
	L2ChainId  = uint8(1)
)
