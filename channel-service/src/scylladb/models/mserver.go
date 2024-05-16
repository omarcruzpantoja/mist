package models

import (
	"net/http"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2/table"
)

// metadata specifies table name and columns it must be in sync with schema.
var MServerMetadata = table.Metadata{
	Name:    "mserver",
	Columns: []string{"id", "name"},
	PartKey: []string{"id"},
	SortKey: []string{"name"},
}

// userTable allows for simple CRUD operations based on user_metadata.
var MServerTable = table.New(MServerMetadata)

// Person represents a row in person table.
// Field names are converted to snake case by default, no need to add special tags.
// A field will not be persisted by adding the `db:"-"` tag or making it unexported.

type MServer struct {
	Id   gocql.UUID `json:"id,omitempty" bson:"id,omitempty"`
	Name string     `json:"name,omitempty" bson:"name,omitempty"`
}

func (mistServer *MServer) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

func (mistServer *MServer) Bind(r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
