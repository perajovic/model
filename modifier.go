package model

import uuidGenerator "github.com/satori/go.uuid"

// Modifier model is a high level model. It should represent a Model object which modify other objects.
// For Example: User & Contact objects are model object. Those can modify values of Ticket object.
// At every moment, you will have linked those infromations -- that Ticket object is modified by Modifier object.
type Modifier struct {
	Uuid           uuidGenerator.UUID
	Firstname      string
	Lastname       string
	Email          string
	PolymorficType string
	Deleted        bool
}

// NewModifier creates a new Modifier from the string.
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

// MarkAsDeleted mark Modifier as deleted.
func (m *Modifier) MarkAsDeleted() {
	m.Deleted = true
}

// Update updates some of Modifier fields.
func (m *Modifier) Update(firstname, lastname, email string) {
	m.Firstname = firstname
	m.Lastname = lastname
	m.Email = email
}
