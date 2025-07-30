package dto

import "time"

type AuthorResponse struct {
	AuthorID   string    `json:"author_id"`
	ExternalID string    `json:"external_id"`
	AuthorName string    `json:"author_name"`
	AuthorBio  *string   `json:"author_bio,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
