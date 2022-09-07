package avalanche

import (
	"context"
	"github.com/bnb-chain/zkbnb-eth-rpc/_rpc"
	"math/big"
)

const (
	NetworkEndPoint = "https://api.avax-test.network/ext/bc/C/rpc"
	SuperAddress    = "0xE9b15a2D396B349ABF60e53ec66Bcf9af262D449"
	SuperSk         = "acbaa269bd7573ff12361be4b97201aef019776ea13384681d4e5ba6a88367d9"
	L2ChainId       = uint8(3)
	AddrFileName    = "contractAddresses.txt"

	ChainConfFileName = "chain_conf_avalanche.json"

	// native chain id
	NativeChainId = uint8(3)
	// native asset id
	NativeAssetId = uint16(4)
	// max pending blocks
	MaxPendingBlocks = uint16(100)

	ChainName = "Avalanche"

	Dir = "D:\\Projects\\mygo\\src\\ZkBNB\\SherLzp\\zkbnb-eth-rpc\\zkbnb\\core\\zkbnb\\avalanche\\"
)

var (
	ChainId            = big.NewInt(43113)
	Cli, _             = _rpc.NewClient(NetworkEndPoint)
	SuggestGasPrice, _ = Cli.SuggestGasPrice(context.Background())
	AuthCli, _         = _rpc.NewAuthClient(Cli, SuperSk, ChainId)

	OnChainOpsMerkleHelper = [4]bool{true, true, false, false}
)
