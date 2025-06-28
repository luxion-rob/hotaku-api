package usecases

import (
	"fmt"
	"hotaku-api/internal/domain/dto"
	"hotaku-api/internal/domain/entities"
	"hotaku-api/internal/domain/interfaces"
)

// AuthUseCaseImpl implements the authentication use cases
type AuthUseCaseImpl struct {
	userRepo     interfaces.UserRepository
	tokenService interfaces.TokenService
}

// NewAuthUseCase creates a new instance of AuthUseCaseImpl
func NewAuthUseCase(userRepo interfaces.UserRepository, tokenService interfaces.TokenService) interfaces.AuthUseCase {
	return &AuthUseCaseImpl{
		userRepo:     userRepo,
		tokenService: tokenService,
	}
}

// Register handles user registration
func (uc *AuthUseCaseImpl) Register(req *dto.RegisterRequest) (*dto.AuthResponse, error) {
	// Check if user already exists
	existingUser, err := uc.userRepo.GetByEmail(req.Email)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("user already exists")
	}

	// Create new user entity
	user := &entities.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	// Hash password
	if err := user.HashPassword(); err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Save user to repository
	if err := uc.userRepo.Create(user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Generate token
	token, err := uc.tokenService.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	// Create response
	userData := &dto.UserData{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return dto.NewAuthResponse("User registered successfully", token, userData), nil
}

// Login handles user login
func (uc *AuthUseCaseImpl) Login(req *dto.LoginRequest) (*dto.AuthResponse, error) {
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
	token, err := uc.tokenService.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	// Create response
	userData := &dto.UserData{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return dto.NewAuthResponse("Login successful", token, userData), nil
}

// GetProfile retrieves user profile
func (uc *AuthUseCaseImpl) GetProfile(userID uint) (*dto.UserResponse, error) {
	user, err := uc.userRepo.GetByID(userID)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	userData := &dto.UserData{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return dto.NewUserResponse("Profile retrieved successfully", userData), nil
}

// UpdateProfile updates user profile
func (uc *AuthUseCaseImpl) UpdateProfile(userID uint, req *dto.UpdateProfileRequest) (*dto.UserResponse, error) {
	user, err := uc.userRepo.GetByID(userID)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	// Update fields if provided
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		// Check if email is already taken by another user
		existingUser, err := uc.userRepo.GetByEmail(req.Email)
		if err == nil && existingUser != nil && existingUser.ID != userID {
			return nil, fmt.Errorf("email already taken")
		}
		user.Email = req.Email
	}

	// Save updated user
	if err := uc.userRepo.Update(user); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	userData := &dto.UserData{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return dto.NewUserResponse("Profile updated successfully", userData), nil
}

// ChangePassword changes user password
func (uc *AuthUseCaseImpl) ChangePassword(userID uint, req *dto.ChangePasswordRequest) error {
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
