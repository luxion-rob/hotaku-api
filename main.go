package main

import (
	"hotaku-api/config"
	"hotaku-api/internal/server"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if it exists (for local development)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Load configuration
	appConfig := config.LoadConfig()

	// Set Gin mode
	gin.SetMode(appConfig.Server.GinMode)

	// Initialize server with all dependencies
	srv := server.InitializeServer()

	if err := srv.Run(appConfig.Server.Port); err != nil {
		log.Printf("Failed to start server: %v", err)
		panic("Failed to start server: " + err.Error())
	}
}
