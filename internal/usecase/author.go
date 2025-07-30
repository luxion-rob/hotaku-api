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

// CreateAuthor handles author creation
func (uc *AuthorUseCaseImpl) CreateAuthor(req *request.CreateAuthorRequest) (*dto.AuthorResponse, error) {
	// Validate request
	if req.AuthorName == "" {
		return nil, fmt.Errorf("author name is required")
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
	// Validate input
	if authorID == "" {
		return nil, fmt.Errorf("author ID is required")
	}

	// Get author from repository
	author, err := uc.authorRepo.GetByID(authorID)
	if err != nil {
		return nil, fmt.Errorf("author not found: %w", err)
	}

	// Create response DTO (simplified version)
	authorDTO := &dto.AuthorDTO{
		AuthorName: author.AuthorName,
		AuthorBio:  author.AuthorBio,
	}

	return authorDTO, nil
}

// UpdateAuthor handles author updates
func (uc *AuthorUseCaseImpl) UpdateAuthor(req *request.UpdateAuthorRequest, authorID string) (*dto.AuthorResponse, error) {
	// Validate input
	if authorID == "" {
		return nil, fmt.Errorf("author ID is required")
	}

	// Get existing author
	author, err := uc.authorRepo.GetByID(authorID)
	if err != nil {
		return nil, fmt.Errorf("author not found: %w", err)
	}

	// Update fields if provided
	if req.AuthorName != "" {
		author.AuthorName = req.AuthorName
	}
	author.AuthorBio = req.AuthorBio

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
	// Validate input
	if authorID == "" {
		return fmt.Errorf("author ID is required")
	}

	// Delete from repository
	if err := uc.authorRepo.Delete(authorID); err != nil {
		return fmt.Errorf("failed to delete author: %w", err)
	}

	return nil
}
