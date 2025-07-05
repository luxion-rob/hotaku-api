package usecaseinf

import (
	"hotaku-api/internal/domain/dto"
	"hotaku-api/internal/domain/request"
)

// AuthUseCase defines the interface for authentication use cases
type AuthUseCase interface {
	// Register creates a new user account and returns authentication response
	Register(req *request.RegisterRequest) (*dto.AuthResponse, error)
	// Login authenticates a user and returns authentication response with token
	Login(req *request.LoginRequest) (*dto.AuthResponse, error)
	// GetProfile retrieves user profile information by user ID
	GetProfile(userID string) (*dto.UserDTO, error)
	// UpdateProfile updates user profile information
	UpdateProfile(userID string, req *request.UpdateProfileRequest) (*dto.UserDTO, error)
	// ChangePassword updates user password after validation
	ChangePassword(userID string, req *request.ChangePasswordRequest) error
}
