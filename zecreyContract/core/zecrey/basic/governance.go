package basic

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
	"math/big"
)

/*
	DeployGovernanceContract: deploy governance
*/
func DeployGovernanceContract(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient,
	gasPrice *big.Int, gasLimit uint64,
) (addr string, txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", "", err
	}
	address, tx, _, err := DeployGovernance(transactOpts, *cli)
	if err != nil {
		return "", "", err
	}
	return address.Hex(), tx.Hash().Hex(), nil
}

/*
	LoadGovernanceInstance: load governance instance if it is already deployed
*/
func LoadGovernanceInstance(cli *_rpc.ProviderClient, addr string) (instance *Governance, err error) {
	instance, err = NewGovernance(common.HexToAddress(addr), *cli)
	return instance, err
}

func PackInitializeGovernanceParams(owner string) ([]byte, error) {
	arguments := abi.Arguments{
		{Type: AddressType},
	}
	hashBytes, err := arguments.Pack(
		common.HexToAddress(owner),
	)
	if err != nil {
		return nil, err
	}
	return hashBytes, nil
}

/*
	GovernanceInitialize: initialize governance contract, set governor
*/
func GovernanceInitialize(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Governance,
	owner string,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	if !IsValidAddress(owner) {
		return "", errors.New("[GovernanceInitialize] invalid address")
	}
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// pack params
	params, err := PackInitializeGovernanceParams(owner)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.Initialize(transactOpts, params)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	GovernanceChangeGovernor: change governor
*/
func GovernanceChangeGovernor(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Governance,
	owner string,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	if !IsValidAddress(owner) {
		return "", errors.New("[GovernanceChangeGovernor] invalid address")
	}
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.ChangeGovernor(transactOpts, common.HexToAddress(owner))
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	GovernanceChangeAssetGovernance: change asset governance
*/
func GovernanceChangeAssetGovernance(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Governance,
	assetGovernanceAddr string,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	if !IsValidAddress(assetGovernanceAddr) {
		return "", errors.New("[GovernanceChangeAssetGovernance] invalid address")
	}
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.ChangeAssetGovernance(transactOpts, common.HexToAddress(assetGovernanceAddr))
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	GovernanceSetCommitter: set committer
*/
func GovernanceSetCommitter(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Governance,
	committer string, isActive bool,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	if !IsValidAddress(committer) {
		return "", errors.New("[GovernanceSetCommitter] invalid address")
	}
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.SetCommitter(transactOpts, common.HexToAddress(committer), isActive)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	GovernanceSetVerifier: set verifier
*/
func GovernanceSetVerifier(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Governance,
	verifier string, isActive bool,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	if !IsValidAddress(verifier) {
		return "", errors.New("[GovernanceSetVerifier] invalid address")
	}
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.SetVerifier(transactOpts, common.HexToAddress(verifier), isActive)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	GovernanceSetMonitor: set monitor
*/
func GovernanceSetMonitor(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Governance,
	monitor string, isActive bool,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	if !IsValidAddress(monitor) {
		return "", errors.New("[GovernanceSetMonitor] invalid address")
	}
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.SetMonitor(transactOpts, common.HexToAddress(monitor), isActive)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	GovernanceValidateAssetAddress: get related asset id of asset address
*/
func GovernanceValidateAssetAddress(
	instance *Governance,
	assetAddr string,
) (uint16, error) {
	if !IsValidAddress(assetAddr) {
		return 0, errors.New("[GovernanceValidateAssetAddress] invalid address")
	}
	assetId, err := instance.ValidateAssetAddress(EmptyCallOpts(), common.HexToAddress(assetAddr))
	return assetId, err
}
