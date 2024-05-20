# Get Started with Golang

## Table of Content

- [About](#about)
- General tips about Go
  - [Simple project setup](#simple-project-setup)
  - [FullCycle - First program](#fullcycle---first-program)
  - [FullCycle - HP, Multithreading e Profiling](#fullcycle---hp-multithreading-e-profiling)
  - [FullCycle - API, gRPC and Profiling](#fullcycle---api-grpc-and-profiling)
  - [Vulnerability check](#vulnerability-check)
- Codes and Sample
  - [Sample web server](#sample-web-server)

## About

How to start a code in Golang?

This own class was made with insights from many sources, exposing steps (there's no imposition) to start a Golang code/project.

Resources and codes on [`resources`](./resources/) directory.

[Return to top](#get-started-with-golang)

## Simple project setup

> Source: [Simple Project Setup in Go](https://dev.to/anuragaffection/simple-project-setup-in-go-43en)
>
> Others:
>
> - [Estrutura de projetos Go](https://dev.to/erick_tmr/estrutura-de-projetos-go-4o7l)

```sh
# Check the Go version
go version 

# Create a folder
mkdir <folder_name>

# Switch to created folder
cd <folder_name>

# Initialize the go project
go mod init <folder_name>

# Create a main.go file
touch main.go
```

- Sample and basic code in main.go file.

```go
package main 

import (
    "fmt"
)

func main () {
    fmt.Println("Hello From Go")
}
```

```sh
## Run the code
go run main.go 
```

[Return to top](#get-started-with-golang)

## Vulnerability check

```sh
go install golang.org/x/vuln/cmd/govulncheck@latest

govulncheck ./...
```

[Return to top](#get-started-with-golang)

## Sample web server

> [Creating Simple Server in Go - 02](https://dev.to/anuragaffection/creating-simple-server-in-go-53b5)

```sh
mkdir <folder_name>
cd <folder_name>
go mod init <folder_name>
# create/edit main.go
go run main.go
```

- `main.go`

```go
package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello From Go \n")
  fmt.Fprintf(w, "You are on %s route", r.URL.Path)
}

func main() {
  fmt.Println("Your Server is running on http://localhost:8000")

  http.HandleFunc("/", handler)
  http.ListenAndServe(":8000", nil)
}
```

[Return to top](#get-started-with-golang)

## FullCycle - First program

> [Golang : Oportunidades e Criando nosso primeiro programa](https://www.youtube.com/watch?v=TGnyHDPe3sA)

```sh
# init project
go mod init fullcycle1st

# Compiling
go build 

GOOS=windows go build -o "hello.exe"

file fullcycle hello.exe
```

[Return to top](#get-started-with-golang)

## FullCycle - HP, Multithreading e Profiling

> [Go: Alta Performance, Multithreading e Profiling na prática](https://www.youtube.com/watch?v=Yk_dp-X7c6s)

```sh
go mod init fullcyclehp

touch main.go

time go run main.go # v1

time go run main.go # v2
```

Steps:

- Process -> Alloc a memory block
- Thread -> Access the memory block
- T1 and T2 -> Access the same memory block
- Race condition

Go Routine:

- Go Routine 1 -> Channel -> Go Routine 2
- Input -> Output

```sh
go run main.go # v3

go run main.go # v4 - deadlock!

time go run main.go # v5

time go run main.go # v6

time go run main.go # v7

time go run main.go # v8
```

[Return to top](#get-started-with-golang)

## FullCycle - API, gRPC and Profiling

> [Go: API sem Framework, gRPC e Profiling + Pós Go Expert](https://www.youtube.com/watch?v=4w1pXqJMoA0)

[Return to top](#get-started-with-golang)

## to be continue

[Return to top](#get-started-with-golang)
