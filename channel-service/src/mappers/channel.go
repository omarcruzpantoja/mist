package mappers

import (
	"context"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"

	"channel-service/src/scylladb/models"
)

func BaseChannelService(dbSession *gocqlx.Session, ctx context.Context) *ScyllaDBMapper[models.Channel] {
	return &ScyllaDBMapper[models.Channel]{dbSession, ctx, models.ChannelTable.Name()}
}

// ** scylladb repo definition **
type ChannelScyllaMapper struct {
	dbSession *gocqlx.Session
	ctx       context.Context
}

func ChannelService(dbSession *gocqlx.Session, ctx context.Context) *ChannelScyllaMapper {
	return &ChannelScyllaMapper{dbSession, ctx}
}

func (m *ChannelScyllaMapper) Create(item *models.Channel) error {

	newId, _ := gocql.RandomUUID()
	item.Id = newId
	currentTime := time.Now().UTC()
	item.CreatedAt = currentTime
	item.UpdatedAt = currentTime
	// TODO: Add check that server exists
	q := m.dbSession.Query(models.ChannelTable.Insert()).BindStruct(item)
	return q.ExecRelease()
}

func (m *ChannelScyllaMapper) Patch(item *models.Channel) error {

	channelService := BaseChannelService(m.dbSession, m.ctx)
	resource := &models.Channel{Id: item.Id, ServerId: item.ServerId}
	err := channelService.GetByKeys(resource, 2, "id", "server_id")
	if err != nil {
		return err
	}

	item.UpdatedAt = time.Now().UTC()
	item.CreatedAt = resource.CreatedAt

	q := m.dbSession.Query(
		qb.
			Update(models.ChannelTable.Name()).
			Set("name").
			Where(qb.Eq("id"), qb.Eq("server_id")).ToCql(),
	).BindStruct(item)

	err = q.ExecRelease()
	return err
}
