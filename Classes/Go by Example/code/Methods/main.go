package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {

	r := rect{width: 10, height: 5}
	fmt.Println("r area   : ", r.area())
	fmt.Println("  perim  : ", r.perim())

	rp := &r
	fmt.Println("rp area  : ", rp.area())
	fmt.Println("   perim : ", rp.perim())
	fmt.Println("r area   : ", r.area())
	fmt.Println("  perim  : ", r.perim())

	rp.height = rp.height * 2
	rp.width = rp.width * 3
	fmt.Println("rp area  : ", rp.area())
	fmt.Println("   perim : ", rp.perim())

	fmt.Println("r area   : ", r.area())
	fmt.Println("  perim  : ", r.perim())
}
