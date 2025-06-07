package main

import (
	"hotaku-api/config"
	"hotaku-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	config.ConnectDatabase()

	r := gin.Default()

	// Health check endpoint
	r.GET("/", controllers.HealthCheck)

	r.Run(":3000")
}
