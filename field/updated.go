package field

import (
	"time"

	"github.com/perajovic/model"
)

// Updated field which can be embedded into high level models.
type Updated struct {
	By *model.Modifier
	At time.Time
}

// NewUpdated creates a new Updated.
func NewUpdated(by *model.Modifier, at time.Time) *Updated {
	return &Updated{by, at}
}
