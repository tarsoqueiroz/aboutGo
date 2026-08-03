package main

import (
	"fmt"
	"time"
)

func dizerOla() {
	fmt.Println("Olá de uma goroutine!")
}

func main() {
	// Inicia uma goroutine
	go dizerOla()

	// Goroutine com função anônima
	go func() {
		fmt.Println("Olá de outra goroutine!")
	}()

	// Dá tempo para as goroutines executarem
	fmt.Println("Disparei duas goroutines e vou parar um pouco...")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Programa principal terminou")
}
