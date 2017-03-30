package field

import (
	"testing"
)

func createStatus() *Status {
	status, _ := NewStatus("active", []string{"pending", "active", "inactive"})

	return status
}

func TestInvalidConstructorStatusValue(t *testing.T) {
	status, err := NewStatus("invalid", []string{"pending", "active", "inactive"})

	if status != nil {
		t.Errorf("Value is incorrect, actual is %v; expected %v", status, nil)
	}

	if err.Error() != "Invalid status value provided." {
		t.Errorf("Value is incorrect, actual is %v; expected %v", err.Error(), "Invalid status value provided.")
	}
}

func TestValidStatus(t *testing.T) {
	status := createStatus()

	if status.IsValid("active") == false {
		t.Errorf("Value is incorrect, actual is %v; expected %v", status.IsValid("active"), true)
	}
}

func TestInvalidStatus(t *testing.T) {
	status := createStatus()

	if status.IsValid("invalid") == true {
		t.Errorf("Value is incorrect, actual is %v; expected %v", status.IsValid("invalid"), false)
	}
}

func TestStatusIsNotChanged(t *testing.T) {
	status := createStatus()

	err := status.Change("invalid")

	if err.Error() != "Invalid status value provided." {
		t.Errorf("Value is incorrect, actual is %v; expected %v", err.Error(), "Invalid status value provided.")
	}
}

func TestStatusIsChanged(t *testing.T) {
	status := createStatus()

	err := status.Change("pending")

	if err != nil {
		t.Errorf("Value is incorrect, actual is %v; expected %v", err, nil)
	}
}
