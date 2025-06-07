package interfaces

import "hotaku-api/models"

// AuthService defines the interface for authentication services
type AuthService interface {
	Register(req *RegisterRequest) (*AuthResponse, error)
	Login(req *LoginRequest) (*AuthResponse, error)
	GetProfile(userID uint) (*UserResponse, error)
	UpdateProfile(userID uint, req *UpdateProfileRequest) (*UserResponse, error)
	ChangePassword(userID uint, req *ChangePasswordRequest) error
}

// UserService defines the interface for user-related services
type UserService interface {
	GetUserByID(id uint) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
	ListUsers(page, perPage int) ([]models.User, *PaginationMeta, error)
}

// TokenService defines the interface for token-related operations
type TokenService interface {
	GenerateToken(userID uint, email string) (string, error)
	ValidateToken(token string) (*TokenClaims, error)
	RefreshToken(token string) (string, error)
}

// DatabaseService defines the interface for database operations
type DatabaseService interface {
	Connect() error
	Close() error
	Migrate() error
	IsHealthy() bool
}

// TokenClaims represents JWT token claims
type TokenClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Exp    int64  `json:"exp"`
}

// Repository interfaces for data access layer
type UserRepository interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
	List(offset, limit int) ([]models.User, int64, error)
}

// Middleware interfaces
type AuthMiddleware interface {
	RequireAuth() func(c interface{})
	OptionalAuth() func(c interface{})
}

// Validator interface for request validation
type Validator interface {
	ValidateStruct(s interface{}) error
	ValidateField(field interface{}, tag string) error
}

// Logger interface for logging
type Logger interface {
	Info(message string, fields ...interface{})
	Error(message string, fields ...interface{})
	Debug(message string, fields ...interface{})
	Warn(message string, fields ...interface{})
}

// Cache interface for caching operations
type Cache interface {
	Set(key string, value interface{}, ttl int) error
	Get(key string) (interface{}, error)
	Delete(key string) error
	Clear() error
} 