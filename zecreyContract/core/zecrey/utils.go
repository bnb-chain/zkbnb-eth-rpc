package zecrey

import (
	"Zecrey-eth-rpc/_rpc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
	transactOpts.Value = big.NewInt(0)
	return transactOpts, nil
}

func SetFixed32Bytes(buf []byte) [32]byte {
	var res [32]byte
	if len(buf) != 0 {
		copy(res[:], buf[:32])
	}
	return res
}

func ConstructStorageBlock(blockNumber uint32, accountRoot []byte, commitment []byte, timestamp int64) StorageBlock {
	return StorageBlock{
		BlockNumber: blockNumber,
		AccountRoot: SetFixed32Bytes(accountRoot),
		Commitment:  SetFixed32Bytes(commitment),
		Timestamp:   big.NewInt(timestamp),
	}
}

func ConstructOperation(txType uint8, assetId uint32, nativeAddr string, zecreyAddr string, amount *big.Int) ZecreyOperation {
	return ZecreyOperation{
		TxType:     txType,
		AssetId:    assetId,
		NativeAddr: common.HexToAddress(nativeAddr),
		ZecreyAddr: zecreyAddr,
		Amount:     amount,
	}
}

func ConstructCommitBlock(blockNumber uint32, accountRoot []byte, publicData []byte, withdrawOperations []ZecreyOperation, timestamp int64) ZecreyCommitBlock {
	return ZecreyCommitBlock{
		BlockNumber:        blockNumber,
		NewAccountRoot:     SetFixed32Bytes(accountRoot),
		PublicData:         publicData,
		WithdrawOperations: withdrawOperations,
		Timestamp:          big.NewInt(timestamp),
	}
}

func ConstructProcessBlock(blockInfo StorageBlock, onchainOperationsIndexes []*big.Int, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int) ZecreyProcessBlock {
	return ZecreyProcessBlock{
		BlockInfo:                blockInfo,
		OnchainOperationsIndexes: onchainOperationsIndexes,
		A:                        a,
		B:                        b,
		C:                        c,
	}
}
