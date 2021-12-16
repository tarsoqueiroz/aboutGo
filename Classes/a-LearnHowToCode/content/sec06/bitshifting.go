package main

import (
	"fmt"
)

func main() {
	x := 4
	fmt.Printf("decimal: %d\t\tbinary: %4b\n", x, x)

	y := x << 1
	fmt.Printf("decimal: %d\t\tbinary: %4b\n", y, y)

	z := x >> 1
	fmt.Printf("decimal: %d\t\tbinary: %4b\n", z, z)

	w := z >> 1
	fmt.Printf("decimal: %d\t\tbinary: %4b\n", w, w)
}
