#!/bin/bash

# Migration script for Hotaku API

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

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
    if [ -f .env ]; then
        print_status "Loading environment variables from .env file"
        
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
        done < .env
        
        print_status "Environment variables loaded successfully"
    else
        print_error "No .env file found in scripts directory"
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
    echo "  status          Show migration status"
    echo "  help            Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0 up"
    echo "  $0 down version=5       # Rollback to version 5"
    echo "  $0 force version=3      # Force to version 3"
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
    if docker compose -f ./docker/docker-compose.yml ps mysql | grep -q "Up"; then
        print_status "Docker MySQL container is running, waiting for it to be ready..."
        max_attempts=30
        attempt=1
        
        while [ $attempt -le $max_attempts ]; do
            if docker compose -f ./docker/docker-compose.yml exec mysql mysqladmin ping -h"localhost" --silent; then
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

# Run migrations up
migrate_up() {
    print_status "Running migrations up..."
    
    # Check if we're running in Docker environment
    if docker compose -f ./docker/docker-compose.yml ps mysql | grep -q "Up"; then
        print_status "Docker MySQL container is running, running migrations inside Docker..."
        check_go
        load_env
        wait_for_db
        
        # Run migration inside the API container
        docker compose -f ./docker/docker-compose.yml exec api go run cmd/migrate/main.go -action=up
    else
        print_status "Docker MySQL container not running, running migrations locally..."
        check_go
        load_env
        wait_for_db
        go run ../cmd/migrate/main.go -action=up
    fi
    
    print_status "Migrations completed!"
}

# Run migrations down
migrate_down() {
    local version=${1:-0}
    print_status "Rolling back to version $version..."
    
    # Check if we're running in Docker environment
    if docker compose -f ./docker/docker-compose.yml ps mysql | grep -q "Up"; then
        print_status "Docker MySQL container is running, running rollback inside Docker..."
        check_go
        wait_for_db
        load_env
        
        # Run migration inside the API container
        docker compose -f ./docker/docker-compose.yml exec api go run cmd/migrate/main.go -action=down -version="$version"
    else
        print_status "Docker MySQL container not running, running rollback locally..."
        check_go
        wait_for_db
        load_env
        go run ../cmd/migrate/main.go -action=down -version="$version"
    fi
    
    print_status "Rollback completed! Target version: $version"
}

# Show migration status
show_status() {
    print_status "Migration status:"
    
    # Check if we're running in Docker environment
    if docker compose -f ./docker/docker-compose.yml ps mysql | grep -q "Up"; then
        print_status "Docker MySQL container is running, checking status inside Docker..."
        check_go
        wait_for_db
        load_env
        
        # Run migration status inside the API container
        docker compose -f ./docker/docker-compose.yml exec api go run cmd/migrate/main.go -action=status
    else
        print_status "Docker MySQL container not running, checking status locally..."
        check_go
        wait_for_db
        load_env
        go run ../cmd/migrate/main.go -action=status
    fi
}

# Force migration version
force_version() {
    local version=$1
    if [ -z "$version" ]; then
        print_error "No version specified. Usage: make migrate-force version=18"
        exit 1
    fi
    print_status "Forcing migration version to $version..."
    if docker compose -f ./docker/docker-compose.yml ps mysql | grep -q "Up"; then
        print_status "Docker MySQL container is running, forcing version inside Docker..."
        check_go
        wait_for_db
        load_env
        docker compose -f ./docker/docker-compose.yml exec api go run cmd/migrate/main.go -action=force -version="$version"
    else
        print_status "Docker MySQL container not running, forcing version locally..."
        check_go
        wait_for_db
        load_env
        go run ../cmd/migrate/main.go -action=force -version="$version"
    fi
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
    "help"|"")
        show_help
        ;;
    *)
        print_error "Unknown command: $1"
        show_help
        exit 1
        ;;
esac 