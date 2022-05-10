package basic

import (
	"context"
	"fmt"
	"github.com/zecrey-labs/zecrey-eth-rpc/_const"
	"github.com/zecrey-labs/zecrey-eth-rpc/zecrey/core/zecrey/aurora"
	"github.com/zecrey-labs/zecrey-eth-rpc/zecrey/core/zecrey/avalanche"
	"github.com/zecrey-labs/zecrey-eth-rpc/zecrey/core/zecrey/bsc"
	"github.com/zecrey-labs/zecrey-eth-rpc/zecrey/core/zecrey/eth"
	"github.com/zecrey-labs/zecrey-eth-rpc/zecrey/core/zecrey/polygon"
	"io/ioutil"
	"testing"
)

func TestInitializeGovernanceContract_Ethereum(t *testing.T) {
	cli, authCli, err := ConstructCliAndAuthCli(eth.ChainId, eth.NetworkEndPoint, SuperSk)
	if err != nil {
		t.Fatal(err)
	}
	SuggestGasPrice, _ := cli.SuggestGasPrice(context.Background())
	addrBytes, err := ioutil.ReadFile(eth.Dir + GovernanceContractFilePrefix + eth.ChainName)
	if err != nil {
		t.Fatal(err)
	}
	instance, err := LoadGovernanceInstance(cli, string(addrBytes))
	owner := SuperAddress
	txHash, err := GovernanceInitialize(
		cli, authCli, instance,
		owner,
		SuggestGasPrice, _const.SuggestHighGasLimit,
	)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(txHash)
}

func TestInitializeGovernanceContract_Aurora(t *testing.T) {
	cli, authCli, err := ConstructCliAndAuthCli(aurora.ChainId, aurora.NetworkEndPoint, SuperSk)
	if err != nil {
		t.Fatal(err)
	}
	SuggestGasPrice, _ := cli.SuggestGasPrice(context.Background())
	addrBytes, err := ioutil.ReadFile(aurora.Dir + GovernanceContractFilePrefix + aurora.ChainName)
	if err != nil {
		t.Fatal(err)
	}
	instance, err := LoadGovernanceInstance(cli, string(addrBytes))
	owner := SuperAddress
	txHash, err := GovernanceInitialize(
		cli, authCli, instance,
		owner,
		SuggestGasPrice, _const.SuggestHighGasLimit,
	)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(txHash)
}

func TestInitializeGovernanceContract_Avalanche(t *testing.T) {
	cli, authCli, err := ConstructCliAndAuthCli(avalanche.ChainId, avalanche.NetworkEndPoint, SuperSk)
	if err != nil {
		t.Fatal(err)
	}
	SuggestGasPrice, _ := cli.SuggestGasPrice(context.Background())
	addrBytes, err := ioutil.ReadFile(avalanche.Dir + GovernanceContractFilePrefix + avalanche.ChainName)
	if err != nil {
		t.Fatal(err)
	}
	instance, err := LoadGovernanceInstance(cli, string(addrBytes))
	owner := SuperAddress
	txHash, err := GovernanceInitialize(
		cli, authCli, instance,
		owner,
		SuggestGasPrice, _const.SuggestHighGasLimit,
	)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(txHash)
}

func TestInitializeGovernanceContract_BSC(t *testing.T) {
	cli, authCli, err := ConstructCliAndAuthCli(bsc.ChainId, bsc.NetworkEndPoint, SuperSk)
	if err != nil {
		t.Fatal(err)
	}
	SuggestGasPrice, _ := cli.SuggestGasPrice(context.Background())
	addrBytes, err := ioutil.ReadFile(bsc.Dir + GovernanceContractFilePrefix + bsc.ChainName)
	if err != nil {
		t.Fatal(err)
	}
	instance, err := LoadGovernanceInstance(cli, string(addrBytes))
	owner := SuperAddress
	txHash, err := GovernanceInitialize(
		cli, authCli, instance,
		owner,
		SuggestGasPrice, _const.SuggestHighGasLimit,
	)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(txHash)
}

func TestInitializeGovernanceContract_Polygon(t *testing.T) {
	cli, authCli, err := ConstructCliAndAuthCli(polygon.ChainId, polygon.NetworkEndPoint, SuperSk)
	if err != nil {
		t.Fatal(err)
	}
	SuggestGasPrice, _ := cli.SuggestGasPrice(context.Background())
	addrBytes, err := ioutil.ReadFile(polygon.Dir + GovernanceContractFilePrefix + polygon.ChainName)
	if err != nil {
		t.Fatal(err)
	}
	instance, err := LoadGovernanceInstance(cli, string(addrBytes))
	owner := SuperAddress
	txHash, err := GovernanceInitialize(
		cli, authCli, instance,
		owner,
		SuggestGasPrice, _const.SuggestHighGasLimit,
	)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(txHash)
}
