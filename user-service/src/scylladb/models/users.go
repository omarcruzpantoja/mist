package models

import (
	"net/http"

	"github.com/scylladb/gocqlx/v2/table"
)

// metadata specifies table name and columns it must be in sync with schema.
var UserMetadata = table.Metadata{
	Name:    "user",
	Columns: []string{"id", "username", "password"},
	PartKey: []string{"id"},
	SortKey: []string{"username"},
}

// userTable allows for simple CRUD operations based on user_metadata.
var UserTable = table.New(UserMetadata)

// Person represents a row in person table.
// Field names are converted to snake case by default, no need to add special tags.
// A field will not be persisted by adding the `db:"-"` tag or making it unexported.
type User struct {
	Id       string
	Username string
	Password string
}

func (user *User) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
