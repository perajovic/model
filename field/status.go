package field

import "errors"

type Status struct {
	Value        string
	StatusConfig []string
}

func NewStatus(value string, config []string) (*Status, error) {
	s := &Status{value, config}

	if s.IsValid(value) {
		return s, nil
	}

	return nil, errors.New("Invalid status value provided.")
}

func (s *Status) IsValid(status string) bool {
	for _, value := range s.StatusConfig {
		if value == status {
			return true
		}
	}

	return false
}
