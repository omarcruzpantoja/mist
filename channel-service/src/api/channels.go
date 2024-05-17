package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2/qb"

	"channel-service/src/scylladb"
	"channel-service/src/scylladb/models"
)

// A completely separate router for administrator routes
func channelRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", createChannel)
	r.Get("/{id}/server/{serverId}", getChannel)
	r.Get("/", getChannels)
	r.Patch("/{id}", patchChannel)
	r.Delete("/{id}/server/{serverId}", deleteChannel)
	return r
}

func createChannel(w http.ResponseWriter, r *http.Request) {
	dbSession := scylladb.GetScyllaSessionFromContext(r)
	channel := &models.Channel{}

	// Bind body content to the channel variable
	if err := render.Bind(r, channel); err != nil {
		render.Render(w, r, ErrInvalidRequest(err, "marshalling error"))
		return
	}
	newId, _ := gocql.RandomUUID()
	channel.Id = newId
	q := dbSession.Query(models.ChannelTable.Insert()).BindStruct(channel)
	if err := q.ExecRelease(); err != nil {
		render.Render(w, r, ErrInvalidRequest(err, "database error"))
		return
	}
	render.Render(w, r, ChannelHTTPResponse(channel))
}

func getChannel(w http.ResponseWriter, r *http.Request) {
	dbSession := *scylladb.GetScyllaSessionFromContext(r)
	channelId, _ := gocql.ParseUUID(chi.URLParam(r, "id"))
	serverId, _ := gocql.ParseUUID(chi.URLParam(r, "serverId"))

	channel := &models.Channel{Mserverid: serverId, Id: channelId}

	q := dbSession.Query(
		qb.Select(models.ChannelTable.Name()).Where(qb.Eq("id"), qb.Eq("mserverid")).ToCql(),
	).BindStruct(channel)

	if err := q.GetRelease(channel); err != nil {
		render.Render(w, r, ErrInvalidRequest(err, "database error"))
		return
	}
	render.Render(w, r, ChannelHTTPResponse(channel))
}

func getChannels(w http.ResponseWriter, r *http.Request) {
	dbSession := scylladb.GetScyllaSessionFromContext(r)
	var channels []*models.Channel
	q := dbSession.Query(models.ChannelTable.SelectAll())
	if err := q.SelectRelease(&channels); err != nil {
		render.Render(w, r, ErrInvalidRequest(err, "database error"))
		return
	}

	render.RenderList(w, r, ChannelListHTTPResponse(channels))
}

func patchChannel(w http.ResponseWriter, r *http.Request) {
	dbSession := *scylladb.GetScyllaSessionFromContext(r)
	channelId, _ := gocql.ParseUUID(chi.URLParam(r, "id"))

	channel := &models.Channel{Id: channelId}

	// Check if channel exists
	q := dbSession.Query(
		qb.Select(models.ChannelTable.Name()).Where(qb.Eq("id")).ToCql(),
	).BindStruct(channel)

	if err := q.GetRelease(channel); err != nil {
		render.Render(w, r, ErrInvalidRequest(err, "database error"))
		return
	}

	// Confirmed channel exists, continue with the update

	// Bind body content to the channel variable
	if err := render.Bind(r, channel); err != nil {
		render.Render(w, r, ErrInvalidRequest(err, "marshalling error"))
		return
	}

	q = dbSession.Query(
		qb.Update(models.ChannelTable.Name()).Set("name").Where(qb.Eq("id"), qb.Eq("mserverid")).ToCql(),
	).BindStruct(channel)
	if err := q.ExecRelease(); err != nil {
		render.Render(w, r, ErrInvalidRequest(err, "database error"))
	}
	render.Render(w, r, ChannelHTTPResponse(channel))
}

func deleteChannel(w http.ResponseWriter, r *http.Request) {
	dbSession := *scylladb.GetScyllaSessionFromContext(r)
	channelId, _ := gocql.ParseUUID(chi.URLParam(r, "id"))
	serverId, _ := gocql.ParseUUID(chi.URLParam(r, "serverId"))

	channel := &models.Channel{Mserverid: serverId, Id: channelId}
	// Check if channel exists
	q := dbSession.Query(
		qb.Select(models.ChannelTable.Name()).Where(qb.Eq("id"), qb.Eq("mserverid")).ToCql(),
	).BindStruct(channel)

	if err := q.GetRelease(channel); err != nil {
		render.Render(w, r, ErrInvalidRequest(err, "database error"))
		return
	}

	// Confirmed channel exists, continue with the delete
	q = dbSession.Query(
		qb.Delete(models.ChannelTable.Name()).Where(qb.Eq("id"), qb.Eq("mserverid")).ToCql(),
	).BindStruct(channel)

	if err := q.ExecRelease(); err != nil {
		render.Render(w, r, ErrInvalidRequest(err, "database error"))
		return
	}

	// TODO: Delete all subscriptions to a channel

	w.WriteHeader(http.StatusNoContent) // send the headers with a 204 response code.
}

// Schemas
type ChannelResponse struct {
	*models.Channel `json:"channel,omitempty"`

	// We add an additional field to the response here.. such as this
	// elapsed computed property
	Elapsed int64 `json:"elapsed"`
}

// Marshalers
func ChannelHTTPResponse(u *models.Channel) *ChannelResponse {
	return &ChannelResponse{Channel: u}
}

func ChannelListHTTPResponse(channels []*models.Channel) []render.Renderer {
	list := []render.Renderer{}
	for _, server := range channels {
		list = append(list, ChannelHTTPResponse(server))
	}
	return list
}
