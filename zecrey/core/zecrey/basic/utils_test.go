package basic

import (
	"encoding/hex"
	"fmt"
	"github.com/zecrey-labs/zecrey-eth-rpc/zecrey/core/zecrey/aurora"
	"github.com/zecrey-labs/zecrey-eth-rpc/zecrey/core/zecrey/avalanche"
	"github.com/zecrey-labs/zecrey-eth-rpc/zecrey/core/zecrey/bsc"
	"github.com/zecrey-labs/zecrey-eth-rpc/zecrey/core/zecrey/eth"
	"github.com/zecrey-labs/zecrey-eth-rpc/zecrey/core/zecrey/polygon"
	"math/big"
	"testing"
)

func TestSetFixed32Bytes(t *testing.T) {
	buf := SetFixed32Bytes([]byte("pool"))
	name := hex.EncodeToString(buf[:])
	fmt.Println(name)
	nameBytes, err := hex.DecodeString(name)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(nameBytes)
	fmt.Println(new(big.Int).SetBytes(buf[:]).String())
}

func TestCheckTxStatus_Ethereum(t *testing.T) {
	cli, _, err := ConstructCliAndAuthCli(eth.ChainId, eth.NetworkEndPoint, SuperSk)
	if err != nil {
		t.Fatal(err)
	}
	txHash := ""
	status, err := cli.WaitingTransactionStatus(txHash)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(status)
}

func TestCheckTxStatus_BSC(t *testing.T) {
	cli, _, err := ConstructCliAndAuthCli(bsc.ChainId, bsc.NetworkEndPoint, SuperSk)
	if err != nil {
		t.Fatal(err)
	}
	txHash := ""
	status, err := cli.WaitingTransactionStatus(txHash)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(status)
}

func TestCheckTxStatus_Avalanche(t *testing.T) {
	cli, _, err := ConstructCliAndAuthCli(avalanche.ChainId, avalanche.NetworkEndPoint, SuperSk)
	if err != nil {
		t.Fatal(err)
	}
	txHash := ""
	status, err := cli.WaitingTransactionStatus(txHash)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(status)
}

func TestCheckTxStatus_Aurora(t *testing.T) {
	cli, _, err := ConstructCliAndAuthCli(aurora.ChainId, aurora.NetworkEndPoint, SuperSk)
	if err != nil {
		t.Fatal(err)
	}
	txHash := ""
	status, err := cli.WaitingTransactionStatus(txHash)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(status)
}

func TestCheckTxStatus_Polygon(t *testing.T) {
	cli, _, err := ConstructCliAndAuthCli(polygon.ChainId, polygon.NetworkEndPoint, SuperSk)
	if err != nil {
		t.Fatal(err)
	}
	txHash := ""
	status, err := cli.WaitingTransactionStatus(txHash)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(status)
}
