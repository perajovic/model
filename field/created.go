package field

import (
	"time"

	"github.com/perajovic/model"
)

type Created struct {
	By *model.Modifier
	At time.Time
}

func NewCreated(by *model.Modifier, at time.Time) *Created {
	return &Created{by, at}
}
