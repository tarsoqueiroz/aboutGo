package main

import "fmt"

// main Code where a pipeline of concurrent operations is orchestrated.
func main() {

	// Step 1: Generate a channel of integers
	c := gen(2, 3, 4, 5)

	// Step 2: Square the numbers concurrently
	out := sq(c)

	// Step 3: Print the results
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
}

// gen Generating Numbers Concurrently
func gen(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

// sq Squaring Numbers Concurrently
func sq(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}
