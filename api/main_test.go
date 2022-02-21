package api

import (
	"log"
	"os"
	"testing"

	"github.com/SajjadManafi/simple-uber/internal/config"
	"github.com/SajjadManafi/simple-uber/internal/store"
)

var TestServer *Server

func TestMain(m *testing.M) {
	config, err := config.LoadConfig("../")
	if err != nil {
		log.Fatalln("cannot load config:", err)
	}

	store, err := store.NewPostgresStore(config)
	if err != nil {
		log.Fatalln("cannot connect to db:", err)
	}

	TestServer, err = NewServer(config, store)
	if err != nil {
		log.Fatalln("cannot create server:", err)
	}

	// gin.SetMode(gin.TestMode)
	os.Exit(m.Run())

}
