.PHONY: help build run migrate-up migrate-down migrate-status docker-up docker-down clean

# Default target
help:
	@echo "Available commands:"
	@echo "  make build         - Build the application"
	@echo "  make run           - Run the application locally"
	@echo "  make docker-up     - Start Docker containers"
	@echo "  make docker-down   - Stop Docker containers"
	@echo "  make migrate-up    - Run database migrations"
	@echo "  make migrate-down  - Rollback 1 migration"
	@echo "  make migrate-status- Show migration status"
	@echo "  make clean         - Clean up Docker containers and volumes"

# Build the application
build:
	go build -o bin/hotaku-api main.go

# Run the application locally (requires local database)
run:
	go run main.go

# Start Docker containers
docker-up:
	docker compose up -d

# Stop Docker containers
docker-down:
	docker compose down

# Run database migrations
migrate-up:
	@chmod +x scripts/migrate.sh
	@./scripts/migrate.sh up

# Rollback database migrations
migrate-down:
	@chmod +x scripts/migrate.sh
	@./scripts/migrate.sh down

# Show migration status
migrate-status:
	@chmod +x scripts/migrate.sh
	@./scripts/migrate.sh status

# Clean up Docker containers and volumes
clean:
	docker compose down -v
	docker system prune -f

# Development setup - start containers and run migrations
dev-setup: docker-up
	@echo "Waiting for containers to be ready..."
	@sleep 10
	@make migrate-up

# Full development start
dev: dev-setup
	docker compose logs -f api 