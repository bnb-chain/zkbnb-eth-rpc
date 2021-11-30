package localchain

import (
	"github.com/zecrey-labs/zecrey-eth-rpc/zecreyContract/core/zecrey/basic"
	"math/big"
)

func ZecreyDepositOrLockNativeAsset(
	opType uint8, accountName string,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	return basic.ZecreyDepositOrLockNativeAsset(
		Cli, AuthCli, ZecreyInstance,
		opType, accountName,
		gasPrice, gasLimit,
	)
}
