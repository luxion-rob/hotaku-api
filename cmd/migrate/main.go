package main

import (
	"flag"
	migration "hotaku-api/infra/migrations"
	"log"
	"os"
)

// main is the entry point for the migration command-line tool.
// It parses command-line flags to determine whether to apply or roll back database migrations.
// For the "down" action, it requires a positive number of steps to roll back.
// Prints usage instructions and exits if an invalid action is provided.
func main() {
	var (
		action = flag.String("action", "up", "Migration action: up, down")
		steps  = flag.Int("steps", 1, "Number of steps for rollback (only used with down action)")
	)
	flag.Parse()

	switch *action {
	case "up":
		if err := migration.RunMigrations(); err != nil {
			log.Fatal("Migration failed:", err)
		}

	case "down":
		if *steps <= 0 {
			log.Fatalf("steps must be > 0 (got %d)", *steps)
		}
		if err := migration.RollbackMigrations(*steps); err != nil {
			log.Fatal("Rollback failed:", err)
		}

	default:
		log.Println("Usage: go run ../cmd/migrate/main.go -action=[up|down] [-steps=n]")
		log.Println("Examples:")
		log.Println("  go run cmd/migrate/main.go -action=up")
		log.Println("  go run cmd/migrate/main.go -action=down -steps=1")
		os.Exit(1)
	}
}
