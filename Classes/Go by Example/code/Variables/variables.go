package main

import "fmt"

func main() {

	var a1 = "initial 1"
	a2 := "initial 2"
	fmt.Println("a1 =", a1)
	fmt.Println("a2 =", a2)
	fmt.Println("")

	var b1, c1 int = 1, 2
	b2, c2 := 3, 4
	fmt.Println("b1 =", b1)
	fmt.Println("b2 =", b2)
	fmt.Println("c1 =", c1)
	fmt.Println("c2 =", c2)
	fmt.Println("")

	var d1 = true
	d2 := false
	fmt.Println("d1 =", d1)
	fmt.Println("d2 =", d2)
	fmt.Println("")

	var e1 int
	e2 := e1
	fmt.Println("e1 =", e1)
	fmt.Println("e2 =", e2)
	fmt.Println("")

}
