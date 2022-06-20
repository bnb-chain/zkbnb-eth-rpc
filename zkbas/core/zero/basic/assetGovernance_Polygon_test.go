package basic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bnb-chain/zkbas-eth-rpc/_const"
	"github.com/bnb-chain/zkbas-eth-rpc/zkbas/core/zero/polygon"
	"io/ioutil"
	"testing"
	"time"
)

func TestAssetGovernanceSetLister_Polygon(t *testing.T) {
	cli, authCli, err := ConstructCliAndAuthCli(polygon.ChainId, polygon.NetworkEndPoint, SuperSk)
	if err != nil {
		t.Fatal(err)
	}
	infoBytes, err := ioutil.ReadFile(polygon.Dir + ContractInfoFileName)
	if err != nil {
		t.Fatal(err)
	}
	var info *ZecreyContracts
	err = json.Unmarshal(infoBytes, &info)
	if err != nil {
		t.Fatal(err)
	}
	instance, err := LoadAssetGovernanceInstance(cli, info.AssetGovernanceAddr)
	if err != nil {
		t.Fatal(err)
	}
	governanceInstance, err := LoadGovernanceInstance(cli, info.GovernanceAddr)
	if err != nil {
		t.Fatal(err)
	}
	governor, err := governanceInstance.NetworkGovernor(EmptyCallOpts())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("governor:", governor)
	SuggestGasPrice, _ := cli.SuggestGasPrice(context.Background())
	gasPrice := SuggestGasPrice
	gasLimit := _const.SuggestHighGasLimit
	txHash, err := AssetGovernanceSetLister(
		cli, authCli, instance, SuperAddress, true, gasPrice, gasLimit,
	)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(txHash)

}

func TestDeployAssetGovernance_Polygon(t *testing.T) {
	waitTime := 10 * time.Second
	cli, authCli, err := ConstructCliAndAuthCli(polygon.ChainId, polygon.NetworkEndPoint, SuperSk)
	if err != nil {
		t.Fatal(err)
	}
	infoBytes, err := ioutil.ReadFile(polygon.Dir + ContractInfoFileName)
	if err != nil {
		t.Fatal(err)
	}
	var info *ZecreyContracts
	err = json.Unmarshal(infoBytes, &info)
	if err != nil {
		t.Fatal(err)
	}
	instance, err := LoadAssetGovernanceInstance(cli, info.AssetGovernanceAddr)
	if err != nil {
		t.Fatal(err)
	}
	SuggestGasPrice, _ := cli.SuggestGasPrice(context.Background())
	gasPrice := SuggestGasPrice
	gasLimit := _const.SuggestHighGasLimit
	// rey
	txHash, err := AssetGovernanceSetAsset(cli, authCli, instance, 0, info.ReyERC20Addr, gasPrice, gasLimit)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(waitTime)
	status, err := cli.WaitingTransactionStatus(txHash)
	if err != nil {
		t.Fatal(err)
	}
	if !status {
		t.Fatal("invalid status")
	}
	// polygon
	//txHash, err = AssetGovernanceSetAsset(cli, authCli, instance, 1, info.EthERC20Addr, gasPrice, gasLimit)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//time.Sleep(waitTime)
	//status, err = cli.WaitingTransactionStatus(txHash)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//if !status {
	//	t.Fatal("invalid status")
	//}
	// matic
	txHash, err = AssetGovernanceSetAsset(cli, authCli, instance, 2, info.MaticERC20Addr, gasPrice, gasLimit)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(waitTime)
	status, err = cli.WaitingTransactionStatus(txHash)
	if err != nil {
		t.Fatal(err)
	}
	if !status {
		t.Fatal("invalid status")
	}
	// near
	txHash, err = AssetGovernanceSetAsset(cli, authCli, instance, 3, info.NearERC20Addr, gasPrice, gasLimit)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(waitTime)
	status, err = cli.WaitingTransactionStatus(txHash)
	if err != nil {
		t.Fatal(err)
	}
	if !status {
		t.Fatal("invalid status")
	}
	// avax
	txHash, err = AssetGovernanceSetAsset(cli, authCli, instance, 4, info.AvaxERC20Addr, gasPrice, gasLimit)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(waitTime)
	status, err = cli.WaitingTransactionStatus(txHash)
	if err != nil {
		t.Fatal(err)
	}
	if !status {
		t.Fatal("invalid status")
	}
	// bit
	txHash, err = AssetGovernanceSetAsset(cli, authCli, instance, 5, info.BitERC20Addr, gasPrice, gasLimit)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(waitTime)
	status, err = cli.WaitingTransactionStatus(txHash)
	if err != nil {
		t.Fatal(err)
	}
	if !status {
		t.Fatal("invalid status")
	}
	// usdt
	txHash, err = AssetGovernanceSetAsset(cli, authCli, instance, 6, info.UsdtERC20Addr, gasPrice, gasLimit)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(waitTime)
	status, err = cli.WaitingTransactionStatus(txHash)
	if err != nil {
		t.Fatal(err)
	}
	if !status {
		t.Fatal("invalid status")
	}
	// usdc
	txHash, err = AssetGovernanceSetAsset(cli, authCli, instance, 7, info.UsdcERC20Addr, gasPrice, gasLimit)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(waitTime)
	status, err = cli.WaitingTransactionStatus(txHash)
	if err != nil {
		t.Fatal(err)
	}
	if !status {
		t.Fatal("invalid status")
	}
	// dai
	txHash, err = AssetGovernanceSetAsset(cli, authCli, instance, 8, info.DaiERC20Addr, gasPrice, gasLimit)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(waitTime)
	status, err = cli.WaitingTransactionStatus(txHash)
	if err != nil {
		t.Fatal(err)
	}
	if !status {
		t.Fatal("invalid status")
	}
	// bnb
	txHash, err = AssetGovernanceSetAsset(cli, authCli, instance, 9, info.BnbERC20Addr, gasPrice, gasLimit)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(waitTime)
	status, err = cli.WaitingTransactionStatus(txHash)
	if err != nil {
		t.Fatal(err)
	}
	if !status {
		t.Fatal("invalid status")
	}
}
