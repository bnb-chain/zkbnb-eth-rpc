package basic

func IsValidAddress(addr string) bool {
	if len(addr) != 42 {
		return false
	}
	return true
}
