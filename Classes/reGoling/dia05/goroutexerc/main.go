package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Pod - struct representando um pod
type Pod struct {
	Name          string
	Namespace     string
	CPURequest    int
	MemoryRequest int
}

// Stats - estatísticas de processamento
type Stats struct {
	Total     int
	Success   int
	Failed    int
	StartTime time.Time
	EndTime   time.Time
}

// PodProcessor - processador de pods
type PodProcessor struct {
	numWorkers int
	pods       []Pod
	stats      Stats
	mu         sync.Mutex
}

// NewPodProcessor - cria novo processador
func NewPodProcessor(numWorkers int, pods []Pod) *PodProcessor {
	return &PodProcessor{
		numWorkers: numWorkers,
		pods:       pods,
		stats: Stats{
			StartTime: time.Now(),
		},
	}
}

// processPod - processa um único pod com possibilidade de falha
func (pp *PodProcessor) processPod(pod Pod, workerID int) error {
	fmt.Printf("[Worker %d] Processando pod %s/%s (CPU: %d, Mem: %d)\n",
		workerID, pod.Namespace, pod.Name, pod.CPURequest, pod.MemoryRequest)

	// Simula tempo de processamento variável
	processTime := time.Duration(50+rand.Intn(100)) * time.Millisecond
	time.Sleep(processTime)

	// Simula falha aleatória (10% de chance)
	if rand.Float32() < 0.1 {
		return fmt.Errorf("falha ao processar pod %s/%s", pod.Namespace, pod.Name)
	}

	return nil
}

// processPodWithRetry - processa com retry
func (pp *PodProcessor) processPodWithRetry(pod Pod, workerID int) error {
	maxRetries := 3
	for attempt := 0; attempt < maxRetries; attempt++ {
		err := pp.processPod(pod, workerID)
		if err == nil {
			return nil
		}

		if attempt < maxRetries-1 {
			fmt.Printf("[Worker %d] ⚠️  Falha no pod %s (tentativa %d), retentando...\n",
				workerID, pod.Name, attempt+1)
			time.Sleep(100 * time.Millisecond)
		} else {
			return fmt.Errorf("falha após %d tentativas: %v", maxRetries, err)
		}
	}
	return nil
}

// Run - executa o processamento
func (pp *PodProcessor) Run() {
	pp.stats.Total = len(pp.pods)

	var wg sync.WaitGroup
	jobs := make(chan Pod, len(pp.pods))

	// Iniciar workers
	for i := 0; i < pp.numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for pod := range jobs {
				// Processa com retry
				err := pp.processPodWithRetry(pod, workerID)

				pp.mu.Lock()
				if err != nil {
					pp.stats.Failed++
					fmt.Printf("[Worker %d] ❌ Falha no pod %s: %v\n",
						workerID, pod.Name, err)
				} else {
					pp.stats.Success++
				}
				pp.mu.Unlock()
			}
		}(i)
	}

	// Enviar jobs
	for _, pod := range pp.pods {
		jobs <- pod
	}
	close(jobs)

	wg.Wait()
	pp.stats.EndTime = time.Now()
}

// PrintStats - imprime estatísticas
func (pp *PodProcessor) PrintStats() {
	duration := pp.stats.EndTime.Sub(pp.stats.StartTime)
	avgTime := duration / time.Duration(pp.stats.Total)

	fmt.Println("\n=== ESTATÍSTICAS DE PROCESSAMENTO ===")
	fmt.Printf("📊 Total de pods: %d\n", pp.stats.Total)
	fmt.Printf("✅ Sucessos: %d (%.1f%%)\n",
		pp.stats.Success, float64(pp.stats.Success)/float64(pp.stats.Total)*100)
	fmt.Printf("❌ Falhas: %d (%.1f%%)\n",
		pp.stats.Failed, float64(pp.stats.Failed)/float64(pp.stats.Total)*100)
	fmt.Printf("⏱️  Tempo total: %v\n", duration)
	fmt.Printf("📈 Tempo médio por pod: %v\n", avgTime)
	fmt.Printf("🚀 Workers utilizados: %d\n", pp.numWorkers)
}

func main() {
	// Seed do random
	rand.Seed(time.Now().UnixNano())

	// Criar 1000 pods
	pods := make([]Pod, 1000)
	for i := 0; i < 1000; i++ {
		pods[i] = Pod{
			Name:          fmt.Sprintf("pod-%d", i),
			Namespace:     "default",
			CPURequest:    100 + rand.Intn(900),
			MemoryRequest: 256 + rand.Intn(1024),
		}
	}

	fmt.Printf("📦 Criados %d pods para processar\n", len(pods))
	fmt.Println("=== INICIANDO PROCESSAMENTO ===")

	// Processar com 10 workers
	processor := NewPodProcessor(10, pods)
	processor.Run()
	processor.PrintStats()
}
