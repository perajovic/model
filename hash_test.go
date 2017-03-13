package model_fields

import (
	"testing"
)

func TestHashIsCreatedFromString(t *testing.T) {
	h := NewHashFromString("123abc")

	if h.Value != "123abc" {
		t.Errorf("Value is incorrect, actual is %s; expected %s", h.Value, "123abc")
	}
}

func TestHashIsCreated(t *testing.T) {
	h := NewHash(1)
	if len(h.Value) != 36 {
		t.Errorf("Value lenght is incorrect, actual is %d; expected %d", len(h.Value), 36)
	}

	h = NewHash(2)
	if len(h.Value) != 72 {
		t.Errorf("Value lenght is incorrect, actual is %d; expected %d", len(h.Value), 72)
	}

	h = NewHash(3)
	if len(h.Value) != 108 {
		t.Errorf("Value lenght is incorrect, actual is %d; expected %d", len(h.Value), 108)
	}
}
