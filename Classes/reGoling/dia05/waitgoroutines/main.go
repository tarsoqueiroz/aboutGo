package main

import (
	"fmt"
	"sync"
	"time"
)

func processarItem(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrementa o contador quando terminar

	fmt.Printf("Iniciando item %d\n", id)
	time.Sleep(time.Second) // Simula trabalho
	fmt.Printf("Finalizando item %d\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Incrementa o contador
		go processarItem(i, &wg)
	}

	wg.Wait() // Espera todas as goroutines terminarem
	fmt.Println("Todos os itens processados!")
}
