# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/portfolio .

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the built binary
COPY --from=builder /app/portfolio .

# Expose the port your SSH server runs on
EXPOSE 23234

# Ensure the app can write its generated SSH keys inside the container
RUN mkdir -p /app/.ssh

# Run the app
CMD ["./portfolio"]
