package zecreyLegend

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
)

func LoadGovernanceInstance(cli *_rpc.ProviderClient, addr string) (instance *Governance, err error) {
	instance, err = NewGovernance(common.HexToAddress(addr), *cli)
	return instance, err
}
