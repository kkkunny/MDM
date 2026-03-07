package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

type BytesSeq interface {
	string | []byte | []rune
}

func MD5[T BytesSeq](data T) T {
	b := md5.Sum([]byte(string(data)))
	return T(hex.EncodeToString(b[:]))
}

func ToJson[T BytesSeq](data any) T {
	b, _ := json.Marshal(data)
	return T(string(b))
}
