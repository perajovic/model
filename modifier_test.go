package model

import (
	"testing"

	"github.com/satori/go.uuid"
)

func createModifier() *Modifier {
	return NewModifier(
		uuid.NewV4(),
		"John",
		"Doe",
		"john@doe.com",
		"user",
		false,
	)
}

func TestModifierIsCreated(t *testing.T) {
	modifier := createModifier()

	if len(modifier.Uuid.Bytes()) != 16 {
		t.Errorf("Uuid is incorrect, actual is %d; expected %d", len(modifier.Uuid.Bytes()), 16)
	}

	if modifier.Firstname != "John" {
		t.Errorf("Firstname is incorrect, actual is %s; expected %s", modifier.Firstname, "John")
	}

	if modifier.Lastname != "Doe" {
		t.Errorf("Lastname is incorrect, actual is %s; expected %s", modifier.Lastname, "Doe")
	}

	if modifier.Email != "john@doe.com" {
		t.Errorf("Email is incorrect, actual is %s; expected %s", modifier.Email, "john@doe.com")
	}

	if modifier.PolymorficType != "user" {
		t.Errorf("PolymorficType is incorrect, actual is %s; expected %s", modifier.PolymorficType, "user")
	}

	if modifier.Deleted {
		t.Errorf("Deleted is incorrect, actual is %t; expected %t", modifier.Deleted, false)
	}
}

func TestModifierIsMarkedAsDeleted(t *testing.T) {
	modifier := createModifier()

	modifier.MarkAsDeleted()

	if modifier.Deleted == false {
		t.Errorf("Deleted is incorrect, actual is %t; expected %t", modifier.Deleted, true)
	}
}

func TestModifierIsUpdated(t *testing.T) {
	modifier := createModifier()

	modifier.Update("John1", "Doe1", "john1@doe.com")

	if modifier.Firstname != "John1" {
		t.Errorf("Firstname is incorrect, actual is %s; expected %s", modifier.Firstname, "John1")
	}

	if modifier.Lastname != "Doe1" {
		t.Errorf("Lastname is incorrect, actual is %s; expected %s", modifier.Lastname, "Doe1")
	}

	if modifier.Email != "john1@doe.com" {
		t.Errorf("Email is incorrect, actual is %s; expected %s", modifier.Email, "john1@doe.com")
	}
}
