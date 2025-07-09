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

## Environment Configuration

The application uses environment variables for configuration with sensible defaults. You can configure the application by setting these environment variables or creating a `.env` file.

### Environment Variables

#### Database Configuration

- `DB_HOST` - Database host (default: `localhost`)
- `DB_PORT` - Database port (default: `3306`)
- `DB_USER` - Database username (default: `root`)
- `DB_PASSWORD` - Database password (default: `password`)
- `DB_NAME` - Database name (default: `hotaku_db`)

#### Server Configuration

- `PORT` - Server port (default: `3000`)
- `GIN_MODE` - Gin framework mode: `debug`, `release`, or `test` (default: `debug`)

#### Application Configuration

- `APP_NAME` - Application name (default: `Hotaku API`)
- `APP_VERSION` - Application version (default: `1.0.0`)
- `APP_ENV` - Environment: `development`, `staging`, or `production` (default: `development`)

### Creating a .env File

#### Quick Setup (Recommended)

Use the provided setup script to create your `.env` file:

```bash
./scripts/setup-env.sh
```

#### Manual Setup

Alternatively, create a `.env` file manually in the root directory with your custom configuration:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password_here
DB_NAME=hotaku_db

# Server Configuration
PORT=3000
GIN_MODE=debug

# Application Configuration
APP_NAME=Hotaku API
APP_VERSION=1.0.0
APP_ENV=development
```

**Note:** The `.env` file is ignored by git for security reasons. Never commit sensitive information like database passwords to version control.

### Using godotenv (Optional)

For automatic `.env` file loading, you can install the `godotenv` package:

```bash
go get github.com/joho/godotenv
```

Then load it in your `main.go`:

```go
import "github.com/joho/godotenv"

func main() {
    // Load .env file if it exists
    godotenv.Load()
    
    // Rest of your application...
}
```

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

1. Create new migration files in the `infra/migrations/` directory
2. Follow the naming convention:
   - `XXXXXX_description.up.sql` for the forward migration
   - `XXXXXX_description.down.sql` for the rollback migration
3. Example:
   - `000002_add_posts_table.up.sql`
   - `000002_add_posts_table.down.sql`

## API Endpoints

### Public Endpoints

- `GET /health` - Health check endpoint

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
├── controllers/        # HTTP handlers and tests
├── config/             # Configuration files
├── utils/              # Utility functions
├── cmd/                # CLI commands
├── docs/               # Documentation and GitHub Pages
├── .github/workflows/  # GitHub Actions workflows
├── infra/              # Infrastructure and DevOps files
│   ├── docker/         # Docker configurations
│   │   ├── docker-compose.yml      # Development environment
│   │   ├── docker-compose.test.yml # Testing environment
│   │   ├── Dockerfile              # Production image
│   │   └── Dockerfile.test         # Testing image
│   ├── migrations/     # Database migration files
│   │   ├── migrate.go              # Migration utilities
│   │   └── sql/                    # SQL migration files
│   │       ├── 000001_create_users_table.up.sql
│   │       └── 000001_create_users_table.down.sql
│   ├── scripts/        # Infrastructure scripts
│   │   ├── setup-env.sh            # Environment setup
│   │   ├── migrate.sh              # Migration scripts
│   │   ├── env.example             # Environment template
│   │   └── init-test-db.sql        # Test database setup
│   ├── config/         # Infrastructure configurations
│   │   └── .air.toml               # Hot reload configuration
│   └── Makefile        # Development and build commands
```

## Testing

The project includes comprehensive test coverage with both unit tests and integration tests:

### Unit Tests

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

### Docker Testing Environment

A complete testing environment is available using Docker Compose:

```bash
# Start test environment (MySQL + Redis + API)
make test-env

# Run tests in containerized environment
make test-run

# Run integration tests against test API
make test-integration

# Stop test environment
make test-env-down

# Clean up test environment completely
make test-cleanup
```

The test environment includes:

- **MySQL 8.0** on port 3307 (isolated from development)
- **Redis 7** on port 6380 (for caching/sessions)
- **API service** on port 3001 (test configuration)
- **Test runner** service for automated testing

#### Test Environment Configuration

- Database: `hotaku_test_db`
- MySQL User: `root` / Password: `testpassword`
- Test API URL: `http://localhost:3001`
- All services run in isolated `test-network`tecurity**: Gosec security scanning
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
