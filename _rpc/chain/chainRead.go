package chain

import (
	"context"
	"crypto/ecdsa"
	"eva-go-rpc/_rpc"
	"eva-go-rpc/_utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

// get balance
func GetBalance(address string) (balance *big.Int, err error) {
	// check address
	isValid := _utils.IsValidEthAddress(address)
	if !isValid {
		return nil, _rpc.InvalidAddress
	}
	// get address
	account := common.HexToAddress(address)
	// get connection
	cli, err := _rpc.GetConnection()
	if err != nil {
		return nil, err
	}
	defer cli.Close()
	// get balance
	return cli.BalanceAt(context.Background(), account, nil)
}

// get balacne by ether
func GetEtherBalance(address string) (etherBalance *big.Float, err error) {
	balance, err := GetBalance(address)
	if err != nil {
		return nil, err
	}
	etherBalance = _utils.WeiToEther(balance)
	return etherBalance, nil
}

// check if the address is a contract address
// @address: address of ethereum
func IsContract(address string) (isContract bool, err error) {
	isValidEthAddress := _utils.IsValidEthAddress(address)
	if !isValidEthAddress {
		return false, _rpc.InvalidAddress
	}
	cli, err := _rpc.GetConnection()
	if err != nil {
		return false, err
	}
	defer cli.Close()
	contract := common.HexToAddress(address)
	bytecode, err := cli.CodeAt(context.Background(), contract, nil)
	if err != nil {
		return false, err
	}
	isContract = len(bytecode) > 0
	return isContract, nil
}

// get latest block header info
func GetLatestBlockHeader() (header *types.Header, err error) {
	return GetBlockHeaderByHash("")
}

// get block header by hash
func GetBlockHeaderByHash(blockHash string) (header *types.Header, err error) {
	cli, err := _rpc.GetConnection()
	if err != nil {
		return nil, err
	}
	defer cli.Close()
	var hash common.Hash
	if blockHash != "" {
		hash = common.HexToHash(blockHash)
	}
	return cli.HeaderByHash(context.Background(), hash)
}

// get latest block info by hash
func GetLatestBlockInfo() (blockInfo *types.Block, err error) {
	return GetBlockInfoByHash("")
}

// get block info by hash
func GetBlockInfoByHash(blockHash string) (blockInfo *types.Block, err error) {
	cli, err := _rpc.GetConnection()
	if err != nil {
		return nil, err
	}
	defer cli.Close()
	// transfer to hash
	var hash common.Hash
	if blockHash != "" {
		hash = common.HexToHash(blockHash)
	}
	return cli.BlockByHash(context.Background(), hash)
}

// get block header by number
func GetBlockHeaderByNumber(height *big.Int) (header *types.Header, err error) {
	cli, err := _rpc.GetConnection()
	if err != nil {
		return nil, err
	}
	defer cli.Close()
	return cli.HeaderByNumber(context.Background(), height)
}

// get block info by number
func GetBlockInfoByNumber(height *big.Int) (blockInfo *types.Block, err error) {
	cli, err := _rpc.GetConnection()
	if err != nil {
		return nil, err
	}
	defer cli.Close()
	// get block info
	return cli.BlockByNumber(context.Background(), height)
}

// get block height
func GetHeight() (height uint64, err error) {
	cli, err := _rpc.GetConnection()
	if err != nil {
		return 0, err
	}
	defer cli.Close()
	// get height
	return cli.BlockNumber(context.Background())
}

// get transaction info by hash
func GetTransactionByHash(transactionHash string) (tx *types.Transaction, isPending bool, err error) {
	// validate hash value
	if !_utils.IsValidHashValue(transactionHash) {
		return nil, false, _rpc.InvalidHashValue
	}
	cli, err := _rpc.GetConnection()
	if err != nil {
		return nil, false, err
	}
	defer cli.Close()
	// get hash value
	hash := common.HexToHash(transactionHash)
	return cli.TransactionByHash(context.Background(), hash)
}

// get transaction receipt
func GetTransactionReceipt(transactionHash string) (receipt *types.Receipt, err error) {
	// validate hash value
	if !_utils.IsValidHashValue(transactionHash) {
		return nil, _rpc.InvalidHashValue
	}
	cli, err := _rpc.GetConnection()
	if err != nil {
		return nil, err
	}
	defer cli.Close()
	// get hash value
	hash := common.HexToHash(transactionHash)
	return cli.TransactionReceipt(context.Background(), hash)
}

func GetPendingNonce(address string) (nonce uint64, err error) {
	// validate address
	isValidEthAddress := _utils.IsValidEthAddress(address)
	if !isValidEthAddress {
		return 0, _rpc.InvalidAddress
	}
	// transfer to address
	account := common.HexToAddress(address)
	// get connection
	cli, err := _rpc.GetConnection()
	if err != nil {
		return 0, err
	}
	defer cli.Close()
	// get nonce
	return cli.PendingNonceAt(context.Background(), account)
}

// transfer private key to address
func PrivateKeyToAddress(privateKey *ecdsa.PrivateKey) (address common.Address, err error) {
	pubKey := privateKey.Public()
	publicKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, _rpc.InvalidPrivateKey
	}
	// get address
	address = crypto.PubkeyToAddress(*publicKeyECDSA)
	return address, nil
}
