package controllers

import (
	"hotaku-api/internal/domain/request"
	"hotaku-api/internal/domain/response"
	"hotaku-api/internal/usecaseinf"
	"hotaku-api/internal/validation"
	"net/http"

	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid request data", err.Error()))
		return
	}

	// Call use case
	body, err := ac.authUseCase.Register(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Registration failed", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "User registered successfully", body))
}

// Login handles user login
func (ac *AuthController) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid request data", err.Error()))
		return
	}

	// Call use case
	body, err := ac.authUseCase.Login(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Login failed", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Login successful", body))
}

// Profile retrieves user profile
func (ac *AuthController) Profile(c *gin.Context) {
	userID := c.GetString("user_id")

	// Validate UUID format
	if err := validation.ValidateUUID(userID, "user ID"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid user ID", err.Error()))
		return
	}

	// Call use case
	body, err := ac.authUseCase.GetProfile(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to get profile", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Profile retrieved successfully", body))
}

// UpdateProfile updates user profile
func (ac *AuthController) UpdateProfile(c *gin.Context) {
	userID := c.GetString("user_id")

	// Validate UUID format
	if err := validation.ValidateUUID(userID, "user ID"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid user ID", err.Error()))
		return
	}

	var req request.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid request data", err.Error()))
		return
	}

	// Call use case
	body, err := ac.authUseCase.UpdateProfile(userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to update profile", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Profile updated successfully", body))
}

// ChangePassword changes user password
func (ac *AuthController) ChangePassword(c *gin.Context) {
	userID := c.GetString("user_id")

	// Validate UUID format
	if err := validation.ValidateUUID(userID, "user ID"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid user ID", err.Error()))
		return
	}

	var req request.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid request data", err.Error()))
		return
	}

	// Call use case
	err := ac.authUseCase.ChangePassword(userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to change password", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Password changed successfully", nil))
}
