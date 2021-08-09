package zecrey

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
	"github.com/zecrey-labs/zecrey-eth-rpc/_utils"
	"math/big"
)

func DeployZecreyContract(cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, verifierAddr string, governanceAddr string, gasPrice *big.Int, gasLimit uint64) (addr string, txHash string, err error) {
	// construct transaction ops
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", "", err
	}
	// check if the verifier address is valid
	isValidVerifier, err := cli.IsContract(verifierAddr)
	if err != nil {
		return "", "", err
	}
	if !isValidVerifier {
		return "", "", ErrNotContractAddr
	}
	// check if the governance address is valid
	isValidGovernance, err := cli.IsContract(governanceAddr)
	if err != nil {
		return "", "", err
	}
	if !isValidGovernance {
		return "", "", ErrNotContractAddr
	}
	// transfer to common.Address
	verifier := common.HexToAddress(verifierAddr)
	governance := common.HexToAddress(governanceAddr)
	// deploy zecrey contract
	address, tx, _, err := DeployZecrey(transactOpts, *cli, verifier, governance)
	if err != nil {
		return "", "", err
	}
	return address.Hex(), tx.Hash().Hex(), nil
}

/*
	LoadZecreyInstance: load zecrey instance if it is already deployed
*/
func LoadZecreyInstance(cli *_rpc.ProviderClient, addr string) (instance *Zecrey, err error) {
	instance, err = NewZecrey(common.HexToAddress(addr), *cli)
	return instance, err
}

func Deposit(cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Zecrey, assetId uint32, zecreyAddr string, amount *big.Int, gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	transactOpts.From = authCli.Address
	transactOpts.Value = amount
	tx, err := instance.Deposit(transactOpts, assetId, zecreyAddr, amount)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func SignTransaction(chainID int64, tx *types.Transaction, privateKeyStr string) (string, error) {
	privateKey, err := _utils.DecodePrivateKeyBigInt(privateKeyStr)
	if err != nil {
		return "", err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(chainID)), privateKey)
	//signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		return "", nil
	}

	return EncodeSignedTx(signedTx)
}

/*
	ConstructDepositSignedTx: helper function, construct deposit signed tx
*/
func ConstructDepositSignedTx(
	skStr string,
	chainId int64,
	nonce uint64,
	toContractAddress string,
	assetId uint32,
	zecreyAddr string,
	amount *big.Int,
	gasPrice *big.Int,
	gasLimit uint64,
) (signedTx string, err error) {
	// define types
	data, err := zecreyABI.Pack("deposit", assetId, zecreyAddr, amount)
	if err != nil {
		return "", err
	}
	// construct tx
	tx := types.NewTransaction(nonce, common.HexToAddress(toContractAddress), amount, gasLimit, gasPrice, data)
	// sign transaction
	signedTx, err = SignTransaction(chainId, tx, skStr)
	if err != nil {
		return "", err
	}
	return signedTx, nil
}

func EncodeSignedTx(signedTx *types.Transaction) (string, error) {
	txBytes, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(txBytes), nil
}

func DecodeSignedTx(txStr string) (signedTx *types.Transaction, err error) {
	txBytes, err := hex.DecodeString(txStr)
	if err != nil {
		return nil, err
	}
	err = rlp.DecodeBytes(txBytes, &signedTx)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}

func CommitBlocks(cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Zecrey, storageBlock StorageBlock, commitBlocks []ZecreyCommitBlock, gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	tx, err := instance.CommitBlocks(transactOpts, storageBlock, commitBlocks)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func ProcessBlocks(cli *_rpc.ProviderClient, authCli *_rpc.AuthClient, instance *Zecrey, processBlocks []ZecreyProcessBlock, gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	transactOpts, err := ConstructTransactOpts(cli, authCli, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	tx, err := instance.ProcessBlocks(transactOpts, processBlocks)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func GetTotalDeposited(cli *_rpc.ProviderClient, instance *Zecrey, assetId uint32) (*big.Int, error) {
	deposited, err := instance.GetTotalDeposited(EmptyCallOpts(), assetId)
	return deposited, err
}

func GetTotalWithdrawn(cli *_rpc.ProviderClient, instance *Zecrey, index uint64) (*big.Int, error) {
	withdrawn, err := instance.TotalWithdrawn(EmptyCallOpts(), big.NewInt(int64(index)))
	return withdrawn, err
}

func GetTotalBlocksCommitted(cli *_rpc.ProviderClient, instance *Zecrey) (uint32, error) {
	committed, err := instance.TotalBlocksCommitted(EmptyCallOpts())
	return committed, err
}

func GetTotalBlocksProcessed(cli *_rpc.ProviderClient, instance *Zecrey) (uint32, error) {
	processed, err := instance.TotalBlocksProcessed(EmptyCallOpts())
	return processed, err
}
