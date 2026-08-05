package main

import (
	"fmt"
	"time"
)

func produtor(ch chan<- int) {
	fmt.Println("Entrando no produtor")
	for i := 1; i <= 5; i++ {
		fmt.Println("Enviando:", i)
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch) // Fecha o canal quando terminar
	fmt.Println("Saindo do produtor")
}

func consumidor(ch <-chan int) {
	fmt.Println("Entrando no consumidor")
	for valor := range ch { // Loop até canal ser fechado
		fmt.Println("Recebido:", valor)
	}
	fmt.Println("Canal fechado!")
	fmt.Println("Saindo do consumidor")
}

func main() {
	fmt.Println("range Channel iniciado!")
	ch := make(chan int)

	fmt.Println("Chamando produtor")
	go produtor(ch)
	fmt.Println("Chamando consumidor")
	consumidor(ch)
}
