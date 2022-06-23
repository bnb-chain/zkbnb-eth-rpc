package legend

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/bnb-chain/zkbas-eth-rpc/_rpc"
)

func LoadGovernanceInstance(cli *_rpc.ProviderClient, addr string) (instance *Governance, err error) {
	instance, err = NewGovernance(common.HexToAddress(addr), *cli)
	return instance, err
}