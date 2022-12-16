package ksuid

import (
	"github.com/segmentio/ksuid"
)

type ksuIdentifier struct {
}

func (i *ksuIdentifier) GenerateId() (string, error) {
	id, err := ksuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func NewIdentifier() *ksuIdentifier {
	return &ksuIdentifier{}
}
