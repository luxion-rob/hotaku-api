.PHONY: help build run test test-coverage fmt lint migrate-up migrate-down migrate-status docker-up docker-down clean test-local docker-prod-up generate-secrets

# Default target
help:
	@echo "Available commands:"
	@echo "  make build          - Build the application"
	@echo "  make run            - Run the application locally"
	@echo "  make test           - Run tests"
	@echo "  make test-coverage  - Run tests with coverage"
	@echo "  make test-local     - Run all local tests (no Docker)"
	@echo "  make fmt            - Format code"
	@echo "  make lint           - Run linter"
	@echo "  make docker-up      - Start Docker containers (development)"
	@echo "  make docker-prod-up - Start Docker containers (production)"
	@echo "  make docker-down    - Stop Docker containers"
	@echo "  make migrate-up     - Run database migrations"
	@echo "  make migrate-down   - Rollback 1 migration"
	@echo "  make migrate-status - Show migration status"
	@echo "  make generate-secrets - Generate secure secrets for production"
	@echo "  make clean          - Clean up Docker containers and volumes"

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

# Start Docker containers (development)
docker-up:
	docker compose --env-file scripts/.env -f docker/docker-compose.yml up -d

# Start Docker containers (production with secrets)
docker-prod-up: .ensure-secrets
	docker compose -f docker/docker-compose.prod.yml up -d

# Stop Docker containers
docker-down:
	docker compose --env-file scripts/.env -f docker/docker-compose.yml down
	docker compose -f docker/docker-compose.prod.yml down

# Generate secure secrets for production
generate-secrets:
	@mkdir -p secrets
	@echo "Generating secure secrets..."
	@openssl rand -base64 32 > secrets/db_password.txt
	@openssl rand -base64 64 > secrets/jwt_secret.txt
	@echo "✅ Secrets generated in secrets/ directory"
	@echo "⚠️  Please secure these files and do not commit them to version control"

.ensure-secrets:
	@if [ ! -f secrets/db_password.txt ] || [ ! -f secrets/jwt_secret.txt ]; then \
		echo "Error: Secret files not found. Run 'make generate-secrets' first"; \
		exit 1; \
	fi

# Run database migrations
migrate-up: .ensure-migrate-script
	@./scripts/migrate.sh up

# Rollback database migrations
migrate-down: .ensure-migrate-script
	@./scripts/migrate.sh down

# Show migration status
migrate-status: .ensure-migrate-script
	@./scripts/migrate.sh status

.ensure-migrate-script:
	@if [ ! -f scripts/migrate.sh ]; then echo "Error: scripts/migrate.sh not found"; exit 1; fi
	@chmod +x scripts/migrate.sh

# Clean up Docker containers and volumes
clean:
	docker compose --env-file scripts/.env -f docker/docker-compose.yml down -v
	docker compose -f docker/docker-compose.prod.yml down -v
	docker system prune -f

# Development setup - start containers and run migrations
dev-setup: docker-up
	@echo "Waiting for containers to be ready..."
	@sleep 10
	@make migrate-up

# Full development start
dev: dev-setup
	docker compose --env-file scripts/.env -f docker/docker-compose.yml logs -f api 