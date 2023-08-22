# How to build a CLI tool with Go and Cobra

## About

- [Part I](https://dev.to/divrhino/building-a-command-line-tool-with-go-and-cobra-3mjd)
- [Part II](https://dev.to/divrhino/adding-flags-to-a-command-line-tool-built-with-go-and-cobra-34f1)

## Part I

```sh
# Go init
go mod init dadjoke

# Cobra init
~/go/bin/cobra-cli init

# Additional Go struct path
mkdir bin pkg

# First tests
go run main.go 
go run main.go --help
go run main.go -h
go run main.go -v
go run main.go completion bash

# Adding command
~/go/bin/cobra-cli add random

# Second tests
go run main.go -h
go run main.go random -h

# API calls
curl -H "Accept: text/html" https://icanhazdadjoke.com/
curl -H "Accept: text/plain" https://icanhazdadjoke.com/
curl -H "Accept: application/json" https://icanhazdadjoke.com/

# Final test
go run main.go random 
go run main.go random 
```

## Part II

