package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/gocql/gocql"

	"channel-service/src/scylladb"
	"channel-service/src/scylladb/models"
)

// A completely separate router for administrator routes
func channelRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", getChannels)
	r.Get("/{id}", getChannel)
	r.Post("/", createChannel)
	return r
}

func getChannels(w http.ResponseWriter, r *http.Request) {
	db_session := scylladb.GetScyllaSessionFromContext(r)
	var channels []*models.Channel
	q := db_session.Query(models.ChannelTable.SelectAll())
	if err := q.SelectRelease(&channels); err != nil {
		render.Render(w, r, ErrInvalidRequest(err, "database error"))
	}

	render.RenderList(w, r, ChannelListHTTPResponse(channels))
}

func getChannel(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get channel"))
}

func createChannel(w http.ResponseWriter, r *http.Request) {
	db_session := scylladb.GetScyllaSessionFromContext(r)
	channel := &models.Channel{}
	if err := render.Bind(r, channel); err != nil {
		render.Render(w, r, ErrInvalidRequest(err, "marshalling error"))
		return
	}
	newId, _ := gocql.RandomUUID()
	channel.Id = newId
	q := db_session.Query(models.ChannelTable.Insert()).BindStruct(channel)
	if err := q.ExecRelease(); err != nil {
		render.Render(w, r, ErrInvalidRequest(err, "database error"))
	}
	render.Render(w, r, ChannelHTTPResponse(channel))
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
