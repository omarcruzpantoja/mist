package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"

	"go-template/src/scylladb"
	"go-template/src/scylladb/models"
)

// A completely separate router for administrator routes
func userRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", getUsers)
	r.Get("/{id}", getUser)
	r.Post("/", createUser)
	return r
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	db_session := scylladb.GetScyllaSessionFromContext(r)
	var users []*models.User
	q := db_session.Query(models.UserTable.SelectAll())
	if err := q.SelectRelease(&users); err != nil {
		fmt.Printf("query error: %v", err)
	}

	render.RenderList(w, r, UserListHTTPResponse(users))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user"))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	db_session := scylladb.GetScyllaSessionFromContext(r)

	u := models.User{Id: uuid.New().String(), Username: "user1", Password: "password1"}
	q := db_session.Query(models.UserTable.Insert()).BindStruct(u)
	if err := q.ExecRelease(); err != nil {
		fmt.Printf("query error: %v", err)
	}
	render.Render(w, r, UserHTTPResponse(&u))
}

// Render helpers

type UserResponse struct {
	*models.User `json:"user,omitempty"`

	// We add an additional field to the response here.. such as this
	// elapsed computed property
	Elapsed int64 `json:"elapsed"`
}

func UserHTTPResponse(u *models.User) *UserResponse {
	return &UserResponse{User: u}
}

func UserListHTTPResponse(users []*models.User) []render.Renderer {
	list := []render.Renderer{}
	for _, user := range users {
		list = append(list, UserHTTPResponse(user))
	}
	return list
}
