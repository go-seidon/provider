package crypto

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type cryptoRandomizer struct {
	dictionary string
}

func (r *cryptoRandomizer) String(n int) (s string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("failed generate random string")
		}
	}()

	res := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(r.dictionary))))
		if err != nil {
			return "", err
		}
		res[i] = r.dictionary[num.Int64()]
	}

	return string(res), nil
}

func NewRandomizer(opts ...RandomizerOption) *cryptoRandomizer {
	p := RandomizerParam{
		Dictionary: "0123456789abcdefghijklmnopqrstuvwxyz",
	}
	for _, opt := range opts {
		opt(&p)
	}

	return &cryptoRandomizer{
		dictionary: p.Dictionary,
	}
}
