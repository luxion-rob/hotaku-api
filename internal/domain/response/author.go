package response

import (
	"hotaku-api/internal/domain/dto"
	"time"
)

// AuthorResponse represents the response structure for a single author
// Extends AuthorDTO with timestamp information
type AuthorResponse struct {
	dto.AuthorDTO
	CreatedAt time.Time `json:"created_at"` // Timestamp when the author was created
	UpdatedAt time.Time `json:"updated_at"` // Timestamp when the author was last updated
}

// AuthorListResponse represents the response structure for a paginated list of authors
type AuthorListResponse struct {
	Authors []dto.AuthorDTO `json:"authors"`
	Total   int64           `json:"total"`
	Offset  int             `json:"offset"`
	Limit   int             `json:"limit"`
}
