#!/bin/bash

# Setup environment variables for Hotaku API
# This script helps create a .env file from the template

set -e

PROJECT_ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
TEMPLATE_FILE="$PROJECT_ROOT/env.example"
ENV_FILE="$PROJECT_ROOT/.env"

echo "🔧 Setting up environment for Hotaku API..."

# Check if env.example exists at root
if [ ! -f "$TEMPLATE_FILE" ]; then
    echo "❌ env.example file not found at project root ($TEMPLATE_FILE)."
    exit 1
fi

# Check if .env already exists at root
if [ -f "$ENV_FILE" ]; then
    read -p "⚠️  .env file already exists at ($PROJECT_ROOT). Do you want to overwrite it? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        echo "❌ Aborted. Keeping existing .env file at root ($PROJECT_ROOT)."
        exit 0
    fi
fi

# Copy the template to root
cp "$TEMPLATE_FILE" "$ENV_FILE"
echo "✅ Created .env file from template at project root ($ENV_FILE)"

echo ""
echo "📝 Please edit the .env file at root ($PROJECT_ROOT) to customize your configuration:"
echo "   - Update database credentials"
echo "   - Set your preferred port"
echo "   - Configure environment settings"
echo ""
echo "🔐 Remember: Never commit the .env file to version control!"
echo ""
echo "🚀 You can now run: make dev-setup" 
