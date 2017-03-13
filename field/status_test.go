package field

import (
	"testing"
)

func createNewStatus() *Status {
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

func TestStatusIsValid(t *testing.T) {
	s := createNewStatus()

	if s.IsValid("active") == false {
		t.Errorf("Value is incorrect, actual is %v; expected %v", s.IsValid("active"), true)
	}
}

func TestStatusIsNotValid(t *testing.T) {
	s := createNewStatus()

	if s.IsValid("invalid") == true {
		t.Errorf("Value is incorrect, actual is %v; expected %v", s.IsValid("invalid"), false)
	}
}
