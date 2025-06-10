package main

import (
	"fmt"
	"hotaku-api/config"
	"hotaku-api/internal/interfaces"
	"log"
	"net/http"

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

	// Connect to database
	config.ConnectDatabase()

	r := gin.Default()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		response := interfaces.NewHealthResponse()
		c.JSON(http.StatusOK, response)
	})

	serverAddr := ":" + appConfig.Server.Port
	fmt.Printf("Starting %s v%s on port %s in %s mode\n",
		appConfig.App.Name, appConfig.App.Version, appConfig.Server.Port, appConfig.App.Env)

	if err := r.Run(serverAddr); err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
