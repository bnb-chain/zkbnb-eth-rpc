package rpc

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/bnb-chain/zkbnb-eth-rpc/utils"
)

func (cli *ProviderClient) GetBalance(address string) (balance *big.Int, err error) {
	// check address
	isValid := utils.IsValidEthAddress(address)
	if !isValid {
		return nil, ErrInvalidAddress
	}
	// get address
	account := common.HexToAddress(address)
	// get balance
	return cli.BalanceAt(context.Background(), account, nil)
}

func (cli *ProviderClient) IsContract(address string) (isContract bool, err error) {
	isValidEthAddress := utils.IsValidEthAddress(address)
	if !isValidEthAddress {
		return false, ErrInvalidAddress
	}
	contract := common.HexToAddress(address)
	bytecode, err := cli.CodeAt(context.Background(), contract, nil)
	if err != nil {
		return false, err
	}
	isContract = len(bytecode) > 0
	return isContract, nil
}

func (cli *ProviderClient) GetLatestBlockHeader() (header *types.Header, err error) {
	return cli.GetBlockHeaderByHash("")
}

func (cli *ProviderClient) GetBlockHeaderByHash(blockHash string) (header *types.Header, err error) {
	var hash common.Hash
	if blockHash != "" {
		hash = common.HexToHash(blockHash)
	}
	return cli.HeaderByHash(context.Background(), hash)
}

func (cli *ProviderClient) GetLatestBlockInfo() (blockInfo *types.Block, err error) {
	return cli.GetBlockInfoByHash("")
}

func (cli *ProviderClient) GetBlockInfoByHash(blockHash string) (blockInfo *types.Block, err error) {
	// transfer to hash
	var hash common.Hash
	if blockHash != "" {
		hash = common.HexToHash(blockHash)
	}
	return cli.BlockByHash(context.Background(), hash)
}

func (cli *ProviderClient) GetBlockHeaderByNumber(height *big.Int) (header *types.Header, err error) {
	return cli.HeaderByNumber(context.Background(), height)
}

func (cli *ProviderClient) GetBlockInfoByNumber(height *big.Int) (blockInfo *types.Block, err error) {
	// get block info
	return cli.BlockByNumber(context.Background(), height)
}

func (cli *ProviderClient) GetHeight() (height uint64, err error) {
	// get height
	return cli.BlockNumber(context.Background())
}

func (cli *ProviderClient) GetTransactionByHash(transactionHash string) (tx *types.Transaction, isPending bool, err error) {
	// validate hash value
	if !utils.IsValidHashValue(transactionHash) {
		return nil, false, ErrInvalidHashValue
	}
	// get hash value
	hash := common.HexToHash(transactionHash)
	return cli.TransactionByHash(context.Background(), hash)
}

func (cli *ProviderClient) GetTransactionReceipt(transactionHash string) (receipt *types.Receipt, err error) {
	// validate hash value
	if !utils.IsValidHashValue(transactionHash) {
		return nil, ErrInvalidHashValue
	}
	// get hash value
	hash := common.HexToHash(transactionHash)
	return cli.TransactionReceipt(context.Background(), hash)
}

func (cli *ProviderClient) GetPendingNonce(address string) (nonce uint64, err error) {
	// validate address
	isValidEthAddress := utils.IsValidEthAddress(address)
	if !isValidEthAddress {
		return 0, ErrInvalidAddress
	}
	// transfer to address
	account := common.HexToAddress(address)
	// get nonce
	return cli.PendingNonceAt(context.Background(), account)
}

func PrivateKeyToAddress(privateKey *ecdsa.PrivateKey) (address common.Address, err error) {
	pubKey := privateKey.Public()
	publicKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, ErrInvalidPrivateKey
	}
	// get address
	address = crypto.PubkeyToAddress(*publicKeyECDSA)
	return address, nil
}
