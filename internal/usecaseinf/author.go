package usecaseinf

import (
	"hotaku-api/internal/domain/dto"
)

// AuthorUseCase defines the interface for author business logic operations
// Implementations should provide business logic, validation, and orchestration for author operations
type AuthorUseCase interface {
	// CreateAuthor handles author creation business logic
	// Validates input data and enforces business rules before saving
	// Returns the created author DTO or error if validation/operation fails
	CreateAuthor(authorDTO *dto.AuthorDTO) (*dto.AuthorDTO, error)

	// GetAuthor retrieves an author by ID with business logic validation
	// Validates UUID format and handles repository errors
	// Returns author DTO or error if not found/validation fails
	GetAuthor(authorID string) (*dto.AuthorDTO, error)

	// UpdateAuthor handles author updates with business logic validation
	// Validates existence, applies partial updates, and enforces business rules
	// Returns updated author DTO or error if validation/operation fails
	UpdateAuthor(authorDTO *dto.AuthorDTO, authorID string) (*dto.AuthorDTO, error)

	// DeleteAuthor handles author deletion with business logic validation
	// Validates UUID format and ensures author exists before deletion
	// Returns error if validation/operation fails
	DeleteAuthor(authorID string) error

	// ListAuthors retrieves a paginated list of all authors
	// Handles pagination parameters and transforms entities to DTOs
	// Returns slice of author DTOs, total count, and error if operation fails
	ListAuthors(offset, limit int) ([]dto.AuthorDTO, int64, error)
}
