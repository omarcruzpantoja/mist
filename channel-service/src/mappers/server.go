package mappers

import (
	"context"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"

	"channel-service/src/scylladb/models"
)

func BaseServerService(dbSession *gocqlx.Session, ctx context.Context) *ScyllaDBMapper[models.Server] {
	return &ScyllaDBMapper[models.Server]{dbSession, ctx, models.ServerTable.Name()}
}

// ** scylladb repo definition **
type ServerScyllaMapper struct {
	dbSession *gocqlx.Session
	ctx       context.Context
}

func ServerService(dbSession *gocqlx.Session, ctx context.Context) *ServerScyllaMapper {
	return &ServerScyllaMapper{dbSession, ctx}
}

func (m *ServerScyllaMapper) Create(item *models.ServerCreate) error {

	newId, _ := gocql.RandomUUID()
	item.Id = newId
	currentTime := time.Now().UTC()
	item.CreatedAt = currentTime
	item.UpdatedAt = currentTime

	// TODO: Add check that server exists
	q := m.dbSession.Query(models.ServerTable.Insert()).BindStruct(item)
	return q.ExecRelease()
}

func (m *ServerScyllaMapper) Patch(item *models.Server) error {

	serverService := BaseServerService(m.dbSession, m.ctx)
	resource := &models.Server{Id: item.Id}
	err := serverService.GetByKeys(resource, 1, "id")
	if err != nil {
		return err
	}

	item.UpdatedAt = time.Now().UTC()
	item.CreatedAt = resource.CreatedAt
	q := m.dbSession.Query(
		qb.
			Update(models.ServerTable.Name()).
			Set("name", "updated_at").
			Where(qb.Eq("id")).ToCql(),
	).BindStruct(item)

	err = q.ExecRelease()
	return err
}
