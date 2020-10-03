package eutils

import "errors"

var (
	InvalidPrivateKey      = errors.New("Input private key is not valid! ")
	InvalidPublicKey       = errors.New("Input nil public key is not valid! ")
	InvalidVerifySigPubKey = errors.New("Input public key is not valid for the signature! ")
)
