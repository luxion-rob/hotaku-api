package config

import (
	"errors"
	"log"
	"os"
	"strconv"
)

// Config holds all configuration for the application
type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	App      AppConfig
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port    int
	GinMode string
}

// AppConfig holds application configuration
type AppConfig struct {
	Name    string
	Version string
	Env     string
}

// LoadConfig loads configuration from environment variables with defaults
func LoadConfig() *Config {
	config := &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvAsInt("DB_PORT", 3306),
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", "rootpassword"),
			Name:     getEnv("DB_NAME", "hotaku_db"),
		},
		Server: ServerConfig{
			Port:    getEnvAsInt("PORT", 3000),
			GinMode: getEnv("GIN_MODE", "debug"),
		},
		App: AppConfig{
			Name:    getEnv("APP_NAME", "Hotaku API"),
			Version: getEnv("APP_VERSION", "1.0.0"),
			Env:     getEnv("APP_ENV", "development"),
		},
	}

	log.Printf("Configuration loaded for environment: %s", config.App.Env)
	if err := config.Validate(); err != nil {
		log.Fatalf("Configuration validation failed: %v", err)
	}

	return config
}

// ValidateConfig validates the loaded configuration
func (c *Config) Validate() error {
	if c.Database.Host == "" {
		return errors.New("database host is required")
	}
	if c.Database.Port < 1 || c.Database.Port > 65535 {
		return errors.New("database port must be between 1 and 65535")
	}
	if c.Server.Port < 1 || c.Server.Port > 65535 {
		return errors.New("server port must be between 1 and 65535")
	}
	return nil
}

// getEnv gets environment variable with fallback to default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt gets environment variable as integer with fallback to default value
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
		log.Printf("Warning: Invalid integer value for %s: %s, using default: %d", key, value, defaultValue)
	}
	return defaultValue
}
