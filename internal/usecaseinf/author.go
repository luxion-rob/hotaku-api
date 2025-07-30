package usecaseinf

import (
	"hotaku-api/internal/domain/dto"
	"hotaku-api/internal/domain/request"
)

type AuthorUseCase interface {
	CreateAuthor(req *request.CreateAuthorRequest) (*dto.AuthorResponse, error)
	GetAuthor(authorID string) (*dto.AuthorResponse, error)
	UpdateAuthor(req *request.UpdateAuthorRequest, authorID string) (*dto.AuthorResponse, error)
	DeleteAuthor(authorID string) error
}
