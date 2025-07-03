package serviceinf

// TokenService defines the interface for token-related operations
type TokenService interface {
	GenerateToken(userID uint, email string) (string, error)
	ValidateToken(token string) (*TokenClaims, error)
	RefreshToken(token string) (string, error)
}

// TokenClaims represents JWT token claims
type TokenClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Exp    int64  `json:"exp"`
}
