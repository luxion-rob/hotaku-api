package main

import (
	"net/http"

	"hotaku-api/config"
	"hotaku-api/controllers"
	"hotaku-api/interfaces"
	"hotaku-api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	config.ConnectDatabase()

	r := gin.Default()

	// Health check endpoint
	r.GET("/", func(c *gin.Context) {
		response := interfaces.NewHealthResponse()
		c.JSON(http.StatusOK, response)
	})

	r.Run(":3000")
} 