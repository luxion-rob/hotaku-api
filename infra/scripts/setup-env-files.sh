#!/bin/bash

# Setup environment files script
# Copies env.example to multiple directories

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # print_status prints an informational message in green with an [INFO] tag.

print_status() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

# print_warning prints a warning message to stdout with a yellow [WARNING] prefix.
print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

# print_error prints an error message prefixed with a red [ERROR] tag.
print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Get the directory where this script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$(dirname "$SCRIPT_DIR")")"

print_status "Setting up environment files..."
print_status "Script directory: $SCRIPT_DIR"
print_status "Project root: $PROJECT_ROOT"

# Check if env.example exists
if [ ! -f "$SCRIPT_DIR/env.example" ]; then
    print_error "env.example not found in $SCRIPT_DIR"
    exit 1
fi

# copy_env_file copies env.example from the script directory to a target directory as .env, prompting the user before overwriting if the file already exists.
copy_env_file() {
    local target_dir="$1"
    local target_file="$target_dir/.env"
    
    if [ -f "$target_file" ]; then
        print_warning "File $target_file already exists."
        while true; do
            read -r -p "Do you want to overwrite it? (y/n): " yn
            case $yn in
                [Yy]* ) 
                    print_status "Overwriting $target_file..."
                    cp "$SCRIPT_DIR/env.example" "$target_file"
                    print_status "✅ Copied env.example to $target_file"
                    break
                    ;;
                [Nn]* ) 
                    print_status "Skipping $target_file"
                    break
                    ;;
                * ) 
                    echo "Please answer y or n."
                    ;;
            esac
        done
    else
        cp "$SCRIPT_DIR/env.example" "$target_file"
        print_status "✅ Copied env.example to $target_file"
    fi
}

# Copy to scripts directory (current directory)
print_status "Copying to scripts directory..."
copy_env_file "$SCRIPT_DIR"

# Copy to docker directory
print_status "Copying to docker directory..."
copy_env_file "$PROJECT_ROOT/infra/docker"

# Copy to infra directory
print_status "Copying to infra directory..."
copy_env_file "$PROJECT_ROOT/infra"

# Copy to project root
print_status "Copying to project root..."
copy_env_file "$PROJECT_ROOT"

print_status "🎉 Environment files setup completed!"
print_status ""
print_status "Files created:"
print_status "  - $SCRIPT_DIR/.env"
print_status "  - $PROJECT_ROOT/infra/docker/.env"
print_status "  - $PROJECT_ROOT/infra/.env"
print_status "  - $PROJECT_ROOT/.env"
print_status ""
print_status "⚠️  Remember to:"
print_status "  1. Update the database password in each .env file"
print_status "  2. Update the JWT_SECRET for security"
print_status "  3. Never commit .env files to version control" 