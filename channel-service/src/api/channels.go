package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"

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
		fmt.Printf("query error: %v", err)
	}

	render.RenderList(w, r, ChannelListHTTPResponse(channels))
}

func getChannel(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get channel"))
}

func createChannel(w http.ResponseWriter, r *http.Request) {
	db_session := scylladb.GetScyllaSessionFromContext(r)

	u := models.Channel{Id: uuid.New().String(), Channelname: "channel1", Password: "password1"}
	q := db_session.Query(models.ChannelTable.Insert()).BindStruct(u)
	if err := q.ExecRelease(); err != nil {
		fmt.Printf("query error: %v", err)
	}
	render.Render(w, r, ChannelHTTPResponse(&u))
}

// Render helpers

type ChannelResponse struct {
	*models.Channel `json:"channel,omitempty"`

	// We add an additional field to the response here.. such as this
	// elapsed computed property
	Elapsed int64 `json:"elapsed"`
}

func ChannelHTTPResponse(u *models.Channel) *ChannelResponse {
	return &ChannelResponse{Channel: u}
}

func ChannelListHTTPResponse(channels []*models.Channel) []render.Renderer {
	list := []render.Renderer{}
	for _, channel := range channels {
		list = append(list, ChannelHTTPResponse(channel))
	}
	return list
}
