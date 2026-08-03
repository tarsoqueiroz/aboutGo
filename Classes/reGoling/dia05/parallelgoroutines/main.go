package main

import (
	"fmt"
	"sync"
	"time"
)

// Pod - struct simples
type Pod struct {
	Name      string
	Namespace string
}

// processarPod - simula processamento de um pod
func processarPod(p Pod, workerID int) {
	fmt.Printf("[Worker %d] Processando pod %s/%s\n",
		workerID, p.Namespace, p.Name)
	time.Sleep(100 * time.Millisecond) // Simula trabalho
}

func main() {
	// Criando 1000 pods para processar
	pods := make([]Pod, 1000)
	for i := 0; i < 1000; i++ {
		pods[i] = Pod{
			Name:      fmt.Sprintf("pod-%d", i),
			Namespace: "default",
		}
	}

	var wg sync.WaitGroup
	maxWorkers := 10
	semaphore := make(chan struct{}, maxWorkers) // Controle de concorrência

	start := time.Now()

	for i, pod := range pods {
		wg.Add(1)
		semaphore <- struct{}{} // Ocupa um slot do pool

		go func(p Pod, workerID int) {
			defer wg.Done()
			defer func() { <-semaphore }() // Libera o slot

			processarPod(p, workerID)
		}(pod, i%maxWorkers)
	}

	wg.Wait()
	elapsed := time.Since(start)

	fmt.Printf("\nProcessados %d pods em %v\n", len(pods), elapsed)
	fmt.Printf("Média: %v por pod\n", elapsed/time.Duration(len(pods)))
}
