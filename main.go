package main

import (
	"log"

	"github.com/SajjadManafi/simple-uber/api"
	"github.com/SajjadManafi/simple-uber/internal/store"
)

const (
	dbSource      = "postgres://root:password@localhost:5432/simple_uber?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	store, err := store.NewPostgresStore(dbSource)
	if err != nil {
		log.Fatalln("cannot connect to db:", err)
	}

	server, err := api.NewServer(store)
	if err != nil {
		log.Fatalln("cannot create server:", err)
	}

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatalln("cannot start server:", err)
	}

}
