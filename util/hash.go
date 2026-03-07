package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

type Bytes interface {
	string | []byte
}

func MD5[T Bytes](data T) T {
	b := md5.Sum([]byte(data))
	return T(hex.EncodeToString(b[:]))
}

func ToJson[T Bytes](data any) T {
	b, _ := json.Marshal(data)
	return T(b)
}
