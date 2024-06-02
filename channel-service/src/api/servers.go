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
func serverRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", createServer)
	r.Get("/{id}", getServer)
	r.Get("/", getServers)
	r.Patch("/{id}", patchServer)
	r.Delete("/{id}", deleteServer)
	return r
}

func createServer(w http.ResponseWriter, r *http.Request) {
	dbSession := scylladb.GetScyllaSessionFromContext(r)
	serverService := mappers.ServerService(dbSession, r.Context())

	server := &models.ServerCreate{}

	// Bind body content to the server variable
	if err := render.Bind(r, server); err != nil {
		render.JSON(w, r, ErrInvalidRequest(500, err, "marshalling error"))
		return
	}

	if err := serverService.Create(server); err != nil {
		render.JSON(w, r, ErrInvalidRequest(400, err, "database error"))
		return
	}

	render.JSON(w, r, server)
}

func getServer(w http.ResponseWriter, r *http.Request) {
	dbSession := scylladb.GetScyllaSessionFromContext(r)
	serverService := mappers.BaseServerService(dbSession, r.Context())

	serverId, _ := gocql.ParseUUID(chi.URLParam(r, "id"))

	server := &models.Server{Id: serverId}

	err := serverService.GetByKeys(server, 1, "id")
	if err != nil {
		render.JSON(w, r, ErrInvalidRequest(404, err, "database error"))
		return
	}
	render.JSON(w, r, server)
}

func getServers(w http.ResponseWriter, r *http.Request) {
	dbSession := scylladb.GetScyllaSessionFromContext(r)
	var servers []models.Server

	serverService := mappers.BaseServerService(dbSession, r.Context())

	server := &models.Server{}
	servers, err := serverService.Filter(server, 0)
	if err != nil {
		render.JSON(w, r, ErrInvalidRequest(400, err, "database error"))
		return
	}
	render.JSON(w, r, servers)
}

func patchServer(w http.ResponseWriter, r *http.Request) {
	dbSession := scylladb.GetScyllaSessionFromContext(r)
	serverId, _ := gocql.ParseUUID(chi.URLParam(r, "id"))
	serverService := mappers.ServerService(dbSession, r.Context())

	server := &models.Server{Id: serverId}

	// Bind body content to the server variable
	if err := render.Bind(r, server); err != nil {
		render.JSON(w, r, ErrInvalidRequest(500, err, "marshalling error"))
		return
	}

	if err := serverService.Patch(server); err != nil {
		render.JSON(w, r, ErrInvalidRequest(400, err, "database error"))
		return
	}
	render.JSON(w, r, server)
}

func deleteServer(w http.ResponseWriter, r *http.Request) {
	dbSession := scylladb.GetScyllaSessionFromContext(r)

	serverId, _ := gocql.ParseUUID(chi.URLParam(r, "id"))
	server := &models.Server{Id: serverId}

	serverService := mappers.BaseServerService(dbSession, r.Context())

	err := serverService.Delete(server, 1, "id")
	if err != nil {
		render.JSON(w, r, ErrInvalidRequest(400, err, "database error"))
		return
	}

	// TODO: Delete all subscriptions to a server
	// TODO: Delete all channels
	w.WriteHeader(http.StatusNoContent) // send the headers with a 204 response code.
}
