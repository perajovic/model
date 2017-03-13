package field

import (
	"time"

	"github.com/perajovic/model"
)

// Created field which can be embedded into high level models.
type Created struct {
	By *model.Modifier
	At time.Time
}

// NewCreated creates a new Created.
func NewCreated(by *model.Modifier, at time.Time) *Created {
	return &Created{by, at}
}
