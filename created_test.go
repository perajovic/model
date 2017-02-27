package model_fields

import (
	"testing"
	"time"

	"github.com/satori/go.uuid"
)

func createCreated() *Created {
	return NewCreated(
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

func TestCreatedIsCreated(t *testing.T) {
	c := createCreated()

	t.Logf("Field %T is created.", c)
}
