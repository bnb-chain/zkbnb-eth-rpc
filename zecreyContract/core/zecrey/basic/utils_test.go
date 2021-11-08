package basic

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"
)

func TestSetFixed32Bytes(t *testing.T) {
	buf := SetFixed32Bytes([]byte("1"))
	fmt.Println(hex.EncodeToString(buf[:]))
	fmt.Println(new(big.Int).SetBytes(buf[:]).String())
}
