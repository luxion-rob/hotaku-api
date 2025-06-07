package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Setup router
	r := gin.Default()
	r.GET("/health", HealthCheck)

	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/health", nil)
	assert.NoError(t, err)

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	var response HealthResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Verify response fields
	assert.Equal(t, "healthy", response.Status)
	assert.Equal(t, "API is running smoothly", response.Message)
	assert.Equal(t, "1.0.0", response.Version)
	assert.NotZero(t, response.Timestamp)

	// Verify content type
	assert.Equal(t, "application/json; charset=utf-8", rr.Header().Get("Content-Type"))
}

func TestHealthCheckResponse(t *testing.T) {
	// Test the response structure
	response := HealthResponse{
		Status:    "healthy",
		Message:   "API is running smoothly",
		Timestamp: 1640995200,
		Version:   "1.0.0",
	}

	assert.Equal(t, "healthy", response.Status)
	assert.Equal(t, "API is running smoothly", response.Message)
	assert.Equal(t, "1.0.0", response.Version)
	assert.Equal(t, int64(1640995200), response.Timestamp)
}
