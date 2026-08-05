package main

import "fmt"

func main() {
	// Criando um channel
	ch := make(chan string) // Unbuffered (sem buffer)

	// Enviar e receber em goroutines separadas
	go func() {
		ch <- "mensagem" // Envia
	}()

	msg := <-ch // Recebe
	fmt.Println(msg)
}
