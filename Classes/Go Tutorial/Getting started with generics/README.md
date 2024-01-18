# Tutorial: Getting started with generics

## About

> `https://go.dev/doc/tutorial/generics`

This tutorial introduces the basics of generics in Go. With generics, you can declare and use functions or types that are written to work with any of a set of types provided by calling code.

## Prerequisites

- An installation of Go 1.18 or later.
- A tool to edit your code.
- A command terminal.

## Folder for your code

> Current folder (`aboutGo/Classes/Go Tutorial/Getting started with generics`)

```sh
go mod init example/generics
```

## Add non-generic functions

In this step, you’ll add two functions that each add together the values of a map and return the total.

> [`main.go.v1`](./releases/main.go.v1)

```sh
go run .
```

## Add a generic function to handle multiple types

In this section, you’ll add a single generic function that can receive a map containing either integer or float values, effectively replacing the two functions you just wrote with a single function.

> [`main.go.v2`](./releases/main.go.v2)

```sh
go run .
```

## Remove type arguments when calling the generic function

In this section, you’ll add a modified version of the generic function call, making a small change to simplify the calling code. You’ll remove the type arguments, which aren’t needed in this case.

> [`main.go.v3`](./releases/main.go.v3)

```sh
go run .
```

## Declare a type constraint

In this last section, you’ll move the constraint you defined earlier into its own interface so that you can reuse it in multiple places. Declaring constraints in this way helps streamline code, such as when a constraint is more complex.

> [`main.go.v4`](./releases/main.go.v4)

```sh
go run .
```

## Conclusion

Nicely done! You’ve just introduced yourself to generics in Go.
