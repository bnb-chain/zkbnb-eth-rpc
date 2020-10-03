package chain

import (
	"context"
	"eva-go-sdk/econst"
	"eva-go-sdk/erpc"
	"eva-go-sdk/eutils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"io/ioutil"
	"math/big"
	"strings"
	"time"
)

// transfer eth
func Transfer(authCli *erpc.AuthClient, to string, amount *big.Int, data []byte, gasLimit uint64) (txHash string, err error) {
	// validate amount,it can't less than zero
	if amount.Cmp(new(big.Int).SetUint64(0)) < 0 {
		return "", erpc.AmountLessThanZero
	}
	// check address
	isValidTo := eutils.IsValidEthAddress(to)
	if !isValidTo {
		return "", erpc.InvalidAddress
	}
	// get connection
	cli, err := erpc.GetConnection()
	if err != nil {
		return "", err
	}
	defer cli.Close()
	// transfer private key to address
	fromAddress, err := PrivateKeyToAddress(authCli.PrivateKey)
	if err != nil {
		return "", err
	}
	// get nonce of from address
	nonce, err := cli.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}
	// get suggested gas price
	gasPrice, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	toAddress := common.HexToAddress(to)
	tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, data)
	signedTx, err := signTx(authCli, tx)
	if err != nil {
		return "", err
	}
	err = cli.SendTransaction(context.Background(), signedTx)
	return signedTx.Hash().String(), err
}

// sign transaction
func signTx(authCli *erpc.AuthClient, tx *types.Transaction) (signedTx *types.Transaction, err error) {
	// check if it is valid params
	if authCli.ChainId == nil || authCli.PrivateKey == nil {
		return nil, erpc.InvalidAuthClientParams
	}
	return types.SignTx(tx, types.NewEIP155Signer(authCli.ChainId), authCli.PrivateKey)
}

// deploy contract
func DeployContract(authCli *erpc.AuthClient, gasPrice *big.Int, abiPath string, binPath string, params []interface{}) (contractAddress common.Address, txHash common.Hash, err error) {
	// check authCli
	if authCli == nil || abiPath == "" || binPath == "" {
		return contractAddress, txHash, erpc.InvalidParams
	}
	// transfer private key to address
	address, err := PrivateKeyToAddress(authCli.PrivateKey)
	if err != nil {
		return contractAddress, txHash, err
	}
	// get pending nonce
	nonce, err := GetPendingNonce(address.String())
	// get connection
	cli, err := erpc.GetConnection()
	if err != nil {
		return contractAddress, txHash, err
	}
	defer cli.Close()
	// get gas price
	if gasPrice == nil {
		// get suggest gas price
		gasPrice, err = cli.SuggestGasPrice(context.Background())
		if err != nil {
			return contractAddress, txHash, err
		}
	}
	// create authentication
	auth := bind.NewKeyedTransactor(authCli.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = econst.SuggestContractGasLimit
	auth.GasPrice = gasPrice
	// get abi
	abiData, err := ioutil.ReadFile(abiPath)
	if err != nil {
		return contractAddress, txHash, err
	}
	// get bin data
	data, err := ioutil.ReadFile(binPath)
	if err != nil {
		return contractAddress, txHash, err
	}
	binData := "0x" + string(data)
	// parse abi
	parsed, err := abi.JSON(strings.NewReader(string(abiData)))
	if err != nil {
		return contractAddress, txHash, err
	}
	address, tx, _, err := bind.DeployContract(auth, parsed, common.FromHex(binData), cli, params...)
	if err != nil {
		return contractAddress, txHash, err
	}
	return address, tx.Hash(), nil
}

// wait until transaction is completed
func WaitingTransactionStatus(txHash string) (status bool, err error) {
	// set ticker
	ticker := time.NewTicker(econst.TryTimeInterval)
	defer ticker.Stop()
	count := 0
	for {
		if count > econst.MaxTryTimes {
			return false, erpc.ErrorGetBlockStatus
		}
		// get transaction info
		_, isPending, err := GetTransactionByHash(txHash)
		if err != nil {
			return false, err
		}
		// if transaction execution is completed
		if !isPending {
			// get receipt of transaction
			receipt, err := GetTransactionReceipt(txHash)
			if err != nil {
				return false, err
			}
			if receipt.Status == econst.TransactionSuccess {
				status = true
			} else if receipt.Status == econst.TransactionFail {
				status = false
			}
			return status, nil
		}
		count++
		<-ticker.C
	}
}

// deploy smart contract until it is completed
func DeployContractUntil(authCli *erpc.AuthClient, gasPrice *big.Int, abiPath string, binPath string, params []interface{}) (status bool, contractAddress common.Address, txHash common.Hash, err error) {
	contractAddress, txHash, err = DeployContract(authCli, gasPrice, abiPath, binPath, params)
	if err != nil {
		return false, contractAddress, txHash, err
	}
	// waiting transaction
	status, err = WaitingTransactionStatus(txHash.String())
	if err != nil {
		return false, contractAddress, txHash, err
	}
	return status, contractAddress, txHash, nil
}
