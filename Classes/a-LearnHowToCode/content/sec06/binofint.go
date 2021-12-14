package main

import "fmt"

func main() {

	var number int8

	number = 127
	fmt.Printf("%4d = %9b\n", number, number)
	number = 126
	fmt.Printf("%4d = %9b\n", number, number)
	number = 2
	fmt.Printf("%4d = %9b\n", number, number)
	number = 1
	fmt.Printf("%4d = %9b\n", number, number)
	number = 0
	fmt.Printf("%4d = %9b\n", number, number)
	number = -1
	fmt.Printf("%4d = %9b\n", number, number)
	number = -2
	fmt.Printf("%4d = %9b\n", number, number)
	number = -126
	fmt.Printf("%4d = %9b\n", number, number)
	number = -127
	fmt.Printf("%4d = %9b\n", number, number)
	number = -128
	fmt.Printf("%4d = %9b\n", number, number)
}
