package contract

import (
	"eva-go-rpc/_const"
	"fmt"
	"testing"
)

func TestGetErc20Balance(t *testing.T) {
	balance, err := GetErc20EtherBalance(_const.SuperAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("eva token balance:", balance)
}
