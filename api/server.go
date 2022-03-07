package api

import (
	"fmt"

	"github.com/SajjadManafi/simple-uber/contract"
	"github.com/SajjadManafi/simple-uber/internal/config"
	"github.com/SajjadManafi/simple-uber/internal/token"
	wsstore "github.com/SajjadManafi/simple-uber/internal/wsStore"
	"github.com/gin-gonic/gin"
)

// // Server serves HTTP requests for our service.
type Server struct {
	Config     config.Config
	store      contract.Store
	router     *gin.Engine
	tokenMaker token.Maker
	driverWSS  *wsstore.MapGrid
}

func NewServer(config config.Config, store contract.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	// tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("could not create token maker: %w", err)
	}
	server := &Server{
		Config:     config,
		store:      store,
		tokenMaker: tokenMaker,
		driverWSS:  wsstore.NewMapGrid(),
	}

	server.SetupRouter()
	return server, nil

}

// SetupRouter sets up the router for the server.
func (server *Server) SetupRouter() {

	router := gin.Default()

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// users
	router.POST("/api/users", server.createUser)
	authRoutes.GET("/api/users/:id", server.getUser)
	router.PATCH("/api/users/:id/balance", server.addUserBalance)
	authRoutes.DELETE("/api/users/:id", server.deleteUser)

	// drivers
	router.POST("/api/drivers", server.createDriver)
	authRoutes.GET("/api/drivers/:id", server.getDriver)
	authRoutes.PATCH("/api/drivers/withdraw", server.driverWithdraw)
	authRoutes.PATCH("/api/drivers/setcab", server.setCab)

	// cabs
	authRoutes.POST("/api/cabs", server.createCab)

	// trips
	authRoutes.GET("/api/trips", server.createTrip)

	//login
	router.POST("/api/login", server.login)

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
