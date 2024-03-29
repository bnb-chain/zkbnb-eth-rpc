package rpc

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/bnb-chain/zkbnb-eth-rpc/constants"
)

func SignTx(authCli *AuthClient, tx *types.Transaction) (signedTx *types.Transaction, err error) {
	// check if it is valid params
	if authCli.ChainId == nil || authCli.PrivateKey == nil {
		return nil, ErrInvalidAuthClientParams
	}
	return types.SignTx(tx, types.NewEIP155Signer(authCli.ChainId), authCli.PrivateKey)
}

func CreateAuthentication(privateKey *ecdsa.PrivateKey, nonce *big.Int, value *big.Int, gasPrice *big.Int) *bind.TransactOpts {
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = nonce
	auth.Value = value
	auth.GasLimit = constants.SuggestContractGasLimit
	auth.GasPrice = gasPrice
	return auth
}
