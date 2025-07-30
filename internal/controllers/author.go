package controllers

import (
	"errors"
	"fmt"
	"hotaku-api/internal/domain/apperrors"
	"hotaku-api/internal/domain/request"
	"hotaku-api/internal/domain/response"
	"hotaku-api/internal/usecaseinf"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AuthorController handles author-related HTTP requests
type AuthorController struct {
	authorUseCase usecaseinf.AuthorUseCase
}

// NewAuthorController creates a new instance of AuthorController
func NewAuthorController(authorUseCase usecaseinf.AuthorUseCase) *AuthorController {
	return &AuthorController{
		authorUseCase: authorUseCase,
	}
}

// validateAuthorID validates that the authorID is a valid UUID format
func validateAuthorID(authorID string) error {
	if authorID == "" {
		return fmt.Errorf("author ID is empty")
	}
	if _, err := uuid.Parse(authorID); err != nil {
		return fmt.Errorf("invalid author ID format: %w", err)
	}
	return nil
}

// CreateAuthor handles author creation
func (ac *AuthorController) CreateAuthor(c *gin.Context) {
	var req request.CreateAuthorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid request data", err.Error()))
		return
	}

	// Call use case
	body, err := ac.authorUseCase.CreateAuthor(&req)
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
	if err := validateAuthorID(authorID); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid author ID", err.Error()))
		return
	}

	// Call use case
	data, err := ac.authorUseCase.GetAuthor(authorID)
	if err != nil {
		if errors.Is(err, apperrors.ErrAuthorNotFound) {
			c.JSON(http.StatusNotFound, response.ErrorResponse(http.StatusNotFound, "Author not found", err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to get author", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Author retrieved successfully", data))
}

// UpdateAuthor handles author updates
func (ac *AuthorController) UpdateAuthor(c *gin.Context) {
	authorID := c.Param("author_id")

	// Validate UUID format
	if err := validateAuthorID(authorID); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid author ID", err.Error()))
		return
	}

	var req request.UpdateAuthorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid request data", err.Error()))
		return
	}

	// Call use case
	data, err := ac.authorUseCase.UpdateAuthor(&req, authorID)
	if err != nil {
		if errors.Is(err, apperrors.ErrAuthorNotFound) {
			c.JSON(http.StatusNotFound, response.ErrorResponse(http.StatusNotFound, "Author not found", err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to update author", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Author updated successfully", data))
}

// DeleteAuthor handles author deletion
func (ac *AuthorController) DeleteAuthor(c *gin.Context) {
	authorID := c.Param("author_id")

	// Validate UUID format
	if err := validateAuthorID(authorID); err != nil {
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

func (ac *AuthorController) ListAuthors(c *gin.Context) {
	// Parse query parameters
	offset := 0
	limit := 10

	if offsetStr := c.Query("offset"); offsetStr != "" {
		if parsed, err := strconv.Atoi(offsetStr); err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid offset parameter", "offset must be a valid integer"))
			return
		} else if parsed < 0 {
			c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid offset parameter", "offset must be >= 0"))
			return
		} else {
			offset = parsed
		}
	}

	if limitStr := c.Query("limit"); limitStr != "" {
		if parsed, err := strconv.Atoi(limitStr); err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid limit parameter", "limit must be a valid integer"))
			return
		} else if parsed <= 0 {
			c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid limit parameter", "limit must be > 0"))
			return
		} else if parsed > 100 {
			c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid limit parameter", "limit must be <= 100"))
			return
		} else {
			limit = parsed
		}
	}

	// Call use case
	data, err := ac.authorUseCase.ListAuthors(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to retrieve authors", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Authors retrieved successfully", data))
}
