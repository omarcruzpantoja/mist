package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/gocql/gocql"

	"channel-service/src/mappers"
	"channel-service/src/scylladb"
	"channel-service/src/scylladb/models"
)

// A completely separate router for administrator routes
func channelRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", createChannel)
	r.Get("/{id}/server/{serverid}", getChannel)
	r.Get("/", getChannels)
	r.Patch("/{id}/server/{serverid}", patchChannel)
	r.Delete("/{id}/server/{serverid}", deleteChannel)
	return r
}

func createChannel(w http.ResponseWriter, r *http.Request) {
	dbSession := scylladb.GetScyllaSessionFromContext(r)
	channelService := mappers.ChannelService(dbSession, r.Context())

	channel := &models.Channel{}

	// Bind body content to the channel variable
	if err := render.Bind(r, channel); err != nil {
		render.JSON(w, r, ErrInvalidRequest(500, err, "marshalling error"))
		return
	}

	if err := channelService.Create(channel); err != nil {
		render.JSON(w, r, ErrInvalidRequest(400, err, "database error"))
		return
	}

	render.JSON(w, r, channel)
}

func getChannel(w http.ResponseWriter, r *http.Request) {
	dbSession := scylladb.GetScyllaSessionFromContext(r)
	channelService := mappers.BaseChannelService(dbSession, r.Context())

	channelId, _ := gocql.ParseUUID(chi.URLParam(r, "id"))
	serverId, _ := gocql.ParseUUID(chi.URLParam(r, "serverid"))

	channel := &models.Channel{Serverid: serverId, Id: channelId}

	err := channelService.GetByKeys(channel, 2, "id", "serverid")
	if err != nil {
		render.JSON(w, r, ErrInvalidRequest(404, err, "database error"))
		return
	}
	render.JSON(w, r, channel)
}

func getChannels(w http.ResponseWriter, r *http.Request) {
	dbSession := scylladb.GetScyllaSessionFromContext(r)
	var channels []models.Channel

	channelService := mappers.BaseChannelService(dbSession, r.Context())

	channel := &models.Channel{}
	channels, err := channelService.Filter(channel, 0)
	if err != nil {
		render.JSON(w, r, ErrInvalidRequest(400, err, "database error"))
		return
	}
	render.JSON(w, r, channels)
}

func patchChannel(w http.ResponseWriter, r *http.Request) {
	dbSession := scylladb.GetScyllaSessionFromContext(r)
	channelId, _ := gocql.ParseUUID(chi.URLParam(r, "id"))
	serverId, _ := gocql.ParseUUID(chi.URLParam(r, "serverid"))
	channelService := mappers.ChannelService(dbSession, r.Context())

	channel := &models.Channel{Serverid: serverId, Id: channelId}

	// Bind body content to the channel variable
	if err := render.Bind(r, channel); err != nil {
		render.JSON(w, r, ErrInvalidRequest(500, err, "marshalling error"))
		return
	}

	if err := channelService.Patch(channel); err != nil {
		render.JSON(w, r, ErrInvalidRequest(400, err, "database error"))
		return
	}
	render.JSON(w, r, channel)
}

func deleteChannel(w http.ResponseWriter, r *http.Request) {
	dbSession := scylladb.GetScyllaSessionFromContext(r)

	channelId, _ := gocql.ParseUUID(chi.URLParam(r, "id"))
	serverId, _ := gocql.ParseUUID(chi.URLParam(r, "serverid"))
	channel := &models.Channel{Serverid: serverId, Id: channelId}

	channelService := mappers.BaseChannelService(dbSession, r.Context())

	err := channelService.Delete(channel, 2, "id", "serverid")
	if err != nil {
		render.JSON(w, r, ErrInvalidRequest(400, err, "database error"))
		return
	}

	// TODO: Delete all subscriptions to a channel
	w.WriteHeader(http.StatusNoContent) // send the headers with a 204 response code.
}
