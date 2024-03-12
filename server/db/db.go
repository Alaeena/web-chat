package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"server/db/database"
)

var client *database.Queries

func Connect() {
	connStr := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	client = database.New(db)
}

func Client() *database.Queries {
	return client
}
