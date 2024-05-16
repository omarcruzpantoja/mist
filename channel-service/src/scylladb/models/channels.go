package models

import (
	"net/http"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2/table"
)

// metadata specifies table name and columns it must be in sync with schema.
var ChannelMetadata = table.Metadata{
	Name:    "channel",
	Columns: []string{"mserverid", "id", "name"},
	PartKey: []string{"mserverid"},
	SortKey: []string{"id", "name"},
}

// userTable allows for simple CRUD operations based on user_metadata.
var ChannelTable = table.New(ChannelMetadata)

// Person represents a row in person table.
// Field names are converted to snake case by default, no need to add special tags.
// A field will not be persisted by adding the `db:"-"` tag or making it unexported.
type Channel struct {
	Mserverid gocql.UUID `json:"mserverid,omitempty" bson:"mserverid,omitempty"`
	Id        gocql.UUID `json:"id,omitempty" bson:"id,omitempty"`
	Name      string     `json:"name,omitempty" bson:"name,omitempty"`
}

func (channel *Channel) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

func (channel *Channel) Bind(r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
