#!/bin/bash

# Migration script for Hotaku API

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color
PROJECT_ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
ENV_FILE="$PROJECT_ROOT/.env"
DOCKER_COMPOSE_FILE="$PROJECT_ROOT/infra/docker/docker-compose.yml"

# Function to print colored output
print_status() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Load environment variables from .env file
load_env() {
    if [ -f $ENV_FILE ]; then
        print_status "Loading environment variables from ($PROJECT_ROOT) file"
        
        # Read .env file line by line and export variables safely
        while IFS= read -r line; do
            # Skip empty lines and comments
            if [[ -z "$line" || "$line" =~ ^[[:space:]]*# ]]; then
                continue
            fi
            
            # Check if line contains = and export safely
            if [[ "$line" =~ ^[[:space:]]*[A-Za-z_][A-Za-z0-9_]*[[:space:]]*=[[:space:]]* ]]; then
                # Remove leading/trailing whitespace and export
                clean_line=$(echo "$line" | sed 's/^[[:space:]]*//;s/[[:space:]]*$//')
                eval "export $clean_line"
            fi
        done < $ENV_FILE
        
        print_status "Environment variables loaded successfully"
    else
        print_error "No .env file found in project root directory ($PROJECT_ROOT)"
    fi
}

# Help function
show_help() {
    echo "Usage: $0 [COMMAND] [OPTIONS]"
    echo ""
    echo "Commands:"
    echo "  up              Run all pending migrations"
    echo "  down [version]  Rollback to specific version (default: 0)"
    echo "  force [version] Force migration to specific version"
    echo "  refresh         Rollback all migrations and run them from start"
    echo "  status          Show migration status"
    echo "  help            Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0 up"
    echo "  $0 down 5              # Rollback to version 5"
    echo "  $0 force 3             # Force to version 3"
    echo "  $0 refresh             # Reset and run all migrations"
    echo "  $0 status"
}

# Check if Go is installed
check_go() {
    if ! command -v go &> /dev/null; then
        print_error "Go is not installed or not in PATH"
        exit 1
    fi
}

# Wait for database to be ready
wait_for_db() {
    print_status "Waiting for database to be ready..."
    
    # Check if we're running in Docker environment
    if docker compose --env-file "$ENV_FILE" -f "$DOCKER_COMPOSE_FILE" ps mysql | grep -q "Up"; then
        print_status "Docker MySQL container is running, waiting for it to be ready..."
        max_attempts=30
        attempt=1
        
        while [ $attempt -le $max_attempts ]; do
            if docker compose --env-file "$ENV_FILE" -f "$DOCKER_COMPOSE_FILE" exec mysql mysqladmin ping -h"localhost" --silent; then
                print_status "Database is ready!"
                break
            fi

            if [ $attempt -eq $max_attempts ]; then
                print_error "Database is not ready after $max_attempts attempts"
                exit 1
            fi

            print_warning "Attempt $attempt/$max_attempts: Database not ready, waiting..."
            sleep 2
            ((attempt++))
        done
    else
        print_status "Docker MySQL container not running, using local database settings"
        # Keep DB_HOST as localhost for local development
        export DB_HOST="localhost"
    fi
}

# Helper function to run migration commands (Docker vs local)
run_migration_cmd() {
    local action=$1
    shift
    local description=$1
    shift
    local extra_args=("$@")
    
    print_status "$description..."
    
    # Check if we're running in Docker environment
    if docker compose --env-file "$ENV_FILE" -f "$DOCKER_COMPOSE_FILE" ps mysql | grep -q "Up"; then
        print_status "Docker MySQL container is running, running $action inside Docker..."
        check_go
        wait_for_db
        load_env
        
        # Run migration inside the API container
        docker compose --env-file "$ENV_FILE" -f "$DOCKER_COMPOSE_FILE" exec api go run cmd/migrate/main.go -action="$action" "${extra_args[@]}"
    else
        print_status "Docker MySQL container not running, running $action locally..."
        check_go
        wait_for_db
        load_env
        go run ../cmd/migrate/main.go -action="$action" "${extra_args[@]}"
    fi
}

# Run migrations up
migrate_up() {
    run_migration_cmd "up" "Running migrations up"
    print_status "Migrations completed!"
}

# Run migrations down
migrate_down() {
    local version=${1:-0}
    run_migration_cmd "down" "Rolling back to version $version" -version="$version"
    print_status "Rollback completed! Target version: $version"
}

# Show migration status
show_status() {
    run_migration_cmd "status" "Checking migration status"
}

# Force migration version
force_version() {
    local version=$1
    if [ -z "$version" ]; then
        print_error "No version specified. Usage: make migrate-force version=18"
        exit 1
    fi
    run_migration_cmd "force" "Forcing migration version to $version" -version="$version"
}

# Refresh migrations (rollback all and run from start)
migrate_refresh() {
    print_status "Starting migration refresh (rollback all + run from start)..."
    
    # First, rollback to version 0 (beginning)
    print_status "Step 1: Rolling back all migrations to version 0..."
    run_migration_cmd "down" "Rolling back all migrations" -version="0"
    
    # Then, run all migrations up
    print_status "Step 2: Running all migrations from start..."
    run_migration_cmd "up" "Running all migrations from start"
    
    print_status "Migration refresh completed successfully!"
}

# Main script logic
case "$1" in
    "up")
        migrate_up
        ;;
    "down")
        migrate_down "$2"
        ;;
    "status")
        show_status
        ;;
    "force")
        force_version "$2"
        ;;
    "refresh")
        migrate_refresh
        ;;
    "help"|"")
        show_help
        ;;
    *)
        print_error "Unknown command: $1"
        show_help
        exit 1
        ;;
esac 
