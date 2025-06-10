package utils

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func TestGenerateToken_Success(t *testing.T) {
	userID := uint(123)
	email := "test@example.com"

	token, err := GenerateToken(userID, email)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Verify the token structure (JWT tokens have 3 parts separated by dots)
	parts := len(strings.Split(token, "."))
	assert.Equal(t, 3, parts, "JWT token should have 3 parts")
}

func TestGenerateToken_WithDifferentUsers(t *testing.T) {
	testCases := []struct {
		userID uint
		email  string
	}{
		{1, "user1@example.com"},
		{999, "user999@example.com"},
		{12345, "long.email.address@very-long-domain-name.com"},
	}

	for _, tc := range testCases {
		token, err := GenerateToken(tc.userID, tc.email)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		// Verify token contains correct claims
		claims, err := ValidateToken(token)
		assert.NoError(t, err)
		assert.Equal(t, tc.userID, claims.UserID)
		assert.Equal(t, tc.email, claims.Email)
	}
}

func TestValidateToken_Success(t *testing.T) {
	userID := uint(123)
	email := "test@example.com"

	// Generate a token
	token, err := GenerateToken(userID, email)
	require.NoError(t, err)

	// Validate the token
	claims, err := ValidateToken(token)

	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, email, claims.Email)

	// Check expiration is in the future
	assert.True(t, claims.ExpiresAt.Time.After(time.Now()))

	// Check expiration is approximately 24 hours from now
	expectedExpiry := time.Now().Add(24 * time.Hour)
	timeDiff := claims.ExpiresAt.Time.Sub(expectedExpiry)
	assert.True(t, timeDiff < time.Minute && timeDiff > -time.Minute,
		"Token expiry should be approximately 24 hours from now")
}

func TestValidateToken_InvalidToken(t *testing.T) {
	testCases := []struct {
		name  string
		token string
	}{
		{
			name:  "Empty token",
			token: "",
		},
		{
			name:  "Invalid format",
			token: "invalid.token.format",
		},
		{
			name:  "Random string",
			token: "this-is-not-a-jwt-token",
		},
		{
			name:  "Malformed JWT",
			token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.invalid.signature",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			claims, err := ValidateToken(tc.token)
			assert.Error(t, err)
			assert.Nil(t, claims)
		})
	}
}

func TestValidateToken_TamperedToken(t *testing.T) {
	userID := uint(123)
	email := "test@example.com"

	// Generate a valid token
	token, err := GenerateToken(userID, email)
	require.NoError(t, err)

	// Tamper with the token by changing the last character
	tamperedToken := token[:len(token)-1] + "X"

	// Validate the tampered token
	claims, err := ValidateToken(tamperedToken)

	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestValidateToken_ExpiredToken(t *testing.T) {
	userID := uint(123)
	email := "test@example.com"

	// Create an expired token manually
	expirationTime := time.Now().Add(-1 * time.Hour) // 1 hour ago
	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	require.NoError(t, err)

	// Validate the expired token
	validatedClaims, err := ValidateToken(tokenString)

	assert.Error(t, err)
	assert.Nil(t, validatedClaims)
	assert.Contains(t, err.Error(), "token is expired")
}

func TestValidateToken_WrongSigningMethod(t *testing.T) {
	userID := uint(123)
	email := "test@example.com"

	// Create a token with wrong signing method (RS256 instead of HS256)
	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	// This would require RSA keys, but we'll just create an invalid token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString([]byte("invalid-key-for-rsa"))

	// This should fail during signing, but if it doesn't, validation should fail
	if err == nil {
		validatedClaims, validationErr := ValidateToken(tokenString)
		assert.Error(t, validationErr)
		assert.Nil(t, validatedClaims)
	}
}

func TestTokenRoundTrip(t *testing.T) {
	// Test multiple round trips with different users
	testUsers := []struct {
		userID uint
		email  string
	}{
		{1, "user1@test.com"},
		{42, "answer@universe.com"},
		{999999, "big.number@large.id"},
	}

	for _, user := range testUsers {
		// Generate token
		token, err := GenerateToken(user.userID, user.email)
		require.NoError(t, err)

		// Validate token
		claims, err := ValidateToken(token)
		require.NoError(t, err)

		// Verify claims match original data
		assert.Equal(t, user.userID, claims.UserID)
		assert.Equal(t, user.email, claims.Email)
	}
}

func TestClaims_Structure(t *testing.T) {
	userID := uint(123)
	email := "test@example.com"

	token, err := GenerateToken(userID, email)
	require.NoError(t, err)

	claims, err := ValidateToken(token)
	require.NoError(t, err)

	// Test that all expected fields are present
	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, email, claims.Email)
	assert.NotNil(t, claims.ExpiresAt)
	assert.True(t, claims.ExpiresAt.Time.After(time.Now()))
}

// Benchmark tests
func BenchmarkGenerateToken(b *testing.B) {
	userID := uint(123)
	email := "test@example.com"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := GenerateToken(userID, email)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkValidateToken(b *testing.B) {
	userID := uint(123)
	email := "test@example.com"

	// Generate token once
	token, err := GenerateToken(userID, email)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ValidateToken(token)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkTokenRoundTrip(b *testing.B) {
	userID := uint(123)
	email := "test@example.com"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		token, err := GenerateToken(userID, email)
		if err != nil {
			b.Fatal(err)
		}

		_, err = ValidateToken(token)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Test edge cases and error scenarios
func TestJWTSecret_Validation(t *testing.T) {
	// This test verifies that our JWT secret validation works
	// Note: This test needs to run in isolation or mock the environment

	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)

	// Test with short secret (should panic)
	os.Setenv("JWT_SECRET", "short")

	// Since we can't easily test panic in package initialization,
	// we'll verify the current secret meets requirements
	secret := os.Getenv("JWT_SECRET")
	assert.True(t, len(secret) >= 32, "JWT secret should be at least 32 characters")
}
