package controllers

import (
	"hotaku-api/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	response := interfaces.NewHealthResponse()
	c.JSON(http.StatusOK, response)
}
