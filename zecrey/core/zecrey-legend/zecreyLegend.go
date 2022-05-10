package zecreyLegend

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
)

/*
	LoadZecreyLegendInstance: load zecrey legend instance if it is already deployed
*/
func LoadZecreyLegendInstance(cli *_rpc.ProviderClient, addr string) (instance *ZecreyLegend, err error) {
	instance, err = NewZecreyLegend(common.HexToAddress(addr), *cli)
	return instance, err
}
