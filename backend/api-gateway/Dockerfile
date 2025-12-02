FROM golang:1.25.4-alpine3.22 AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o micro-service ./cmd

# Final stage
FROM scratch
WORKDIR /app
COPY --from=builder /app/micro-service .
CMD ["./micro-service"]