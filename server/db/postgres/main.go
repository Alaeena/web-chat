package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"server/db/postgres/postgresdb"
)

var queries *postgresdb.Queries

func Connect() {
	connStr := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	queries = postgresdb.New(db)
}

func Queries() *postgresdb.Queries {
	return queries
}
