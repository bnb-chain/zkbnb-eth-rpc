package bsc

import (
	"context"
	"github.com/bnb-chain/zkbas-eth-rpc/_rpc"
	"math/big"
)

const (
	NetworkEndPoint = "https://data-seed-prebsc-1-s1.binance.org:8545/"
	SuperAddress    = "0xE9b15a2D396B349ABF60e53ec66Bcf9af262D449"
	SuperSk         = "acbaa269bd7573ff12361be4b97201aef019776ea13384681d4e5ba6a88367d9"
	L2ChainId       = uint8(4)
	AddrFileName    = "contractAddresses.txt"

	ChainConfFileName = "chain_conf_bsc.json"

	// native chain id
	NativeChainId = uint8(4)
	// native asset id
	NativeAssetId = uint16(9)
	// max pending blocks
	MaxPendingBlocks = uint16(100)

	ChainName = "BSC"
	Dir       = "D:\\Projects\\mygo\\src\\Zecrey\\SherLzp\\zecrey-eth-rpc\\zecrey\\core\\zecrey\\bsc\\"
)

var (
	ChainId            = big.NewInt(97)
	Cli, _             = _rpc.NewClient(NetworkEndPoint)
	SuggestGasPrice, _ = Cli.SuggestGasPrice(context.Background())
	AuthCli, _         = _rpc.NewAuthClient(Cli, SuperSk, ChainId)
)
