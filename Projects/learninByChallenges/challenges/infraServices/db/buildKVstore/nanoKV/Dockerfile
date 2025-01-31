##
## Stage 1: Build the Go application
##
FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY main.go .
COPY kvstore/ ./kvstore/

RUN go mod init nanokv

RUN go get github.com/gofiber/fiber/v2

RUN go build -o /app/nanokv

# Stage 2: Create a minimal runtime environment
FROM scratch

COPY --from=builder /app/nanokv /app/nanokv

# Port exposing the app
EXPOSE 3456

# Command to run the application
CMD ["/app/nanokv"]
