package legend

import (
	"github.com/bnb-chain/zkbnb-eth-rpc/_rpc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

func ConstructTransactOpts(cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, gasPrice *big.Int, gasLimit uint64) (transactOpts *bind.TransactOpts, err error) {
	transactOpts, err = bind.NewKeyedTransactorWithChainID(authCli.PrivateKey, authCli.ChainId)
	if err != nil {
		return nil, err
	}
	nonce, err := cli.GetPendingNonce(authCli.Address.Hex())
	if err != nil {
		return nil, err
	}
	transactOpts.Nonce = big.NewInt(int64(nonce))
	transactOpts.GasPrice = gasPrice
	transactOpts.GasLimit = gasLimit
	transactOpts.From = authCli.Address
	transactOpts.Value = big.NewInt(0)
	return transactOpts, nil
}
