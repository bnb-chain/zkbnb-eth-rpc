package legend

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/bnb-chain/zkbnb-eth-rpc/rpc"
)

func LoadERC20(cli *rpc.ProviderClient, addr string) (instance *Erc20, err error) {
	instance, err = NewErc20(common.HexToAddress(addr), *cli)
	return instance, err
}
