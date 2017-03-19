package field

import (
	"testing"
	"time"

	"github.com/perajovic/model"
	"github.com/satori/go.uuid"
)

func createUpdated() *Updated {
	return NewUpdated(
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

func TestUpdatedIsCreated(t *testing.T) {
	createUpdated()
}
