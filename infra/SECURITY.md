# Security Best Practices

## Environment Variables and Secrets Management

### ⚠️ NEVER copy .env files into Docker images

**Problem:** The original Dockerfile contained this vulnerable line:
```dockerfile
COPY infra/scripts/.env* ./
```

This copies sensitive environment files directly into the Docker image, making secrets extractable by anyone with access to the image.

### ✅ Recommended Approaches

#### 1. Development Environment
Use `env_file` in docker-compose.yml to load environment variables at runtime:

```yaml
services:
  api:
    env_file:
      - ../scripts/.env
    environment:
      - DB_PASSWORD=${DB_PASSWORD}  # No defaults for secrets
      - JWT_SECRET=${JWT_SECRET}    # No defaults for secrets
```

#### 2. Production Environment
Use Docker secrets for sensitive data:

```yaml
services:
  api:
    environment:
      - DB_PASSWORD_FILE=/run/secrets/db_password
      - JWT_SECRET_FILE=/run/secrets/jwt_secret
    secrets:
      - db_password
      - jwt_secret

secrets:
  db_password:
    external: true  # Managed externally (e.g., Docker Swarm, Kubernetes)
  jwt_secret:
    external: true
```

#### 3. Cloud Deployment
- **AWS:** Use AWS Secrets Manager, Parameter Store, or IAM roles
- **GCP:** Use Secret Manager or Workload Identity
- **Azure:** Use Key Vault or Managed Identity
- **Kubernetes:** Use native Secrets or external secret operators

### Security Hardening Applied

#### Container Security
- **Non-root user:** Application runs as user ID 1001
- **Capability dropping:** Removes unnecessary Linux capabilities
- **Read-only filesystem:** Consider enabling where possible
- **Network isolation:** Uses custom networks

#### Image Security
- **Multi-stage builds:** Reduces final image size and attack surface
- **Alpine base:** Smaller, more secure base image
- **CA certificates:** Included for HTTPS communications

### File Structure for Secrets

```
infra/
├── secrets/           # Local development secrets (gitignored)
│   ├── db_password.txt
│   └── jwt_secret.txt
├── scripts/
│   ├── .env           # Development env vars (gitignored)
│   └── env.example    # Template for required variables
└── docker/
    ├── docker-compose.yml      # Development
    └── docker-compose.prod.yml # Production with secrets
```

### Best Practices

1. **Never commit secrets to version control**
2. **Use different secrets for each environment**
3. **Rotate secrets regularly**
4. **Use strong, randomly generated passwords**
5. **Monitor secret access and usage**
6. **Use least-privilege principles**

### Application Code Updates Needed

To support Docker secrets, update your configuration loading to check for `*_FILE` environment variables:

```go
func loadSecretFromFile(envVar string) string {
    if fileVar := os.Getenv(envVar + "_FILE"); fileVar != "" {
        content, err := os.ReadFile(fileVar)
        if err != nil {
            log.Printf("Warning: could not read secret file %s: %v", fileVar, err)
            return os.Getenv(envVar)
        }
        return strings.TrimSpace(string(content))
    }
    return os.Getenv(envVar)
}

// Usage:
dbPassword := loadSecretFromFile("DB_PASSWORD")
jwtSecret := loadSecretFromFile("JWT_SECRET")
```

### Migration Steps

1. **Remove .env copying from Dockerfile** ✅ Done
2. **Update docker-compose.yml to use env_file** ✅ Done
3. **Create production docker-compose with secrets** ✅ Done
4. **Update application code to handle secret files** (Recommended)
5. **Add secrets directory to .gitignore**
6. **Document environment setup for team**

### Verification

To verify secrets are not in your image:
```bash
# Check image layers
docker history your-image-name

# Check for files in image
docker run --rm your-image-name find / -name "*.env*" 2>/dev/null
``` 