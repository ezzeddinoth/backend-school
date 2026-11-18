# Stage 1: Build the binary
FROM golang:1.27-alpine AS builder
WORKDIR /app

# Copy dependency files first to cache downloads
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /backend-school ./cmd/main.go

# Stage 2: Run the binary in a minimal image
FROM alpine:latest
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /backend-school .

# Expose your application port (if it's a web server)
EXPOSE 8080

# Command to run when the container starts
CMD ["./backend-school"]


# podman build -t backend-school .

# podman run -it --rm  --network school-net -p 8080:8080 -v ./cmd/.env:/root/.env:Z --name running-backend-school backend-school

# podman logs running-backend-school