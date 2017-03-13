package field

import (
	"testing"
	"time"

	"github.com/perajovic/model"
	"github.com/satori/go.uuid"
)

func createCreated() *Created {
	return NewCreated(
		model.NewModifier(
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
