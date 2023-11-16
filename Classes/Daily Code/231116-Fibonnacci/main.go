package main

import "fmt"

func main() {

	fib := []int{0, 1}
	fmt.Println("v1 fib initial:", fib)
	for i, f1 := range fib {
		f2 := fib[i+1]
		fib = append(fib, f1+f2)
		// fmt.Println("f1:", f1, "f2:", f2, "f1+f2:", f1+f2, "fib:", fib)
		if f1+f2 > 100 {
			break
		}
	}
	fmt.Println("v1 fib final  :", fib)

	fib = []int{0, 1}
	fmt.Println("v2 fib initial:", fib)
	for i := 0; i < len(fib); i++ {
		f1, f2 := fib[i], fib[i+1]
		fib = append(fib, f1+f2)
		// fmt.Println("f1:", f1, "f2:", f2, "f1+f2:", f1+f2, "fib:", fib)
		if f1+f2 > 100 {
			break
		}
	}
	fmt.Println("v2 fib final  :", fib)

}
