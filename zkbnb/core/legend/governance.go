package legend

import (
	"github.com/bnb-chain/zkbnb-eth-rpc/_rpc"
	"github.com/ethereum/go-ethereum/common"
)

func LoadGovernanceInstance(cli *_rpc.ProviderClient, addr string) (instance *Governance, err error) {
	instance, err = NewGovernance(common.HexToAddress(addr), *cli)
	return instance, err
}
