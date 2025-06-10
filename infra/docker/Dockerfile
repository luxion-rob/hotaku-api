FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY infra/scripts/.env* ./

COPY . .
RUN go build -o main .

FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 3000

CMD ["./main"] 