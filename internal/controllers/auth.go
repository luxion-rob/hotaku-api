package controllers

import (
	"net/http"

	"hotaku-api/config"
	"hotaku-api/internal/interfaces"
	"hotaku-api/internal/models"
	"hotaku-api/utils"

	"github.com/gin-gonic/gin"
)

// Register handles user registration by validating input, creating a new user, and returning an authentication token upon success.
func Register(c *gin.Context) {
	var req interfaces.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response := interfaces.NewErrorResponse("Invalid request data", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Validate request
	if err := req.Validate(); err != nil {
		if validationErr, ok := err.(*interfaces.ValidationError); ok {
			response := interfaces.NewValidationErrorResponse("Validation failed", []interfaces.ValidationError{*validationErr})
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := interfaces.NewErrorResponse("Validation failed", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Check if user already exists
	var existingUser models.User
	if err := config.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		response := interfaces.NewErrorResponse("User already exists", "A user with this email already exists")
		c.JSON(http.StatusConflict, response)
		return
	}

	// Create new user
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := user.HashPassword(); err != nil {
		response := interfaces.NewErrorResponse("Failed to process password", "Internal server error")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if err := config.DB.Create(&user).Error; err != nil {
		response := interfaces.NewErrorResponse("Failed to create user", "Internal server error")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		response := interfaces.NewErrorResponse("Failed to generate token", "Internal server error")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	userData := &interfaces.UserData{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	response := interfaces.NewAuthResponse("User registered successfully", token, userData)
	c.JSON(http.StatusCreated, response)
}

// Login authenticates a user with email and password, returning an authentication token and user profile data on success.
func Login(c *gin.Context) {
	var req interfaces.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response := interfaces.NewErrorResponse("Invalid request data", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Validate request
	if err := req.Validate(); err != nil {
		if validationErr, ok := err.(*interfaces.ValidationError); ok {
			response := interfaces.NewValidationErrorResponse("Validation failed", []interfaces.ValidationError{*validationErr})
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := interfaces.NewErrorResponse("Validation failed", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		response := interfaces.NewErrorResponse("Invalid credentials", "Email or password is incorrect")
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	if !user.CheckPassword(req.Password) {
		response := interfaces.NewErrorResponse("Invalid credentials", "Email or password is incorrect")
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		response := interfaces.NewErrorResponse("Failed to generate token", "Internal server error")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	userData := &interfaces.UserData{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	response := interfaces.NewAuthResponse("Login successful", token, userData)
	c.JSON(http.StatusOK, response)
}

// Profile retrieves the authenticated user's profile and returns it as a JSON response.
// If the user does not exist, responds with a 404 error.
func Profile(c *gin.Context) {
	userID := c.GetUint("user_id")

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		response := interfaces.NewErrorResponse("User not found", "The requested user does not exist")
		c.JSON(http.StatusNotFound, response)
		return
	}

	userData := &interfaces.UserData{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	response := interfaces.NewUserResponse("Profile retrieved successfully", userData)
	c.JSON(http.StatusOK, response)
}

// UpdateProfile updates the authenticated user's profile information, including name and email, if provided.
// Returns a JSON response with the updated profile data or an error message if validation fails, the user is not found, the email is already taken, or a database error occurs.
func UpdateProfile(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req interfaces.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response := interfaces.NewErrorResponse("Invalid request data", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Validate request
	if err := req.Validate(); err != nil {
		if validationErr, ok := err.(*interfaces.ValidationError); ok {
			response := interfaces.NewValidationErrorResponse("Validation failed", []interfaces.ValidationError{*validationErr})
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := interfaces.NewErrorResponse("Validation failed", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		response := interfaces.NewErrorResponse("User not found", "The requested user does not exist")
		c.JSON(http.StatusNotFound, response)
		return
	}

	// Update fields if provided
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		// Check if email is already taken by another user
		var existingUser models.User
		if err := config.DB.Where("email = ? AND id != ?", req.Email, userID).First(&existingUser).Error; err == nil {
			response := interfaces.NewErrorResponse("Email already taken", "This email is already in use by another user")
			c.JSON(http.StatusConflict, response)
			return
		}
		user.Email = req.Email
	}

	if err := config.DB.Save(&user).Error; err != nil {
		response := interfaces.NewErrorResponse("Failed to update profile", "Internal server error")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	userData := &interfaces.UserData{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	response := interfaces.NewUserResponse("Profile updated successfully", userData)
	c.JSON(http.StatusOK, response)
}
