# Start from a minimal base image with Go installed
FROM golang:1.22-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Start a new stage with a minimal base image
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /app

# Copy the executable from the builder stage
COPY --from=builder /app/app .

EXPOSE 5000

# Run the Go application
CMD ["./app"]