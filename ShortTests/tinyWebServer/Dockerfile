FROM golang:1.20 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /tinywebinfo main.go

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY ./views/index.html ./views/index.html
COPY --from=build-stage /tinywebinfo /tinywebinfo

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/tinywebinfo"]
