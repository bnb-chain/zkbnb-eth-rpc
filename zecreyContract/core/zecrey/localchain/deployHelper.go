package localchain

import (
	"github.com/zecrey-labs/zecrey-eth-rpc/zecreyContract/core/zecrey/basic"
	"math/big"
)

func FirstDeployment(
	governor string,
	l2ChainId uint8, nativeAssetId uint16, maxPendingBlocks uint16,
	listingFeeTokenAddr string, listingFee *big.Int, listingCap uint16, treasuryAddr string,
	genesisBlockNumber uint32, genesisOnchainOpsRoot []byte, genesisStateRoot []byte, genesisTimestamp *big.Int,
	genesisCommitment []byte, genesisMerkleHelper [OnChainOpsTreeHelperDepth]bool,
	gasPrice *big.Int, gasLimit uint64,
) (addrs *basic.ZecreyContracts, err error) {
	return basic.FirstDeployment(
		Cli, AuthCli,
		governor,
		l2ChainId, nativeAssetId, maxPendingBlocks,
		listingFeeTokenAddr, listingFee, listingCap, treasuryAddr,
		genesisBlockNumber, genesisOnchainOpsRoot, genesisStateRoot, genesisTimestamp,
		genesisCommitment, genesisMerkleHelper,
		gasPrice, gasLimit,
	)
}
