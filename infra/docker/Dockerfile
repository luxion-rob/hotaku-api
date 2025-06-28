FROM golang:1.24-alpine

WORKDIR /app

# Install Air
ARG AIR_VERSION=v1.62.0
RUN wget -qO /tmp/air.tgz "https://github.com/cosmtrek/air/releases/download/${AIR_VERSION}/air_${AIR_VERSION#v}_linux_amd64.tar.gz" && \
    tar -C /usr/local/bin -xzf /tmp/air.tgz air && \
    rm /tmp/air.tgz

# Install build tools
RUN apk --no-cache add git

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

# Expose the app port
EXPOSE 3000

# Use Air with the config in infra/config/.air.toml
CMD ["air", "-c", "infra/config/.air.toml"]
