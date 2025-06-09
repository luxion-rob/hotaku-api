package main

import (
	"fmt"
	"hotaku-api/config"
	"hotaku-api/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	config.ConnectDatabase()

	r := gin.Default()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		response := interfaces.NewHealthResponse()
		c.JSON(http.StatusOK, response)
	})

	if err := r.Run(":3000"); err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
	r.Run(":3000")
}
