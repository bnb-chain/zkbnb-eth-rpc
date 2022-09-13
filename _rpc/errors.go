package _rpc

import (
	"errors"
)

var (
	ErrInvalidAddress          = errors.New("err: input address is not valid! ")
	ErrInvalidAuthClientParams = errors.New("err: auth cli params is not valid! ")
	ErrInvalidParams           = errors.New("err: input params are not valid! ")
	ErrInvalidHashValue        = errors.New("err: input hash value is not valid! ")
	ErrInvalidPrivateKey       = errors.New("err: input private key is not valid! ")
	ErrGetBlockStatus          = errors.New("err: error when get status from chain! ")
	ErrAmountLessThanZero      = errors.New("err: input amount value can not less than 0! ")
)
