package model_fields

import "time"

type Created struct {
	By *Modifier
	At time.Time
}

func NewCreated(by *Modifier, at time.Time) *Created {
	return &Created{by, at}
}
