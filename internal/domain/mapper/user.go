package mapper

import (
	"hotaku-api/internal/domain/dto"
	"hotaku-api/internal/domain/entities"
	"hotaku-api/internal/domain/request"
	"hotaku-api/internal/domain/response"
)

func FromUserDTOToAuthResponse(userDTO *dto.UserDTO, token string) *response.AuthResponse {
	if userDTO == nil {
		return nil
	}

	return &response.AuthResponse{
		Token: token,
		User:  userDTO,
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

func FromUpdateProfileRequestToUserDTO(req *request.UpdateProfileRequest) *dto.UserDTO {
	if req == nil {
		return nil
	}

	return &dto.UserDTO{
		Name:  req.Name,
		Email: req.Email,
	}
}
