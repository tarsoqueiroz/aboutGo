package main

import "fmt"

// interface with two methods
type Vehicle interface {
	Start() string
	Stop() string
}

// implementation for a Car
type Car struct{}

func (c Car) Start() string {
	return "Car started!"
}
func (c Car) Stop() string {
	return "Car stopped!"
}

// implementation for a Mothorcycle
type Mothorcycle struct{}

func (m Mothorcycle) Start() string {
	return "Mothorcycle running!"
}
func (m Mothorcycle) Stop() string {
	return "Mothorcycle stopped!"
}

func main() {
	var car Vehicle = Car{}

	fmt.Println(car.Start())
	fmt.Println(car.Stop())

	var motho Vehicle = Mothorcycle{}
	fmt.Println(motho.Start())
	fmt.Println(motho.Stop())
}
