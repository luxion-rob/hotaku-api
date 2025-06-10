# Hotaku API Documentation

Welcome to the Hotaku API documentation. This is a RESTful API built with Go and the Gin framework.

## 🚀 Quick Start

### Base URL

```
https://your-domain.com/api
```

### Authentication

This API uses JWT (JSON Web Tokens) for authentication. Include the token in the Authorization header:

```
Authorization: Bearer <your-jwt-token>
```

## 📚 API Endpoints

### Health Check

Check if the API is running and healthy.

**Endpoint:** `GET /`

**Response:**

```json
{
  "status": "healthy",
  "message": "API is running smoothly",
  "timestamp": 1640995200,
  "version": "1.0.0"
}
```

## 🏗️ Project Structure

```
hotaku-api/
├── controllers/          # HTTP handlers
├── config/              # Configuration files
├── utils/               # Utility functions
├── cmd/                 # CLI commands
├── docs/                # Documentation
├── .github/workflows/   # GitHub Actions
├── infra/               # Infrastructure and DevOps
│   ├── docker/         # Docker configurations
│   ├── migrations/     # Database migrations
│   ├── scripts/        # Infrastructure scripts
│   ├── config/         # Infrastructure configs
│   └── Makefile        # Build commands
├── main.go              # Application entry point
├── go.mod              # Go module file
└── README.md           # Project README
```

## 🧪 Testing

The API includes comprehensive test coverage:

- Unit tests for controllers
- Integration tests for API endpoints
- Coverage reporting
- Automated testing in CI/CD pipeline

To run tests locally:

```bash
go test -v ./...
```

To run tests with coverage:

```bash
go test -v -race -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

## 🐳 Docker

### Development

```bash
docker compose up --build
```

### Production

```bash
docker build -t hotaku-api .
docker run -p 3000:3000 hotaku-api
```

## 🔧 Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_HOST` | Database host | `localhost` |
| `DB_PORT` | Database port | `3306` |
| `DB_USER` | Database user | `root` |
| `DB_PASSWORD` | Database password | `rootpassword` |
| `DB_NAME` | Database name | `hotaku_db` |
| `JWT_SECRET` | JWT secret key | `your-secret-key` |
| `GIN_MODE` | Gin mode | `debug` |

## 🚀 Deployment

### GitHub Actions CI/CD

This project includes a comprehensive CI/CD pipeline that:

1. **Testing**: Runs tests with MySQL service
2. **Building**: Builds the Go application
3. **Security**: Scans for security vulnerabilities
4. **Docker**: Builds and pushes Docker images
5. **Documentation**: Generates and deploys documentation to GitHub Pages

### Pipeline Status

The pipeline runs on:

- Push to `main` or `develop` branches
- Pull requests to `main`

## 📖 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🔗 Links

- [Repository](https://github.com/your-username/hotaku-api)
- [Docker Hub](https://hub.docker.com/r/your-username/hotaku-api)
- [Coverage Reports](./coverage.html)

---

Documentation generated automatically by GitHub Actions
