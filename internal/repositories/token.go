package repositories

import (
	"fmt"
	"hotaku-api/internal/domain/interfaces"
	"hotaku-api/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TokenServiceImpl implements the token service interface
type TokenServiceImpl struct {
	secretKey string
}

// NewTokenService creates a new instance of TokenServiceImpl
func NewTokenService(secretKey string) interfaces.TokenService {
	return &TokenServiceImpl{secretKey: secretKey}
}

// GenerateToken generates a new JWT token
func (s *TokenServiceImpl) GenerateToken(userID uint, email string) (string, error) {
	return utils.GenerateToken(userID, email)
}

// ValidateToken validates and parses a JWT token
func (s *TokenServiceImpl) ValidateToken(tokenString string) (*interfaces.TokenClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token signing method")
		}
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if token is expired
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return nil, fmt.Errorf("token is invalid or expired")
			}
		}

		userID := uint(claims["user_id"].(float64))
		email := claims["email"].(string)
		exp := int64(claims["exp"].(float64))

		return &interfaces.TokenClaims{
			UserID: userID,
			Email:  email,
			Exp:    exp,
		}, nil
	}

	return nil, fmt.Errorf("token is invalid or expired")
}

// RefreshToken refreshes an existing token
func (s *TokenServiceImpl) RefreshToken(tokenString string) (string, error) {
	claims, err := s.ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	// Generate new token with same user data
	return s.GenerateToken(claims.UserID, claims.Email)
}
