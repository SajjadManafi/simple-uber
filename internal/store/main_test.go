package store

import (
	"log"
	"os"
	"testing"

	"github.com/SajjadManafi/simple-uber/contract"
)

var TestDB contract.Store

func TestMain(m *testing.M) {
	var err error
	TestDB, err = NewPostgresStore("postgres://root:password@localhost:5432/simple_uber?sslmode=disable")
	if err != nil {
		log.Fatalln("cannot connect to db:", err)
	}
	os.Exit(m.Run())
}
