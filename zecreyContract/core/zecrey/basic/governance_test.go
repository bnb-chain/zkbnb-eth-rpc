package basic

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestInitializeGovernance(t *testing.T) {
	var buf bytes.Buffer
	buf.WriteString("0xE9a8Abc30CB4BeB38647aF2C69039878D4044a70")
	buf.WriteString("0xE9a8Abc30CB4BeB38647aF2C69039878D4044a70")
	params, _ := PackInitializeGovernanceParams("0xE9a8Abc30CB4BeB38647aF2C69039878D4044a70")
	fmt.Println(common.Bytes2Hex(params))
}
