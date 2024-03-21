package scylla

import (
	"context"
	"github.com/scylladb/gocqlx/v2/migrate"
	"log"
	"server/db/scylla/cql"
	"server/db/scylla/scylladb"
)

var queries *scylladb.Queries

func Connect() {
	ctx := context.Background()
	manager := scylladb.NewManager()

	err := manager.CreateKeySPace()
	if err != nil {
		panic(err)
	}

	session, err := manager.Connect()
	if err != nil {
		panic(err)
	}

	err = migrate.FromFS(ctx, session, cql.Files)
	if err != nil {
		log.Fatal("Migrate:", err)
	}
	queries = scylladb.New(session, manager.ScyllaKeyspace)
}
func Queries() *scylladb.Queries {
	return queries
}
