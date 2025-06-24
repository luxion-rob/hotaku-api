FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

ARG AIR_VERSION=v1.62.0
RUN go install github.com/air-verse/air@${AIR_VERSION}

COPY . .
RUN go build -o main .

FROM alpine:3.19

# Install ca-certificates for HTTPS requests and envsubst for environment variable substitution
RUN apk --no-cache add ca-certificates=20240111-r0 gettext=0.22.3-r0

WORKDIR /app

# Create non-root user for security
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# COPY --from=builder /go/bin/air /usr/local/bin/air
COPY --from=builder /app/main .
COPY .env.example .env

# Create a script to handle environment variables
RUN cat <<'EOF' > /app/entrypoint.sh
#!/bin/sh
set -euo pipefail
envsubst < .env > .env.processed
exec ./main
EOF
RUN chmod +x /app/entrypoint.sh

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

EXPOSE 3000
HEALTHCHECK --interval=30s --timeout=3s --retries=3 CMD wget -qO- http://localhost:3000/health || exit 1
CMD ["/app/entrypoint.sh"] 