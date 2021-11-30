package basic

import "log"

const(
	OnChainOpsTreeDepth = 5
	OnChainOpsTreeHelperDepth = 4
)

func init() {
	// set log info
	log.SetFlags(log.Llongfile | log.Ldate | log.Lmicroseconds)
}
