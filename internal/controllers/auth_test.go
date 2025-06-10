package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"hotaku-api/config"
	"hotaku-api/internal/interfaces"
	"hotaku-api/internal/middleware"
	"hotaku-api/internal/models"
	"hotaku-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var testDB *gorm.DB

func setupTestDB() {
	// Set test environment variables
	os.Setenv("JWT_SECRET", "test-super-secret-key-for-testing-that-is-at-least-32-characters-long")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "testpassword")
	os.Setenv("DB_NAME", "hotaku_test_db")

	// Load test configuration
	config.LoadConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "testpassword", "localhost", "3306", "hotaku_test_db")

	var err error
	testDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to test database: %v", err))
	}

	// Auto migrate test tables
	testDB.AutoMigrate(&models.User{})

	// Set global DB for testing
	config.DB = testDB
}

func cleanupTestDB() {
	testDB.Exec("DELETE FROM users")
}

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Auth routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", Register)
		auth.POST("/login", Login)
	}

	// Protected routes
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/profile", Profile)
		api.PUT("/profile", UpdateProfile)
	}

	return r
}

func TestMain(m *testing.M) {
	setupTestDB()
	defer cleanupTestDB()

	code := m.Run()

	// Cleanup after all tests
	if testDB != nil {
		if sqlDB, err := testDB.DB(); err == nil {
			sqlDB.Close()
		}
	}

	os.Exit(code)
}

// Helper function to create a test user
func createTestUser(email, password, name string) *models.User {
	user := &models.User{
		Email:    email,
		Password: password,
		Name:     name,
	}
	user.HashPassword()
	testDB.Create(user)
	return user
}

// Helper function to generate test token
func generateTestToken(userID uint, email string) string {
	token, _ := utils.GenerateToken(userID, email)
	return token
}

func TestRegister_Success(t *testing.T) {
	cleanupTestDB()
	router := setupTestRouter()

	payload := interfaces.RegisterRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}

	jsonData, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response interfaces.AuthResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.True(t, response.Success)
	assert.Equal(t, "User registered successfully", response.Message)
	assert.NotEmpty(t, response.Token)
	assert.NotNil(t, response.User)
	assert.Equal(t, "John Doe", response.User.Name)
	assert.Equal(t, "john@example.com", response.User.Email)

	// Verify user exists in database
	var user models.User
	result := testDB.Where("email = ?", "john@example.com").First(&user)
	assert.NoError(t, result.Error)
	assert.Equal(t, "John Doe", user.Name)
}

func TestRegister_InvalidJSON(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response interfaces.BaseResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.False(t, response.Success)
	assert.Equal(t, "Invalid request data", response.Message)
}

func TestRegister_ValidationErrors(t *testing.T) {
	router := setupTestRouter()

	testCases := []struct {
		name     string
		payload  interfaces.RegisterRequest
		expected string
	}{
		{
			name: "Short name",
			payload: interfaces.RegisterRequest{
				Name:     "J",
				Email:    "john@example.com",
				Password: "password123",
			},
			expected: "Name must be at least 2 characters",
		},
		{
			name: "Short password",
			payload: interfaces.RegisterRequest{
				Name:     "John Doe",
				Email:    "john@example.com",
				Password: "123",
			},
			expected: "Password must be at least 6 characters",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			jsonData, _ := json.Marshal(tc.payload)
			req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)

			var response interfaces.ErrorResponse
			err := json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)

			assert.False(t, response.Success)
			assert.Equal(t, "Validation failed", response.Message)
		})
	}
}

func TestRegister_UserAlreadyExists(t *testing.T) {
	cleanupTestDB()
	router := setupTestRouter()

	// Create existing user
	createTestUser("john@example.com", "password123", "John Existing")

	payload := interfaces.RegisterRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}

	jsonData, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)

	var response interfaces.BaseResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.False(t, response.Success)
	assert.Equal(t, "User already exists", response.Message)
}

func TestLogin_Success(t *testing.T) {
	cleanupTestDB()
	router := setupTestRouter()

	// Create test user
	createTestUser("john@example.com", "password123", "John Doe")

	payload := interfaces.LoginRequest{
		Email:    "john@example.com",
		Password: "password123",
	}

	jsonData, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response interfaces.AuthResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.True(t, response.Success)
	assert.Equal(t, "Login successful", response.Message)
	assert.NotEmpty(t, response.Token)
	assert.NotNil(t, response.User)
	assert.Equal(t, "John Doe", response.User.Name)
	assert.Equal(t, "john@example.com", response.User.Email)
}

func TestLogin_InvalidCredentials(t *testing.T) {
	cleanupTestDB()
	router := setupTestRouter()

	// Create test user
	createTestUser("john@example.com", "password123", "John Doe")

	testCases := []struct {
		name     string
		email    string
		password string
	}{
		{
			name:     "Wrong email",
			email:    "wrong@example.com",
			password: "password123",
		},
		{
			name:     "Wrong password",
			email:    "john@example.com",
			password: "wrongpassword",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			payload := interfaces.LoginRequest{
				Email:    tc.email,
				Password: tc.password,
			}

			jsonData, _ := json.Marshal(payload)
			req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusUnauthorized, w.Code)

			var response interfaces.BaseResponse
			err := json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)

			assert.False(t, response.Success)
			assert.Equal(t, "Invalid credentials", response.Message)
		})
	}
}

func TestLogin_ValidationErrors(t *testing.T) {
	router := setupTestRouter()

	testCases := []struct {
		name    string
		payload interfaces.LoginRequest
	}{
		{
			name: "Empty email",
			payload: interfaces.LoginRequest{
				Email:    "",
				Password: "password123",
			},
		},
		{
			name: "Empty password",
			payload: interfaces.LoginRequest{
				Email:    "john@example.com",
				Password: "",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			jsonData, _ := json.Marshal(tc.payload)
			req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)

			var response interfaces.BaseResponse
			err := json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)

			assert.False(t, response.Success)
			assert.Equal(t, "Validation failed", response.Message)
		})
	}
}

func TestProfile_Success(t *testing.T) {
	cleanupTestDB()
	router := setupTestRouter()

	// Create test user
	user := createTestUser("john@example.com", "password123", "John Doe")
	token := generateTestToken(user.ID, user.Email)

	req, _ := http.NewRequest("GET", "/api/profile", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response interfaces.UserResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.True(t, response.Success)
	assert.Equal(t, "Profile retrieved successfully", response.Message)
	assert.NotNil(t, response.User)
	assert.Equal(t, "John Doe", response.User.Name)
	assert.Equal(t, "john@example.com", response.User.Email)
}

func TestProfile_NoAuthHeader(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/api/profile", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, "Authorization header required", response["error"])
}

func TestProfile_InvalidToken(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/api/profile", nil)
	req.Header.Set("Authorization", "Bearer invalid-token")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, "Invalid token", response["error"])
}

func TestProfile_UserNotFound(t *testing.T) {
	cleanupTestDB()
	router := setupTestRouter()

	// Generate token for non-existent user
	token := generateTestToken(999, "nonexistent@example.com")

	req, _ := http.NewRequest("GET", "/api/profile", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response interfaces.BaseResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.False(t, response.Success)
	assert.Equal(t, "User not found", response.Message)
}

func TestUpdateProfile_Success(t *testing.T) {
	cleanupTestDB()
	router := setupTestRouter()

	// Create test user
	user := createTestUser("john@example.com", "password123", "John Doe")
	token := generateTestToken(user.ID, user.Email)

	payload := interfaces.UpdateProfileRequest{
		Name:  "John Updated",
		Email: "john.updated@example.com",
	}

	jsonData, _ := json.Marshal(payload)
	req, _ := http.NewRequest("PUT", "/api/profile", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response interfaces.UserResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.True(t, response.Success)
	assert.Equal(t, "Profile updated successfully", response.Message)
	assert.NotNil(t, response.User)
	assert.Equal(t, "John Updated", response.User.Name)
	assert.Equal(t, "john.updated@example.com", response.User.Email)

	// Verify user is updated in database
	var updatedUser models.User
	result := testDB.First(&updatedUser, user.ID)
	assert.NoError(t, result.Error)
	assert.Equal(t, "John Updated", updatedUser.Name)
	assert.Equal(t, "john.updated@example.com", updatedUser.Email)
}

func TestUpdateProfile_PartialUpdate(t *testing.T) {
	cleanupTestDB()
	router := setupTestRouter()

	// Create test user
	user := createTestUser("john@example.com", "password123", "John Doe")
	token := generateTestToken(user.ID, user.Email)

	// Update only name
	payload := interfaces.UpdateProfileRequest{
		Name: "John Updated Only",
	}

	jsonData, _ := json.Marshal(payload)
	req, _ := http.NewRequest("PUT", "/api/profile", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response interfaces.UserResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.True(t, response.Success)
	assert.Equal(t, "John Updated Only", response.User.Name)
	assert.Equal(t, "john@example.com", response.User.Email) // Email should remain unchanged
}

func TestUpdateProfile_EmailAlreadyTaken(t *testing.T) {
	cleanupTestDB()
	router := setupTestRouter()

	// Create two test users
	user1 := createTestUser("john@example.com", "password123", "John Doe")
	createTestUser("jane@example.com", "password123", "Jane Doe")

	token := generateTestToken(user1.ID, user1.Email)

	// Try to update user1's email to user2's email
	payload := interfaces.UpdateProfileRequest{
		Email: "jane@example.com",
	}

	jsonData, _ := json.Marshal(payload)
	req, _ := http.NewRequest("PUT", "/api/profile", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)

	var response interfaces.BaseResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.False(t, response.Success)
	assert.Equal(t, "Email already taken", response.Message)
}

func TestUpdateProfile_ValidationError(t *testing.T) {
	cleanupTestDB()
	router := setupTestRouter()

	// Create test user
	user := createTestUser("john@example.com", "password123", "John Doe")
	token := generateTestToken(user.ID, user.Email)

	// Try to update with invalid name (too short)
	payload := interfaces.UpdateProfileRequest{
		Name: "J",
	}

	jsonData, _ := json.Marshal(payload)
	req, _ := http.NewRequest("PUT", "/api/profile", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response interfaces.BaseResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.False(t, response.Success)
	assert.Equal(t, "Validation failed", response.Message)
}

func TestUpdateProfile_UserNotFound(t *testing.T) {
	cleanupTestDB()
	router := setupTestRouter()

	// Generate token for non-existent user
	token := generateTestToken(999, "nonexistent@example.com")

	payload := interfaces.UpdateProfileRequest{
		Name: "Updated Name",
	}

	jsonData, _ := json.Marshal(payload)
	req, _ := http.NewRequest("PUT", "/api/profile", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response interfaces.BaseResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.False(t, response.Success)
	assert.Equal(t, "User not found", response.Message)
}

// Benchmark tests
func BenchmarkRegister(b *testing.B) {
	cleanupTestDB()
	router := setupTestRouter()

	payload := interfaces.RegisterRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}

	jsonData, _ := json.Marshal(payload)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Cleanup before each iteration
		testDB.Exec("DELETE FROM users WHERE email = ?", "john@example.com")

		req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
	}
}

func BenchmarkLogin(b *testing.B) {
	cleanupTestDB()
	router := setupTestRouter()

	// Create test user once
	createTestUser("john@example.com", "password123", "John Doe")

	payload := interfaces.LoginRequest{
		Email:    "john@example.com",
		Password: "password123",
	}

	jsonData, _ := json.Marshal(payload)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
	}
}
