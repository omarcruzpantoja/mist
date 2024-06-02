package models

import (
	"net/http"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2/table"
)

// metadata specifies table name and columns it must be in sync with schema.
var ServerMetadata = table.Metadata{
	Name:    "server",
	Columns: []string{"id", "name", "created_at", "updated_at", "deleted_at"},
	PartKey: []string{"id"},
	SortKey: []string{"name"},
}

// userTable allows for simple CRUD operations based on user_metadata.
var ServerTable = table.New(ServerMetadata)

// Person represents a row in person table.
// Field names are converted to snake case by default, no need to add special tags.
// A field will not be persisted by adding the `db:"-"` tag or making it unexported.

type Server struct {
	Id        gocql.UUID `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt time.Time  `json:"deleted_at,omitempty"`
}

type ServerCreate struct {
	Id        gocql.UUID `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt time.Time  `json:"deleted_at,omitempty"`
}

func (mistServer *Server) Bind(r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

func (mistServer *ServerCreate) Bind(r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
