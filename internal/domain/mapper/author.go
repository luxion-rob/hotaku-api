package mapper

import (
	"hotaku-api/internal/domain/dto"
	"hotaku-api/internal/domain/entities"
	"hotaku-api/internal/domain/request"
	"hotaku-api/internal/domain/response"

	"github.com/google/uuid"
)

// FromCreateAuthorRequestToAuthorDTO transforms a CreateAuthorRequest to AuthorDTO
// Generates new UUIDs for AuthorID and ExternalID fields
// Returns nil if the input request is nil
func FromCreateAuthorRequestToAuthorDTO(req *request.CreateAuthorRequest) *dto.AuthorDTO {
	if req == nil {
		return nil
	}

	return &dto.AuthorDTO{
		AuthorID:   uuid.NewString(),
		ExternalID: uuid.NewString(),
		AuthorName: req.AuthorName,
		AuthorBio:  req.AuthorBio,
	}
}

// FromUpdateAuthorRequestToAuthorDTO transforms an UpdateAuthorRequest to AuthorDTO
// Preserves the existing AuthorID from the request
// Returns nil if the input request is nil
func FromUpdateAuthorRequestToAuthorDTO(req *request.UpdateAuthorRequest) *dto.AuthorDTO {
	if req == nil {
		return nil
	}

	updatedAuthor := &dto.AuthorDTO{
		AuthorID:   req.AuthorID,
		AuthorName: req.AuthorName,
		AuthorBio:  req.AuthorBio,
	}

	return updatedAuthor
}

// FromAuthorDTOToAuthorEntity transforms an AuthorDTO to Author entity
// Used when saving author data to the database
// Returns nil if the input DTO is nil
func FromAuthorDTOToAuthorEntity(authorDTO *dto.AuthorDTO) *entities.Author {
	if authorDTO == nil {
		return nil
	}

	return &entities.Author{
		AuthorID:   authorDTO.AuthorID,
		ExternalID: authorDTO.ExternalID,
		AuthorName: authorDTO.AuthorName,
		AuthorBio:  authorDTO.AuthorBio,
	}
}

// FromAuthorEntityToAuthorDTO transforms an Author entity to AuthorDTO
// Used when retrieving author data from the database
// Returns nil if the input entity is nil
func FromAuthorEntityToAuthorDTO(author *entities.Author) *dto.AuthorDTO {
	if author == nil {
		return nil
	}

	return &dto.AuthorDTO{
		AuthorID:   author.AuthorID,
		ExternalID: author.ExternalID,
		AuthorName: author.AuthorName,
		AuthorBio:  author.AuthorBio,
	}
}

// FromAuthorDTOToAuthorListResponse transforms a slice of AuthorDTOs to AuthorListResponse
// Includes pagination metadata (total count, offset, limit)
// Used for paginated list responses
func FromAuthorDTOToAuthorListResponse(authorDTOs []dto.AuthorDTO, total int64, offset, limit int) *response.AuthorListResponse {
	return &response.AuthorListResponse{
		Authors: authorDTOs,
		Total:   total,
		Offset:  offset,
		Limit:   limit,
	}
}
