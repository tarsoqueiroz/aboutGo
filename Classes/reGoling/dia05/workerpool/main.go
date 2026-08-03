package main

import (
	"fmt"
	"sync"
	"time"
)

type Pod struct {
	Name      string
	Namespace string
}

// WorkerPool - estrutura para gerenciar workers
type WorkerPool struct {
	numWorkers int
	jobs       chan Pod
	wg         sync.WaitGroup
}

// NewWorkerPool - cria um pool com N workers
func NewWorkerPool(numWorkers int) *WorkerPool {
	return &WorkerPool{
		numWorkers: numWorkers,
		jobs:       make(chan Pod, 100), // Buffer para jobs
	}
}

// Start - inicia os workers
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.numWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

// worker - processa jobs do canal
func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	for pod := range wp.jobs {
		fmt.Printf("[Worker %d] Processando %s/%s\n",
			id, pod.Namespace, pod.Name)
		time.Sleep(50 * time.Millisecond)
	}
}

// Submit - adiciona um job à fila
func (wp *WorkerPool) Submit(pod Pod) {
	wp.jobs <- pod
}

// Wait - espera todos os jobs terminarem
func (wp *WorkerPool) Wait() {
	close(wp.jobs) // Fecha o canal para workers pararem
	wp.wg.Wait()
}

func main() {
	// Criar 1000 pods
	pods := make([]Pod, 1000)
	for i := 0; i < 1000; i++ {
		pods[i] = Pod{
			Name:      fmt.Sprintf("pod-%d", i),
			Namespace: "default",
		}
	}

	// Criar pool com 10 workers
	pool := NewWorkerPool(10)

	start := time.Now()

	// Iniciar workers
	pool.Start()

	// Enviar jobs
	for _, pod := range pods {
		pool.Submit(pod)
	}

	// Esperar finalizar
	pool.Wait()

	elapsed := time.Since(start)
	fmt.Printf("\n✅ Processados %d pods em %v\n", len(pods), elapsed)
	fmt.Printf("📊 Média: %v por pod\n", elapsed/time.Duration(len(pods)))
}
