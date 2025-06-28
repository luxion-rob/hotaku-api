package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HealthController handles health check HTTP requests
type HealthController struct{}

// NewHealthController creates a new instance of HealthController
func NewHealthController() *HealthController {
	return &HealthController{}
}

// HealthCheck performs health check
func (hc *HealthController) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"timestamp": time.Now().Unix(),
	})
}
