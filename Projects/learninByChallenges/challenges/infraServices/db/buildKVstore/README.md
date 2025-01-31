# How to Build a Concurrent Key-Value Store in Go

## About

> **Source:** [`dev.to/siddheshk02/how-to-build-a-concurrent-key-value-store-in-go-3pep`](https://dev.to/siddheshk02/how-to-build-a-concurrent-key-value-store-in-go-3pep)

Key-Value stores are one of the simplest yet most effective forms of data storage used in modern applications.

In this tutorial, we will walk through building NanoKV, a minimalist in-memory key-value store written in Go, using the Fiber web framework.

By the end of this tutorial, you’ll understand:

- How key-value stores work under the hood
- Implementing basic CRUD operations in Go
- Using concurrency and TTL (Time-To-Live) for ephemeral data storage
- Running the application in a Docker container

## Initialize project

```sh
# init project
mkdir nanoKV && cd nanoKV
go mod init nanokv

# install dependencies
go get github.com/gofiber/fiber/v2

# create project path & files
mkdir kvstore
touch kvstore/kvstore.go
touch main.go
touch Dockerfile
```

## Implementing the KV store

1. Defining the store structure
1. Implementing CRUD operations with TTL support

## Exposing an HTTP API with Fiber

To make our key-value store accessible via HTTP, we set up a simple REST API using Fiber.

1. Setting up the HTTP server
1. Run the app

Test it with:

- Curl

```sh
curl -X POST http://localhost:3000/set/foo/bar
curl -X GET http://localhost:3000/get/foo
curl -X DELETE http://localhost:3000/delete/foo
```

- [`restApiCall.http`](./restApiCall.http)

## Containerizing the app with Docker

- [`Dockerfile`](./nanoKV/Dockerfile)

```sh
docker build -t nanokv .

docker images

docker run --rm -d -p 3456:3456 --name nanokv nanokv

docker ps

docker stop nanokv
```
