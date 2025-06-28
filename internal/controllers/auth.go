package controllers

import (
	"hotaku-api/internal/domain/dto"
	"hotaku-api/internal/domain/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthController handles authentication-related HTTP requests
type AuthController struct {
	authUseCase interfaces.AuthUseCase
}

// NewAuthController creates a new instance of AuthController
func NewAuthController(authUseCase interfaces.AuthUseCase) *AuthController {
	return &AuthController{
		authUseCase: authUseCase,
	}
}

// Register handles user registration
func (ac *AuthController) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	// Call use case
	response, err := ac.authUseCase.Register(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(response.GetStatus(), response)
}

// Login handles user login
func (ac *AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	// Call use case
	response, err := ac.authUseCase.Login(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(response.GetStatus(), response)
}

// Profile retrieves user profile
func (ac *AuthController) Profile(c *gin.Context) {
	userID := c.GetUint("user_id")

	// Call use case
	response, err := ac.authUseCase.GetProfile(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(response.GetStatus(), response)
}

// UpdateProfile updates user profile
func (ac *AuthController) UpdateProfile(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	// Call use case
	response, err := ac.authUseCase.UpdateProfile(userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(response.GetStatus(), response)
}

// ChangePassword changes user password
func (ac *AuthController) ChangePassword(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req dto.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	// Call use case
	err := ac.authUseCase.ChangePassword(userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.NewSuccessResponse("Password changed successfully", nil)
	c.JSON(response.GetStatus(), response)
}
