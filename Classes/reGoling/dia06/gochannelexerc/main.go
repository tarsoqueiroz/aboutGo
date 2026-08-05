package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Log - struct representando uma entrada de log
type Log struct {
	Timestamp time.Time
	Level     string
	Message   string
}

// Stage 1: Coletor de Logs (Produtor)
func coletorLogs(numLogs int) <-chan Log {
	out := make(chan Log, numLogs)
	levels := []string{"INFO", "WARN", "ERROR", "INFO", "INFO"}
	messages := []string{
		"Serviço iniciado",
		"Conexão estabelecida",
		"Timeout na requisição",
		"Cache atualizado",
		"Usuário autenticado",
		"Erro no banco de dados",
		"Processamento concluído",
		"Health check OK",
	}

	go func() {
		defer close(out)
		for i := 0; i < numLogs; i++ {
			log := Log{
				Timestamp: time.Now(),
				Level:     levels[i%len(levels)],
				Message:   fmt.Sprintf("%s [%d]", messages[i%len(messages)], i),
			}
			fmt.Printf("📝 Coletor: [%s] %s\n", log.Level, log.Message)
			out <- log
			time.Sleep(50 * time.Millisecond) // Simula coleta
		}
		fmt.Println("📝 Coletor finalizado")
	}()
	return out
}

// Stage 2: Filtro e Processador (com múltiplos workers)
func filtroProcessador(in <-chan Log, nivelFiltro string, numWorkers int) <-chan string {
	out := make(chan string, 100)
	var wg sync.WaitGroup

	// Worker function
	worker := func(id int) {
		defer wg.Done()
		for log := range in {
			// Filtra pelo nível
			if log.Level != nivelFiltro {
				continue
			}

			// Processa: converte mensagem para maiúsculas
			processed := fmt.Sprintf("[Worker %d] %s | %s",
				id,
				log.Timestamp.Format("15:04:05"),
				strings.ToUpper(log.Message))

			fmt.Printf("⚙️  Worker %d processou: %s\n", id, log.Message)
			out <- processed
			time.Sleep(20 * time.Millisecond) // Simula processamento
		}
	}

	// Inicia workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i)
	}

	// Fecha o canal de saída quando todos os workers terminarem
	go func() {
		wg.Wait()
		close(out)
		fmt.Println("⚙️  Processador finalizado")
	}()

	return out
}

// Stage 3: Consumidor
func consumidor(in <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	var totalRecebidos, totalProcessados int
	var mutex sync.Mutex

	var innerWg sync.WaitGroup
	innerWg.Add(1)

	go func() {
		defer innerWg.Done()
		for processed := range in {
			mutex.Lock()
			totalRecebidos++
			totalProcessados++
			mutex.Unlock()

			fmt.Printf("✅ Consumidor: %s\n", processed)
		}
	}()

	// Timeout para mostrar estatísticas mesmo se pipeline demorar
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("⏰ Timeout no consumidor")
	case <-func() <-chan struct{} {
		ch := make(chan struct{})
		go func() {
			innerWg.Wait()
			close(ch)
		}()
		return ch
	}():
		// Pipeline completou
	}

	fmt.Printf("📊 Estatísticas: %d logs processados\n", totalProcessados)
}

func main() {
	fmt.Println("=== PIPELINE DE PROCESSAMENTO DE LOGS ===")
	fmt.Println()

	// Stage 1: Coletor (100 logs)
	logs := coletorLogs(20)

	// Stage 2: Processador (filtra INFO, 3 workers)
	processed := filtroProcessador(logs, "INFO", 3)

	// Stage 3: Consumidor
	var wg sync.WaitGroup
	wg.Add(1)
	consumidor(processed, &wg)

	wg.Wait()
	fmt.Println("\n✅ Pipeline completo!")
}
