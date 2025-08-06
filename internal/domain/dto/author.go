package dto

// AuthorDTO represents the data transfer object for author information
// Used for transferring author data between different layers of the application
type AuthorDTO struct {
	AuthorID   string  `json:"author_id"`
	ExternalID string  `json:"external_id"`
	AuthorName string  `json:"author_name"`
	AuthorBio  *string `json:"author_bio,omitempty"`
}
