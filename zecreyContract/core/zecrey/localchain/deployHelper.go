package localchain

import (
	"bytes"
	"fmt"
	"github.com/zecrey-labs/zecrey-eth-rpc/_const"
	"github.com/zecrey-labs/zecrey-eth-rpc/zecreyContract/core/zecrey/basic"
	"io/ioutil"
	"log"
	"math/big"
)

func FirstDeployment() {
	addrs, err := basic.FirstDeployment(
		Cli, AuthCli,
		L2ChainId, 1, 50,
		"0x5ee4C56aab8F8A1E3dC524A890E06D2f0c9073cE", big.NewInt(100), 100, "0x5ee4C56aab8F8A1E3dC524A890E06D2f0c9073cE",
		SuggestGasPrice, _const.SuggestHighGasLimit)
	if err != nil {
		log.Println("err:", err)
	}
	var buf bytes.Buffer
	for i, addr := range addrs {
		buf.WriteString(addr)
		if i != len(addrs)-1 {
			buf.WriteString("\n")
		}
	}
	err = ioutil.WriteFile("contractsAddresses.txt", buf.Bytes(), 666)
	if err != nil {
		panic(err)
	}
	fmt.Println(addrs)
}
