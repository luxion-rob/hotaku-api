package controllers

import (
	"errors"
	"hotaku-api/internal/domain/apperrors"
	"hotaku-api/internal/domain/mapper"
	"hotaku-api/internal/domain/request"
	"hotaku-api/internal/domain/response"
	"hotaku-api/internal/usecaseinf"
	"hotaku-api/internal/validation"
	"hotaku-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

const MaxLimit = 100

// AuthorController handles all HTTP requests related to author management
type AuthorController struct {
	authorUseCase usecaseinf.AuthorUseCase
}

// NewAuthorController creates a new instance of AuthorController
func NewAuthorController(authorUseCase usecaseinf.AuthorUseCase) *AuthorController {
	return &AuthorController{
		authorUseCase: authorUseCase,
	}
}

// CreateAuthor handles author creation
func (ac *AuthorController) CreateAuthor(c *gin.Context) {
	var req request.CreateAuthorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid request data", err.Error()))
		return
	}

	// Map to DTO
	authorDTO := mapper.FromCreateAuthorRequestToAuthorDTO(&req)

	// Call use case
	body, err := ac.authorUseCase.CreateAuthor(authorDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Create author failed", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse(http.StatusCreated, "Author created successfully", body))
}

// GetAuthor retrieves an author by ID
func (ac *AuthorController) GetAuthor(c *gin.Context) {
	authorID := c.Param("author_id")

	// Validate UUID format
	if err := validation.ValidateUUID(authorID, "author ID"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid author ID", err.Error()))
		return
	}

	// Call use case
	body, err := ac.authorUseCase.GetAuthor(authorID)
	if err != nil {
		if errors.Is(err, apperrors.ErrAuthorNotFound) {
			c.JSON(http.StatusNotFound, response.ErrorResponse(http.StatusNotFound, "Author not found", err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to get author", err.Error()))
		return
	}

	c.JSON(http.StatusOK, body)
}

// UpdateAuthor handles author updates
func (ac *AuthorController) UpdateAuthor(c *gin.Context) {
	authorID := c.Param("author_id")

	// Validate UUID format
	if err := validation.ValidateUUID(authorID, "author ID"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid author ID", err.Error()))
		return
	}

	// request from client must have authorID
	var req request.UpdateAuthorRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid request data", err.Error()))
		return
	}

	// Map to DTO
	authorUpdate := mapper.FromUpdateAuthorRequestToAuthorDTO(&req)

	// Call use case
	_, err := ac.authorUseCase.UpdateAuthor(authorUpdate, authorID)
	if err != nil {
		if errors.Is(err, apperrors.ErrAuthorNotFound) {
			c.JSON(http.StatusNotFound, response.ErrorResponse(http.StatusNotFound, "Author not found", err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to update author", err.Error()))
		return
	}

	c.Status(http.StatusNoContent)
}

// DeleteAuthor handles author deletion
func (ac *AuthorController) DeleteAuthor(c *gin.Context) {
	authorID := c.Param("author_id")

	// Validate UUID format
	if err := validation.ValidateUUID(authorID, "author ID"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid author ID", err.Error()))
		return
	}

	// Call use case
	err := ac.authorUseCase.DeleteAuthor(authorID)
	if err != nil {
		if errors.Is(err, apperrors.ErrAuthorNotFound) {
			c.JSON(http.StatusNotFound, response.ErrorResponse(http.StatusNotFound, "Author not found", err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to delete author", err.Error()))
		return
	}

	c.Status(http.StatusNoContent)
}

// ListAuthors retrieves a list author
func (ac *AuthorController) ListAuthors(c *gin.Context) {
	// Parse and validate pagination parameters
	pagination, ok := utils.ParsePagination(c, MaxLimit)
	if !ok {
		return
	}

	// Call use case
	authorDTOs, total, err := ac.authorUseCase.ListAuthors(pagination.Offset, pagination.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to retrieve authors", err.Error()))
		return
	}

	// Map to AuthorListResponse
	body := mapper.FromAuthorDTOToAuthorListResponse(authorDTOs, total, pagination.Offset, pagination.Limit)

	c.JSON(http.StatusOK, body)
}
