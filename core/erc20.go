package core

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"

	"github.com/bnb-chain/zkbnb-eth-rpc/rpc"
)

func LoadERC20(cli *rpc.ProviderClient, addr string) (instance *ERC20, err error) {
	instance, err = NewERC20(common.HexToAddress(addr), *cli)
	return instance, err
}

func BalanceOf(instance *ERC20, address common.Address, assetAddr common.Address) (*big.Int, error) {
	amount, err := instance.BalanceOf(&bind.CallOpts{From: address}, assetAddr)
	if err != nil {
		return nil, err
	}
	return amount, nil
}
