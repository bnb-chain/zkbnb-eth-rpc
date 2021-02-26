package _utils

import (
	"hash"
)

func CalHash(msg []byte, hFunc func() hash.Hash) (res []byte, err error) {
	h := hFunc()
	_, err = h.Write(msg)
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}
