package erpc

import (
	"context"
	"crypto/ecdsa"
	"eva-go-rpc/econst"
	"eva-go-rpc/eutils"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

// get connection of ethereum or rinkeby network
func GetConnection() (cli *ethclient.Client, err error) {
	return ethclient.Dial(econst.InfuraRinkebyNetwork)
}

// auth of user
type AuthClient struct {
	PrivateKey *ecdsa.PrivateKey `json:"private_key"`
	ChainId    *big.Int          `json:"chain_id"`
}

// create a new auth client
func NewAuthClient(priKey string) (authCli *AuthClient, err error) {
	// validate private key
	if !eutils.IsValidPrivateKey(priKey) {
		return nil, InvalidPrivateKey
	}
	// transfer private key str to private key
	privateKey, err := eutils.DecodePrivateKey(priKey)
	if err != nil {
		return nil, err
	}
	// create connection
	cli, err := GetConnection()
	if err != nil {
		return nil, err
	}
	chainID, err := cli.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	return &AuthClient{PrivateKey: privateKey, ChainId: chainID}, nil
}
