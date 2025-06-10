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
}

# Run migrations up
migrate_up() {
    print_status "Running migrations up..."
    check_go
    wait_for_db
    pwd
    go run ../cmd/migrate/main.go -action=up
    print_status "Migrations completed!"
}

# Run migrations down
migrate_down() {
    local steps=${1:-1}
    print_status "Rolling back $steps migration(s)..."
    check_go
    wait_for_db
    go run ../cmd/migrate/main.go -action=down -steps=$steps
    print_status "Rollback completed!"
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