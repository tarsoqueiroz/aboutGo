package main

import "fmt"

func Equals[T comparable](a, b T) bool {

	return a == b
}

type Pessoa struct {
	Name string
}

func main() {
	fmt.Println("Equals(\"a\", \"a\")...........................:", Equals("a", "a"))
	fmt.Println("Equals(\"a\", \"b\")...........................:", Equals("a", "b"))
	fmt.Println("Equals(1, 1)...............................:", Equals(1, 1))
	fmt.Println("Equals(1, 2)...............................:", Equals(1, 2))
	fmt.Println("Equals(true, true).........................:", Equals(true, true))
	fmt.Println("Equals(true, false)........................:", Equals(true, false))
	fmt.Println("Equals(Pessoa{\"Tarso\"}, Pessoa{\"Tarso\"})...:", Equals(Pessoa{"Tarso"}, Pessoa{"Tarso"}))
	fmt.Println("Equals(Pessoa{\"Tarso\"}, Pessoa{\"Queiroz\"}).:", Equals(Pessoa{"Tarso"}, Pessoa{"Queiroz"}))
}
