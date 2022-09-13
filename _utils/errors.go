package _utils

import (
	"errors"
)

var (
	ErrInvalidPrivateKey      = errors.New("err: input private key is not valid! ")
	ErrInvalidPublicKey       = errors.New("err: input nil public key is not valid! ")
	ErrInvalidVerifySigPubKey = errors.New("err: input public key is not valid for the signature! ")
)
