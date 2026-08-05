package main

import (
	"fmt"
	"time"
)

// Unbuffered - comunicação síncrona
func unbufferedExample() {
	ch := make(chan int) // Sem buffer

	fmt.Println("Chamando func inline pra enviar via channel")
	go func() {
		ch <- 42 // BLOQUEIA até alguém receber
		fmt.Println("Enviado!")
	}()

	fmt.Println("Dando um tempo pra ver o sync")
	time.Sleep(2 * time.Second) // Simula atraso

	valor := <-ch // BLOQUEIA até alguém enviar
	fmt.Println("Recebido:", valor)
}

func main() {
	fmt.Println("Call unbuffered channel function")
	unbufferedExample()
	fmt.Println("Voltei do call da unbuffered channel function")
}
