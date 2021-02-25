package _rpc

import "github.com/ethereum/go-ethereum/core/types"

// sign transaction
func SignTx(authCli *AuthClient, tx *types.Transaction) (signedTx *types.Transaction, err error) {
	// check if it is valid params
	if authCli.ChainId == nil || authCli.PrivateKey == nil {
		return nil, InvalidAuthClientParams
	}
	return types.SignTx(tx, types.NewEIP155Signer(authCli.ChainId), authCli.PrivateKey)
}
