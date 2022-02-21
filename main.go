package main

import (
	"log"

	"github.com/SajjadManafi/simple-uber/api"
	"github.com/SajjadManafi/simple-uber/internal/config"
	"github.com/SajjadManafi/simple-uber/internal/store"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalln("cannot load config:", err)
	}

	store, err := store.NewPostgresStore(config)
	if err != nil {
		log.Fatalln("cannot connect to db:", err)
	}

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatalln("cannot create server:", err)
	}

	go server.ListenToWsChannel()

	err = server.Start()
	if err != nil {
		log.Fatalln("cannot start server:", err)
	}

}
