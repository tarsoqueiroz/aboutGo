package main

import "fmt"

func printValue(v interface{}) {
	fmt.Println("Received value:", v)
}

func main() {

	printValue(42)
	printValue("Hello, Go!")
	printValue(3.14)
	printValue(true)

}
