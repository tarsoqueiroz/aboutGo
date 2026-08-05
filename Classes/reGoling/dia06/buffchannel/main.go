package main

import (
	"fmt"
)

// Buffered - comunicação assíncrona
func bufferedExample() {
	ch := make(chan int, 3) // Buffer de 3

	// Pode enviar 3 mensagens sem receptor
	fmt.Println("Enviando msg 1")
	ch <- 1
	fmt.Println("Enviando msg 2")
	ch <- 2
	fmt.Println("Enviando msg 3")
	ch <- 3
	fmt.Println("Enviadas")

	// Agora recebe
	fmt.Println("Recebendo msg 1")
	fmt.Println(<-ch) // 1
	fmt.Println("Recebendo msg 2")
	fmt.Println(<-ch) // 2
	fmt.Println("Recebendo msg 3")
	fmt.Println(<-ch) // 3
	fmt.Println("Recebidas")
}

func main() {
	fmt.Println("Call buffered channel function")
	bufferedExample()
	fmt.Println("Voltei do call da buffered channel function")
}
