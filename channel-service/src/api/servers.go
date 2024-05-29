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
func serverRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", getServers)
	r.Get("/{id}", getServer)
	r.Post("/", createServer)
	return r
}

func getServers(w http.ResponseWriter, r *http.Request) {
	dbSession := scylladb.GetScyllaSessionFromContext(r)
	var servers []*models.Server
	q := dbSession.Query(models.ServerTable.SelectAll())
	if err := q.SelectRelease(&servers); err != nil {
		render.Render(w, r, ErrInvalidRequest(400, err, "database error"))
	}

	render.JSON(w, r, servers)
}

func getServer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get server"))
}

func createServer(w http.ResponseWriter, r *http.Request) {
	dbSession := scylladb.GetScyllaSessionFromContext(r)
	server := &models.Server{}
	if err := render.Bind(r, server); err != nil {
		render.JSON(w, r, ErrInvalidRequest(500, err, "marshalling error"))
		return
	}
	newId, _ := gocql.RandomUUID()
	server.Id = newId
	q := dbSession.Query(models.ServerTable.Insert()).BindStruct(server)
	if err := q.ExecRelease(); err != nil {
		render.JSON(w, r, ErrInvalidRequest(400, err, "database error"))
	}
	render.JSON(w, r, server)
}
