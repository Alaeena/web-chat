package main

import (
	"github.com/joho/godotenv"
	"server/db"
	"server/router"
)

func main() {
	godotenv.Load()

	db.Connect()
	router.Listen()
}
