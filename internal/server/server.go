package server

import (
	"fmt"
	"hotaku-api/internal/controllers"
	"hotaku-api/internal/domain/interfaces"
	"log"

	"github.com/gin-gonic/gin"
)

// Server represents the HTTP server
type Server struct {
	router           *gin.Engine
	authController   *controllers.AuthController
	healthController *controllers.HealthController
	authMiddleware   gin.HandlerFunc
}

// NewServer creates a new server instance
func NewServer(
	authController *controllers.AuthController,
	healthController *controllers.HealthController,
	tokenService interfaces.TokenService,
) *Server {
	router := gin.Default()

	server := &Server{
		router:           router,
		authController:   authController,
		healthController: healthController,
	}

	// Setup middleware
	server.setupMiddleware()
	server.setupAuthMiddleware(tokenService)

	// Setup routes
	server.setupRoutes()

	return server
}

// Run starts the HTTP server
func (s *Server) Run(port int) error {
	serverAddr := fmt.Sprintf(":%d", port)
	log.Printf("Server starting on port %d", port)
	return s.router.Run(serverAddr)
}

// GetRouter returns the gin router (useful for testing)
func (s *Server) GetRouter() *gin.Engine {
	return s.router
}
