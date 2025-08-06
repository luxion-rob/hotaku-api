package usecaseinf

import (
	"hotaku-api/internal/domain/dto"
)

// AuthUseCase defines the interface for authentication use cases
type AuthUseCase interface {
	// Register creates a new user account and returns authentication response
	Register(authDTO *dto.AuthDTO) (*dto.UserDTO, *string, error)
	// Login authenticates a user and returns authentication response with token
	Login(loginDTO *dto.LoginDTO) (*dto.UserDTO, *string, error)
	// GetProfile retrieves user profile information by user ID
	GetProfile(userID string) (*dto.UserDTO, error)
	// UpdateProfile updates user profile information
	UpdateProfile(updateUserDTO *dto.UserDTO, userID string) (*dto.UserDTO, error)
	// ChangePassword updates user password after validation
	ChangePassword(changePasswordDTO *dto.ChangePasswordDTO, userID string) error
}
