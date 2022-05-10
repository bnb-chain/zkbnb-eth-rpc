package aurora

import (
	"context"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
	"math/big"
)

const (
	NetworkEndPoint = "https://testnet.aurora.dev"
	SuperAddress    = "0xE9b15a2D396B349ABF60e53ec66Bcf9af262D449"
	SuperSk         = "acbaa269bd7573ff12361be4b97201aef019776ea13384681d4e5ba6a88367d9"
	L2ChainId       = uint8(2)
	AddrFileName    = "contractAddresses.txt"

	ChainConfFileName = "chain_conf_aurora.json"

	// native chain id
	NativeChainId = uint8(2)
	// native asset id
	NativeAssetId = uint16(1)
	// max pending blocks
	MaxPendingBlocks = uint16(100)

	ChainName = "Aurora"
	Dir       = "D:\\Projects\\mygo\\src\\Zecrey\\SherLzp\\zecrey-eth-rpc\\zecrey\\core\\zecrey\\aurora\\"
)

var (
	ChainId            = big.NewInt(1313161555)
	Cli, _             = _rpc.NewClient(NetworkEndPoint)
	SuggestGasPrice, _ = Cli.SuggestGasPrice(context.Background())
	AuthCli, _         = _rpc.NewAuthClient(Cli, SuperSk, ChainId)

	OnChainOpsMerkleHelper = [4]bool{false, true, false, false}
)
