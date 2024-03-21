package main

import (
	"github.com/joho/godotenv"
	"server/db/mongo"
	"server/db/scylla"
	"server/router"
)

func main() {
	godotenv.Load()

	//postgres.Connect()
	scylla.Connect()
	mongo.Connect()

	router.Listen()
}
