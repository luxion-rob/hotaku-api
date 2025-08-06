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

func FromLoginRequestToLoginDTO(req *request.LoginRequest) *dto.LoginDTO {
	if req == nil {
		return nil
	}

	return &dto.LoginDTO{
		Email:    req.Email,
		Password: req.Password,
	}
}

func FromChangePasswordRequestToChangePasswordDTO(req *request.ChangePasswordRequest) *dto.ChangePasswordDTO {
	if req == nil {
		return nil
	}

	return &dto.ChangePasswordDTO{
		CurrentPassword: req.CurrentPassword,
		NewPassword:     req.NewPassword,
		ConfirmPassword: req.ConfirmPassword,
	}
}
