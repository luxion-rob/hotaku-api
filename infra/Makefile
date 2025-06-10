.PHONY: help build run test test-coverage fmt lint migrate-up migrate-down migrate-status docker-up docker-down clean test-env test-env-down test-run test-cleanup

# Default target
help:
	@echo "Available commands:"
	@echo "  make build         - Build the application"
	@echo "  make run           - Run the application locally"
	@echo "  make test          - Run tests"
	@echo "  make test-coverage - Run tests with coverage"
	@echo "  make fmt           - Format code"
	@echo "  make lint          - Run linter"
	@echo "  make docker-up     - Start Docker containers"
	@echo "  make docker-down   - Stop Docker containers"
	@echo "  make migrate-up    - Run database migrations"
	@echo "  make migrate-down  - Rollback 1 migration"
	@echo "  make migrate-status- Show migration status"
	@echo "  make clean         - Clean up Docker containers and volumes"
	@echo "  make test-env      - Start test environment"
	@echo "  make test-env-down - Stop test environment"
	@echo "  make test-run      - Run tests in test environment"
	@echo "  make test-cleanup  - Clean up test environment completely"

# Build the application
build:
	cd .. && go build -o bin/hotaku-api main.go

# Run the application locally (requires local database)
run:
	cd .. && go run main.go

# Run tests
test:
	cd .. && go test -v -race ./...

# Run tests with coverage
test-coverage:
	cd .. && go test -v -race -coverprofile=coverage.out ./...
	cd .. && go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Format code
fmt:
	cd .. && go fmt ./...

# Run linter
lint:
	cd .. && go vet ./...

# Start Docker containers
docker-up:
	docker compose -f docker/docker-compose.yml up -d

# Stop Docker containers
docker-down:
	docker compose -f docker/docker-compose.yml down

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
	docker compose -f docker/docker-compose.yml down -v
	docker system prune -f

# Development setup - start containers and run migrations
dev-setup: docker-up
	@echo "Waiting for containers to be ready..."
	@sleep 10
	@make migrate-up

# Full development start
dev: dev-setup
	docker compose -f docker/docker-compose.yml logs -f api

# Test environment targets
test-env:
	docker compose -f docker/docker-compose.test.yml up -d
	@echo "Waiting for test environment to be ready..."
	@sleep 15
	@echo "Test environment is ready!"

# Stop test environment
test-env-down:
	docker compose -f docker/docker-compose.test.yml down

# Run tests in containerized environment
test-run: test-env
	docker compose -f docker/docker-compose.test.yml --profile test run --rm test-runner
	@make test-env-down

# Clean up test environment completely
test-cleanup:
	docker compose -f docker/docker-compose.test.yml down -v
	docker system prune -f

# Run integration tests against test API
test-integration: test-env
	@echo "Running integration tests against test API..."
	@sleep 5
	# Add your integration test commands here
	# Example: curl tests, API tests, etc.
	@make test-env-down 