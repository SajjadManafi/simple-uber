package api

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/SajjadManafi/simple-uber/contract"
	"github.com/SajjadManafi/simple-uber/internal/config"
	"github.com/SajjadManafi/simple-uber/internal/store"
	"github.com/SajjadManafi/simple-uber/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

var TestServer *Server

func NewTestServer(t *testing.T, store contract.Store) *Server {
	config := config.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	config, err := config.LoadConfig("../")
	if err != nil {
		log.Fatalln("cannot load config:", err)
	}

	config.AccessTokenDuration = time.Minute
	config.TokenSymmetricKey = util.RandomString(32)

	store, err := store.NewPostgresStore(config)
	if err != nil {
		log.Fatalln("cannot connect to db:", err)
	}

	TestServer, err = NewServer(config, store)
	if err != nil {
		log.Fatalln("cannot create server:", err)
	}

	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())

}
