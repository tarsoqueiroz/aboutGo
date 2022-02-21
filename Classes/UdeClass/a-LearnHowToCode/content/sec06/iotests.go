package main

import (
	"fmt"
)

const (
	a = 1 << iota
	b = 1
	c = 1 << iota
	d = 1
	e = 1 << iota
)

func main() {
	fmt.Printf("%5d = %08b\n", a, a)
	fmt.Printf("%5d = %08b\n", b, b)
	fmt.Printf("%5d = %08b\n", c, c)
	fmt.Printf("%5d = %08b\n", d, d)
	fmt.Printf("%5d = %08b\n", e, e)
}
