package model_fields

import (
	"testing"
	"time"

	"github.com/satori/go.uuid"
)

func createUpdated() *Updated {
	return NewUpdated(
		NewModifier(
			uuid.NewV4(),
			"John",
			"Doe",
			"john@doe.com",
			"user",
			false,
		),
		time.Now(),
	)
}

func TestUpdatedIsCreated(t *testing.T) {
	u := createUpdated()

	t.Logf("Field %T is created.", u)
}
