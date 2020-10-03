package chain

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestGetBalance(t *testing.T) {
	toAddress := "0xE9b15a2D396B349ABF60e53ec66Bcf9af262D449"
	balance, err := GetBalance(toAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("balance:", balance)

	etherBalance, err := GetEtherBalance(toAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("ether balance:", etherBalance)
}

func TestIsContract(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name           string
		args           args
		wantIsContract bool
		wantErr        bool
	}{
		{
			name: "valid",
			args: args{
				// smart contract address
				address: "0x8b2a865c5856571bc7f9951fee16215a6b2250e1",
			},
			wantIsContract: true,
			wantErr:        false,
		},
		{
			name: "invalid",
			args: args{
				// eth account
				address: "0xfef01c3494e9fbad65a5d12b3852ca87361bc9b2",
			},
			wantIsContract: false,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIsContract, err := IsContract(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsContract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIsContract != tt.wantIsContract {
				t.Errorf("IsContract() = %v, want %v", gotIsContract, tt.wantIsContract)
			}
		})
	}
}

func TestGetBlockHeaderByNumber(t *testing.T) {
	header, err := GetBlockHeaderByNumber(big.NewInt(7298401))
	if err != nil {
		t.Error(err)
	}
	fmt.Println("header info:", header)
}

func TestGetBlockInfoByNumber(t *testing.T) {
	blockInfo, err := GetBlockInfoByNumber(big.NewInt(7298401))
	if err != nil {
		t.Error(err)
	}
	fmt.Println("block info:", blockInfo)
}

func TestGetHeight(t *testing.T) {
	height, err := GetHeight()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("height:", height)
}

func TestGetTransactionByHash(t *testing.T) {
	hash := "0xf900253477a50a1cd808f61058f68eb2e73afcb0161c31e82ecafa034d7c8eec"
	tx, isPending, err := GetTransactionByHash(hash)
	if err != nil {
		t.Error(err)
	}
	txBytes, err := tx.MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("tx:", string(txBytes))
	fmt.Println("isPending:", isPending)
}

func TestGetTransactionReceipt(t *testing.T) {
	successHash := "0xf900253477a50a1cd808f61058f68eb2e73afcb0161c31e82ecafa034d7c8eec"
	receipt, err := GetTransactionReceipt(successHash)
	if err != nil {
		t.Error(err)
	}
	receiptBytes, err := receipt.MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("success receipt:", string(receiptBytes))
	errorHash := "0x782d3bd436bea25d6d687daa8527abb2487b4d73eb5f9e1350ab70a11adf13b3"
	receipt, err = GetTransactionReceipt(errorHash)
	if err != nil {
		t.Error(err)
	}
	receiptBytes, err = receipt.MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("error receipt:", string(receiptBytes))
}

func TestPrivateKeyToAddress(t *testing.T) {
	type args struct {
		privateKey *ecdsa.PrivateKey
	}
	sk, err := crypto.GenerateKey()
	if err != nil {
		t.Error(err)
	}
	pubKey := sk.PublicKey
	address := crypto.PubkeyToAddress(pubKey)
	tests := []struct {
		name        string
		args        args
		wantAddress common.Address
		wantErr     bool
	}{
		{
			name: "private key",
			args: args{
				sk,
			},
			wantAddress: address,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAddress, err := PrivateKeyToAddress(tt.args.privateKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrivateKeyToAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAddress, tt.wantAddress) {
				t.Errorf("PrivateKeyToAddress() = %v, want %v", gotAddress, tt.wantAddress)
			}
		})
	}
}
