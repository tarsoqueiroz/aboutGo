# Melange and a Docker Image with Apko

> `https://dev.to/patrickdomnick/building-a-go-package-with-melange-and-a-docker-image-with-apko-141c`

In this tutorial, we will learn how to create a Go package with Melange, build a Docker with Apko image, and showcase examples using GitLab CI.

## Prerequisites

Before we begin, make sure you have the following tools installed on your system:

- [Docker](https://docs.docker.com/get-docker/)
- [Golang](https://go.dev/doc/install)
- [Melange](https://github.com/chainguard-dev/melange)
- [Apko](https://github.com/chainguard-dev/apko)

Using Golang version `go version go1.23.0 linux/amd64` managed by `gvm`.

Melange install:

```sh
go install chainguard.dev/melange@latest
```

Apki install:

```sh
go install chainguard.dev/apko@latest
```

## Keyfile generation

Before generate it's important include on `.gitignore` two entrances:

```txt
*.rsa
*.rsa.pub
```

We need to create a private and public key pair to sign our artifacts:

```sh
melange keygen
```

## Creating a Go Package with Melange

Initialize a Go module:

```sh
go mod init gitlab.com/tarsoqueiroz/golang-apko-example
```

Create your Go source code files and write your package logic.

- File [`main.go`](./main.go): simple web server.

Create a [`melange.yaml`](./melange.yaml) file in the root of your package directory. This file will define the build configuration for your package.

Build your package locally with Melange:

```sh
melange build --signing-key melange.rsa --runner docker melange.yaml
```

This will generate an `APKINDEX.json`, `APKINDEX.tar.gz` and `.apk` file in your packages directory.

## Building a Docker Image with `apko.yaml`

Create an [`apko.yaml`](./apko.yaml) file in the root of your package directory. This file will define the Docker image build configuration.

The important line is the golang-apko-example@local package signaling a local package should be used instead of the wolfi packages.

Build the Image with your local repository appended:

```sh
apko build --log-level debug --sbom-path ./sbom/ --repository-append $(pwd)/packages --keyring-append=melange.rsa.pub apko.yaml golang-apko-example:latest image.tar
```

This will build a Docker image as an `image.tar` with our local package and create SBOM files for our image.

To actually use the image locally we have to load and execute it:

```sh
docker load --input image.tar
docker run --rm --name apk -p 8910:8910 -d golang-apko-example:latest-amd64
```
