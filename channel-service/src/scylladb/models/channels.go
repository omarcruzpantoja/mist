package models

import (
	"net/http"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2/table"
)

// metadata specifies table name and columns it must be in sync with schema.
var ChannelMetadata = table.Metadata{
	Name:    "channel",
	Columns: []string{"server_id", "id", "name", "created_at", "updated_at"},
	PartKey: []string{"server_id"},
	SortKey: []string{"id", "name"},
}

var ChannelTable = table.New(ChannelMetadata)

// Schemas

type Channel struct {
	ServerId  gocql.UUID `json:"server_id,omitempty"`
	Id        gocql.UUID `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
}

/***** HELPERS *****/

// TODO: add in here binder for CRUD objets
func (channel *Channel) Bind(r *http.Request) error {
	// Marsalling payload into a channel object
	return nil
}
