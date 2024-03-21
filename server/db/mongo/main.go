package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"server/db/mongo/mongodb"
)

var queries *mongodb.Queries

func Connect() {
	connStr := os.Getenv("DB_URL")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connStr))
	if err != nil {
		log.Fatal(err)
	}
	queries = mongodb.New(client)
	queries.InitUsers()
}

func Queries() *mongodb.Queries {
	return queries
}
