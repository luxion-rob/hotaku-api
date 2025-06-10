package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"hotaku-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// Set test JWT secret
	os.Setenv("JWT_SECRET", "test-super-secret-key-for-testing-that-is-at-least-32-characters-long")

	code := m.Run()

	// Cleanup
	os.Unsetenv("JWT_SECRET")

	os.Exit(code)
}

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Protected route
	protected := r.Group("/protected")
	protected.Use(AuthMiddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			userID := c.GetUint("user_id")
			userEmail := c.GetString("user_email")
			c.JSON(http.StatusOK, gin.H{
				"user_id":    userID,
				"user_email": userEmail,
				"message":    "Success",
			})
		})
	}

	return r
}

func TestAuthMiddleware_Success(t *testing.T) {
	router := setupTestRouter()

	// Generate valid token
	userID := uint(123)
	email := "test@example.com"
	token, err := utils.GenerateToken(userID, email)
	require.NoError(t, err)

	req, _ := http.NewRequest("GET", "/protected/profile", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response to verify user data was set correctly
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, float64(userID), response["user_id"])
	assert.Equal(t, email, response["user_email"])
	assert.Equal(t, "Success", response["message"])
}

func TestAuthMiddleware_NoAuthHeader(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/protected/profile", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, "Authorization header required", response["error"])
}

func TestAuthMiddleware_EmptyAuthHeader(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/protected/profile", nil)
	req.Header.Set("Authorization", "")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, "Authorization header required", response["error"])
}

func TestAuthMiddleware_InvalidAuthHeaderFormat(t *testing.T) {
	router := setupTestRouter()

	testCases := []struct {
		name   string
		header string
	}{
		{
			name:   "Missing Bearer prefix",
			header: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
		},
		{
			name:   "Wrong prefix",
			header: "Basic token-here",
		},
		{
			name:   "Missing token after Bearer",
			header: "Bearer",
		},
		{
			name:   "Multiple spaces",
			header: "Bearer  token-here",
		},
		{
			name:   "Extra parts",
			header: "Bearer token-here extra-part",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/protected/profile", nil)
			req.Header.Set("Authorization", tc.header)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusUnauthorized, w.Code)

			var response map[string]string
			err := json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)

			assert.Equal(t, "Invalid authorization header format", response["error"])
		})
	}
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	router := setupTestRouter()

	testCases := []struct {
		name  string
		token string
	}{
		{
			name:  "Invalid token format",
			token: "invalid-token",
		},
		{
			name:  "Expired token signature",
			token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.invalid",
		},
		{
			name:  "Empty token",
			token: "",
		},
		{
			name:  "Malformed JWT",
			token: "not.a.jwt.token.at.all",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/protected/profile", nil)
			req.Header.Set("Authorization", "Bearer "+tc.token)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusUnauthorized, w.Code)

			var response map[string]string
			err := json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)

			assert.Equal(t, "Invalid token", response["error"])
		})
	}
}

func TestAuthMiddleware_TamperedToken(t *testing.T) {
	router := setupTestRouter()

	// Generate valid token
	userID := uint(123)
	email := "test@example.com"
	token, err := utils.GenerateToken(userID, email)
	require.NoError(t, err)

	// Tamper with the token by changing the last character
	tamperedToken := token[:len(token)-1] + "X"

	req, _ := http.NewRequest("GET", "/protected/profile", nil)
	req.Header.Set("Authorization", "Bearer "+tamperedToken)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, "Invalid token", response["error"])
}

func TestAuthMiddleware_ContextValues(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Create a test route that checks context values
	r.GET("/test", AuthMiddleware(), func(c *gin.Context) {
		userID := c.GetUint("user_id")
		userEmail := c.GetString("user_email")

		// Verify the values are set correctly
		assert.Equal(t, uint(456), userID)
		assert.Equal(t, "context@test.com", userEmail)

		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Generate token with specific values
	userID := uint(456)
	email := "context@test.com"
	token, err := utils.GenerateToken(userID, email)
	require.NoError(t, err)

	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAuthMiddleware_MultipleRequests(t *testing.T) {
	router := setupTestRouter()

	// Test with multiple different users
	testUsers := []struct {
		userID uint
		email  string
	}{
		{1, "user1@test.com"},
		{999, "user999@test.com"},
		{42, "answer@universe.com"},
	}

	for _, user := range testUsers {
		t.Run(user.email, func(t *testing.T) {
			token, err := utils.GenerateToken(user.userID, user.email)
			require.NoError(t, err)

			req, _ := http.NewRequest("GET", "/protected/profile", nil)
			req.Header.Set("Authorization", "Bearer "+token)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)

			var response map[string]interface{}
			err = json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)

			assert.Equal(t, float64(user.userID), response["user_id"])
			assert.Equal(t, user.email, response["user_email"])
		})
	}
}

// Test middleware chaining
func TestAuthMiddleware_WithOtherMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Add a custom middleware before auth
	customMiddleware := func(c *gin.Context) {
		c.Set("custom_value", "test")
		c.Next()
	}

	r.GET("/chain", customMiddleware, AuthMiddleware(), func(c *gin.Context) {
		userID := c.GetUint("user_id")
		customValue := c.GetString("custom_value")

		c.JSON(http.StatusOK, gin.H{
			"user_id":      userID,
			"custom_value": customValue,
		})
	})

	userID := uint(123)
	email := "test@example.com"
	token, err := utils.GenerateToken(userID, email)
	require.NoError(t, err)

	req, _ := http.NewRequest("GET", "/chain", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, float64(userID), response["user_id"])
	assert.Equal(t, "test", response["custom_value"])
}

// Benchmark tests
func BenchmarkAuthMiddleware_ValidToken(b *testing.B) {
	router := setupTestRouter()

	userID := uint(123)
	email := "test@example.com"
	token, err := utils.GenerateToken(userID, email)
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("GET", "/protected/profile", nil)
		req.Header.Set("Authorization", "Bearer "+token)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			b.Fatalf("Expected 200, got %d", w.Code)
		}
	}
}

func BenchmarkAuthMiddleware_InvalidToken(b *testing.B) {
	router := setupTestRouter()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("GET", "/protected/profile", nil)
		req.Header.Set("Authorization", "Bearer invalid-token")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusUnauthorized {
			b.Fatalf("Expected 401, got %d", w.Code)
		}
	}
}
