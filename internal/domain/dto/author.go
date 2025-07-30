package dto

import "time"

type AuthorDTO struct {
	AuthorID   string  `json:"author_id"`
	AuthorName string  `json:"author_name"`
	AuthorBio  *string `json:"author_bio,omitempty"`
}

type AuthorResponse struct {
	AuthorID   string    `json:"author_id"`
	ExternalID string    `json:"external_id"`
	AuthorName string    `json:"author_name"`
	AuthorBio  *string   `json:"author_bio,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type AuthorListResponse struct {
	Authors []AuthorDTO `json:"authors"`
	Total   int64       `json:"total"`
	Offset  int         `json:"offset"`
	Limit   int         `json:"limit"`
}
