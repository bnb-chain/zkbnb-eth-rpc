package rpc

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/bnb-chain/zkbnb-eth-rpc/utils"
)

type AuthClient struct {
	PrivateKey *ecdsa.PrivateKey `json:"private_key"`
	PublicKey  *ecdsa.PublicKey  `json:"public_key"`
	Address    common.Address    `json:"address"`
	ChainId    *big.Int          `json:"chain_id"`
}

func NewAuthClient(priKey string, chainId *big.Int) (authCli *AuthClient, err error) {
	// validate private key
	if !utils.IsValidPrivateKey(priKey) {
		return nil, ErrInvalidPrivateKey
	}
	// transfer private key str to private key
	privateKey, err := utils.DecodePrivateKey(priKey)
	if err != nil {
		return nil, err
	}
	// get public key
	publicKey := privateKey.PublicKey
	address := crypto.PubkeyToAddress(publicKey)
	return &AuthClient{PrivateKey: privateKey, PublicKey: &publicKey, Address: address, ChainId: chainId}, nil
}

func (c *AuthClient) ConstructTransactOpts(cli *ProviderClient, gasPrice *big.Int, gasLimit uint64) (transactOpts *bind.TransactOpts, err error) {
	transactOpts, err = bind.NewKeyedTransactorWithChainID(c.PrivateKey, c.ChainId)
	if err != nil {
		return nil, err
	}
	nonce, err := cli.GetPendingNonce(c.Address.Hex())
	if err != nil {
		return nil, err
	}
	transactOpts.Nonce = big.NewInt(int64(nonce))
	transactOpts.GasPrice = gasPrice
	transactOpts.GasLimit = gasLimit
	transactOpts.From = c.Address
	transactOpts.Value = big.NewInt(0)
	return transactOpts, nil
}

func (c *AuthClient) ConstructTransactOptsWithNonce(gasPrice *big.Int, gasLimit uint64, nonce uint64) (transactOpts *bind.TransactOpts, err error) {
	transactOpts, err = bind.NewKeyedTransactorWithChainID(c.PrivateKey, c.ChainId)
	if err != nil {
		return nil, err
	}
	transactOpts.Nonce = big.NewInt(int64(nonce))
	transactOpts.GasPrice = gasPrice
	transactOpts.GasLimit = gasLimit
	transactOpts.From = c.Address
	transactOpts.Value = big.NewInt(0)
	return transactOpts, nil
}
