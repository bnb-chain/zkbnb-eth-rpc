package contract

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"golang.org/x/crypto/sha3"
	"math/big"
	"zksneak-eth-rpc/_const"
	"zksneak-eth-rpc/_rpc"
	"zksneak-eth-rpc/_utils"
	"zksneak-eth-rpc/contract/_interface/erc20"
)

// get erc20 token balance
func GetErc20Balance(cli *_rpc.ProviderClient, address string, tokenAddress string) (balance *big.Int, err error) {
	// validate the address
	if !_utils.IsValidEthAddress(address) {
		return nil, _rpc.InvalidAddress
	}
	account := common.HexToAddress(address)
	tokenContractAddress := common.HexToAddress(tokenAddress)
	token, err := erc20.NewToken(tokenContractAddress, cli)
	if err != nil {
		return nil, err
	}
	return token.BalanceOf(&bind.CallOpts{}, account)
}

func GetErc20EtherBalance(cli *_rpc.ProviderClient, address string, tokenAddress string) (balance *big.Float, err error) {
	_balance, err := GetErc20Balance(cli, address, tokenAddress)
	if err != nil {
		return nil, err
	}
	balance = _utils.WeiToEther(_balance)
	return balance, nil
}

func TransferErc20(cli *_rpc.ProviderClient, authClient *_rpc.AuthClient, toAddress string, amount *big.Int, gasPrice *big.Int, tokenAddress string) (txHash string, err error) {
	// validate auth cli
	isContract, err := cli.IsContract(tokenAddress)
	if err != nil {
		return "", _rpc.InvalidParams
	}
	if authClient == nil || !_utils.IsValidEthAddress(toAddress) || !isContract {
		return "", _rpc.InvalidParams
	}
	// get account and token address
	toAccount := common.HexToAddress(toAddress)
	tokenContractAddress := common.HexToAddress(tokenAddress)
	// construct transfer function signature hash value
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodId := hash.Sum(nil)[:4]
	// pad address
	paddedToAddress := common.LeftPadBytes(toAccount.Bytes(), 32)
	// pad transfer value
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	var data []byte
	data = append(data, methodId...)
	data = append(data, paddedToAddress...)
	data = append(data, paddedAmount...)
	// estimate gas limit
	gasLimit, err := cli.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAccount,
		Data: data,
	})
	if gasLimit < _const.SuggestFunctionGasLimit {
		gasLimit = _const.SuggestFunctionGasLimit
	}
	// get nonce of address
	nonce, err := cli.GetPendingNonce(authClient.Address.String())
	if err != nil {
		return "", err
	}
	// set default gasPrice
	if gasPrice == nil {
		gasPrice, err = cli.SuggestGasPrice(context.Background())
		if err != nil {
			return "", err
		}
	}
	// construct transaction
	transaction := types.NewTransaction(nonce, tokenContractAddress, big.NewInt(0), gasLimit, gasPrice, data)
	// sign transaction
	signedTx, err := _rpc.SignTx(authClient, transaction)
	if err != nil {
		return "", err
	}
	// send transaction
	err = cli.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}
	return signedTx.Hash().String(), nil
}
