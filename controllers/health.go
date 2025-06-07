package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthResponse struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
	Version   string `json:"version"`
}

func HealthCheck(c *gin.Context) {
	response := HealthResponse{
		Status:    "healthy",
		Message:   "API is running smoothly",
		Timestamp: time.Now().Unix(),
		Version:   "1.0.0",
	}
	c.JSON(http.StatusOK, response)
}
