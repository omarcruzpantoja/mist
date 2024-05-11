package models

import (
	"net/http"

	"github.com/scylladb/gocqlx/v2/table"
)

// metadata specifies table name and columns it must be in sync with schema.
var ChannelMetadata = table.Metadata{
	Name:    "user",
	Columns: []string{"id", "username", "password"},
	PartKey: []string{"id"},
	SortKey: []string{"username"},
}

// userTable allows for simple CRUD operations based on user_metadata.
var ChannelTable = table.New(ChannelMetadata)

// Person represents a row in person table.
// Field names are converted to snake case by default, no need to add special tags.
// A field will not be persisted by adding the `db:"-"` tag or making it unexported.
type Channel struct {
	Id          string
	Channelname string
	Password    string
}

func (user *Channel) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
