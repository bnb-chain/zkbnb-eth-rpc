package basic

import (
	"errors"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
	"log"
	"math/big"
	"time"
)

/*
	FirstDeployment: helper method to deploy all contracts
*/
func FirstDeployment(
	cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, waitTime time.Duration,
	governor string,
	l2ChainId uint8, nativeAssetId uint16, maxPendingBlocks uint16,
	listingFeeTokenAddr string, listingFee *big.Int, listingCap uint16, treasuryAddr string,
	genesisBlockNumber uint32, genesisOnchainOpsRoot []byte, genesisStateRoot []byte, genesisTimestamp *big.Int,
	genesisCommitment []byte, genesisMerkleHelper [OnChainOpsTreeHelperDepth]bool,
	gasPrice *big.Int, gasLimit uint64,
) (addrs *ZecreyContracts, err error) {
	// deploy REY ERC20
	reyErc20Addr, txHash, err := DeployErc20Contract(cli, authCli, ERC20_REY_Name, ERC20_REY_Symbol, gasPrice, gasLimit)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	time.Sleep(waitTime)
	// wait status of deployment
	status, err := cli.WaitingTransactionStatus(txHash)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	if !status {
		log.Println("[FirstDeployment] invalid tx, cannot get status")
		return nil, errors.New("[FirstDeployment] invalid tx, cannot get status")
	}
	// deploy ETH ERC20
	ethErc20Addr, txHash, err := DeployErc20Contract(cli, authCli, ERC20_ETH_Name, ERC20_ETH_Symbol, gasPrice, gasLimit)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	time.Sleep(waitTime)
	// wait status of deployment
	status, err = cli.WaitingTransactionStatus(txHash)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	if !status {
		log.Println("[FirstDeployment] invalid tx, cannot get status")
		return nil, errors.New("[FirstDeployment] invalid tx, cannot get status")
	}
	// deploy MATIC ERC20
	maticErc20Addr, txHash, err := DeployErc20Contract(cli, authCli, ERC20_MATIC_Name, ERC20_MATIC_Symbol, gasPrice, gasLimit)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	time.Sleep(waitTime)
	// wait status of deployment
	status, err = cli.WaitingTransactionStatus(txHash)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	if !status {
		log.Println("[FirstDeployment] invalid tx, cannot get status")
		return nil, errors.New("[FirstDeployment] invalid tx, cannot get status")
	}
	// deploy NEAR ERC20
	nearErc20Addr, txHash, err := DeployErc20Contract(cli, authCli, ERC20_NEAR_Name, ERC20_NEAR_Symbol, gasPrice, gasLimit)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	time.Sleep(waitTime)
	// wait status of deployment
	status, err = cli.WaitingTransactionStatus(txHash)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	if !status {
		log.Println("[FirstDeployment] invalid tx, cannot get status")
		return nil, errors.New("[FirstDeployment] invalid tx, cannot get status")
	}
	// deploy AVAX ERC20
	avaxErc20Addr, txHash, err := DeployErc20Contract(cli, authCli, ERC20_AVAX_Name, ERC20_AVAX_Symbol, gasPrice, gasLimit)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	time.Sleep(waitTime)
	// wait status of deployment
	status, err = cli.WaitingTransactionStatus(txHash)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	if !status {
		log.Println("[FirstDeployment] invalid tx, cannot get status")
		return nil, errors.New("[FirstDeployment] invalid tx, cannot get status")
	}
	// deploy BIT ERC20
	bitErc20Addr, txHash, err := DeployErc20Contract(cli, authCli, ERC20_BIT_Name, ERC20_BIT_Symbol, gasPrice, gasLimit)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	time.Sleep(waitTime)
	// wait status of deployment
	status, err = cli.WaitingTransactionStatus(txHash)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	if !status {
		log.Println("[FirstDeployment] invalid tx, cannot get status")
		return nil, errors.New("[FirstDeployment] invalid tx, cannot get status")
	}
	// deploy USDT ERC20
	usdtErc20Addr, txHash, err := DeployErc20Contract(cli, authCli, ERC20_USDT_Name, ERC20_USDT_Symbol, gasPrice, gasLimit)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	time.Sleep(waitTime)
	// wait status of deployment
	status, err = cli.WaitingTransactionStatus(txHash)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	if !status {
		log.Println("[FirstDeployment] invalid tx, cannot get status")
		return nil, errors.New("[FirstDeployment] invalid tx, cannot get status")
	}
	// deploy USDC ERC20
	usdcErc20Addr, txHash, err := DeployErc20Contract(cli, authCli, ERC20_USDC_Name, ERC20_USDC_Symbol, gasPrice, gasLimit)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	time.Sleep(waitTime)
	// wait status of deployment
	status, err = cli.WaitingTransactionStatus(txHash)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	if !status {
		log.Println("[FirstDeployment] invalid tx, cannot get status")
		return nil, errors.New("[FirstDeployment] invalid tx, cannot get status")
	}
	// deploy DAI ERC20
	daiErc20Addr, txHash, err := DeployErc20Contract(cli, authCli, ERC20_DAI_Name, ERC20_DAI_Symbol, gasPrice, gasLimit)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	time.Sleep(waitTime)
	// wait status of deployment
	status, err = cli.WaitingTransactionStatus(txHash)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	if !status {
		log.Println("[FirstDeployment] invalid tx, cannot get status")
		return nil, errors.New("[FirstDeployment] invalid tx, cannot get status")
	}
	// deploy BNB ERC20
	bnbErc20Addr, txHash, err := DeployErc20Contract(cli, authCli, ERC20_BNB_Name, ERC20_BNB_Symbol, gasPrice, gasLimit)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	time.Sleep(waitTime)
	// wait status of deployment
	status, err = cli.WaitingTransactionStatus(txHash)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	if !status {
		log.Println("[FirstDeployment] invalid tx, cannot get status")
		return nil, errors.New("[FirstDeployment] invalid tx, cannot get status")
	}
	// deploy governance contract
	governanceAddr, txHash, err := DeployGovernanceContract(cli, authCli, l2ChainId, nativeAssetId, maxPendingBlocks, gasPrice, gasLimit)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	time.Sleep(waitTime)
	// wait status of deployment
	status, err = cli.WaitingTransactionStatus(txHash)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	if !status {
		log.Println("[FirstDeployment] invalid tx, cannot get status")
		return nil, errors.New("[FirstDeployment] invalid tx, cannot get status")
	}
	// load instance from address
	governanceInstance, err := LoadGovernanceInstance(cli, governanceAddr)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	time.Sleep(waitTime)
	// initialize governance contract
	txHash, err = GovernanceInitialize(cli, authCli, governanceInstance, governor, gasPrice, gasLimit)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	time.Sleep(waitTime)
	// wait status of deployment
	status, err = cli.WaitingTransactionStatus(txHash)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	if !status {
		log.Println("[FirstDeployment] invalid tx, cannot get status")
		return nil, errors.New("[FirstDeployment] invalid tx, cannot get status")
	}
	// deploy asset governance contract
	assetGovernanceAddr, txHash, err := DeployAssetGovernanceContract(
		cli, authCli, governanceAddr,
		listingFeeTokenAddr, listingFee, listingCap, treasuryAddr, gasPrice, gasLimit,
	)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	time.Sleep(waitTime)
	// wait status of deployment
	status, err = cli.WaitingTransactionStatus(txHash)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	if !status {
		log.Println("[FirstDeployment] invalid tx, cannot get status")
		return nil, errors.New("[FirstDeployment] invalid tx, cannot get status")
	}
	//// get instance of asset governance
	//assetGovernanceInstance, err := LoadAssetGovernanceInstance(cli, assetGovernanceAddr)
	//if err != nil {
	//	log.Println("[FirstDeployment] err info:", err)
	//	return nil, err
	//}
	//// set lister
	//txHash, err = AssetGovernanceSetLister(cli, authCli, assetGovernanceInstance, governor, true, gasPrice, gasLimit)
	//if err != nil {
	//	log.Println("[FirstDeployment] err info:", err)
	//	return nil, err
	//}
	//time.Sleep(waitTime)
	//// wait status of deployment
	//status, err = cli.WaitingTransactionStatus(txHash)
	//if err != nil {
	//	log.Println("[FirstDeployment] err info:", err)
	//	return nil, err
	//}
	//if !status {
	//	log.Println("[FirstDeployment] invalid tx, cannot get status")
	//	return nil, errors.New("[FirstDeployment] invalid tx, cannot get status")
	//}
	// deploy verifier contract
	verifierAddr, txHash, err := DeployVerifierContract(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	time.Sleep(waitTime)
	// wait status of deployment
	status, err = cli.WaitingTransactionStatus(txHash)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	if !status {
		log.Println("[FirstDeployment] invalid tx, cannot get status")
		return nil, errors.New("[FirstDeployment] invalid tx, cannot get status")
	}
	// deploy zecrey contract
	zecreyAddr, txHash, err := DeployZecreyContract(cli, authCli, l2ChainId, nativeAssetId, maxPendingBlocks, gasPrice, gasLimit)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	time.Sleep(waitTime)
	// wait status of deployment
	status, err = cli.WaitingTransactionStatus(txHash)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	if !status {
		log.Println("[FirstDeployment] invalid tx, cannot get status")
		return nil, errors.New("[FirstDeployment] invalid tx, cannot get status")
	}
	// load zecrey instance
	zecreyInstance, err := LoadZecreyInstance(cli, zecreyAddr)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	// initialize zecrey contract
	txHash, err = ZecreyInitialize(
		cli, authCli, zecreyInstance,
		governanceAddr,
		verifierAddr,
		genesisBlockNumber,
		genesisOnchainOpsRoot,
		genesisStateRoot,
		genesisTimestamp,
		genesisCommitment,
		genesisMerkleHelper,
		gasPrice, gasLimit)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	time.Sleep(waitTime)
	// wait status of deployment
	status, err = cli.WaitingTransactionStatus(txHash)
	if err != nil {
		log.Println("[FirstDeployment] err info:", err)
		return nil, err
	}
	if !status {
		return nil, errors.New("[FirstDeployment] invalid tx, cannot get status")
	}
	return &ZecreyContracts{
		GovernanceAddr:      governanceAddr,
		AssetGovernanceAddr: assetGovernanceAddr,
		VerifierAddr:        verifierAddr,
		ZecreyAddr:          zecreyAddr,
		ReyERC20Addr:        reyErc20Addr,
		EthERC20Addr:        ethErc20Addr,
		MaticERC20Addr:      maticErc20Addr,
		NearERC20Addr:       nearErc20Addr,
		AvaxERC20Addr:       avaxErc20Addr,
		BitERC20Addr:        bitErc20Addr,
		UsdtERC20Addr:       usdtErc20Addr,
		UsdcERC20Addr:       usdcErc20Addr,
		DaiERC20Addr:        daiErc20Addr,
		BnbERC20Addr:        bnbErc20Addr,
	}, nil
}
