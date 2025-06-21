#!/bin/bash

# Setup environment variables for Hotaku API
# This script helps create a .env file from the template

set -e

echo "🔧 Setting up environment for Hotaku API..."

# Check if .env already exists
if [ -f .env ]; then
    read -p "⚠️  .env file already exists. Do you want to overwrite it? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        echo "❌ Aborted. Keeping existing .env file."
        exit 0
    fi
fi

# Check if env.example exists
if [ ! -f env.example ]; then
    echo "❌ env.example file not found. Please ensure you're in the project scripts directory."
    exit 1
fi

# Copy the template
cp env.example .env
echo "✅ Created .env file from template"

# Make it clear that they should customize it
echo ""
echo "📝 Please edit the .env file to customize your configuration:"
echo "   - Update database credentials"
echo "   - Set your preferred port"
echo "   - Configure environment settings"
echo ""
echo "🔐 Remember: Never commit the .env file to version control!"
echo ""
echo "🚀 You can now run: make dev-setup" 