FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

ARG AIR_VERSION=v1.62.0
RUN set -eu; \
    url="https://github.com/cosmtrek/air/releases/download/${AIR_VERSION}/air_${AIR_VERSION#v}_linux_amd64.tar.gz"; \
    sha256="3322883bd2c12c56ac7df605bd4ad9d9f99ca71c8e8d3024a73c28d47a365491"; \
    wget -qO /tmp/air.tgz "${url}" && \
    echo "${sha256}  /tmp/air.tgz" | sha256sum -c - && \
    tar -C /usr/local/bin -xzf /tmp/air.tgz air && \
    rm /tmp/air.tgz

COPY . .
RUN go build -o main .

FROM alpine:3.19

# Install ca-certificates for HTTPS requests and envsubst for environment variable substitution
RUN apk --no-cache add ca-certificates=20240111-r0 gettext=0.22.3-r0

WORKDIR /app

# Create non-root user for security
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

COPY --from=builder /app/main .
COPY .env.example .env

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

EXPOSE 3000
HEALTHCHECK --interval=30s --timeout=3s --retries=3 CMD wget -qO- http://localhost:3000/health || exit 1
CMD ["/app/entrypoint.sh"] 