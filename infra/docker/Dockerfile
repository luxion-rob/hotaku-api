FROM golang:1.24-alpine

WORKDIR /app

# Install Air
ARG AIR_VERSION=v1.62.0
RUN wget -qO /tmp/air.tgz "https://github.com/cosmtrek/air/releases/download/${AIR_VERSION}/air_${AIR_VERSION#v}_linux_amd64.tar.gz" && \
    tar -C /usr/local/bin -xzf /tmp/air.tgz air && \
    rm /tmp/air.tgz

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Expose the app port
EXPOSE 3000

# Use Air with the config in infra/config/.air.toml
CMD ["air", "-c", "infra/config/.air.toml"]
