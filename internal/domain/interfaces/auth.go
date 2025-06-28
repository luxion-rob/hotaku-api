package interfaces

import (
	"hotaku-api/internal/domain/dto"
	"hotaku-api/internal/domain/entities"
)

// AuthUseCase defines the interface for authentication use cases
type AuthUseCase interface {
	Register(req *dto.RegisterRequest) (*dto.AuthResponse, error)
	Login(req *dto.LoginRequest) (*dto.AuthResponse, error)
	GetProfile(userID uint) (*dto.UserResponse, error)
	UpdateProfile(userID uint, req *dto.UpdateProfileRequest) (*dto.UserResponse, error)
	ChangePassword(userID uint, req *dto.ChangePasswordRequest) error
}

// UserRepository defines the interface for user data access
type UserRepository interface {
	Create(user *entities.User) error
	GetByID(id uint) (*entities.User, error)
	GetByEmail(email string) (*entities.User, error)
	Update(user *entities.User) error
	Delete(id uint) error
	List(offset, limit int) ([]entities.User, int64, error)
}

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
