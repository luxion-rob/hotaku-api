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
        print_status "Current directory: $(pwd)"
        
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
                print_status "Loading: $clean_line"
                export "$clean_line"
            else
                print_warning "Skipping malformed line: $line"
            fi
        done < .env
        
        print_status "Environment variables loaded successfully"
    else
        print_warning "No .env file found in scripts directory"
        print_status "Using system environment variables"
    fi
}

# Help function
show_help() {
    echo "Usage: $0 [COMMAND] [OPTIONS]"
    echo ""
    echo "Commands:"
    echo "  up              Run all pending migrations"
    echo "  down [steps]    Rollback migrations (default: 1 step)"
    echo "  status          Show migration status"
    echo "  help            Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0 up"
    echo "  $0 down 2"
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
                # Update DB_HOST to use Docker service name
                export DB_HOST="mysql"
                print_status "Updated DB_HOST to 'mysql' for Docker environment"
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
    check_go
    wait_for_db
    load_env
    go run ../cmd/migrate/main.go -action=up
    print_status "Migrations completed!"
}

# Run migrations down
migrate_down() {
    local steps=${1:-1}
    print_status "Rolling back $steps migration(s)..."
    check_go
    wait_for_db
    load_env
    go run ../cmd/migrate/main.go -action=down -steps="$steps"
    print_status "Rollback completed! $steps"
}

# Show migration status
show_status() {
    print_status "Migration status:"
    # This would require additional implementation to show current migration version
    print_warning "Status command not implemented yet"
}

# Main script logic
case "$1" in
    "up")
        migrate_up
        ;;
    "down")
        migrate_down $2
        ;;
    "status")
        show_status
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