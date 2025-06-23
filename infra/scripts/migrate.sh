#!/bin/bash

# Migration script for Hotaku API

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# print_status prints an informational message in green color to stdout.
print_status() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

# print_warning prints a warning message in yellow to stdout.
print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

# print_error prints an error message in red to stderr.
print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# load_env loads environment variables from a .env file in the current directory, exporting each valid key-value pair to the environment. If the file is missing, it prints an error message.
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

# show_help prints usage instructions and available commands for the migration script.
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

# check_go verifies that the Go programming language is installed and available in the system PATH, exiting with an error if not found.
check_go() {
    if ! command -v go &> /dev/null; then
        print_error "Go is not installed or not in PATH"
        exit 1
    fi
}

# wait_for_db waits for the MySQL database to become ready, either in a Docker container or locally, and sets DB_HOST accordingly.
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

# migrate_up runs all pending database migrations in the "up" direction, using either Docker or local environment as appropriate.
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

# migrate_down rolls back database migrations by the specified number of steps, running the rollback either inside a Docker container or locally depending on the environment.
migrate_down() {
    local steps=${1:-1}
    print_status "Rolling back $steps migration(s)..."
    
    # Check if we're running in Docker environment
    if docker compose -f ./docker/docker-compose.yml ps mysql | grep -q "Up"; then
        print_status "Docker MySQL container is running, running rollback inside Docker..."
        check_go
        wait_for_db
        load_env
        
        # Run migration inside the API container
        docker compose -f ./docker/docker-compose.yml exec api go run cmd/migrate/main.go -action=down -steps="$steps"
    else
        print_status "Docker MySQL container not running, running rollback locally..."
        check_go
        wait_for_db
        load_env
    go run ../cmd/migrate/main.go -action=down -steps="$steps"
    fi
    
    print_status "Rollback completed! $steps"
}

# show_status prints a placeholder message indicating that the migration status command is not yet implemented.
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