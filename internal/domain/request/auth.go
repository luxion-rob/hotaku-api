package request

import (
	"fmt"
	"hotaku-api/internal/domain/entities"

	"github.com/google/uuid"
)

// RegisterRequest represents user registration request
type RegisterRequest struct {
	RoleID   string `json:"role_id" binding:"required"`
	Name     string `json:"name" binding:"required,min=2,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=100"`
}

// LoginRequest represents user login request
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// UpdateProfileRequest represents profile update request
type UpdateProfileRequest struct {
	Name  string `json:"name,omitempty" binding:"omitempty,min=2,max=100"`
	Email string `json:"email,omitempty" binding:"omitempty,email"`
}

// ChangePasswordRequest represents password change request
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required,min=6,max=100"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=NewPassword"`
}

func (req *RegisterRequest) ToUserEntity() (*entities.User, error) {
	if req == nil {
		return nil, fmt.Errorf("request empty")
	}

	// Hash password
	tempUser := entities.User{Password: req.Password}
	if err := tempUser.HashPassword(); err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := entities.User{
		UserID:   uuid.NewString(),
		RoleID:   req.RoleID,
		Email:    req.Email,
		Password: tempUser.Password,
		Name:     req.Name,
	}

	return &user, nil
}

func (req *UpdateProfileRequest) ToUserEntity(existingUser *entities.User) *entities.User {
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
