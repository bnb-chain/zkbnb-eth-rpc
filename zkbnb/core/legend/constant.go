package legend

import "log"

func init() {
	// set log info
	log.SetFlags(log.Llongfile | log.Ldate | log.Lmicroseconds)
}
