package mapper

import (
	"fmt"
	"hotaku-api/internal/domain/dto"
	"hotaku-api/internal/domain/entities"
	"hotaku-api/internal/domain/request"

	"github.com/google/uuid"
)

func FromAuthDTOToUserEntity(authDTO *dto.AuthDTO) *entities.User {
	if authDTO == nil {
		return nil
	}

	// Password in AuthDTO should already be hashed
	return &entities.User{
		UserID:   uuid.NewString(),
		RoleID:   authDTO.RoleID,
		Email:    authDTO.Email,
		Password: authDTO.Password,
		Name:     authDTO.Name,
	}
}

func FromAuthDTOToUserDTO(authDTO *dto.AuthDTO) *dto.UserDTO {
	if authDTO == nil {
		return nil
	}

	return &dto.UserDTO{
		UserID: authDTO.UserID,
		RoleID: authDTO.RoleID,
		Name:   authDTO.Name,
		Email:  authDTO.Email,
	}
}

func FromAuthDTOToAuthResponse(authDTO *dto.AuthDTO, token string) *dto.AuthResponse {
	if authDTO == nil {
		return nil
	}

	return &dto.AuthResponse{
		Token: token,
		User:  FromAuthDTOToUserDTO(authDTO),
	}
}

func FromRegisterRequestToAuthDTO(req *request.RegisterRequest) (*dto.AuthDTO, error) {
	if req == nil {
		return nil, fmt.Errorf("request empty")
	}

	// Hash password
	tempUser := entities.User{Password: req.Password}
	if err := tempUser.HashPassword(); err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := dto.AuthDTO{
		UserID:   uuid.NewString(),
		RoleID:   req.RoleID,
		Email:    req.Email,
		Password: tempUser.Password,
		Name:     req.Name,
	}

	return &user, nil
}

func FromUpdateProfileRequestToUserEntity(req *request.UpdateProfileRequest, existingUser *entities.User) *entities.User {
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

func FromLoginRequestToLoginDTO(req *request.LoginRequest) *dto.LoginDTO {
	if req == nil {
		return nil
	}

	return &dto.LoginDTO{
		Email:    req.Email,
		Password: req.Password,
	}
}
