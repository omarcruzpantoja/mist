package mappers

import (
	"context"

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

	// TODO: Add check that server exists
	q := m.dbSession.Query(models.ChannelTable.Insert()).BindStruct(item)
	return q.ExecRelease()
}

func (m *ChannelScyllaMapper) Patch(item *models.Channel) error {

	channelService := BaseChannelService(m.dbSession, m.ctx)
	resource := &models.Channel{Id: item.Id, Serverid: item.Serverid}
	err := channelService.GetByKeys(resource, 2, "id", "serverid")
	if err != nil {
		return err
	}

	q := m.dbSession.Query(
		qb.
			Update(models.ChannelTable.Name()).
			Set("name").
			Where(qb.Eq("id"), qb.Eq("serverid")).ToCql(),
	).BindStruct(item)

	err = q.ExecRelease()
	return err
}
