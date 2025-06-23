package controllers

import (
	"hotaku-api/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck handles HTTP requests for health checks and responds with a JSON payload indicating service health.
func HealthCheck(c *gin.Context) {
	response := interfaces.NewHealthResponse()
	c.JSON(http.StatusOK, response)
}
