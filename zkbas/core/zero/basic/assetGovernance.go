package basic

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/bnb-chain/zkbas-eth-rpc/_rpc"
	"log"
	"math/big"
)

/*
	DeployGovernanceContract: deploy governance
*/
func DeployAssetGovernanceContract(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient,
	governanceAddr string, listingFeeAssetAddr string, listingFee *big.Int, listingCap uint16, treasuryAddr string,
	gasPrice *big.Int, gasLimit uint64,
) (addr string, txHash string, err error) {
	if !IsValidAddress(governanceAddr) || !IsValidAddress(listingFeeAssetAddr) || !IsValidAddress(treasuryAddr) {
		log.Println("[DeployAssetGovernanceContract] invalid address")
		return "", "", errors.New("[DeployAssetGovernanceContract] invalid address")
	}
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", "", err
	}
	address, tx, _, err := DeployAssetGovernance(
		transactOpts, *cli)
	if err != nil {
		return "", "", err
	}
	return address.Hex(), tx.Hash().Hex(), nil
}

/*
	LoadGovernanceInstance: load governance instance if it is already deployed
*/
func LoadAssetGovernanceInstance(cli *_rpc.ProviderClient, addr string) (instance *AssetGovernance, err error) {
	instance, err = NewAssetGovernance(common.HexToAddress(addr), *cli)
	return instance, err
}

/*
	AssetGovernanceSetAsset: set asset
*/
func AssetGovernanceSetAsset(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *AssetGovernance,
	assetId uint16, assetAddr string,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	if !IsValidAddress(assetAddr) {
		return "", errors.New("[AssetGovernanceSetAsset] invalid asset address")
	}
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.SetAsset(transactOpts, assetId, common.HexToAddress(assetAddr))
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	AssetGovernanceSetLister: set lister
*/
func AssetGovernanceSetLister(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *AssetGovernance,
	listerAddr string, isActive bool,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	if !IsValidAddress(listerAddr) {
		return "", errors.New("[AssetGovernanceSetTreasury] invalid treasury address")
	}
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.SetLister(transactOpts, common.HexToAddress(listerAddr), isActive)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}
