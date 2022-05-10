package _rpc

import (
	"github.com/zecrey-labs/zecrey-eth-rpc/_utils"
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

// get balance
func (cli *ProviderClient) GetBalance(address string) (balance *big.Int, err error) {
	// check address
	isValid := _utils.IsValidEthAddress(address)
	if !isValid {
		return nil, ErrInvalidAddress
	}
	// get address
	account := common.HexToAddress(address)
	// get balance
	return cli.BalanceAt(context.Background(), account, nil)
}

// check if the address is a zecrey address
// @address: address of ethereum
func (cli *ProviderClient) IsContract(address string) (isContract bool, err error) {
	isValidEthAddress := _utils.IsValidEthAddress(address)
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

// get latest block header info
func (cli *ProviderClient) GetLatestBlockHeader() (header *types.Header, err error) {
	return cli.GetBlockHeaderByHash("")
}

// get block header by hash
func (cli *ProviderClient) GetBlockHeaderByHash(blockHash string) (header *types.Header, err error) {
	var hash common.Hash
	if blockHash != "" {
		hash = common.HexToHash(blockHash)
	}
	return cli.HeaderByHash(context.Background(), hash)
}

// get latest block info by hash
func (cli *ProviderClient) GetLatestBlockInfo() (blockInfo *types.Block, err error) {
	return cli.GetBlockInfoByHash("")
}

// get block info by hash
func (cli *ProviderClient) GetBlockInfoByHash(blockHash string) (blockInfo *types.Block, err error) {
	// transfer to hash
	var hash common.Hash
	if blockHash != "" {
		hash = common.HexToHash(blockHash)
	}
	return cli.BlockByHash(context.Background(), hash)
}

// get block header by number
func (cli *ProviderClient) GetBlockHeaderByNumber(height *big.Int) (header *types.Header, err error) {
	return cli.HeaderByNumber(context.Background(), height)
}

// get block info by number
func (cli *ProviderClient) GetBlockInfoByNumber(height *big.Int) (blockInfo *types.Block, err error) {
	// get block info
	return cli.BlockByNumber(context.Background(), height)
}

// get block height
func (cli *ProviderClient) GetHeight() (height uint64, err error) {
	// get height
	return cli.BlockNumber(context.Background())
}

// get transaction info by hash
func (cli *ProviderClient) GetTransactionByHash(transactionHash string) (tx *types.Transaction, isPending bool, err error) {
	// validate hash value
	if !_utils.IsValidHashValue(transactionHash) {
		return nil, false, ErrInvalidHashValue
	}
	// get hash value
	hash := common.HexToHash(transactionHash)
	return cli.TransactionByHash(context.Background(), hash)
}

// get transaction receipt
func (cli *ProviderClient) GetTransactionReceipt(transactionHash string) (receipt *types.Receipt, err error) {
	// validate hash value
	if !_utils.IsValidHashValue(transactionHash) {
		return nil, ErrInvalidHashValue
	}
	// get hash value
	hash := common.HexToHash(transactionHash)
	return cli.TransactionReceipt(context.Background(), hash)
}

func (cli *ProviderClient) GetPendingNonce(address string) (nonce uint64, err error) {
	// validate address
	isValidEthAddress := _utils.IsValidEthAddress(address)
	if !isValidEthAddress {
		return 0, ErrInvalidAddress
	}
	// transfer to address
	account := common.HexToAddress(address)
	// get nonce
	return cli.PendingNonceAt(context.Background(), account)
}

// transfer private key to address
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
