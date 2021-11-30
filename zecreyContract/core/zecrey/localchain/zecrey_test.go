package localchain

import (
	"fmt"
	"github.com/zecrey-labs/zecrey-eth-rpc/_const"
	"testing"
)

func TestZecreyDepositOrLockNativeAsset(t *testing.T) {
	txHash, err := ZecreyDepositOrLockNativeAsset(
		0, "sher", SuggestGasPrice, _const.SuggestHighGasLimit,
	)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(txHash)
}
