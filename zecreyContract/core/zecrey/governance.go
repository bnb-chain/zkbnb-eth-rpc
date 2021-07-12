package zecrey

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
	"math/big"
)

/*
	DeployGovernanceContract: deploy governance
*/
func DeployGovernanceContract(cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, gasPrice *big.Int, gasLimit uint64) (addr string, txHash string, err error) {
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

/*
	AddAsset: add supported assets
*/
func AddAsset(cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Governance, assetAddressStr string, gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	assetAddress := common.HexToAddress(assetAddressStr)
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	tx, err := instance.AddAsset(transactOpts, assetAddress)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

/*
	UpdateRollupProvider: update operator status
*/
func UpdateRollupProvider(cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Governance, providerAddr string, status bool, gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	provider := common.HexToAddress(providerAddr)
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	tx, err := instance.UpdateRollupProvider(transactOpts, provider, status)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

/*
	UpdateAssetStatus: update asset status
*/
func UpdateAssetStatus(cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Governance, assetId uint32, status bool, gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	tx, err := instance.UpdateAssetStatus(transactOpts, assetId, status)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

/*
	GetSupportedAsset: get assetAddr by assetId
*/
func GetSupportedAsset(instance *Governance, assetId uint32) (addr string, err error) {
	assetAddr, err := instance.GetSupportedAsset(EmptyCallOpts(), assetId)
	if err != nil {
		return "", err
	}
	return assetAddr.Hex(), err
}

/*
	IsValidAsset: check if the asset is valid or not
*/
func IsValidAsset(instance *Governance, assetId uint32) (bool, error) {
	isValidAsset, err := instance.IsValidAsset(EmptyCallOpts(), assetId)
	return isValidAsset, err
}

/*
	IsValidOperator: check if the operator is valid or not
*/
func IsValidOperator(instance *Governance, operatorAddr string) (bool, error) {
	operator := common.HexToAddress(operatorAddr)
	isValidOperator, err := instance.IsValidOperator(EmptyCallOpts(), operator)
	return isValidOperator, err
}
