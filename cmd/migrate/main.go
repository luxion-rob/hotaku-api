package main

import (
	"flag"
	migration "hotaku-api/infra/migrations"
	"log"
	"os"
)

func main() {
	var (
		action = flag.String("action", "up", "Migration action: up, down")
		steps  = flag.Int("steps", 1, "Number of steps for rollback (only used with down action)")
	)
	flag.Parse()

	// Set default environment variables if not set
	setDefaultEnvVars()

	switch *action {
	case "up":
		if err := migration.RunMigrations(); err != nil {
			log.Fatal("Migration failed:", err)
		}
		log.Println("Migrations completed successfully!")

	case "down":
		if err := migration.RollbackMigrations(*steps); err != nil {
			log.Fatal("Rollback failed:", err)
		}
		log.Printf("Rolled back %d migrations successfully!\n", *steps)

	default:
		log.Println("Usage: go run ../cmd/migrate/main.go -action=[up|down] [-steps=n]")
		log.Println("Examples:")
		log.Println("  go run cmd/migrate/main.go -action=up")
		log.Println("  go run cmd/migrate/main.go -action=down -steps=1")
		os.Exit(1)
	}
}

func setDefaultEnvVars() {
	envVars := map[string]string{
		"DB_HOST":     "localhost",
		"DB_PORT":     "3306",
		"DB_USER":     "root",
		"DB_PASSWORD": "rootpassword",
		"DB_NAME":     "hotaku_db",
	}

	for key, defaultValue := range envVars {
		if os.Getenv(key) == "" {
			os.Setenv(key, defaultValue)
		}
	}
}
