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

// Custom logger for migrations
type migrationLogger struct{}

func (l *migrationLogger) Printf(format string, v ...interface{}) {
	log.Printf("[MIGRATION] "+format, v...)
}

func (l *migrationLogger) Verbose() bool {
	return true
}

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

	// Create migrate instance with custom logger
	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"mysql",
		driver,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create migrate instance: %v", err)
	}

	// Set custom logger
	m.Log = &migrationLogger{}

	return m, nil
}

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

	log.Printf("[MIGRATION] Starting migration process...")

	// Run migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %v", err)
	}

	if err == migrate.ErrNoChange {
		log.Printf("[MIGRATION] No new migrations to apply")
	} else {
		log.Printf("[MIGRATION] All migrations completed successfully")
	}

	return nil
}

// RollbackMigrations rolls back migrations to a specific version
func RollbackMigrations(targetVersion int) error {
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

	// Get current version
	currentVersion, dirty, err := m.Version()
	if err != nil {
		if err == migrate.ErrNilVersion {
			log.Printf("[MIGRATION] No migrations have been applied yet")
			return nil
		}
		return fmt.Errorf("failed to get current migration version: %v", err)
	}

	if dirty {
		return fmt.Errorf("database is in dirty state, please fix it first using force command")
	}

	log.Printf("[MIGRATION] Current version: %d, Target version: %d", currentVersion, targetVersion)

	if int(currentVersion) <= targetVersion {
		log.Printf("[MIGRATION] Already at or below target version %d", targetVersion)
		return nil
	}

	// Calculate how many steps to rollback
	stepsToRollback := int(currentVersion) - targetVersion
	log.Printf("[MIGRATION] Rolling back %d migration(s) to version %d...", stepsToRollback, targetVersion)

	// Rollback migrations
	if err := m.Steps(-stepsToRollback); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to rollback migrations: %v", err)
	}

	if err == migrate.ErrNoChange {
		log.Printf("[MIGRATION] No migrations to rollback")
	} else {
		log.Printf("[MIGRATION] Successfully rolled back to version %d", targetVersion)
	}

	return nil
}

// ForceVersion forces the migration version (useful for fixing dirty state)
func ForceVersion(version int) error {
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

	log.Printf("[MIGRATION] Forcing migration version to %d...", version)

	// Force the version
	if err := m.Force(version); err != nil {
		return fmt.Errorf("failed to force version %d: %v", version, err)
	}

	log.Printf("[MIGRATION] Successfully forced migration version to %d", version)
	return nil
}

// ShowMigrationStatus prints the current migration version and dirty state
func ShowMigrationStatus() error {
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

	version, dirty, err := m.Version()
	if err != nil {
		if err == migrate.ErrNilVersion {
			fmt.Println("No migrations have been applied yet.")
			return nil
		}
		return fmt.Errorf("failed to get migration version: %v", err)
	}

	fmt.Printf("Current migration version: %d\n", version)
	fmt.Printf("Dirty state: %v\n", dirty)
	return nil
}
