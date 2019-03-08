package util

import (
	"bytes"
	"compress/zlib"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandStringBytes get random string
func RandStringBytes(n uint) string {
	if n < 0 {
		return ""
	}
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// ZlibCompress be data compress to zlib string
func ZlibCompress(src []byte) (data string, err error) {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	_, err = w.Write(src)
	if err != nil {
		return
	}
	err = w.Close()
	if err != nil {
		return
	}
	return in.String(), nil
}
