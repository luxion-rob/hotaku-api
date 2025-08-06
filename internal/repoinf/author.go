package repoinf

import "hotaku-api/internal/domain/entities"

// AuthorRepository defines the interface for author data access operations
// Implementations should provide CRUD operations and pagination for author entities
type AuthorRepository interface {
	// Create saves a new author to the database
	// Returns error if the operation fails or if there are constraint violations
	Create(author *entities.Author) error

	// GetByID retrieves an author by ID from the database
	// Returns ErrAuthorNotFound if no author exists with the given ID
	// Returns error if database operation fails
	GetByID(id string) (*entities.Author, error)

	// Update updates an existing author in the database
	// Returns ErrAuthorNotFound if no author exists with the given ID
	// Returns error if database operation fails
	Update(author *entities.Author) error

	// Delete removes an author by ID from the database (hard delete)
	// Returns ErrAuthorNotFound if no author exists with the given ID
	// Returns error if database operation fails
	Delete(id string) error

	// List retrieves a paginated list of all authors from the database
	// Returns authors ordered by author_name, total count, and error if operation fails
	List(offset, limit int) ([]entities.Author, int64, error)
}
