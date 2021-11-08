package basic

import (
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
	"math/big"
)

/*
	FirstDeployment: helper method to deploy all contracts
*/
func FirstDeployment(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient,
	l2ChainId uint8, nativeAssetId uint16, maxPendingBlocks uint16,
	listingFeeTokenAddr string, listingFee *big.Int, listingCap uint16, treasuryAddr string,
	gasPrice *big.Int, gasLimit uint64,
) (addrs []string, err error) {
	governanceAddr, _, err := DeployGovernanceContract(cli, authCli, l2ChainId, nativeAssetId, maxPendingBlocks, gasPrice, gasLimit)
	if err != nil {
		return nil, err
	}
	assetGovernanceAddr, _, err := DeployAssetGovernanceContract(cli, authCli, governanceAddr, listingFeeTokenAddr, listingFee, listingCap, treasuryAddr, gasPrice, gasLimit)
	if err != nil {
		return nil, err
	}
	verifierAddr, _, err := DeployVerifierContract(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return nil, err
	}
	zecreyAddr, _, err := DeployZecreyContract(cli, authCli, l2ChainId, nativeAssetId, maxPendingBlocks, gasPrice, gasLimit)
	if err != nil {
		return nil, err
	}
	return []string{governanceAddr, assetGovernanceAddr, verifierAddr, zecreyAddr}, nil
}
