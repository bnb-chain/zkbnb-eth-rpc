package basic

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/bnb-chain/zkbas-eth-rpc/_rpc"
	"math/big"
)

func EmptyCallOpts() *bind.CallOpts {
	return &bind.CallOpts{}
}

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

func SetFixed32Bytes(buf []byte) [32]byte {
	newBuf := new(big.Int).SetBytes(buf).FillBytes(make([]byte, Bytes32Len))
	var res [32]byte
	if len(buf) != 0 {
		copy(res[:], newBuf[:])
	}
	return res
}
