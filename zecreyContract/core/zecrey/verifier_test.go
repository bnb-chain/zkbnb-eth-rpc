package zecrey

import (
	"fmt"
	"github.com/zecrey-labs/zecrey-eth-rpc/_const"
	"math/big"
	"testing"
	"time"
)

const ZecreyVerifierAddr = "0x9E19EE8FfCB5C3810975724eb4e772144f348a65"
const RinkebyZecreyVerifierAddr = "0xFb2Bf454615C7eb6485C323Fc1a961736efb3485"
const AuroraVerifierAddr = "0x3C9191954AD100C9dE6A18B8A5691916c4F704C4"
const ArbitrumVerifierAddr = "0x3C9191954AD100C9dE6A18B8A5691916c4F704C4"
const PolyTestVerifierAddr = "0x3C9191954AD100C9dE6A18B8A5691916c4F704C4"

func TestDeployZecreyVerifierContract(t *testing.T) {
	elapse := time.Now()
	addr, txHash, err := DeployZecreyVerifierContract(polyTestCli, polyTestAuthCli, big.NewInt(1000000000), _const.SuggestHighGasLimit)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(time.Since(elapse))
	fmt.Println("contract address:", addr)
	fmt.Println("deploy contract tx hash:", txHash)
}
