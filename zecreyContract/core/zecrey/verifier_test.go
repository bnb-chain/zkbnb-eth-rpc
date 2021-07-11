package zecrey

import (
	"Zecrey-eth-rpc/_const"
	"fmt"
	"testing"
	"time"
)

const ZecreyVerifierAddr = "0x9E19EE8FfCB5C3810975724eb4e772144f348a65"
const RinkebyZecreyVerifierAddr = "0xFb2Bf454615C7eb6485C323Fc1a961736efb3485"

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
