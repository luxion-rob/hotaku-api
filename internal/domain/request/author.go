package request

type CreateAuthorRequest struct {
	AuthorName string  `json:"author_name" binding:"required,min=1,max=50"`
	AuthorBio  *string `json:"author_bio,omitempty"`
}

type UpdateAuthorRequest struct {
	AuthorName string  `json:"author_name" binding:"min=1,max=50"`
	AuthorBio  *string `json:"author_bio,omitempty"`
}
