package field

import (
	"testing"
)

func TestHashIsCreated(t *testing.T) {
	hash := NewHash("123abc")

	if hash.Value != "123abc" {
		t.Errorf("Value is incorrect, actual is %s; expected %s", hash.Value, "123abc")
	}
}

func TestRandomHashIsCreated(t *testing.T) {
	hash := NewRandomHash(1)
	if len(hash.Value) != 36 {
		t.Errorf("Value lenght is incorrect, actual is %d; expected %d", len(hash.Value), 36)
	}

	hash = NewRandomHash(2)
	if len(hash.Value) != 72 {
		t.Errorf("Value lenght is incorrect, actual is %d; expected %d", len(hash.Value), 72)
	}

	hash = NewRandomHash(3)
	if len(hash.Value) != 108 {
		t.Errorf("Value lenght is incorrect, actual is %d; expected %d", len(hash.Value), 108)
	}
}
