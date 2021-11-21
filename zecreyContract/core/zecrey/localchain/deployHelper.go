package localchain

import (
	"encoding/json"
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
		return
	}
	addrBytes, err := json.Marshal(addrs)
	if err != nil {
		log.Println("err:", err)
		return
	}
	err = ioutil.WriteFile(AddrFileName, addrBytes, 666)
	if err != nil {
		panic(err)
	}
}
