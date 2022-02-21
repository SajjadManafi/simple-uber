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

	// users
	router.POST("/api/users", server.createUser)
	router.GET("/api/users/:id", server.getUser)
	router.PATCH("/api/users/:id/balance", server.addUserBalance)
	router.DELETE("/api/users/:id", server.deleteUser)

	// drivers
	router.POST("/api/drivers", server.createDriver)
	router.GET("/api/drivers/:id", server.getDriver)
	router.PATCH("/api/drivers/withdraw", server.driverWithdraw)
	router.PATCH("/api/drivers/setcab", server.setCab)

	// cabs
	router.POST("/api/cabs", server.createCab)

	// trips
	router.GET("/api/trips", server.createTrip)

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
