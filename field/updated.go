package field

import (
	"time"

	"github.com/perajovic/model"
)

type Updated struct {
	By *model.Modifier
	At time.Time
}

func NewUpdated(by *model.Modifier, at time.Time) *Updated {
	return &Updated{by, at}
}
