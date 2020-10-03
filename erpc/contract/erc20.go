package contract

import (
	"eva-go-sdk/econst"
	"eva-go-sdk/erpc"
	"eva-go-sdk/erpc/contract/einterface/erc20"
	"eva-go-sdk/eutils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// get erc20 token balance
func GetErc20Balance(address string) (balance *big.Int, err error) {
	// validate the address
	if !eutils.IsValidEthAddress(address) {
		return nil, erpc.InvalidAddress
	}
	cli, err := erpc.GetConnection()
	if err != nil {
		return nil, err
	}
	defer cli.Close()
	account := common.HexToAddress(address)
	evaTokenAddress := common.HexToAddress(econst.EVATokenAddress)
	token, err := erc20.NewToken(evaTokenAddress, cli)
	if err != nil {
		return nil, err
	}
	return token.BalanceOf(&bind.CallOpts{}, account)
}

func GetErc20EtherBalance(address string) (balance *big.Float, err error) {
	_balance, err := GetErc20Balance(address)
	if err != nil {
		return nil, err
	}
	balance = eutils.WeiToEther(_balance)
	return balance, nil
}
