package _rpc

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/zecrey-labs/zecrey-eth-rpc/_const"
	"github.com/zecrey-labs/zecrey-eth-rpc/_utils"
	"io/ioutil"
	"math/big"
	"strings"
	"time"
)

// transfer eth
func (cli *ProviderClient) Transfer(authCli *AuthClient, to string, amount *big.Int, data []byte, gasLimit uint64) (txHash string, err error) {
	// validate amount,it can't less than zero
	if amount.Cmp(new(big.Int).SetUint64(0)) < 0 {
		return "", ErrAmountLessThanZero
	}
	// check address
	isValidTo := _utils.IsValidEthAddress(to)
	if !isValidTo {
		return "", ErrInvalidAddress
	}
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
	signedTx, err := SignTx(authCli, tx)
	if err != nil {
		return "", err
	}
	err = cli.SendTransaction(context.Background(), signedTx)
	return signedTx.Hash().String(), err
}

// deploy zecrey
func (cli *ProviderClient) DeployContract(authCli *AuthClient, gasPrice *big.Int, abiPath string, binPath string, params []interface{}) (contractAddress common.Address, txHash common.Hash, err error) {
	// check authCli
	if authCli == nil || abiPath == "" || binPath == "" {
		return contractAddress, txHash, ErrInvalidParams
	}
	// transfer private key to address
	address, err := PrivateKeyToAddress(authCli.PrivateKey)
	if err != nil {
		return contractAddress, txHash, err
	}
	// get pending nonce
	nonce, err := cli.GetPendingNonce(address.String())
	// get gas price
	if gasPrice == nil {
		// get suggest gas price
		gasPrice, err = cli.SuggestGasPrice(context.Background())
		if err != nil {
			return contractAddress, txHash, err
		}
	}
	// create authentication
	auth := CreateAuthentication(
		authCli.PrivateKey,
		big.NewInt(int64(nonce)),
		big.NewInt(0),
		_const.SuggestContractGasLimit,
		gasPrice,
	)
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
func (cli *ProviderClient) WaitingTransactionStatus(txHash string) (status bool, err error) {
	// set ticker
	ticker := time.NewTicker(_const.TryTimeInterval)
	defer ticker.Stop()
	count := 0
	for {
		if count > _const.MaxTryTimes {
			return false, ErrGetBlockStatus
		}
		// get transaction info
		_, isPending, err := cli.GetTransactionByHash(txHash)
		if err != nil {
			return false, err
		}
		// if transaction execution is completed
		if !isPending {
			// get receipt of transaction
			receipt, err := cli.GetTransactionReceipt(txHash)
			if err != nil {
				return false, err
			}
			if receipt.Status == _const.TransactionSuccess {
				status = true
			} else if receipt.Status == _const.TransactionFail {
				status = false
			}
			return status, nil
		}
		count++
		<-ticker.C
	}
}

// deploy smart zecrey until it is completed
func (cli *ProviderClient) DeployContractUntil(authCli *AuthClient, gasPrice *big.Int, abiPath string, binPath string, params []interface{}) (status bool, contractAddress common.Address, txHash common.Hash, err error) {
	contractAddress, txHash, err = cli.DeployContract(authCli, gasPrice, abiPath, binPath, params)
	if err != nil {
		return false, contractAddress, txHash, err
	}
	// waiting transaction
	status, err = cli.WaitingTransactionStatus(txHash.String())
	if err != nil {
		return false, contractAddress, txHash, err
	}
	return status, contractAddress, txHash, nil
}
