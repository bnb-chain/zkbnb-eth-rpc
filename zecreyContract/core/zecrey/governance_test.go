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
const AuroraTestNetGovernanceAddr = "0x1142982cd9f7d292Ff1B94Ae33c53087D5Cec3E4"
const AribitrumTestNetGovernanceAddr = "0x1142982cd9f7d292Ff1B94Ae33c53087D5Cec3E4"
const PolyTestNetGovernanceAddr = "0x1142982cd9f7d292Ff1B94Ae33c53087D5Cec3E4"

var (
	localCli         *_rpc.ProviderClient
	localAuthCli     *_rpc.AuthClient
	rinkebyCli       *_rpc.ProviderClient
	rinkebyAuthCli   *_rpc.AuthClient
	auroraCli        *_rpc.ProviderClient
	auroraAuthCli    *_rpc.AuthClient
	arbitrumCli      *_rpc.ProviderClient
	arbitrumAuthCli  *_rpc.AuthClient
	polyTestCli      *_rpc.ProviderClient
	polyTestAuthCli  *_rpc.AuthClient
	localGasPrice    *big.Int
	auroraGasPrice   *big.Int
	arbitrumGasPrice *big.Int
	polyGasPrice     *big.Int
)

func init() {
	localCli, _ = _rpc.NewClient(_const.LocalNetwork)
	localAuthCli, _ = _rpc.NewAuthClient(localCli, _const.LocalSuperSk, LocalChainId)
	rinkebyCli, _ = _rpc.NewClient(_const.InfuraRinkebyNetwork)
	rinkebyAuthCli, _ = _rpc.NewAuthClient(rinkebyCli, _const.RinkebySuperSk, RinkebyChainId)
	auroraCli, _ = _rpc.NewClient(_const.AuroraTestNetwork)
	auroraAuthCli, _ = _rpc.NewAuthClient(auroraCli, _const.RinkebySuperSk, AuroraChainId)
	arbitrumCli, _ = _rpc.NewClient(_const.ArbitrumTestNetwork)
	arbitrumAuthCli, _ = _rpc.NewAuthClient(arbitrumCli, _const.RinkebySuperSk, ArbitrumChainId)
	polyTestCli, _ = _rpc.NewClient(_const.PolyTestNetwork)
	polyTestAuthCli, _ = _rpc.NewAuthClient(polyTestCli, _const.RinkebySuperSk, PolyChainId)
	localGasPrice, _ = localCli.SuggestGasPrice(context.Background())
	auroraGasPrice, _ = auroraCli.SuggestGasPrice(context.Background())
	arbitrumGasPrice, _ = arbitrumCli.SuggestGasPrice(context.Background())
	polyGasPrice, _ = polyTestCli.SuggestGasPrice(context.Background())
}

func TestDeployGovernanceContract(t *testing.T) {
	elapse := time.Now()
	//addr, txHash, err := DeployGovernanceContract(auroraCli, auroraAuthCli, localGasPrice, _const.SuggestHighGasLimit)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//addr, txHash, err := DeployGovernanceContract(localCli, localAuthCli, localGasPrice, _const.SuggestHighGasLimit)
	//if err != nil {
	//	t.Fatal(err)
	//}
	addr, txHash, err := DeployGovernanceContract(rinkebyCli, rinkebyAuthCli, localGasPrice, _const.SuggestHighGasLimit)
	if err != nil {
		t.Fatal(err)
	}
	//addr, txHash, err := DeployGovernanceContract(arbitrumCli, arbitrumAuthCli, big.NewInt(20000000000000), _const.SuggestHighGasLimit)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//addr, txHash, err := DeployGovernanceContract(polyTestCli, polyTestAuthCli, big.NewInt(1000000000), _const.SuggestHighGasLimit)
	//if err != nil {
	//	t.Fatal(err)
	//}
	fmt.Println(time.Since(elapse))
	fmt.Println("contract address:", addr)
	fmt.Println("deploy contract tx hash:", txHash)
}

func LoadGovernance() *Governance {
	instance, _ := LoadGovernanceInstance(rinkebyCli, RinkebyGovernanceAddr)
	return instance
}

func TestAddAsset(t *testing.T) {
	governance := LoadGovernance()
	txHash, err := AddAsset(rinkebyCli, rinkebyAuthCli, governance, RinkebyZecreyTokenAddr, localGasPrice, _const.SuggestHighGasLimit)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(txHash)
}

func TestUpdateRollupProvider(t *testing.T) {
	instance := LoadGovernance()
	txHash, err := UpdateRollupProvider(localCli, localAuthCli, instance, _const.LocalSuperAddress, true, localGasPrice, _const.SuggestHighGasLimit)
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
