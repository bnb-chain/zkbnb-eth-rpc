package localchain

import (
	"encoding/json"
	"fmt"
	"github.com/zecrey-labs/zecrey-eth-rpc/zecreyContract/core/zecrey/basic"
	"io/ioutil"
	"testing"
)

func TestFirstDeployment(t *testing.T) {
	FirstDeployment()
	addrBytes, err := ioutil.ReadFile(AddrFileName)
	if err != nil {
		t.Fatal(err)
	}
	var addrs basic.ZecreyContracts
	err = json.Unmarshal(addrBytes, &addrs)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(addrs.ZecreyAddr)
}
