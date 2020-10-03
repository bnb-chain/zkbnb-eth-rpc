package _rpc

import (
	"context"
	"crypto/ecdsa"
	"eva-go-rpc/_const"
	"eva-go-rpc/_utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

// get connection of ethereum or rinkeby network
func GetConnection() (cli *ethclient.Client, err error) {
	return ethclient.Dial(_const.InfuraRinkebyNetwork)
}

// auth of user
type AuthClient struct {
	PrivateKey *ecdsa.PrivateKey `json:"private_key"`
	PublicKey  *ecdsa.PublicKey  `json:"public_key"`
	Address    common.Address    `json:"address"`
	ChainId    *big.Int          `json:"chain_id"`
}

// create a new auth client
func NewAuthClient(priKey string) (authCli *AuthClient, err error) {
	// validate private key
	if !_utils.IsValidPrivateKey(priKey) {
		return nil, InvalidPrivateKey
	}
	// transfer private key str to private key
	privateKey, err := _utils.DecodePrivateKey(priKey)
	if err != nil {
		return nil, err
	}
	// get public key
	publicKey := privateKey.PublicKey
	address := crypto.PubkeyToAddress(publicKey)
	// create connection
	cli, err := GetConnection()
	if err != nil {
		return nil, err
	}
	chainID, err := cli.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	return &AuthClient{PrivateKey: privateKey, PublicKey: &publicKey, Address: address, ChainId: chainID}, nil
}
