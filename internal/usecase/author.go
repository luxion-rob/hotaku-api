package usecase

import (
	"fmt"
	"hotaku-api/internal/domain/dto"
	"hotaku-api/internal/domain/entities"
	"hotaku-api/internal/domain/request"
	"hotaku-api/internal/repoinf"
	"hotaku-api/internal/usecaseinf"

	"github.com/google/uuid"
)

// AuthorUseCaseImpl implements the author use cases
type AuthorUseCaseImpl struct {
	authorRepo repoinf.AuthorRepository
}

// NewAuthorUseCase creates a new instance of AuthorUseCaseImpl
func NewAuthorUseCase(authorRepo repoinf.AuthorRepository) usecaseinf.AuthorUseCase {
	return &AuthorUseCaseImpl{
		authorRepo: authorRepo,
	}
}

// validateAuthorID validates that the authorID is a valid UUID format
func validateAuthorID(authorID string) error {
	if authorID == "" {
		return fmt.Errorf("author ID is required")
	}
	if _, err := uuid.Parse(authorID); err != nil {
		return fmt.Errorf("invalid author ID format: %w", err)
	}
	return nil
}

// CreateAuthor handles author creation
func (uc *AuthorUseCaseImpl) CreateAuthor(req *request.CreateAuthorRequest) (*dto.AuthorResponse, error) {
	// Validate request
	if req.AuthorName == "" {
		return nil, fmt.Errorf("author name is required")
	}

	// Enforce maximum length for optional bio
	if req.AuthorBio != nil && len(*req.AuthorBio) > 1000 {
		return nil, fmt.Errorf("author bio must not exceed 1000 characters")
	}

	// Create author entity
	author := &entities.Author{
		AuthorID:   uuid.New().String(),
		ExternalID: uuid.New().String(),
		AuthorName: req.AuthorName,
		AuthorBio:  req.AuthorBio,
	}

	// Save to repository
	if err := uc.authorRepo.Create(author); err != nil {
		return nil, fmt.Errorf("failed to create author: %w", err)
	}

	// Create response DTO
	authorResponse := &dto.AuthorResponse{
		AuthorID:   author.AuthorID,
		ExternalID: author.ExternalID,
		AuthorName: author.AuthorName,
		AuthorBio:  author.AuthorBio,
		CreatedAt:  author.CreatedAt,
		UpdatedAt:  author.UpdatedAt,
	}

	return authorResponse, nil
}

// GetAuthor retrieves an author by ID
func (uc *AuthorUseCaseImpl) GetAuthor(authorID string) (*dto.AuthorDTO, error) {
	// Validate authorID
	if err := validateAuthorID(authorID); err != nil {
		return nil, err
	}

	// Get author from repository
	author, err := uc.authorRepo.GetByID(authorID)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	// Create response DTO (simplified version)
	authorDTO := &dto.AuthorDTO{
		AuthorID:   author.AuthorID,
		AuthorName: author.AuthorName,
		AuthorBio:  author.AuthorBio,
	}

	return authorDTO, nil
}

// UpdateAuthor handles author updates
func (uc *AuthorUseCaseImpl) UpdateAuthor(req *request.UpdateAuthorRequest, authorID string) (*dto.AuthorResponse, error) {
	// Validate authorID
	if err := validateAuthorID(authorID); err != nil {
		return nil, err
	}

	// Get existing author
	author, err := uc.authorRepo.GetByID(authorID)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	// Update fields if provided
	if req.AuthorName != "" {
		author.AuthorName = req.AuthorName
	}
	if req.AuthorBio != nil {
		author.AuthorBio = req.AuthorBio
	}

	// Save updates to repository
	if err := uc.authorRepo.Update(author); err != nil {
		return nil, fmt.Errorf("failed to update author: %w", err)
	}

	// Create response DTO
	authorResponse := &dto.AuthorResponse{
		AuthorID:   author.AuthorID,
		ExternalID: author.ExternalID,
		AuthorName: author.AuthorName,
		AuthorBio:  author.AuthorBio,
		CreatedAt:  author.CreatedAt,
		UpdatedAt:  author.UpdatedAt,
	}

	return authorResponse, nil
}

// DeleteAuthor handles author deletion
func (uc *AuthorUseCaseImpl) DeleteAuthor(authorID string) error {
	// Validate authorID
	if err := validateAuthorID(authorID); err != nil {
		return err
	}

	// Delete from repository
	if err := uc.authorRepo.Delete(authorID); err != nil {
		return fmt.Errorf("failed to delete author: %w", err)
	}

	return nil
}

// ListAuthors retrieves a paginated list of all authors
func (uc *AuthorUseCaseImpl) ListAuthors(offset, limit int) (*dto.AuthorListResponse, error) {
	// Validate pagination parameters
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 10 // Default limit
	}
	if limit > 100 {
		limit = 100 // Maximum limit
	}

	// Get authors from repository
	authors, total, err := uc.authorRepo.List(offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve authors: %w", err)
	}

	// Convert entities to DTOs
	authorDTOs := make([]dto.AuthorDTO, len(authors))
	for i, author := range authors {
		authorDTOs[i] = dto.AuthorDTO{
			AuthorID:   author.AuthorID,
			AuthorName: author.AuthorName,
			AuthorBio:  author.AuthorBio,
		}
	}

	// Create response
	response := &dto.AuthorListResponse{
		Authors: authorDTOs,
		Total:   total,
		Offset:  offset,
		Limit:   limit,
	}

	return response, nil
}
