# Hotaku API

A simple Hello World API built with Go and Gin framework.

## Prerequisites

- Docker
- Docker Compose

## Getting Started

1. Start the database and application:
```bash
make dev-setup
```

Or manually:
```bash
# Start containers
docker compose up -d

# Wait for database to be ready, then run migrations
make migrate-up
```

2. The API will be available at `http://localhost:3000`. The application supports hot reloading using Air - any changes you make to the Go files will automatically trigger a rebuild.

## Database Migrations

The project uses a proper migration system for database schema management.

### Available Migration Commands

```bash
# Run all pending migrations
make migrate-up

# Rollback the last migration
make migrate-down

# Show migration status
make migrate-status

# Using the script directly
./scripts/migrate.sh up
./scripts/migrate.sh down 2  # rollback 2 migrations
```

### Creating New Migrations

1. Create new migration files in the `migrations/` directory
2. Follow the naming convention: 
   - `XXXXXX_description.up.sql` for the forward migration
   - `XXXXXX_description.down.sql` for the rollback migration
3. Example:
   - `000002_add_posts_table.up.sql`
   - `000002_add_posts_table.down.sql`

## API Endpoints

### Public Endpoints
- `GET /` - Health check endpoint

### Authentication Endpoints
- `POST /auth/register` - User registration
- `POST /auth/login` - User login

### Protected Endpoints (require Bearer token)
- `GET /api/profile` - Get user profile
- `PUT /api/profile` - Update user profile

## Development

The project uses:
- Go 1.22.2
- Gin web framework v1.10.1
- Air for hot reloading
- Docker with development and production configurations
- GitHub Actions for CI/CD
- Comprehensive test suite with coverage reporting

## Project Structure

```
.
├── main.go              # Main application code
├── main_test.go         # Main integration tests
├── go.mod              # Go module definition
├── .air.toml           # Air configuration for hot reload
├── Dockerfile          # Docker configuration
├── compose.yml         # Docker Compose configuration
├── migrations/         # Database migration files
├── controllers/        # HTTP handlers and tests
├── config/             # Configuration files
├── utils/              # Utility functions
├── cmd/                # CLI commands
├── scripts/            # Shell scripts
├── docs/               # Documentation and GitHub Pages
├── .github/workflows/  # GitHub Actions workflows
└── Makefile            # Development commands
```

## Testing

The project includes comprehensive test coverage:

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Format code
make fmt

# Run linter
make lint
```

## GitHub Actions CI/CD

This project includes a comprehensive CI/CD pipeline with multiple workflows:

### 🔄 Main CI/CD Pipeline (`.github/workflows/ci-cd.yml`)
Runs on pushes to `main`/`develop` branches:
- **Testing**: Unit tests with MySQL service, coverage reporting
- **Building**: Go application compilation
- **Security**: Gosec security scanning
- **Docker**: Multi-platform container builds
- **Documentation**: Automatic GitHub Pages deployment

### 🔍 Pull Request Checks (`.github/workflows/pr-check.yml`)
Lightweight validation for pull requests:
- Code formatting checks
- Tests execution
- Build verification

### 🚀 Release Workflow (`.github/workflows/release.yml`)
Triggered on version tags (`v*.*.*`):
- Multi-platform binary builds (Linux, Windows, macOS)
- Automated changelog generation
- GitHub release creation
- Docker image publishing

### 📚 GitHub Pages
Documentation is automatically deployed to GitHub Pages:
- API documentation
- Test coverage reports
- Project structure and usage guides

## Building for Production

To build a production version:

```bash
docker build -t hotaku-api .
```

To run the production container:

```bash
docker run -p 3000:3000 hotaku-api
```

## API Usage Examples

### Register User
```bash
curl -X POST http://localhost:3000/auth/register \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com","password":"password123"}'
```

### Login
```bash
curl -X POST http://localhost:3000/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","password":"password123"}'
```

### Get Profile
```bash
curl -X GET http://localhost:3000/api/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Update Profile
```bash
curl -X PUT http://localhost:3000/api/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"John Updated","email":"john.updated@example.com"}'
```

## Response Format

All API responses follow a consistent format:

### Success Response
```json
{
  "success": true,
  "message": "Operation successful",
  "data": {...},
  "timestamp": 1640995200
}
```

### Error Response
```json
{
  "success": false,
  "message": "Error message",
  "error": "Detailed error information",
  "timestamp": 1640995200
}
```

### Validation Error Response
```json
{
  "success": false,
  "message": "Validation failed",
  "details": [
    {
      "field": "email",
      "message": "Email is required"
    }
  ],
  "timestamp": 1640995200
}
``` 