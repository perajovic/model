package field

import "errors"

var ErrInvalidStatus = errors.New("Invalid status value provided.")

// Status field which can be embedded into high level models.
type Status struct {
	Value  string
	Config []string
}

// NewStatus creates a new Status.
func NewStatus(value string, config []string) (*Status, error) {
	s := &Status{value, config}

	if s.IsValid(value) == false {
		return nil, ErrInvalidStatus
	}

	return s, nil
}

// IsValid checks if valid value is provided as a Status value.
func (s *Status) IsValid(status string) bool {
	for _, value := range s.Config {
		if value == status {
			return true
		}
	}

	return false
}
