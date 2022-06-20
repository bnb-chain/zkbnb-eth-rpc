package _rpc

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/bnb-chain/zkbas-eth-rpc/_const"
	"math/big"
)

// sign transaction
func SignTx(authCli *AuthClient, tx *types.Transaction) (signedTx *types.Transaction, err error) {
	// check if it is valid params
	if authCli.ChainId == nil || authCli.PrivateKey == nil {
		return nil, ErrInvalidAuthClientParams
	}
	return types.SignTx(tx, types.NewEIP155Signer(authCli.ChainId), authCli.PrivateKey)
}

func CreateAuthentication(privateKey *ecdsa.PrivateKey, nonce *big.Int, value *big.Int, gasLimit uint64, gasPrice *big.Int) *bind.TransactOpts {
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = nonce
	auth.Value = value
	auth.GasLimit = _const.SuggestContractGasLimit
	auth.GasPrice = gasPrice
	return auth
}
