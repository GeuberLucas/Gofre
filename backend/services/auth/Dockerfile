FROM golang:1.25.4-alpine3.22 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o auth-microservice ./cmd

# Final stage
FROM scratch
WORKDIR /app
COPY --from=builder /app/auth-microservice .
CMD ["./auth-microservice"]