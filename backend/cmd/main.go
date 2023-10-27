package main

import (
	"log"

	"github.com/markraiter/chat/internal/storage/postgres"
)

func main() {
	_, err := postgres.NewDB()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s\n", err.Error())
	}
}
