package main

import "fmt"

// interface definition
type Animal interface {
	Speak() string
}

// structs implementing the interface
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	var a Animal

	a = Dog{}
	fmt.Println("Dog says:", a.Speak())

	a = Cat{}
	fmt.Println("Cat says:", a.Speak())
}
