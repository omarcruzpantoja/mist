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

var ChannelTable = table.New(ChannelMetadata)

type Channel struct {
	Mserverid gocql.UUID `json:"mserverid,omitempty" bson:"mserverid,omitempty"`
	Id        gocql.UUID `json:"id,omitempty" bson:"id,omitempty"`
	Name      string     `json:"name,omitempty" bson:"name,omitempty"`
}

/***** HELPERS *****/

// TODO: add in here binder for CRUD objets
func (channel *Channel) Bind(r *http.Request) error {
	// Marsalling payload into a channel object
	return nil
}
