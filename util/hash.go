package util

import (
	"crypto/md5"
	"encoding/hex"
)

type Bytes interface {
	string | []byte
}

func MD5[T Bytes](data T) T {
	b := md5.Sum([]byte(data))
	return T(hex.EncodeToString(b[:]))
}
