# Understanding Iterators in Go: A Fun Dive!

> `https://dev.to/tuna99/understanding-iterators-in-go-a-fun-dive-1c57`

If you're a Go programmer, you've probably heard about iterators many times in Go 1.22, and especially in Go 1.23. But maybe you're still scratching your head, wondering why they're useful or when you should use them. Well, you're in the right place! Let's start by looking at how iterators work in Go and why they can be so useful.

## Starting

```sh
go mod init iteratoringo

touch main.go
```

## A Simple Transformation: No Iterator Yet

- Code for [`main.go`](./_versions/v1.noiterator/main.go)

```sh
go run main.go

mkdir -p _versions/v1.noiterator
cp main.go _versions/v1.noiterator/.
```

## Enter the Iterator

- Code for [`main.go`]()

```sh
go run main.go

mkdir -p _versions/v2.theiterator
cp main.go _versions/v2.theiterator/.
```

## That's all

...folks!!!
