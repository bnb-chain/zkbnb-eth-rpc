package core

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/bnb-chain/zkbnb-eth-rpc/rpc"
)

func LoadERC20(cli *rpc.ProviderClient, addr string) (instance *ERC20, err error) {
	instance, err = NewERC20(common.HexToAddress(addr), *cli)
	return instance, err
}
