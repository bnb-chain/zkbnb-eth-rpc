package eth

import (
	"context"
	"github.com/bnb-chain/zkbnb-eth-rpc/_rpc"
	"math/big"
)

const (
	NetworkEndPoint = "wss://rinkeby.infura.io/ws/v3/787bc04c3f044d77a538d519ef26b53e"
	SuperAddress    = "0xE9b15a2D396B349ABF60e53ec66Bcf9af262D449"
	SuperSk         = "acbaa269bd7573ff12361be4b97201aef019776ea13384681d4e5ba6a88367d9"
	L2ChainId       = uint8(0)
	AddrFileName    = "contractAddresses.txt"

	ChainConfFileName = "chain_conf_ethereum.json"

	// native chain id
	NativeChainId = uint8(0)
	// native asset id
	NativeAssetId = uint16(1)
	// max pending blocks
	MaxPendingBlocks = uint16(100)

	ChainName = "Ethereum"
	Dir       = "D:\\Projects\\mygo\\src\\ZkBNB\\SherLzp\\zkbnb-eth-rpc\\zkbnb\\core\\zkbnb\\eth\\"
)

var (
	ChainId            = big.NewInt(4)
	Cli, _             = _rpc.NewClient(NetworkEndPoint)
	SuggestGasPrice, _ = Cli.SuggestGasPrice(context.Background())
	AuthCli, _         = _rpc.NewAuthClient(Cli, SuperSk, ChainId)

	OnChainOpsMerkleHelper = [4]bool{false, false, false, false}
)
