FROM golang:1.24-alpine

WORKDIR /app

# Install Air
ARG AIR_VERSION=v1.62.0
RUN wget -qO /tmp/air.tgz "https://github.com/cosmtrek/air/releases/download/${AIR_VERSION}/air_${AIR_VERSION#v}_linux_amd64.tar.gz" && \
    tar -C /usr/local/bin -xzf /tmp/air.tgz air && \
    rm /tmp/air.tgz

# Install build tools and wget for healthcheck
RUN apk --no-cache add git wget

# Create non-root user for security
RUN adduser -D -s /bin/sh -u 1001 appuser

# Set Go module cache to a directory the user can write to
ENV GOMODCACHE=/app/.go-mod-cache
ENV GOPATH=/app/.go
ENV PATH=$PATH:/app/.go/bin

# Prepare writable cache directories
RUN mkdir -p /app/.go-mod-cache /app/.go

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy only the necessary source code directories
COPY cmd/ ./cmd/
COPY config/ ./config/
COPY internal/ ./internal/
COPY utils/ ./utils/
COPY main.go ./

# Copy Air configuration
COPY infra/config/.air.toml ./infra/config/.air.toml

# Create Go cache directories and set permissions
RUN mkdir -p /app/.go-mod-cache /app/.go && \
    chown -R appuser:appuser /app

# Configure Git to work with the mounted repository
RUN git config --global --add safe.directory /app

# Expose the app port
EXPOSE 3000

# Switch to non-root user
USER appuser

# Health check to monitor container health
HEALTHCHECK --interval=60s --timeout=3s --start-period=5s --retries=3 \
    CMD wget -qO- http://localhost:3000/health || exit 1

# Use Air with the config in infra/config/.air.toml
CMD ["air", "-c", "infra/config/.air.toml"]
