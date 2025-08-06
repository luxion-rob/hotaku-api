package utils

import (
	"fmt"
	"hotaku-api/internal/domain/request"
	"hotaku-api/internal/domain/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	DefaultOffset = 0
	DefaultLimit  = 10
)

func ValidatePaginationParams(offset, limit, maxLimit int) error {
	if offset < 0 {
		return fmt.Errorf("offset must be >= 0")
	}
	if limit <= 0 {
		return fmt.Errorf("limit must be > 0")
	}
	if limit > maxLimit {
		return fmt.Errorf("limit must be <= %d", maxLimit)
	}
	return nil
}

func ParsePagination(c *gin.Context, maxLimit int) (*request.Pagination, bool) {
	offset := DefaultOffset
	limit := DefaultLimit

	if offsetStr := c.Query("offset"); offsetStr != "" {
		parsed, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid offset parameter", "offset must be a valid integer"))
			return nil, false
		}
		offset = parsed
	}

	if limitStr := c.Query("limit"); limitStr != "" {
		parsed, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid limit parameter", "limit must be a valid integer"))
			return nil, false
		}
		limit = parsed
	}

	if err := ValidatePaginationParams(offset, limit, maxLimit); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid pagination parameter", err.Error()))
		return nil, false
	}

	return &request.Pagination{Offset: offset, Limit: limit}, true
}
