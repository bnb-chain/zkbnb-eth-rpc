package basic

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"log"
)

const (
	Bytes32Len = 32
)

var (
	AddressType, _      = abi.NewType("address", "", nil)
	Bytes32Type, _      = abi.NewType("bytes32", "", nil)
	Uint32Type, _       = abi.NewType("uint32", "", nil)
	Uint256Type, _      = abi.NewType("uint256", "", nil)
	MerkleHelperType, _ = abi.NewType("bool[4]", "", nil)
)

func init() {
	// set log info
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
}