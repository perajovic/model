package model_fields

import "time"

type Updated struct {
	By *Modifier
	At time.Time
}

func NewUpdated(by *Modifier, at time.Time) *Updated {
	return &Updated{by, at}
}
