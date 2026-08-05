package main

import (
	"fmt"
	"sync"
	"time"
)

// Estágio 1: Produtor - gera números
func produtor(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			fmt.Printf("📦 Produtor enviou: %d\n", n)
			out <- n
			time.Sleep(100 * time.Millisecond)
		}
		close(out)
		fmt.Println("📦 Produtor finalizado")
	}()
	return out
}

// Estágio 2: Processador - dobra os números
func processador(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			resultado := n * 2
			fmt.Printf("⚙️  Processador: %d → %d\n", n, resultado)
			out <- resultado
			time.Sleep(50 * time.Millisecond)
		}
		close(out)
		fmt.Println("⚙️  Processador finalizado")
	}()
	return out
}

// Estágio 3: Consumidor - imprime os resultados
func consumidor(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for resultado := range in {
		fmt.Printf("✅ Consumidor recebeu: %d\n", resultado)
	}
	fmt.Println("✅ Consumidor finalizado")
}

// Pipeline com múltiplos workers em paralelo
func pipelineParalelo() {
	fmt.Println("\n=== PIPELINE COM WORKERS PARALELOS ===")

	// Entrada
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Estágio 1: Produtor
	prod := produtor(numeros...)

	// Estágio 2: Processadores paralelos (fan-out)
	numWorkers := 3
	processadores := make([]<-chan int, numWorkers)
	for i := 0; i < numWorkers; i++ {
		processadores[i] = processador(prod)
	}

	// Estágio 3: Consumidor (fan-in)
	consumidorChan := make(chan int)
	var wg sync.WaitGroup

	// Fan-in: junta os resultados dos processadores
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, ch := range processadores {
			for v := range ch {
				consumidorChan <- v
			}
		}
		close(consumidorChan)
	}()

	// Consumidor final
	wg.Add(1)
	go consumidor(consumidorChan, &wg)

	wg.Wait()
}

// Pipeline com contexto (cancellation)
func pipelineComContexto() {
	// Veremos amanhã com Context
	fmt.Println("\n=== PIPELINE COM CONTEXTO (amanhã) ===")
}

func main() {
	fmt.Println("=== PIPELINE SIMPLES ===")
	// Pipeline simples
	prod := produtor(1, 2, 3, 4, 5)
	proc := processador(prod)
	var wg sync.WaitGroup
	wg.Add(1)
	consumidor(proc, &wg)
	wg.Wait()

	// Pipeline com workers paralelos
	pipelineParalelo()
}
