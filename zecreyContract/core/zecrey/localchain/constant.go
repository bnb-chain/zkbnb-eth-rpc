package localchain

import (
	"context"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
	"math/big"
)

const (
	NetworkEndPoint = "http://localhost:8545"
	SuperAddress    = "0x5ee4C56aab8F8A1E3dC524A890E06D2f0c9073cE"
	SuperSk         = "f2c99690a0c10dcdfefeffcb8d1e088533261bc47b73ea790eb7275b03afcaeb"
	L2ChainId       = uint8(0)
	AddrFileName    = "contractAddresses.txt"

	// native chain id
	NativeChainId = uint8(0)
	// native asset id
	NativeAssetId = uint16(1)
	// max pending blocks
	MaxPendingBlocks = uint16(100)

	Dir = "D:\\Projects\\mygo\\src\\Zecrey\\SherLzp\\zecrey-eth-rpc\\zecreyContract\\core\\zecrey\\localchain\\"
)

var (
	ChainId            = big.NewInt(5777)
	Cli, _             = _rpc.NewClient(NetworkEndPoint)
	SuggestGasPrice, _ = Cli.SuggestGasPrice(context.Background())
	AuthCli, _         = _rpc.NewAuthClient(Cli, SuperSk, ChainId)

	OnChainOpsMerkleHelper = [4]bool{false, false, false, false}
)
