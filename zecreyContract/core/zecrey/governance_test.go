package zecrey

import (
	"context"
	"fmt"
	"github.com/zecrey-labs/zecrey-eth-rpc/_const"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
	"math/big"
	"testing"
	"time"
)

const GovernanceAddr = "0x0668F761a5428B4c650cCeEF5410277a262407e1"
const RinkebyGovernanceAddr = "0xdAdf3DE3B613efd0a93282DE4151Aab8062c5f77"

var (
	localCli     *_rpc.ProviderClient
	localAuthCli *_rpc.AuthClient
	rinkebyCli     *_rpc.ProviderClient
	rinkebyAuthCli *_rpc.AuthClient
	gasPrice     *big.Int
)

func init() {
	localCli, _ = _rpc.NewClient(_const.LocalNetwork)
	localAuthCli, _ = _rpc.NewAuthClient(localCli, _const.LocalSuperSk, LocalChainId)
	rinkebyCli, _ = _rpc.NewClient(_const.InfuraRinkebyNetwork)
	rinkebyAuthCli, _ = _rpc.NewAuthClient(localCli, _const.RinkebySuperSk, RinkebyChainId)
	gasPrice, _ = localCli.SuggestGasPrice(context.Background())
}

func TestDeployGovernanceContract(t *testing.T) {
	elapse := time.Now()
	addr, txHash, err := DeployGovernanceContract(localCli, localAuthCli, gasPrice, _const.SuggestHighGasLimit)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(time.Since(elapse))
	fmt.Println("contract address:", addr)
	fmt.Println("deploy contract tx hash:", txHash)
}

func LoadGovernance() *Governance {
	instance, _ := LoadGovernanceInstance(localCli, GovernanceAddr)
	return instance
}

func TestAddAsset(t *testing.T) {
	governance := LoadGovernance()
	txHash, err := AddAsset(localCli, localAuthCli, governance, RinkebyZecreyAddr, gasPrice, _const.SuggestHighGasLimit)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(txHash)
}

func TestUpdateRollupProvider(t *testing.T) {
	instance := LoadGovernance()
	txHash, err := UpdateRollupProvider(localCli, localAuthCli, instance, _const.RinkebySuperAddress, true, gasPrice, _const.SuggestHighGasLimit)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(txHash)
}

func TestGetSupportedAsset(t *testing.T) {
	instance := LoadGovernance()
	elapse := time.Now()
	assetAddr, err := GetSupportedAsset(instance, 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(time.Since(elapse))
	fmt.Println(assetAddr)
}

func TestIsValidAsset(t *testing.T) {
	instance := LoadGovernance()
	isValid, err := IsValidAsset(instance, 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(isValid)
}
