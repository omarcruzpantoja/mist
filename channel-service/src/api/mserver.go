package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/gocql/gocql"

	"channel-service/src/scylladb"
	"channel-service/src/scylladb/models"
)

// A completely separate router for administrator routes
func mserverRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", getMServers)
	r.Get("/{id}", getMServer)
	r.Post("/", createMServer)
	return r
}

func getMServers(w http.ResponseWriter, r *http.Request) {
	db_session := scylladb.GetScyllaSessionFromContext(r)
	var servers []*models.MServer
	q := db_session.Query(models.MServerTable.SelectAll())
	if err := q.SelectRelease(&servers); err != nil {
		render.Render(w, r, ErrInvalidRequest(err, "database error"))
	}

	render.RenderList(w, r, MServerListHTTPResponse(servers))
}

func getMServer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get server"))
}

func createMServer(w http.ResponseWriter, r *http.Request) {
	db_session := scylladb.GetScyllaSessionFromContext(r)
	mserver := &models.MServer{}
	if err := render.Bind(r, mserver); err != nil {
		render.Render(w, r, ErrInvalidRequest(err, "marshalling error"))
		return
	}
	newId, _ := gocql.RandomUUID()
	mserver.Id = newId
	q := db_session.Query(models.MServerTable.Insert()).BindStruct(mserver)
	if err := q.ExecRelease(); err != nil {
		render.Render(w, r, ErrInvalidRequest(err, "database error"))
	}
	render.Render(w, r, MServerHTTPResponse(mserver))
}

// Schemas
// type MServerRequestCreate struct {
// 	Name string `json:"name,omitempty" bson:"name,omitempty"`
// }

type MServerResponse struct {
	*models.MServer `json:"mserver,omitempty"`

	// We add an additional field to the response here.. such as this
	// elapsed computed property
	Elapsed int64 `json:"elapsed"`
}

// Marshalers
func MServerHTTPResponse(u *models.MServer) *MServerResponse {
	return &MServerResponse{MServer: u}
}

func MServerListHTTPResponse(servers []*models.MServer) []render.Renderer {
	list := []render.Renderer{}
	for _, server := range servers {
		list = append(list, MServerHTTPResponse(server))
	}
	return list
}

// Todo: moved somewhere else
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func ErrInvalidRequest(err error, message string) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     fmt.Sprintf("Invalid request: %s", message),
		ErrorText:      err.Error(),
	}
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}
