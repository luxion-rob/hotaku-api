package mapper

import (
	"hotaku-api/internal/domain/dto"
	"hotaku-api/internal/domain/entities"
)

func FromUserDTOToAuthResponse(userDTO *dto.UserDTO, token string) *dto.AuthResponse {
	if userDTO == nil {
		return nil
	}

	return &dto.AuthResponse{
		Token: token,
		User:  userDTO,
	}
}

func FromUserEntityToAuthResponse(user *entities.User, token string) *dto.AuthResponse {
	if user == nil {
		return nil
	}

	return &dto.AuthResponse{
		Token: token,
		User:  FromUserEntityToUserDTO(user),
	}
}

func FromUserEntityToUserDTO(user *entities.User) *dto.UserDTO {
	if user == nil {
		return nil
	}

	return &dto.UserDTO{
		UserID:    user.UserID,
		RoleID:    user.RoleID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
