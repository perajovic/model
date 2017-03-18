package field

import (
	"testing"
)

func createStatus() *Status {
	s, _ := NewStatus("active", []string{"pending", "active", "inactive"})

	return s
}

func TestInvalidConstructorStatusValue(t *testing.T) {
	s, err := NewStatus("invalid", []string{"pending", "active", "inactive"})

	if s != nil {
		t.Errorf("Value is incorrect, actual is %v; expected %v", s, nil)
	}

	if err.Error() != "Invalid status value provided." {
		t.Errorf("Value is incorrect, actual is %v; expected %v", err.Error(), "Invalid status value provided.")
	}
}

func TestValidStatus(t *testing.T) {
	s := createStatus()

	if s.IsValid("active") == false {
		t.Errorf("Value is incorrect, actual is %v; expected %v", s.IsValid("active"), true)
	}
}

func TestInvalidStatus(t *testing.T) {
	s := createStatus()

	if s.IsValid("invalid") == true {
		t.Errorf("Value is incorrect, actual is %v; expected %v", s.IsValid("invalid"), false)
	}
}

func TestStatusIsNotChanged(t *testing.T) {
	s := createStatus()

	err := s.Change("invalid")

	if err.Error() != "Invalid status value provided." {
		t.Errorf("Value is incorrect, actual is %v; expected %v", err.Error(), "Invalid status value provided.")
	}
}

func TestStatusIsChanged(t *testing.T) {
	s := createStatus()

	err := s.Change("pending")

	if err != nil {
		t.Errorf("Value is incorrect, actual is %v; expected %v", err, nil)
	}
}
