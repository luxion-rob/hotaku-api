package migration

import (
	"fmt"
	"hotaku-api/config"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func createDatabaseConnection() (*migrate.Migrate, error) {
	// Create database connection
	config.ConnectDatabase()
	// Get underlying *sql.DB from GORM
	sqlDB, err := config.DB.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying *sql.DB: %v", err)
	}

	// Create driver instance
	driver, err := mysql.WithInstance(sqlDB, &mysql.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to create migration driver: %v", err)
	}

	// Get current working directory
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working directory: %v", err)
	}

	// Point to the migrations/sql directory where SQL files are located
	migrationsPath := fmt.Sprintf("file://%s/migrations/sql", wd)

	// Create migrate instance
	return migrate.NewWithDatabaseInstance(
		migrationsPath,
		"mysql",
		driver,
	)
}

func RunMigrations() error {
	m, err := createDatabaseConnection()
	if err != nil {
		return fmt.Errorf("failed to create database connection: %v", err)
	}

	// Run migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %v", err)
	}

	log.Println("Migrations completed successfully")
	return nil
}

func RollbackMigrations(steps int) error {
	m, err := createDatabaseConnection()
	if err != nil {
		return fmt.Errorf("failed to create database connection: %v", err)
	}

	// Rollback migrations
	if err := m.Steps(-steps); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to rollback migrations: %v", err)
	}

	log.Printf("Rolled back %d migrations successfully", steps)
	return nil
}
