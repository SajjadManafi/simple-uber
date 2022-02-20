package api

import (
	"github.com/SajjadManafi/simple-uber/contract"
	"github.com/SajjadManafi/simple-uber/internal/config"
	"github.com/gin-gonic/gin"
)

// // Server serves HTTP requests for our service.
type Server struct {
	Config config.Config
	store  contract.Store
	router *gin.Engine
}

func NewServer(config config.Config, store contract.Store) (*Server, error) {
	server := &Server{
		Config: config,
		store:  store,
	}

	server.SetupRouter()
	return server, nil

}

// SetupRouter sets up the router for the server.
func (server *Server) SetupRouter() {

	router := gin.Default()

	router.POST("/api/users", server.createUser)

	server.router = router
}

// Start runs HTTP server on a specific address.
func (server *Server) Start() error {
	return server.router.Run(server.Config.ServerAddress)
}

// errorResponse represents a response with an error.
func errorResponse(err error) gin.H {
	return gin.H{"error:": err.Error()}
}
