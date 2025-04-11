# Use the official Golang image as the base image
FROM golang:1.24

# Set the working directory inside the container
WORKDIR /app

# Copy all Go source files.
COPY . .

# Set environment variables for port
ENV PORT=8080

# Run the Go application directly
CMD ["go", "run", "main.go"]
