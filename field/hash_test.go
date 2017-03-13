package field

import (
	"testing"
)

func TestHashIsCreated(t *testing.T) {
	h := NewHash("123abc")

	if h.Value != "123abc" {
		t.Errorf("Value is incorrect, actual is %s; expected %s", h.Value, "123abc")
	}
}

func TestRandomHashIsCreated(t *testing.T) {
	h := NewRandomHash(1)
	if len(h.Value) != 36 {
		t.Errorf("Value lenght is incorrect, actual is %d; expected %d", len(h.Value), 36)
	}

	h = NewRandomHash(2)
	if len(h.Value) != 72 {
		t.Errorf("Value lenght is incorrect, actual is %d; expected %d", len(h.Value), 72)
	}

	h = NewRandomHash(3)
	if len(h.Value) != 108 {
		t.Errorf("Value lenght is incorrect, actual is %d; expected %d", len(h.Value), 108)
	}
}
