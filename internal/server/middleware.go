package server

import (
	"hotaku-api/internal/middleware"
	"hotaku-api/internal/serviceinf"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// setupMiddleware configures all middleware for the server
func (s *Server) setupMiddleware() {
	// Global middleware
	s.router.Use(gin.Recovery())
	s.router.Use(gin.Logger())

	// CORS middleware (if needed)
	s.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // TODO: Change to specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

// setupAuthMiddleware creates the authentication middleware
func (s *Server) setupAuthMiddleware(tokenService serviceinf.TokenService) {
	s.authMiddleware = middleware.AuthMiddleware(tokenService)
}
