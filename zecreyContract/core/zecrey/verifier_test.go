package zecrey

import (
	"Zecrey-eth-rpc/_const"
	"fmt"
	"testing"
	"time"
)

const ZecreyVerifierAddr = "0x7473A7322c188583b1E11Df6352680f6bA7c8866"

func TestDeployZecreyVerifierContract(t *testing.T) {
	elapse := time.Now()
	addr, txHash, err := DeployZecreyTokenContract(cli, authCli, InitialTokenSupply, gasPrice, _const.SuggestHighGasLimit)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(time.Since(elapse))
	fmt.Println("contract address:", addr)
	fmt.Println("deploy contract tx hash:", txHash)
}
