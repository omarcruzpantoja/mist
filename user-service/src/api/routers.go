package api

import (
	"fmt"
	"net/http"
	"time"

	"user-service/src/scylladb"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/docgen"
	"github.com/scylladb/gocqlx/v2"
)

func StartService() {

	scyllaCluster := scylladb.InitializeScyllaConnection()
	scyllaSession, err := scyllaCluster.CreateSession()

	// var scyllaSession gocqlx.Session
	for scyllaSession == nil {
		fmt.Println("ScyllaDB connection failed, retrying...")
		time.Sleep(3 * time.Second)
		scyllaSession, err = scyllaCluster.CreateSession()
	}

	fmt.Println("Connection to ScyllaDB completed!")

	gocqlxScyllaSession, _ := gocqlx.WrapSession(scyllaSession, err)
	defer scyllaSession.Close()

	r := chi.NewRouter()
	// SETUP MIDDDLEWARES
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(scylladb.SetScyllaSession(&gocqlxScyllaSession))

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
