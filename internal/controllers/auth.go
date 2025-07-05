package controllers

import (
	"fmt"
	"hotaku-api/internal/domain/request"
	"hotaku-api/internal/domain/response"
	"hotaku-api/internal/usecaseinf"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AuthController handles authentication-related HTTP requests
type AuthController struct {
	authUseCase usecaseinf.AuthUseCase
}

// NewAuthController creates a new instance of AuthController
func NewAuthController(authUseCase usecaseinf.AuthUseCase) *AuthController {
	return &AuthController{
		authUseCase: authUseCase,
	}
}

// Register handles user registration
func (ac *AuthController) Register(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request data", err.Error()))
		return
	}

	// Call use case
	body, err := ac.authUseCase.Register(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Registration failed", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse("User registered successfully", body))
}

// Login handles user login
func (ac *AuthController) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request data", err.Error()))
		return
	}

	// Call use case
	body, err := ac.authUseCase.Login(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Login failed", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse("Login successful", body))
}

// validateUserID validates that the userID is a valid UUID format
func validateUserID(userID string) error {
	if userID == "" {
		return fmt.Errorf("user ID is empty")
	}
	if _, err := uuid.Parse(userID); err != nil {
		return fmt.Errorf("invalid user ID format: %w", err)
	}
	return nil
}

// Profile retrieves user profile
func (ac *AuthController) Profile(c *gin.Context) {
	userID := c.GetString("user_id")

	// Validate UUID format
	if err := validateUserID(userID); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid user ID", err.Error()))
		return
	}

	// Call use case
	body, err := ac.authUseCase.GetProfile(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to get profile", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse("Profile retrieved successfully", body))
}

// UpdateProfile updates user profile
func (ac *AuthController) UpdateProfile(c *gin.Context) {
	userID := c.GetString("user_id")

	// Validate UUID format
	if err := validateUserID(userID); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid user ID", err.Error()))
		return
	}

	var req request.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request data", err.Error()))
		return
	}

	// Call use case
	body, err := ac.authUseCase.UpdateProfile(userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to update profile", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse("Profile updated successfully", body))
}

// ChangePassword changes user password
func (ac *AuthController) ChangePassword(c *gin.Context) {
	userID := c.GetString("user_id")

	// Validate UUID format
	if err := validateUserID(userID); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid user ID", err.Error()))
		return
	}

	var req request.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request data", err.Error()))
		return
	}

	// Call use case
	err := ac.authUseCase.ChangePassword(userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to change password", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse("Password changed successfully", nil))
}
