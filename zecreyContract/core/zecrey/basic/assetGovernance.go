package basic

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
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
		transactOpts, *cli,
		common.HexToAddress(governanceAddr), common.HexToAddress(listingFeeAssetAddr),
		listingFee, listingCap, common.HexToAddress(treasuryAddr))
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
	AssetGovernanceSetListingFeeToken: set listing fee token
*/
func AssetGovernanceSetListingFeeToken(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *AssetGovernance,
	listingFeeTokenAddr string, listingFee *big.Int,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	if !IsValidAddress(listingFeeTokenAddr) {
		return "", errors.New("[AssetGovernanceSetListingFeeToken] invalid asset address")
	}
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.SetListingFeeToken(transactOpts, common.HexToAddress(listingFeeTokenAddr), listingFee)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	AssetGovernanceSetListingFee: set listing fee
*/
func AssetGovernanceSetListingFee(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *AssetGovernance,
	listingFee *big.Int,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.SetListingFee(transactOpts, listingFee)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	AssetGovernanceSetListingCap: set listing cap
*/
func AssetGovernanceSetListingCap(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *AssetGovernance,
	listingCap uint16,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.SetListingCap(transactOpts, listingCap)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	AssetGovernanceSetTreasury: set treasury addr
*/
func AssetGovernanceSetTreasury(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *AssetGovernance,
	treasuryAddr string,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	if !IsValidAddress(treasuryAddr) {
		return "", errors.New("[AssetGovernanceSetTreasury] invalid treasury address")
	}
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := instance.SetTreasury(transactOpts, common.HexToAddress(treasuryAddr))
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
