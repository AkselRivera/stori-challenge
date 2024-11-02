FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . . 

WORKDIR /app/cmd/api

RUN go mod tidy

RUN go build -o /app/migration_service

# Stage 2

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/migration_service /app/migration_service

RUN touch /app/.env

EXPOSE 8080

CMD ["/app/migration_service"]

