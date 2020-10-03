package contract

import (
	"eva-go-rpc/econst"
	"fmt"
	"testing"
)

func TestGetErc20Balance(t *testing.T) {
	balance, err := GetErc20EtherBalance(econst.SuperAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("eva token balance:", balance)
}
