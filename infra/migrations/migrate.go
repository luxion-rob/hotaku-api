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

// createDatabaseConnection initializes and returns a migration instance configured for the MySQL database using the golang-migrate library.
// It establishes a database connection, creates a migration driver, and sets the migration source path to the SQL files directory.
// Returns a configured *migrate.Migrate instance or an error if any step fails.
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
	migrationsPath := fmt.Sprintf("file://%s/infra/migrations/sql", wd)

	// Create migrate instance
	return migrate.NewWithDatabaseInstance(
		migrationsPath,
		"mysql",
		driver,
	)
}

// RunMigrations applies all pending database migrations using the configured migration source.
// Returns an error if the migration process fails, except when there are no changes to apply.
func RunMigrations() error {
	m, err := createDatabaseConnection()
	if err != nil {
		return fmt.Errorf("failed to create database connection: %v", err)
	}

	defer func() {
		serr, dberr := m.Close()
		if serr != nil {
			log.Printf("Warning: failed to close migration instance: %v", serr)
		}
		if dberr != nil {
			log.Printf("Warning: failed to close database connection: %v", dberr)
		}
	}()

	// Run migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %v", err)
	}

	return nil
}

// RollbackMigrations rolls back the specified number of database migration steps.
//
// If no migrations are available to roll back, the function completes without error.
// Returns an error if the rollback operation fails for reasons other than no change.
func RollbackMigrations(steps int) error {
	m, err := createDatabaseConnection()
	if err != nil {
		return fmt.Errorf("failed to create database connection: %v", err)
	}

	defer func() {
		serr, dberr := m.Close()
		if serr != nil {
			log.Printf("Warning: failed to close migration instance: %v", serr)
		}
		if dberr != nil {
			log.Printf("Warning: failed to close database connection: %v", dberr)
		}
	}()

	// Rollback migrations
	if err := m.Steps(-steps); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to rollback migrations: %v", err)
	}

	return nil
}
