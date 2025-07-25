package server

import (
	"hotaku-api/config"
	"hotaku-api/internal/controllers"
	"hotaku-api/internal/repo"
	"hotaku-api/internal/service"
	"hotaku-api/internal/usecase"
	"os"
)

// InitializeServer creates and configures all dependencies and returns a configured server
func InitializeServer() *Server {
	// Connect to database
	config.ConnectDatabase()

	// Initialize repositories
	userRepo := repo.NewUserRepository(config.DB)

	// Get JWT secret from environment
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		panic("JWT_SECRET environment variable is required")
	}
	tokenService := service.NewTokenService(jwtSecret)

	// Initialize MinIO service
	appConfig := config.LoadConfig()
	minioService := InitializeMinioService(appConfig)

	// Initialize use cases
	authUseCase := usecase.NewAuthUseCase(userRepo, tokenService)

	// Initialize controllers
	authController := controllers.NewAuthController(authUseCase)
	healthController := controllers.NewHealthController()
	uploadController := controllers.NewUploadController(minioService)

	// Initialize and return server
	return NewServer(authController, healthController, uploadController, tokenService)
}

// InitializeServerWithConfig creates and configures all dependencies with custom config
func InitializeServerWithConfig(appConfig *config.Config) *Server {
	// Connect to database
	config.ConnectDatabase()

	// Initialize repositories
	userRepo := repo.NewUserRepository(config.DB)

	// Get JWT secret from environment
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		panic("JWT_SECRET environment variable is required")
	}
	tokenService := service.NewTokenService(jwtSecret)

	// Initialize MinIO service
	minioService := InitializeMinioService(appConfig)

	// Initialize use cases
	authUseCase := usecase.NewAuthUseCase(userRepo, tokenService)

	// Initialize controllers
	authController := controllers.NewAuthController(authUseCase)
	healthController := controllers.NewHealthController()
	uploadController := controllers.NewUploadController(minioService)

	// Initialize and return server
	return NewServer(authController, healthController, uploadController, tokenService)
}

// InitializeMinioService initializes the MinIO service
func InitializeMinioService(appConfig *config.Config) *service.MinIOService {
	minioService, err := service.NewMinIOService(appConfig)
	if err != nil {
		panic("Failed to initialize MinIO service: " + err.Error())
	}
	return minioService
}
