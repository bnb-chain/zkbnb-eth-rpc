package basic

import (
	"errors"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
	"math/big"
)

/*
	FirstDeployment: helper method to deploy all contracts
*/
func FirstDeployment(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient,
	governor string,
	l2ChainId uint8, nativeAssetId uint16, maxPendingBlocks uint16,
	listingFeeTokenAddr string, listingFee *big.Int, listingCap uint16, treasuryAddr string,
	gasPrice *big.Int, gasLimit uint64,
) (addrs *ZecreyContracts, err error) {
	// deploy governance contract
	governanceAddr, txHash, err := DeployGovernanceContract(cli, authCli, l2ChainId, nativeAssetId, maxPendingBlocks, gasPrice, gasLimit)
	if err != nil {
		return nil, err
	}
	// wait status of deployment
	status, err := cli.WaitingTransactionStatus(txHash)
	if err != nil {
		return nil, err
	}
	if !status {
		return nil, errors.New("[FirstDeployment] invalid tx, cannot get status")
	}
	// load instance from address
	governanceInstance, err := LoadGovernanceInstance(cli, governanceAddr)
	if err != nil {
		return nil, err
	}
	// initialize governance contract
	txHash, err = GovernanceInitialize(cli, authCli, governanceInstance, governor, gasPrice, gasLimit)
	if err != nil {
		return nil, err
	}
	// deploy asset governance contract
	assetGovernanceAddr, txHash, err := DeployAssetGovernanceContract(cli, authCli, governanceAddr, listingFeeTokenAddr, listingFee, listingCap, treasuryAddr, gasPrice, gasLimit)
	if err != nil {
		return nil, err
	}
	// wait status of deployment
	status, err = cli.WaitingTransactionStatus(txHash)
	if err != nil {
		return nil, err
	}
	if !status {
		return nil, errors.New("[FirstDeployment] invalid tx, cannot get status")
	}
	// deploy verifier contract
	verifierAddr, txHash, err := DeployVerifierContract(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return nil, err
	}
	// wait status of deployment
	status, err = cli.WaitingTransactionStatus(txHash)
	if err != nil {
		return nil, err
	}
	if !status {
		return nil, errors.New("[FirstDeployment] invalid tx, cannot get status")
	}
	// deploy zecrey contract
	zecreyAddr, txHash, err := DeployZecreyContract(cli, authCli, l2ChainId, nativeAssetId, maxPendingBlocks, gasPrice, gasLimit)
	if err != nil {
		return nil, err
	}
	// wait status of deployment
	status, err = cli.WaitingTransactionStatus(txHash)
	if err != nil {
		return nil, err
	}
	if !status {
		return nil, errors.New("[FirstDeployment] invalid tx, cannot get status")
	}

	return &ZecreyContracts{governanceAddr, assetGovernanceAddr, verifierAddr, zecreyAddr}, nil
}
