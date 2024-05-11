package api

import (
	"net/http"

	"user-service/src/scylladb"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/docgen"
	"github.com/scylladb/gocqlx/v2"
)

func StartService() {

	scyllaCluster := scylladb.InitializeScyllaConnection()
	scyllaSession, _ := gocqlx.WrapSession(scyllaCluster.CreateSession())
	defer scyllaSession.Close()

	r := chi.NewRouter()
	// SETUP MIDDDLEWARES
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(scylladb.SetScyllaSession(&scyllaSession))

	r.Get("/", func(w http.ResponseWriter, request *http.Request) {
		w.Write([]byte("welcome"))
	})

	// Mount the user router
	r.Mount("/users", userRouter())

	r.Get("/routes.json", func(w http.ResponseWriter, request *http.Request) {
		w.Write([]byte(docgen.JSONRoutesDoc(r)))
	})
	http.ListenAndServe(":3000", r)
}
