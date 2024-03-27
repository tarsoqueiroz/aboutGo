# Simple HTTP Server in Go

## About

> `https://dev.to/kartikmehta8/building-a-rest-api-in-go-4p8a`

This example demonstrates how to create a basic HTTP server in Go that responds with "Hello, World from Go!" when accessed. It showcases the simplicity and power of Go's standard library for web applications.

## Tasks

```sh
docker build -t tarsoqueiroz/tinywebinfo:1.0 .

docker run --publish 8080:8080 tarsoqueiroz/tinywebinfo:1.0

docker image push tarsoqueiroz/tinywebinfo:1.0
docker image push tarsoqueiroz/tinywebinfo:2.0
docker image push tarsoqueiroz/tinywebinfo:3.0 
```
