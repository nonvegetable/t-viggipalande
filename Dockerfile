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

# Expose the port your SSH server runs on
EXPOSE 23234

# Ensure the app can write its generated SSH keys inside the container
RUN mkdir -p /app/.ssh

# Install openssh-client so that the ssh-keygen utility is available
RUN apk add --no-cache openssh-client

# Generate a static ED25519 host key file with no passphrase
RUN ssh-keygen -t ed25519 -f /app/.ssh/id_ed25519 -N ""

# Copy your built binary and run it
COPY --from=builder /app/portfolio .
CMD ["./portfolio"]