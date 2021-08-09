package zecrey

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
	"math/big"
	"strings"
)

var (
	InitialTokenSupply, _ = new(big.Int).SetString("100000000000000000000000000", 10)
	RinkebyChainId        = _rpc.RinkebyChainId
	AuroraChainId         = _rpc.AuroraChainId
	ArbitrumChainId       = _rpc.ArbitrumChainId
	PolyChainId           = big.NewInt(80001)
	LocalChainId          = big.NewInt(5777)
	zecreyABI, _          = abi.JSON(strings.NewReader(ZecreyABI))
)
