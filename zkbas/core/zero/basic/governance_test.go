package basic

import (
	"context"
	"fmt"
	"github.com/bnb-chain/zkbas-eth-rpc/_const"
	"github.com/bnb-chain/zkbas-eth-rpc/zkbas/core/zero/aurora"
	"github.com/bnb-chain/zkbas-eth-rpc/zkbas/core/zero/avalanche"
	"github.com/bnb-chain/zkbas-eth-rpc/zkbas/core/zero/bsc"
	"github.com/bnb-chain/zkbas-eth-rpc/zkbas/core/zero/eth"
	"github.com/bnb-chain/zkbas-eth-rpc/zkbas/core/zero/polygon"
	"io/fs"
	"io/ioutil"
	"testing"
)

func TestDeployGovernanceContract_Ethereum(t *testing.T) {
	cli, authCli, err := ConstructCliAndAuthCli(eth.ChainId, eth.NetworkEndPoint, SuperSk)
	if err != nil {
		t.Fatal(err)
	}
	SuggestGasPrice, _ := cli.SuggestGasPrice(context.Background())
	addr, txHash, err := DeployGovernanceContract(
		cli, authCli, SuggestGasPrice, _const.SuggestHighGasLimit,
	)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("tx hash:", txHash)
	err = ioutil.WriteFile(eth.Dir+GovernanceContractFilePrefix+eth.ChainName, []byte(addr), fs.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeployGovernanceContract_BSC(t *testing.T) {
	cli, authCli, err := ConstructCliAndAuthCli(bsc.ChainId, bsc.NetworkEndPoint, SuperSk)
	if err != nil {
		t.Fatal(err)
	}
	SuggestGasPrice, _ := cli.SuggestGasPrice(context.Background())
	addr, txHash, err := DeployGovernanceContract(
		cli, authCli, SuggestGasPrice, _const.SuggestHighGasLimit,
	)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("tx hash:", txHash)
	err = ioutil.WriteFile(bsc.Dir+GovernanceContractFilePrefix+bsc.ChainName, []byte(addr), fs.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeployGovernanceContract_Avalanche(t *testing.T) {
	cli, authCli, err := ConstructCliAndAuthCli(avalanche.ChainId, avalanche.NetworkEndPoint, SuperSk)
	if err != nil {
		t.Fatal(err)
	}
	SuggestGasPrice, _ := cli.SuggestGasPrice(context.Background())
	addr, txHash, err := DeployGovernanceContract(
		cli, authCli, SuggestGasPrice, _const.SuggestHighGasLimit,
	)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("tx hash:", txHash)
	err = ioutil.WriteFile(avalanche.Dir+GovernanceContractFilePrefix+avalanche.ChainName, []byte(addr), fs.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}
func TestDeployGovernanceContract_Polygon(t *testing.T) {
	cli, authCli, err := ConstructCliAndAuthCli(polygon.ChainId, polygon.NetworkEndPoint, SuperSk)
	if err != nil {
		t.Fatal(err)
	}
	SuggestGasPrice, _ := cli.SuggestGasPrice(context.Background())
	addr, txHash, err := DeployGovernanceContract(
		cli, authCli, SuggestGasPrice, _const.SuggestHighGasLimit,
	)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("tx hash:", txHash)
	err = ioutil.WriteFile(polygon.Dir+GovernanceContractFilePrefix+polygon.ChainName, []byte(addr), fs.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeployGovernanceContract_Aurora(t *testing.T) {
	cli, authCli, err := ConstructCliAndAuthCli(aurora.ChainId, aurora.NetworkEndPoint, SuperSk)
	if err != nil {
		t.Fatal(err)
	}
	SuggestGasPrice, _ := cli.SuggestGasPrice(context.Background())
	addr, txHash, err := DeployGovernanceContract(
		cli, authCli, SuggestGasPrice, _const.SuggestHighGasLimit,
	)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("tx hash:", txHash)
	err = ioutil.WriteFile(aurora.Dir+GovernanceContractFilePrefix+aurora.ChainName, []byte(addr), fs.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}
