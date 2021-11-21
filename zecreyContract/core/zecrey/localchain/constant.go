package localchain

import (
	"context"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
	"math/big"
)

const (
	NetworkEndPoint = "http://localhost:8545"
	SuperAddress    = "0x5ee4C56aab8F8A1E3dC524A890E06D2f0c9073cE"
	SuperSk         = "59b99397f23d9ee2b816a875d56f41b28e4e036d82aa0801b72fc2453f41f532"
	L2ChainId       = uint8(1)
	AddrFileName    = "contractAddresses.txt"
)

var (
	ChainId            = big.NewInt(5777)
	Cli, _             = _rpc.NewClient(NetworkEndPoint)
	SuggestGasPrice, _ = Cli.SuggestGasPrice(context.Background())
	AuthCli, _         = _rpc.NewAuthClient(Cli, SuperSk, ChainId)
)
