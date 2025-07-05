package main

import (
	"flag"
	migration "hotaku-api/infra/migrations"
	"log"
	"os"
)

func main() {
	var (
		action  = flag.String("action", "up", "Migration action: up, down, force, status")
		version = flag.Int("version", 0, "Target version for down action, or version to force")
	)
	flag.Parse()

	switch *action {
	case "up":
		if err := migration.RunMigrations(); err != nil {
			log.Fatal("Migration failed:", err)
		}

	case "down":
		if *version < 0 {
			log.Fatalf("version must be >= 0 (got %d)", *version)
		}
		if err := migration.RollbackMigrations(*version); err != nil {
			log.Fatal("Rollback failed:", err)
		}

	case "force":
		if *version < 0 {
			log.Fatalf("version must be >= 0 (got %d)", *version)
		}
		if err := migration.ForceVersion(*version); err != nil {
			log.Fatal("Force version failed:", err)
		}

	case "status":
		if err := migration.ShowMigrationStatus(); err != nil {
			log.Fatal("Status failed:", err)
		}

	default:
		os.Exit(1)
	}
}
