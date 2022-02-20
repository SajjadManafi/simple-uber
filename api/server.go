package api

import (
	"github.com/SajjadManafi/simple-uber/contract"
	"github.com/gin-gonic/gin"
)

// // Server serves HTTP requests for our service.
type Server struct {
	store  contract.Store
	router *gin.Engine
}

func NewServer(store contract.Store) (*Server, error) {
	server := &Server{
		store: store,
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
func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}

// errorResponse represents a response with an error.
func errorResponse(err error) gin.H {
	return gin.H{"error:": err.Error()}
}
