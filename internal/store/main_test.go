package store

import (
	"log"
	"os"
	"testing"

	"github.com/SajjadManafi/simple-uber/contract"
	"github.com/SajjadManafi/simple-uber/internal/config"
)

var TestDB contract.Store

func TestMain(m *testing.M) {

	config, err := config.LoadConfig("../..")
	if err != nil {
		log.Fatalln("cannot load config:", err)
	}

	TestDB, err = NewPostgresStore(config)
	if err != nil {
		log.Fatalln("cannot connect to db:", err)
	}
	os.Exit(m.Run())
}
