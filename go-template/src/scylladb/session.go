package scylladb

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

func InitializeScyllaConnection() *gocql.ClusterConfig {
	// Get the ScyllaDB connection details from the environment variables
	scyllaAddress := fmt.Sprintf("%s:%s", os.Getenv("SCYLLA_HOST"), os.Getenv("SCYLLA_PORT"))
	scylladbCluster := gocql.NewCluster(scyllaAddress)
	scylladbCluster.Keyspace = os.Getenv("SCYLLA_KEYSPACE")
	return scylladbCluster
}

type ScyllaDBSessionKey string

func SetScyllaSession(dbSession *gocqlx.Session) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, ScyllaDBSessionKey("scylladb_session"), dbSession)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}

}

func GetScyllaSessionFromContext(r *http.Request) *gocqlx.Session {
	return r.Context().Value(ScyllaDBSessionKey("scylladb_session")).(*gocqlx.Session)
}
