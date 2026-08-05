package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Duas goroutines enviando em canais diferentes
	go func() {
		fmt.Println("go func 1")
		time.Sleep(2 * time.Second)
		fmt.Println("go func 1: enviando msg")
		ch1 <- "mensagem do canal 1"
		fmt.Println("go func 1: msg enviada")
	}()

	go func() {
		fmt.Println("go func 2")
		time.Sleep(1 * time.Second)
		fmt.Println("go func 2: enviando msg")
		ch2 <- "mensagem do canal 2"
		fmt.Println("go func 2: msg enviada")
	}()

	// Select espera o primeiro canal que receber dados
	for i := 0; i < 2; i++ {
		fmt.Println("Quem mandou msg?")
		select {
		case msg1 := <-ch1:
			fmt.Println("Recebido do canal 1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Recebido do canal 2:", msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout!")
			return
		}
	}
	fmt.Println("Saindo")
}
