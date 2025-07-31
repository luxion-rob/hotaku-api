package mapper

import (
	"hotaku-api/internal/domain/dto"
	"hotaku-api/internal/domain/entities"
	"hotaku-api/internal/domain/request"

	"github.com/google/uuid"
)

// ToAuthorDTO converts Author entity to AuthorDTO
func ToAuthorDTO(author *entities.Author) *dto.AuthorDTO {
	if author == nil {
		return nil
	}

	return &dto.AuthorDTO{
		AuthorID:   author.AuthorID,
		AuthorName: author.AuthorName,
		AuthorBio:  author.AuthorBio,
	}
}

// ToAuthorResponse converts Author entity to AuthorResponse
func ToAuthorResponse(author *entities.Author) *dto.AuthorResponse {
	if author == nil {
		return nil
	}

	authorDTO := ToAuthorDTO(author)

	return &dto.AuthorResponse{
		AuthorDTO:  *authorDTO,
		ExternalID: author.ExternalID,
		CreatedAt:  author.CreatedAt,
		UpdatedAt:  author.UpdatedAt,
	}
}

// ToAuthorListResponse converts slice of Author entities to AuthorListResponse
func ToAuthorListResponse(authors []entities.Author, total int64, offset, limit int) *dto.AuthorListResponse {
	authorDTOs := make([]dto.AuthorDTO, len(authors))
	for i := range authors {
		authorDTOs[i] = *ToAuthorDTO(&authors[i])
	}

	return &dto.AuthorListResponse{
		Authors: authorDTOs,
		Total:   total,
		Offset:  offset,
		Limit:   limit,
	}
}

// ToAuthorEntityFromCreateRequest converts CreateAuthorRequest to Author entity
func ToAuthorEntityFromCreateRequest(req *request.CreateAuthorRequest) *entities.Author {
	if req == nil {
		return nil
	}

	return &entities.Author{
		AuthorID:   uuid.New().String(),
		ExternalID: uuid.New().String(),
		AuthorName: req.AuthorName,
		AuthorBio:  req.AuthorBio,
	}
}

// ToAuthorEntityFromUpdateRequest converts UpdateAuthorRequest to Author entity (for updates)
func ToAuthorEntityFromUpdateRequest(req *request.UpdateAuthorRequest, existingAuthor *entities.Author) *entities.Author {
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
