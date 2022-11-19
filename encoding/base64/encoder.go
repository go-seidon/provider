package base64

import (
	"encoding/base64"
)

type base64Encoder struct {
}

func (e *base64Encoder) Encode(src []byte) (string, error) {
	return base64.StdEncoding.EncodeToString(src), nil
}

func (e *base64Encoder) Decode(d string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(d)
}

func NewEncoder() *base64Encoder {
	e := &base64Encoder{}
	return e
}
