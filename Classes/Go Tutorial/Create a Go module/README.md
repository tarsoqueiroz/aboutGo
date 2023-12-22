# Tutorial: Create a Go module

## About

> `https://go.dev/doc/tutorial/create-module`

This is the first part of a tutorial that introduces a few fundamental features of the Go language. If you're just getting started with Go, be sure to take a look at [Tutorial: Get started with Go](https://go.dev/doc/tutorial/getting-started.html), which introduces the go command, Go modules, and very simple Go code.

For more about developing modules, see [Developing and publishing modules](https://go.dev/doc/modules/developing).

## Hands-on

### Create a Go module

Start by creating a Go module. In a module, you collect one or more related packages for a discrete and useful set of functions. For example, you might create a module with packages that have functions for doing financial analysis so that others writing financial applications can use your work.

```sh
cd <base path>
mkdir greetings
cd greetings

go mod init example.com/greetings
## file `greeting.go`
```

### Call your code from another module

In this section, you'll write code to make calls to the Hello function in the module you just wrote. You'll write code you can execute as an application, and which calls code in the greetings module.

```sh
cd ..
mkdir hello
cd hello

go mod init example.com/hello
## file `hello.go`

go mod edit -replace example.com/greetings=../greetings

go mod tidy

go run .
```

### Return and handle an error

Handling errors is an essential feature of solid code. In this section, you'll add a bit of code to return an error from the greetings module, then handle it in the caller.

```sh
## file `greetings.go`

## file `hello.go`

go run .
```

### Return a random greeting

In this section, you'll change your code so that instead of returning a single greeting every time, it returns one of several predefined greeting messages.

```sh
## file `greetings.go`

## file `hello.go`

go run .

go run .

go run .

go run .

go run .
```

### Return greetings for multiple people

In the last changes you'll make to your module's code, you'll add support for getting greetings for multiple people in one request. In other words, you'll handle a multiple-value input, then pair values in that input with a multiple-value output. To do this, you'll need to pass a set of names to a function that can return a greeting for each of them.

```sh
## file `greetings.go`

## file `hello.go`

go run .
```

### Add a test

Now that you've gotten your code to a stable place (nicely done, by the way), add a test. Testing your code during development can expose bugs that find their way in as you make changes. In this topic, you add a test for the Hello function.

```sh
## file `greetings_test.go`

cd ../greetings/

go test

go test -v

## file `greetings.go`

go test

go test -v
```

### Compile and install the application

In this last topic, you'll learn a couple new go commands. While the go run command is a useful shortcut for compiling and running a program when you're making frequent changes, it doesn't generate a binary executable.

```sh

```

### Conclusion

In this tutorial, you wrote functions that you packaged into two modules: one with logic for sending greetings; the other as a consumer for the first.
