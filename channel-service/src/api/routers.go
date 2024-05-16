package api

import (
	"fmt"
	"net/http"
	"time"

	"channel-service/src/scylladb"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/docgen"
	"github.com/go-chi/render"
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
		w.Write([]byte("healthy"))
	})

	// Mount the user router
	r.Mount("/channels", channelRouter())
	r.Mount("/mservers", mserverRouter())

	r.Get("/routes.json", func(w http.ResponseWriter, request *http.Request) {
		w.Write([]byte(docgen.JSONRoutesDoc(r)))
	})
	http.ListenAndServe(":3000", r)
}

// ErrResponse renders an error response
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
