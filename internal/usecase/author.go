package usecase

import (
	"fmt"
	"hotaku-api/internal/domain/dto"
	"hotaku-api/internal/domain/entities"
	"hotaku-api/internal/domain/request"
	"hotaku-api/internal/repoinf"
	"hotaku-api/internal/usecaseinf"
	"hotaku-api/internal/validation"
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
	// Enforce maximum length for optional bio
	if req.AuthorBio != nil && len(*req.AuthorBio) > 1000 {
		return nil, fmt.Errorf("author bio must not exceed 1000 characters")
	}

	// Create author entity using mapper
	author := req.ToAuthorEntity()

	// Save to repository
	if err := uc.authorRepo.Create(author); err != nil {
		return nil, fmt.Errorf("failed to create author: %w", err)
	}

	return author.ToAuthorResponse(), nil
}

// GetAuthor retrieves an author by ID
func (uc *AuthorUseCaseImpl) GetAuthor(authorID string) (*dto.AuthorDTO, error) {
	// Validate authorID
	if err := validation.ValidateUUID(authorID, "author ID"); err != nil {
		return nil, err
	}

	// Get author from repository
	author, err := uc.authorRepo.GetByID(authorID)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	return author.ToDTO(), nil
}

// UpdateAuthor handles author updates
func (uc *AuthorUseCaseImpl) UpdateAuthor(req *request.UpdateAuthorRequest, authorID string) (*dto.AuthorResponse, error) {
	// Validate authorID
	if err := validation.ValidateUUID(authorID, "author ID"); err != nil {
		return nil, err
	}

	// Get existing author
	author, err := uc.authorRepo.GetByID(authorID)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	// Update author using mapper
	updatedAuthor := req.ToAuthorEntityFromUpdateRequest(author)

	// Save updates to repository
	if err := uc.authorRepo.Update(updatedAuthor); err != nil {
		return nil, fmt.Errorf("failed to update author: %w", err)
	}

	return updatedAuthor.ToAuthorResponse(), nil
}

// DeleteAuthor handles author deletion
func (uc *AuthorUseCaseImpl) DeleteAuthor(authorID string) error {
	// Validate authorID
	if err := validation.ValidateUUID(authorID, "author ID"); err != nil {
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
	var authors entities.AuthorList
	// Get authors from repository
	authors, total, err := uc.authorRepo.List(offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve authors: %w", err)
	}

	return authors.ToAuthorListResponse(total, offset, limit), nil
}
