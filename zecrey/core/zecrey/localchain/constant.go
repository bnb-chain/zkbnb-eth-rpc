package localchain

import (
	"context"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
	"math/big"
)

const (
	NetworkEndPoint = "http://localhost:8545"
	SuperAddress    = "0xDA00601380Bc7aE4fe67dA2EB78f9161570c9EB4"
	SuperSk         = "a36902d14b35e3ed9a288bebd513baa77b3772c6263d6fefff70fadf12fe097a"
	L2ChainId       = uint8(0)
	AddrFileName    = "contractAddresses.txt"

	// native chain id
	NativeChainId = uint8(0)
	// native asset id
	NativeAssetId = uint16(1)
	// max pending blocks
	MaxPendingBlocks = uint16(100)

	Dir = "D:\\Projects\\mygo\\src\\Zecrey\\SherLzp\\zecrey-eth-rpc\\zecrey\\core\\zecrey\\localchain\\"
)

var (
	ChainId            = big.NewInt(5777)
	Cli, _             = _rpc.NewClient(NetworkEndPoint)
	SuggestGasPrice, _ = Cli.SuggestGasPrice(context.Background())
	AuthCli, _         = _rpc.NewAuthClient(Cli, SuperSk, ChainId)

	OnChainOpsMerkleHelper = [4]bool{false, false, false, false}
)
