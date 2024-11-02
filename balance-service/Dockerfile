FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . . 

WORKDIR /app/cmd/api

RUN go mod tidy

RUN go build -o /app/balance_service

# Stage 2

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/balance_service /app/balance_service

RUN touch /app/.env

EXPOSE 8080

CMD ["/app/balance_service"]

