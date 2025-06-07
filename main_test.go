package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"hotaku-api/controllers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/", controllers.HealthCheck)
	return r
}

func TestMainHealthEndpoint(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response controllers.HealthResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "healthy", response.Status)
}

func TestMainEndpointNotFound(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/nonexistent", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestMainHealthEndpointMethod(t *testing.T) {
	router := setupRouter()

	// Test POST method on health endpoint (should return 404)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}
