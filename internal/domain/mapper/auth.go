package mapper

import (
	"fmt"
	"hotaku-api/internal/domain/dto"
	"hotaku-api/internal/domain/entities"
	"hotaku-api/internal/domain/request"

	"github.com/google/uuid"
)

// ToUserDTO converts User entity to UserDTO
func ToUserDTO(user *entities.User) *dto.UserDTO {
	if user == nil {
		return nil
	}

	return &dto.UserDTO{
		UserID: user.UserID,
		RoleID: user.RoleID,
		Name:   user.Name,
		Email:  user.Email,
	}
}

// ToAuthResponse converts User entity and token to AuthResponse
func ToAuthResponse(user *entities.User, token string) *dto.AuthResponse {
	if user == nil {
		return nil
	}

	return &dto.AuthResponse{
		Token: token,
		User:  ToUserDTO(user),
	}
}

// ToUserEntity converts RegisterRequest to User entity
func ToUserEntityFromRegisterRequest(req *request.RegisterRequest) (*entities.User, error) {
	if req == nil {
		return nil, fmt.Errorf("request empty")
	}

	// Hash password
	tempUser := entities.User{Password: req.Password}
	if err := tempUser.HashPassword(); err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := entities.User{
		UserID:   uuid.New().String(),
		RoleID:   req.RoleID,
		Email:    req.Email,
		Password: tempUser.Password,
		Name:     req.Name,
	}

	return &user, nil
}

// ToUserEntityFromUpdateProfileRequest converts UpdateProfileRequest to User entity (for updates)
func ToUserEntityFromUpdateProfileRequest(req *request.UpdateProfileRequest, existingUser *entities.User) *entities.User {
	if req == nil || existingUser == nil {
		return nil
	}

	// Create a copy of the existing user
	updatedUser := *existingUser

	// Update only the fields that are provided in the request
	if req.Name != "" {
		updatedUser.Name = req.Name
	}
	if req.Email != "" {
		updatedUser.Email = req.Email
	}

	return &updatedUser
}
