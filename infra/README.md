# Infrastructure

This directory contains all infrastructure and DevOps related files for the Hotaku API project.

## Directory Structure

```txt
infra/
├── docker/             # Docker configurations
│   ├── docker-compose.yml      # Development environment
│   ├── docker-compose.test.yml # Testing environment
│   ├── Dockerfile              # Production image
│   └── Dockerfile.test         # Testing image
├── migrations/         # Database migration files
│   ├── migrate.go              # Migration utilities
│   └── sql/                    # SQL migration files
│       ├── *.up.sql           # Forward migrations
│       └── *.down.sql         # Rollback migrations
├── scripts/           # Infrastructure scripts
│   ├── setup-env.sh            # Environment setup
│   ├── migrate.sh              # Migration scripts
│   ├── env.example             # Environment template
│   └── init-test-db.sql        # Test database setup
├── config/            # Infrastructure configurations
│   └── .air.toml               # Hot reload configuration
├── Makefile           # Build and development commands
└── README.md          # This file
```

## Usage

### Development Commands

Navigate to the `infra/` directory and use the Makefile:

```bash
cd infra/

# Start development environment
make docker-up

# Run tests
make test

# Run with hot reload
make dev

# Stop environment
make docker-down
```

### Docker Commands

#### Development Environment

```bash
# Start development stack (API + MySQL)
docker compose -f docker/docker-compose.yml up -d

# View logs
docker compose -f docker/docker-compose.yml logs -f

# Stop stack
docker compose -f docker/docker-compose.yml down
```

#### Test Environment

```bash
# Start test stack (API + MySQL + Redis)
docker compose -f docker/docker-compose.test.yml up -d

# Run tests in container
docker compose -f docker/docker-compose.test.yml --profile test run --rm test-runner

# Stop test stack
docker compose -f docker/docker-compose.test.yml down
```

### Database Migrations

```bash
# Run migrations
make migrate-up

# Rollback migrations
make migrate-down

# Check migration status
make migrate-status
```

### Configuration Files

#### Docker Compose

- `docker/docker-compose.yml` - Development environment with MySQL
- `docker/docker-compose.test.yml` - Testing environment with MySQL and Redis

#### Dockerfiles

- `docker/Dockerfile` - Multi-stage production image
- `docker/Dockerfile.test` - Testing image with test dependencies

#### Air Configuration

- `config/.air.toml` - Hot reload configuration for development

### Environment Variables

Copy `scripts/env.example` to your environment file and configure:

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_HOST` | Database host | `localhost` |
| `DB_PORT` | Database port | `3306` |
| `DB_USER` | Database user | `root` |
| `DB_PASSWORD` | Database password | `rootpassword` |
| `DB_NAME` | Database name | `hotaku_db` |
| `JWT_SECRET` | JWT secret key | `your-secret-key` |
| `GIN_MODE` | Gin mode | `debug` |

### Scripts

- `scripts/setup-env.sh` - Environment setup script
- `scripts/migrate.sh` - Database migration script
- `scripts/init-test-db.sql` - Test database initialization

## Best Practices

1. **Always use Docker for development** - Ensures consistent environment
2. **Run tests in isolated environment** - Use test compose file
3. **Version your migrations** - Follow sequential numbering
4. **Document environment variables** - Update env.example when adding new vars
5. **Use Makefile for common tasks** - Simplifies development workflow

## Troubleshooting

### Port Conflicts

- Development: API on 3000, MySQL on 3306
- Testing: API on 3001, MySQL on 3307, Redis on 6380

### Migration Issues

```bash
# Reset migrations (development only)
make test-cleanup
make clean
```

### Docker Issues

```bash
# Clean up everything
docker system prune -f
make clean
```
