package core

import (
	"github.com/bnb-chain/zkbnb-eth-rpc/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

type TransactOptsConstructor interface {
	ConstructTransactOpts(cli *rpc.ProviderClient, gasPrice *big.Int, gasLimit uint64) (transactOpts *bind.TransactOpts, err error)
	ConstructTransactOptsWithNonce(gasPrice *big.Int, gasLimit uint64, nonce uint64) (transactOpts *bind.TransactOpts, err error)
}
