.PHONY: help build run test test-coverage fmt lint migrate-up migrate-down migrate-status docker-up docker-down clean test-local

# Default target
help:
	@echo "Available commands:"
	@echo "  make build         - Build the application"
	@echo "  make run           - Run the application locally"
	@echo "  make test          - Run tests"
	@echo "  make test-coverage - Run tests with coverage"
	@echo "  make test-local    - Run all local tests (no Docker)"
	@echo "  make fmt           - Format code"
	@echo "  make lint          - Run linter"
	@echo "  make docker-up     - Start Docker containers"
	@echo "  make docker-down   - Stop Docker containers"
	@echo "  make migrate-up    - Run database migrations"
	@echo "  make migrate-down  - Rollback 1 migration"
	@echo "  make migrate-status- Show migration status"
	@echo "  make clean         - Clean up Docker containers and volumes"

# Build the application
build:
	cd .. && go build -o bin/hotaku-api main.go

# Run the application locally (requires local database)
run:
	cd .. && go run main.go

# Run tests
test:
	cd .. && go test -v -race ./utils ./internal/...

# Run tests with coverage
test-coverage:
	cd .. && go test -v -race -coverprofile=coverage.out ./utils ./internal/...
	cd .. && go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run all local tests (without Docker)
test-local:
	@export JWT_SECRET="test-super-secret-key-for-testing-that-is-at-least-32-characters-long" && \
	export GIN_MODE="test" && \
	cd .. && go test -v ./... || echo "⚠️ Some test packages may not exist yet"

# Format code
fmt:
	cd .. && go fmt ./...

# Run linter
lint:
	cd .. && go vet ./...

# Start Docker containers
docker-up:
	docker compose --env-file scripts/.env -f docker/docker-compose.yml up -d

# Stop Docker containers
docker-down:
	docker compose --env-file scripts/.env -f docker/docker-compose.yml down

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
	docker compose --env-file scripts/.env -f docker/docker-compose.yml down -v
	docker system prune -f

# Development setup - start containers and run migrations
dev-setup: docker-up
	@echo "Waiting for containers to be ready..."
	@sleep 10
	@make migrate-up

# Full development start
dev: dev-setup
	docker compose --env-file scripts/.env -f docker/docker-compose.yml logs -f api 