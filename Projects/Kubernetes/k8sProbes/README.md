# Understanding Kubernetes Probes By Deploying A Go App

## About

> [`https://dev.to/olymahmud/understanding-kubernetes-probes-by-deploying-a-go-app-716`](https://dev.to/olymahmud/understanding-kubernetes-probes-by-deploying-a-go-app-716)

## Probes in Action: A Simple Project

We will explore a professional implementation of probes in a generic Go project. This involves creating a simple Go API, containerizing it with Docker, and deploying it to Kubernetes with detailed probe configurations.

## Creating the Simple Go API

```sh
# Create a new Go project
go mod init k8s-probes-golang
# Create a main.go file
touch main.go
```

- [main.go](./main.go)

This API runs on port 8080, offering endpoints for liveness (/healthz), readiness (/readyz), and a root check (/), with a 10-second delay before becoming ready.

- Dockerization: Building and Pushing the Image

Use [Dockerfile](./Dockerfile) to containerize the application.

Steps:

```sh
# Build the Docker image
docker build -t tarsoqueiroz/k8s-probes-golang:latest .

# Test the image locally
docker run -d --name k8s-probes-golang --rm -p 8888:8080 tarsoqueiroz/k8s-probes-golang:latest

# Verify the container is running
## Should return "Hello, Kubernetes!"
curl http://localhost:8888/
## Should return "OK"
curl http://localhost:8888/healthz
## Should return "Not Ready", then "Ready" after 10 seconds
curl http://localhost:8888/readyz

# Stop the container
docker stop k8s-probes-golang

# Push the image to Docker Hub
docker push tarsoqueiroz/k8s-probes-golang:latest
```

## That's all

...folks!!!
