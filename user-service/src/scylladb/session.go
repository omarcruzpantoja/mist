package scylladb

import (
	"context"
	"net/http"
	"os"
	"strconv"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

func InitializeScyllaConnection() *gocql.ClusterConfig {
	// Get the ScyllaDB connection details from the environment variables
	scylladbCluster := gocql.NewCluster(os.Getenv("SCYLLA_HOST"))
	scylladbCluster.Keyspace = os.Getenv("SCYLLA_KEYSPACE")
	scylladbCluster.Port, _ = strconv.Atoi(os.Getenv("SCYLLA_PORT"))

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
