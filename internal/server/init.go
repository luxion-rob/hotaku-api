package server

import (
	"hotaku-api/config"
	"hotaku-api/internal/controllers"
	"hotaku-api/internal/repositories"
	"hotaku-api/internal/usecases"
	"os"
)

// InitializeServer creates and configures all dependencies and returns a configured server
func InitializeServer() *Server {
	// Connect to database
	config.ConnectDatabase()

	// Initialize repositories
	userRepo := repositories.NewUserRepository(config.DB)

	// Get JWT secret from environment
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		panic("JWT_SECRET environment variable is required")
	}
	tokenService := repositories.NewTokenService(jwtSecret)

	// Initialize use cases
	authUseCase := usecases.NewAuthUseCase(userRepo, tokenService)

	// Initialize controllers
	authController := controllers.NewAuthController(authUseCase)
	healthController := controllers.NewHealthController()

	// Initialize and return server
	return NewServer(authController, healthController, tokenService)
}

// InitializeServerWithConfig creates and configures all dependencies with custom config
func InitializeServerWithConfig(appConfig *config.Config) *Server {
	// Connect to database
	config.ConnectDatabase()

	// Initialize repositories
	userRepo := repositories.NewUserRepository(config.DB)

	// Get JWT secret from environment
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		panic("JWT_SECRET environment variable is required")
	}
	tokenService := repositories.NewTokenService(jwtSecret)

	// Initialize use cases
	authUseCase := usecases.NewAuthUseCase(userRepo, tokenService)

	// Initialize controllers
	authController := controllers.NewAuthController(authUseCase)
	healthController := controllers.NewHealthController()

	// Initialize and return server
	return NewServer(authController, healthController, tokenService)
}
