package request

// CreateAuthorRequest represents the request structure for creating a new author
type CreateAuthorRequest struct {
	AuthorName string  `json:"author_name" binding:"required,min=1,max=50"`
	AuthorBio  *string `json:"author_bio,omitempty"`
}

// UpdateAuthorRequest represents the request structure for updating an existing author
type UpdateAuthorRequest struct {
	AuthorName string  `json:"author_name" binding:"min=1,max=50"`
	AuthorBio  *string `json:"author_bio,omitempty"`
}
