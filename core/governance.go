package core

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/bnb-chain/zkbnb-eth-rpc/rpc"
)

func LoadGovernanceInstance(cli *rpc.ProviderClient, addr string) (instance *Governance, err error) {
	instance, err = NewGovernance(common.HexToAddress(addr), *cli)
	return instance, err
}
