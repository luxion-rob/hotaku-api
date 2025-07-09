# Hotaku API

A modern, production-ready manga management API built with Go, Gin framework, and MySQL. Features user authentication, file uploads with MinIO, and comprehensive security measures.

## 🚀 Features

- **🔐 JWT Authentication** - Secure user registration, login, and profile management
- **📁 File Management** - MinIO integration for manga images and chapter pages
- **🛡️ Security** - Path traversal protection, input validation, and secure middleware
- **🗄️ Database** - MySQL with comprehensive migration system
- **🐳 Docker** - Complete containerized development and production environments
- **📊 Monitoring** - Health checks and comprehensive logging
- **🧪 Testing** - Unit and integration tests with coverage reporting

## 🏗️ Architecture

The project follows clean architecture principles with clear separation of concerns:

```txt
hotaku-api/
├── cmd/                    # CLI commands and entry points
├── config/                 # Configuration management
├── internal/               # Private application code
│   ├── controllers/        # HTTP request handlers
│   ├── domain/            # Business entities and DTOs
│   ├── middleware/        # HTTP middleware (auth, path sanitization)
│   ├── repo/              # Data access layer
│   ├── service/           # Business logic services
│   ├── server/            # HTTP server setup and routing
│   └── usecase/           # Application use cases
├── infra/                 # Infrastructure and DevOps
│   ├── docker/            # Docker configurations
│   ├── migrations/        # Database migrations and seeds
│   └── scripts/           # Development and deployment scripts
├── utils/                 # Utility functions
└── main.go               # Application entry point
```

## 🛠️ Tech Stack

- **Language**: Go 1.24.0
- **Framework**: Gin v1.10.1
- **Database**: MySQL 8.0
- **ORM**: GORM v1.26.0
- **File Storage**: MinIO
- **Authentication**: JWT
- **Containerization**: Docker & Docker Compose
- **Migrations**: golang-migrate/v4

## 🚀 Quick Start

### Prerequisites

- Docker and Docker Compose
- Go 1.24+ (for local development)
- Make (optional, for convenience commands)

### 1. Clone and Setup

```bash
git clone <repository-url>
cd hotaku-api
```

### 2. Environment Setup

```bash
# Setup environment files
make setup-env-files

# Or manually copy the example
cp infra/scripts/env.example .env
```

### 3. Start Development Environment

```bash
# Start all services (MySQL, MinIO, API)
make dev-setup

# Or step by step:
make docker-up
make migrate-up
make setup-minio
```

### 4. Access the API

- **API**: http://localhost:3000
- **MinIO Console**: http://localhost:9001
- **MySQL**: localhost:3306

## 📚 API Endpoints

### Public Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/health` | Health check |
| `POST` | `/api/v1/auth/register` | User registration |
| `POST` | `/api/v1/auth/login` | User login |
| `GET` | `/api/v1/images/*` | Public image access |

### Protected Endpoints (Require JWT)

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/api/v1/auth/profile` | Get user profile |
| `PUT` | `/api/v1/auth/profile` | Update user profile |
| `PUT` | `/api/v1/auth/change-password` | Change password |
| `POST` | `/api/v1/upload/manga/:id/image` | Upload manga image |
| `POST` | `/api/v1/upload/manga/:id/chapters/:chapter_id/pages` | Upload chapter pages |
| `DELETE` | `/api/v1/upload/files/*` | Delete file |
| `GET` | `/api/v1/upload/files/*` | Get file info |

## 🔧 Configuration

### Environment Variables

Create a `.env` file in the root directory:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME=hotaku_db

# Server Configuration
PORT=3000
GIN_MODE=debug

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key

# MinIO Configuration
MINIO_ENDPOINT=localhost:9000
MINIO_ACCESS_KEY_ID=minioadmin
MINIO_SECRET_ACCESS_KEY=minioadmin
MINIO_USE_SSL=false
MINIO_BUCKET_NAME=manga-images

# Application Configuration
APP_NAME=Hotaku API
APP_VERSION=1.0.0
APP_ENV=development
```

## 🗄️ Database

### Migration Commands

```bash
# Run migrations
make migrate-up

# Rollback migrations
make migrate-down version=5

# Check migration status
make migrate-status

# Refresh migrations (rollback all + run from start)
make migrate-refresh
```

### Database Schema

The application includes a comprehensive manga management schema:

- **Users & Authentication**: User accounts with role-based access
- **Manga Management**: Manga metadata, status, and relationships
- **Content Management**: Chapters, pages, and file storage
- **User Interactions**: Favorites, reading progress, notifications
- **Metadata**: Authors, categories, groups, and statuses

## 🔐 Security Features

### Path Traversal Protection

The API includes middleware that prevents directory traversal attacks:

```go
// Automatically sanitizes wildcard parameters
images.GET("/*object_name", uploadController.GetImage)
```

### Authentication Middleware

JWT-based authentication with secure token validation:

```go
protected.Use(authMiddleware)
```

### Input Validation

Comprehensive validation for all user inputs and file uploads.

## 🐳 Docker

### Development

```bash
# Start development environment
make docker-up

# View logs
docker compose -f infra/docker/docker-compose.yml logs -f

# Stop services
make docker-down
```

### Production

```bash
# Generate production secrets
make generate-secrets

# Start production environment
make docker-prod-up
```

## 🧪 Testing

### Run Tests

```bash
# Unit tests
go test ./...

# Tests with coverage
go test -v -race -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# Integration tests
go test -v ./internal/controllers/...
```

### Test Environment

```bash
# Start test environment
make test-env

# Run tests in containerized environment
make test-run

# Clean up test environment
make test-cleanup
```

## 📦 Development Commands

```bash
# Format code
go fmt ./...

# Run linter
golangci-lint run

# Build application
go build -o bin/hotaku-api main.go

# Run with hot reload (requires Air)
air
```

## 🚀 Deployment

### Production Build

```bash
# Build production image
docker build -f infra/docker/Dockerfile -t hotaku-api .

# Run production container
docker run -p 3000:3000 --env-file .env hotaku-api
```

### Environment Setup

1. Set up MySQL database
2. Configure MinIO storage
3. Set environment variables
4. Run database migrations
5. Start the application

## 📊 Monitoring

### Health Check

```bash
curl http://localhost:3000/health
```

Response:
```json
{
  "status": "healthy",
  "message": "API is running smoothly",
  "timestamp": 1640995200,
  "version": "1.0.0"
}
```

### Logging

The application includes structured logging for:
- Request/response logging
- Error tracking
- Performance monitoring
- Security events

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow Go coding standards
- Write tests for new features
- Update documentation
- Use conventional commit messages
- Ensure all tests pass before submitting PR

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🔗 Links

- [API Documentation](./docs/index.md)
- [Docker Hub](https://hub.docker.com/r/your-username/hotaku-api)
- [Issue Tracker](https://github.com/your-username/hotaku-api/issues)

---

Built with ❤️ using Go and Gin
