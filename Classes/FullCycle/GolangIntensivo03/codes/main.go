package main

import "fmt"

type Order struct {
	ID       string
	Price    float64
	Quantity int
}

func main() {
	strName := "Tarso" // string
	order := Order{
		ID:       "123",
		Price:    10.0,
		Quantity: 5,
	}

	println("Hello world!!!")
	println(strName)
	fmt.Println(order.ID, order.Price, order.Price)
}
