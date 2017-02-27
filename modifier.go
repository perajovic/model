package model_fields

import uuidGenerator "github.com/satori/go.uuid"

type Modifier struct {
	Uuid           uuidGenerator.UUID
	Firstname      string
	Lastname       string
	Email          string
	PolymorficType string
	Deleted        bool
}

func NewModifier(uuid uuidGenerator.UUID, firstname, lastname, email, polymorficType string, deleted bool) *Modifier {
	return &Modifier{
		uuid,
		firstname,
		lastname,
		email,
		polymorficType,
		deleted,
	}
}

func (m *Modifier) MarkAsDeleted() {
	m.Deleted = true
}

func (m *Modifier) Update(firstname, lastname, email string) {
	m.Firstname = firstname
	m.Lastname = lastname
	m.Email = email
}
