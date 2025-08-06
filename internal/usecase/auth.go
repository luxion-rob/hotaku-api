package usecase

import (
	"fmt"
	"hotaku-api/internal/domain/dto"
	"hotaku-api/internal/domain/mapper"
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
func (uc *AuthUseCaseImpl) Register(authDTO *dto.AuthDTO) (*dto.UserDTO, *string, error) {
	// Check if user already exists
	existingUser, err := uc.userRepo.GetByEmail(authDTO.Email)
	if err == nil && existingUser != nil {
		return nil, nil, fmt.Errorf("user already exists")
	}

	user := mapper.FromAuthDTOToUserEntity(authDTO)

	// Save user to repository
	if err := uc.userRepo.Create(user); err != nil {
		return nil, nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Generate token
	token, err := uc.tokenService.GenerateToken(user.UserID, user.Email)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return mapper.FromUserEntityToUserDTO(user), &token, nil
}

// Login handles user login
func (uc *AuthUseCaseImpl) Login(loginDTO *dto.LoginDTO) (*dto.UserDTO, *string, error) {
	// Get user by email
	user, err := uc.userRepo.GetByEmail(loginDTO.Email)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid credentials")
	}

	// Check password
	if !user.CheckPassword(loginDTO.Password) {
		return nil, nil, fmt.Errorf("invalid credentials")
	}

	// Generate token
	token, err := uc.tokenService.GenerateToken(user.UserID, user.Email)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate token: %w", err)
	}

	// Create response using mapper
	return mapper.FromUserEntityToUserDTO(user), &token, nil
}

// GetProfile retrieves user profile
func (uc *AuthUseCaseImpl) GetProfile(userID string) (*dto.UserDTO, error) {
	user, err := uc.userRepo.GetByID(userID)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	// Create response using mapper
	return mapper.FromUserEntityToUserDTO(user), nil
}

// UpdateProfile updates user profile
func (uc *AuthUseCaseImpl) UpdateProfile(updateUserDTO *dto.UserDTO, userID string) (*dto.UserDTO, error) {
	existingUser, err := uc.userRepo.GetByID(userID)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	// Check if email is already taken by another user (if email is being updated)
	if updateUserDTO.Email != "" {
		existingEmailUser, err := uc.userRepo.GetByEmail(updateUserDTO.Email)
		if err == nil && existingEmailUser != nil && existingEmailUser.UserID != userID {
			return nil, fmt.Errorf("email already taken")
		}
	}

	// Create a copy of the existing user
	updatedUser := *existingUser

	// Update only the fields that are provided in the request
	if updateUserDTO.Name != "" {
		updatedUser.Name = updateUserDTO.Name
	}

	if updateUserDTO.Email != "" {
		updatedUser.Email = updateUserDTO.Email
	}

	// Save updated user
	if err := uc.userRepo.Update(&updatedUser); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	// Create response using mapper
	return mapper.FromUserEntityToUserDTO(&updatedUser), nil
}

// ChangePassword changes user password
func (uc *AuthUseCaseImpl) ChangePassword(changePasswordDTO *dto.ChangePasswordDTO, userID string) error {
	user, err := uc.userRepo.GetByID(userID)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	// Verify current password
	if !user.CheckPassword(changePasswordDTO.CurrentPassword) {
		return fmt.Errorf("current password is incorrect")
	}

	// Update password
	user.Password = changePasswordDTO.NewPassword
	if err := user.HashPassword(); err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Save updated user
	if err := uc.userRepo.Update(user); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
