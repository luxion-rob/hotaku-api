package usecase

import (
	"fmt"
	"hotaku-api/internal/domain/dto"
	"hotaku-api/internal/domain/entities"
	"hotaku-api/internal/domain/mapper"
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
func (uc *AuthorUseCaseImpl) CreateAuthor(authorDTO *dto.AuthorDTO) (*dto.AuthorDTO, error) {
	// Enforce maximum length for optional bio
	if authorDTO.AuthorBio != nil && len(*authorDTO.AuthorBio) > 1000 {
		return nil, fmt.Errorf("author bio must not exceed 1000 characters")
	}

	// Map to Author entity
	author := mapper.FromAuthorDTOToAuthorEntity(authorDTO)

	// Save to repository
	if err := uc.authorRepo.Create(author); err != nil {
		return nil, fmt.Errorf("failed to create author: %w", err)
	}

	// Map the created entity back to DTO to include auto-generated fields
	return mapper.FromAuthorEntityToAuthorDTO(author), nil
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

	return mapper.FromAuthorEntityToAuthorDTO(author), nil
}

// UpdateAuthor handles author updates
func (uc *AuthorUseCaseImpl) UpdateAuthor(updateAuthorDTO *dto.AuthorDTO, authorID string) (*dto.AuthorDTO, error) {
	// Get existing author to ensure it exists and for partial update logic
	author, err := uc.authorRepo.GetByID(authorID)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	// Create base DTO from existing author
	existingAuthor := mapper.FromAuthorEntityToAuthorDTO(author)

	// Apply partial updates (only update provided fields)
	updatedAuthor := *existingAuthor

	if updateAuthorDTO.AuthorName != "" {
		updatedAuthor.AuthorName = updateAuthorDTO.AuthorName
	}

	if updateAuthorDTO.AuthorBio != nil {
		updatedAuthor.AuthorBio = updateAuthorDTO.AuthorBio
	}

	// Transform back to entity for database update
	updatedAuthorEntity := mapper.FromAuthorDTOToAuthorEntity(&updatedAuthor)

	// Save updates to repository
	if err := uc.authorRepo.Update(updatedAuthorEntity); err != nil {
		return nil, fmt.Errorf("failed to update author: %w", err)
	}

	return &updatedAuthor, nil
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
func (uc *AuthorUseCaseImpl) ListAuthors(offset, limit int) ([]dto.AuthorDTO, int64, error) {
	var authors []entities.Author

	// Get authors from repository with pagination
	authors, total, err := uc.authorRepo.List(offset, limit)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to retrieve authors: %w", err)
	}

	// Transform entities to DTOs for response
	authorDTOs := make([]dto.AuthorDTO, len(authors))
	for i := 0; i < len(authors); i++ {
		authorDTO := mapper.FromAuthorEntityToAuthorDTO(&authors[i])
		authorDTOs[i] = *authorDTO
	}

	return authorDTOs, total, nil
}
