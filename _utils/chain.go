package _utils

import (
	"bytes"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"math"
	"math/big"
	"regexp"
	"Zecrey-eth-rpc/_const"
)

// transfer wei to ether balance/10^{18}
func WeiToEther(_balance *big.Int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(_balance.String())
	return new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
}

// check if it is a valid hash value
func IsValidHashValue(hash string) bool {
	re := regexp.MustCompile(_const.HashRegexp)
	return re.MatchString(hash)
}

// check if it is a valid address
func IsValidEthAddress(address string) bool {
	re := regexp.MustCompile(_const.AddressRegexp)
	return re.MatchString(address)
}

// check if it is a valid private key
func IsValidPrivateKey(sk string) bool {
	// if sk starts with 0x
	var re *regexp.Regexp
	if len(sk) == 66 && sk[:2] == "0x" {
		re = regexp.MustCompile(_const.SkRegexp0x)
	} else if len(sk) == 64 {
		re = regexp.MustCompile(_const.SkRegexp)
	}
	if re == nil {
		return false
	}
	return re.MatchString(sk)

}

// transfer ether to wei
func EtherToWei(value float64) *big.Int {
	one := math.Pow10(18)
	res := value * one
	result := new(big.Int)
	big.NewFloat(res).Int(result)
	return result
}

// transfer private key to str
func EncodePrivateKey(sk *ecdsa.PrivateKey) (priKey string, err error) {
	if sk == nil {
		return "", ErrInvalidPrivateKey
	}
	// remove "0x"
	return hexutil.Encode(crypto.FromECDSA(sk)), nil
}

// transfer str to private key
func DecodePrivateKey(sk string) (privateKey *ecdsa.PrivateKey, err error) {
	isValidPrivateKey := IsValidPrivateKey(sk)
	if !isValidPrivateKey {
		return nil, ErrInvalidPrivateKey
	}
	// add "0x"
	if sk[:2] != "0x" {
		sk = "0x" + sk
	}
	privateKeyBytes, err := hexutil.Decode(sk)
	if err != nil {
		return nil, err
	}
	return crypto.ToECDSA(privateKeyBytes)
}

// public key to str
func EncodePubKey(pk *ecdsa.PublicKey) (pubKey string, err error) {
	if pk == nil {
		return "", ErrInvalidPublicKey
	}
	return hexutil.Encode(crypto.FromECDSAPub(pk)), nil
}

// decode public key
func DecodePubKey(pk string) (pubKey *ecdsa.PublicKey, err error) {
	pubKeyBytes, err := hexutil.Decode(pk)
	if err != nil {
		return nil, err
	}
	pubKey, err = crypto.UnmarshalPubkey(pubKeyBytes)
	if err != nil {
		return nil, err
	}
	return pubKey, nil
}

// create a signature
func Sign(sk string, data []byte) (sig []byte, err error) {
	hash := crypto.Keccak256Hash(data)
	// transfer sk str to private key
	privateKey, err := DecodePrivateKey(sk)
	if err != nil {
		return nil, err
	}
	return crypto.Sign(hash.Bytes(), privateKey)
}

// verify ecdsa signature
func VerifySig(pk string, sig []byte, data []byte) (res bool, err error) {
	hash := crypto.Keccak256Hash(data)
	// transfer pk str to public key
	pubKey, err := DecodePubKey(pk)
	if err != nil {
		return false, err
	}
	// get public key which was used to create the sig
	sigPubKeyBytes, err := crypto.Ecrecover(hash.Bytes(), sig)
	if err != nil {
		return false, err
	}
	// get public key bytes
	pubKeyBytes := crypto.FromECDSAPub(pubKey)
	// compare two public keys
	isEqual := bytes.Equal(sigPubKeyBytes, pubKeyBytes)
	if !isEqual {
		return false, ErrInvalidVerifySigPubKey
	}
	// get signature recover id
	signatureNoRecoverID := sig[:len(sig)-1] // remove recovery ID
	return crypto.VerifySignature(pubKeyBytes, hash.Bytes(), signatureNoRecoverID), nil
}