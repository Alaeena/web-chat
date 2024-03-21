package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

type Queries struct {
	db *mongo.Database
}

func New(client *mongo.Client) *Queries {
	dbName := os.Getenv("DB_NAME")
	database := client.Database(dbName)

	return &Queries{db: database}
}
