package model_fields

import (
	"github.com/satori/go.uuid"
)

type Hash struct {
	Value string
}

func NewHashFromString(v string) *Hash {
	return &Hash{v}
}

func NewHash(iterations int) *Hash {
	var value string

	for i := 0; i < iterations; i++ {
		value += uuid.NewV4().String()
	}

	return &Hash{value}
}
