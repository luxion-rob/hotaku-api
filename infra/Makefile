.PHONY: help migrate-up migrate-down migrate-status docker-up docker-down clean docker-prod-up generate-secrets setup-env migrate-force dev-setup migrate-refresh all test

# Project root directory
INFRA_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
PROJECT_ROOT := $(abspath $(INFRA_DIR)/..)

# Default target
help: ## 🚀 Show this help message
	@echo "🚀 Hotaku API - Development Commands"
	@echo "====================================="
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo ""
	@echo "💡 Quick Start:"
	@echo "  1. make setup-env-files    # Setup environment files"
	@echo "  2. make dev               # Start full development environment"
	@echo "  3. make test              # Run tests"
	@echo ""
	@echo "📝 For more details, check the README.md file"

# Start Docker containers (development)
docker-up: ## 🐳 Start development containers
	docker compose --env-file "$(PROJECT_ROOT)/.env" -f "$(INFRA_DIR)/docker/docker-compose.yml" up -d

# Start Docker containers (production with secrets)
docker-prod-up: .ensure-secrets ## 🐳 Start production containers (requires secrets)
	docker compose --env-file "$(PROJECT_ROOT)/.env" -f "$(INFRA_DIR)/docker/docker-compose.prod.yml" up -d

# Stop Docker containers
docker-down: ## 🐳 Stop all containers
	docker compose --env-file "$(PROJECT_ROOT)/.env" -f "$(INFRA_DIR)/docker/docker-compose.yml" down
	docker compose --env-file "$(PROJECT_ROOT)/.env" -f "$(INFRA_DIR)/docker/docker-compose.prod.yml" down

# Generate secure secrets for production
generate-secrets: ## 🔐 Generate secure secrets for production
	@echo "Start generate..."
	@mkdir -p "$(INFRA_DIR)/secrets"
	@openssl rand -base64 32 > "$(INFRA_DIR)/secrets/db_password.txt"
	@openssl rand -base64 64 > "$(INFRA_DIR)/secrets/jwt_secret.txt"
	@echo "Done"

.ensure-secrets:
	@if [ ! -f "$(INFRA_DIR)/secrets/db_password.txt" ] || [ ! -f "$(INFRA_DIR)/secrets/jwt_secret.txt" ]; then \
		echo "Error: Secret files not found. Run 'make generate-secrets' first"; \
		exit 1; \
	else \
		echo "✅ Secret files found and ready"; \
	fi

# Run database migrations
migrate-up: .ensure-migrate-script ## 🗄️ Run database migrations
	@bash "$(INFRA_DIR)/scripts/migrate.sh" up

# Rollback database migrations
migrate-down: .ensure-migrate-script ## 🗄️ Rollback to specific version
	@bash "$(INFRA_DIR)/scripts/migrate.sh" down "$(version)"

# Show migration status
migrate-status: .ensure-migrate-script ## 🗄️ Show migration status
	@bash "$(INFRA_DIR)/scripts/migrate.sh" status

# Refresh migrations (rollback all and run from start)
migrate-refresh: .ensure-migrate-script ## 🗄️ Refresh migrations (rollback all + run from start)
	@bash "$(INFRA_DIR)/scripts/migrate.sh" refresh

.ensure-migrate-script:
	@chmod +x "$(INFRA_DIR)/scripts/migrate.sh"

# Force migration version (usage: make migrate-force version=18)
migrate-force: .ensure-migrate-script ## 🗄️ Force migration version (usage: make migrate-force version=18)
	@bash "$(INFRA_DIR)/scripts/migrate.sh" force $(version)

# Clean up Docker containers and volumes
clean: ## 🧹 Clean up containers, volumes, and Docker cache
	docker compose --env-file "$(PROJECT_ROOT)/.env" -f "$(INFRA_DIR)/docker/docker-compose.yml" down -v
	docker compose --env-file "$(PROJECT_ROOT)/.env" -f "$(INFRA_DIR)/docker/docker-compose.prod.yml" down -v
	docker system prune -f

# Development setup - start containers and run migrations
dev-setup: docker-up ## 🛠️ Setup development environment (Docker + DB + Migrations)
	@$(MAKE) -f "$(INFRA_DIR)/Makefile" migrate-up

# Setup environment files
setup-env: ## 🔧 Copy environment files to required locations
	@bash "$(INFRA_DIR)/scripts/setup-env.sh"
