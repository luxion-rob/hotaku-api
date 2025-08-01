package usecase

import (
	"fmt"
	"hotaku-api/internal/domain/dto"
	"hotaku-api/internal/domain/mapper"
	"hotaku-api/internal/domain/request"
	"hotaku-api/internal/repoinf"
	"hotaku-api/internal/serviceinf"
	"hotaku-api/internal/usecaseinf"
)

// AuthUseCaseImpl implements the authentication use cases
type AuthUseCaseImpl struct {
	userRepo     repoinf.UserRepository
	tokenService serviceinf.TokenService
}

// NewAuthUseCase creates a new instance of AuthUseCaseImpl
func NewAuthUseCase(userRepo repoinf.UserRepository, tokenService serviceinf.TokenService) usecaseinf.AuthUseCase {
	return &AuthUseCaseImpl{
		userRepo:     userRepo,
		tokenService: tokenService,
	}
}

// Register handles user registration
func (uc *AuthUseCaseImpl) Register(req *request.RegisterRequest) (*dto.AuthResponse, error) {
	// Check if user already exists
	existingUser, err := uc.userRepo.GetByEmail(req.Email)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("user already exists")
	}

	// Create new user entity using mapper
	user := mapper.ToUserEntityFromRegisterRequest(req)

	if user == nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Save user to repository
	if err := uc.userRepo.Create(user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Generate token
	token, err := uc.tokenService.GenerateToken(user.UserID, user.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	// Create response using mapper
	return mapper.ToAuthResponse(user, token), nil
}

// Login handles user login
func (uc *AuthUseCaseImpl) Login(req *request.LoginRequest) (*dto.AuthResponse, error) {
	// Get user by email
	user, err := uc.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Check password
	if !user.CheckPassword(req.Password) {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Generate token
	token, err := uc.tokenService.GenerateToken(user.UserID, user.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	// Create response using mapper
	return mapper.ToAuthResponse(user, token), nil
}

// GetProfile retrieves user profile
func (uc *AuthUseCaseImpl) GetProfile(userID string) (*dto.UserDTO, error) {
	user, err := uc.userRepo.GetByID(userID)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	// Create response using mapper
	return mapper.ToUserDTO(user), nil
}

// UpdateProfile updates user profile
func (uc *AuthUseCaseImpl) UpdateProfile(userID string, req *request.UpdateProfileRequest) (*dto.UserDTO, error) {
	user, err := uc.userRepo.GetByID(userID)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	// Check if email is already taken by another user (if email is being updated)
	if req.Email != "" {
		existingUser, err := uc.userRepo.GetByEmail(req.Email)
		if err == nil && existingUser != nil && existingUser.UserID != userID {
			return nil, fmt.Errorf("email already taken")
		}
	}

	// Update user using mapper
	updatedUser := mapper.ToUserEntityFromUpdateProfileRequest(req, user)

	// Save updated user
	if err := uc.userRepo.Update(updatedUser); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	// Create response using mapper
	return mapper.ToUserDTO(updatedUser), nil
}

// ChangePassword changes user password
func (uc *AuthUseCaseImpl) ChangePassword(userID string, req *request.ChangePasswordRequest) error {
	user, err := uc.userRepo.GetByID(userID)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	// Verify current password
	if !user.CheckPassword(req.CurrentPassword) {
		return fmt.Errorf("current password is incorrect")
	}

	// Update password
	user.Password = req.NewPassword
	if err := user.HashPassword(); err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Save updated user
	if err := uc.userRepo.Update(user); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
