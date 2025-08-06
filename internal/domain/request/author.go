package request

// CreateAuthorRequest represents the request structure for creating a new author
type CreateAuthorRequest struct {
	AuthorName string  `json:"author_name" binding:"required,min=1,max=50"`
	AuthorBio  *string `json:"author_bio,omitempty" binding:"omitempty,max=1000"`
}

// UpdateAuthorRequest represents the request structure for updating an existing author
type UpdateAuthorRequest struct {
	AuthorName string  `json:"author_name" binding:"omitempty,min=1,max=50"`
	AuthorBio  *string `json:"author_bio,omitempty" binding:"omitempty,max=1000"`
}
