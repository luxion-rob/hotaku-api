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

func ParsePagination(c *gin.Context, maxLimit int) (*request.Pagination, bool) {
	offset := DefaultOffset
	limit := DefaultLimit

	if offsetStr := c.Query("offset"); offsetStr != "" {
		parsed, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid offset parameter", "offset must be a valid integer"))
			return nil, false
		}
		if parsed < 0 {
			c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid offset parameter", "offset must be >= 0"))
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
		if parsed <= 0 {
			c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid limit parameter", "limit must be > 0"))
			return nil, false
		}
		if parsed > maxLimit {
			c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid limit parameter", fmt.Sprintf("limit must be <= %d", maxLimit)))
			return nil, false
		}
		limit = parsed
	}

	return &request.Pagination{Offset: offset, Limit: limit}, true
}
