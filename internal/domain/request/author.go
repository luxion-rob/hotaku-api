package request

import (
	"hotaku-api/internal/domain/entities"

	"github.com/google/uuid"
)

type CreateAuthorRequest struct {
	AuthorName string  `json:"author_name" binding:"required,min=1,max=50"`
	AuthorBio  *string `json:"author_bio,omitempty"`
}

type UpdateAuthorRequest struct {
	AuthorName string  `json:"author_name" binding:"min=1,max=50"`
	AuthorBio  *string `json:"author_bio,omitempty"`
}

func (req *CreateAuthorRequest) ToAuthorEntity() *entities.Author {
	return &entities.Author{
		AuthorID:   uuid.NewString(),
		ExternalID: uuid.NewString(),
		AuthorName: req.AuthorName,
		AuthorBio:  req.AuthorBio,
	}
}

func (req *UpdateAuthorRequest) ToAuthorEntityFromUpdateRequest(existingAuthor *entities.Author) *entities.Author {
	if req == nil || existingAuthor == nil {
		return nil
	}

	// Create a copy of the existing author
	updatedAuthor := *existingAuthor

	// Update only the fields that are provided in the request
	if req.AuthorName != "" {
		updatedAuthor.AuthorName = req.AuthorName
	}

	if req.AuthorBio != nil {
		updatedAuthor.AuthorBio = req.AuthorBio
	}

	return &updatedAuthor
}
