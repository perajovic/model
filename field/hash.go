package field

import (
	"github.com/satori/go.uuid"
)

// Hash field which can be embedded into high level models.
// It represents some random unique value.
type Hash struct {
	Value string
}

// NewHash creates a new Hash from the string.
func NewHash(value string) *Hash {
	return &Hash{value}
}

// NewHash creates a new Hash. A length of Hash.Value depends on the iterations number.
// If iteration is equal to 1 resulting Hash.Value will be 36. If it is 2, resulting Hash.Value will be 72, end so on.
func NewRandomHash(iterations int) *Hash {
	var value string

	for i := 0; i < iterations; i++ {
		value += uuid.NewV4().String()
	}

	return NewHash(value)
}
