package erpc

import "errors"

var (
	InvalidAddress          = errors.New("Input address is not valid! ")
	InvalidAuthClientParams = errors.New("Auth client params is not valid! ")
	InvalidParams           = errors.New("Input params are not valid! ")
	InvalidHashValue        = errors.New("Input hash value is not valid! ")
	InvalidPrivateKey       = errors.New("Input private key is not valid! ")
	ErrorGetBlockStatus     = errors.New("Error when get status from chain! ")
	AmountLessThanZero      = errors.New("Input amount value can not less than 0! ")
)
