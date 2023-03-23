package core

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/bnb-chain/zkbnb-eth-rpc/rpc"
)

func LoadGovernanceInstance(cli *rpc.ProviderClient, addr string) (instance *Governance, err error) {
	instance, err = NewGovernance(common.HexToAddress(addr), *cli)
	return instance, err
}

func ValidateAssetAddress(instance *Governance, assetAddr common.Address,
) (uint16, error) {
	uint16, err := instance.ValidateAssetAddress(&bind.CallOpts{}, assetAddr)
	if err != nil {
		return 0, err
	}
	return uint16, nil
}
