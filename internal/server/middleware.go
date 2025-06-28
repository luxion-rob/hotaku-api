package server

import (
	"hotaku-api/internal/domain/interfaces"
	"hotaku-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

// setupMiddleware configures all middleware for the server
func (s *Server) setupMiddleware() {
	// Global middleware
	s.router.Use(gin.Logger())
	s.router.Use(gin.Recovery())

	// CORS middleware (if needed)
	// s.router.Use(cors.Default())
}

// setupAuthMiddleware creates the authentication middleware
func (s *Server) setupAuthMiddleware(tokenService interfaces.TokenService) {
	s.authMiddleware = middleware.AuthMiddleware(tokenService)
}
