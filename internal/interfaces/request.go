package interfaces

// AuthRequest defines the interface for authentication-related requests
type AuthRequest interface {
	Validate() error
}

// RegisterRequest represents user registration request
type RegisterRequest struct {
	Name     string `json:"name" binding:"required" validate:"required,min=2,max=100"`
	Email    string `json:"email" binding:"required,email" validate:"required,email"`
	Password string `json:"password" binding:"required,min=6" validate:"required,min=6,max=100"`
}

// LoginRequest represents user login request
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" validate:"required,email"`
	Password string `json:"password" binding:"required" validate:"required"`
}

// UpdateProfileRequest represents profile update request
type UpdateProfileRequest struct {
	Name  string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Email string `json:"email,omitempty" validate:"omitempty,email"`
}

// ChangePasswordRequest represents password change request
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required" validate:"required"`
	NewPassword     string `json:"new_password" binding:"required,min=6" validate:"required,min=6,max=100"`
	ConfirmPassword string `json:"confirm_password" binding:"required" validate:"required,eqfield=NewPassword"`
}

// Validate implements AuthRequest interface for RegisterRequest
func (r *RegisterRequest) Validate() error {
	if len(r.Name) < 2 {
		return &ValidationError{Field: "name", Message: "Name must be at least 2 characters"}
	}
	if len(r.Password) < 6 {
		return &ValidationError{Field: "password", Message: "Password must be at least 6 characters"}
	}
	return nil
}

// Validate implements AuthRequest interface for LoginRequest
func (r *LoginRequest) Validate() error {
	if r.Email == "" {
		return &ValidationError{Field: "email", Message: "Email is required"}
	}
	if r.Password == "" {
		return &ValidationError{Field: "password", Message: "Password is required"}
	}
	return nil
}

// Validate method for UpdateProfileRequest
func (r *UpdateProfileRequest) Validate() error {
	if r.Name != "" && len(r.Name) < 2 {
		return &ValidationError{Field: "name", Message: "Name must be at least 2 characters"}
	}
	return nil
}

// Validate method for ChangePasswordRequest
func (r *ChangePasswordRequest) Validate() error {
	if len(r.NewPassword) < 6 {
		return &ValidationError{Field: "new_password", Message: "New password must be at least 6 characters"}
	}
	if r.NewPassword != r.ConfirmPassword {
		return &ValidationError{Field: "confirm_password", Message: "Passwords do not match"}
	}
	return nil
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *ValidationError) Error() string {
	return e.Message
}
