package contract

import (
	"eva-go-rpc/_const"
	"eva-go-rpc/_rpc"
	"eva-go-rpc/_rpc/contract/_interface/erc20"
	"eva-go-rpc/_utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// get erc20 token balance
func GetErc20Balance(address string) (balance *big.Int, err error) {
	// validate the address
	if !_utils.IsValidEthAddress(address) {
		return nil, _rpc.InvalidAddress
	}
	cli, err := _rpc.GetConnection()
	if err != nil {
		return nil, err
	}
	defer cli.Close()
	account := common.HexToAddress(address)
	evaTokenAddress := common.HexToAddress(_const.EVATokenAddress)
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
	balance = _utils.WeiToEther(_balance)
	return balance, nil
}
