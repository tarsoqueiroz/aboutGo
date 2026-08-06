package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Erro de conexão customizado
type ConnectionError struct {
	Address string
	Attempt int
	Cause   error
}

func (e ConnectionError) Error() string {
	return fmt.Sprintf("conexão com %s falhou na tentativa %d: %v",
		e.Address, e.Attempt, e.Cause)
}

// Simula conexão com servidor (sucesso ou falha aleatória)
func conectarServidor(addr string) error {
	// Simula latência de rede
	time.Sleep(50 * time.Millisecond)

	// 70% de chance de falha
	if rand.Float32() < 0.7 {
		return fmt.Errorf("servidor %s não respondeu", addr)
	}
	return nil
}

// Função com retry e backoff exponencial
func conectarComRetry(addr string, maxRetries int) error {
	var lastErr error

	for attempt := 0; attempt <= maxRetries; attempt++ {
		// Tentativa de conexão
		err := conectarServidor(addr)
		if err == nil {
			fmt.Printf("✅ Conexão com %s estabelecida na tentativa %d\n",
				addr, attempt+1)
			return nil
		}

		// Erro na tentativa
		lastErr = ConnectionError{
			Address: addr,
			Attempt: attempt + 1,
			Cause:   err,
		}

		// Se for a última tentativa, retorna o erro
		if attempt == maxRetries {
			break
		}

		// Calcula backoff exponencial: 1s, 2s, 4s, 8s...
		backoff := time.Duration(1<<uint(attempt)) * time.Second
		// Adiciona jitter (variação aleatória) para evitar "thundering herd"
		jitter := time.Duration(rand.Intn(500)) * time.Millisecond
		waitTime := backoff + jitter

		fmt.Printf("⚠️  Tentativa %d falhou. Aguardando %v antes de tentar novamente...\n",
			attempt+1, waitTime)
		time.Sleep(waitTime)
	}

	return fmt.Errorf("falha ao conectar com %s após %d tentativas: %w",
		addr, maxRetries+1, lastErr)
}

// Versão com jitter e backoff mais sofisticado
func conectarComRetryAvancado(addr string, maxRetries int) error {
	var lastErr error
	baseDelay := 1 * time.Second
	maxDelay := 30 * time.Second

	for attempt := 0; attempt <= maxRetries; attempt++ {
		err := conectarServidor(addr)
		if err == nil {
			fmt.Printf("✅ Conexão com %s estabelecida na tentativa %d\n",
				addr, attempt+1)
			return nil
		}

		lastErr = ConnectionError{
			Address: addr,
			Attempt: attempt + 1,
			Cause:   err,
		}

		if attempt == maxRetries {
			break
		}

		// Backoff exponencial com jitter
		// Formula: min(maxDelay, baseDelay * 2^attempt)
		backoff := baseDelay * time.Duration(1<<uint(attempt))
		if backoff > maxDelay {
			backoff = maxDelay
		}

		// Jitter: adiciona +/- 20% de variação
		jitterRange := float64(backoff) * 0.2
		jitter := time.Duration(rand.Float64()*jitterRange*2 - jitterRange)
		waitTime := backoff + jitter

		fmt.Printf("⚠️  Tentativa %d falhou. Aguardando %v antes de tentar novamente...\n",
			attempt+1, waitTime)
		time.Sleep(waitTime)
	}

	return fmt.Errorf("falha ao conectar com %s após %d tentativas: %w",
		addr, maxRetries+1, lastErr)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("=== RETRY COM BACKOFF EXPONENCIAL ===")
	err := conectarComRetry("api.k8s.local", 4)
	if err != nil {
		fmt.Printf("❌ Erro final: %v\n", err)

		// Verifica se é ConnectionError
		var connErr ConnectionError
		if errors.As(err, &connErr) {
			fmt.Printf("   Última tentativa: %d, Endereço: %s\n",
				connErr.Attempt, connErr.Address)
		}
	}

	fmt.Println("\n=== RETRY AVANÇADO COM JITTER ===")
	err = conectarComRetryAvancado("database.k8s.local", 5)
	if err != nil {
		fmt.Printf("❌ Erro final: %v\n", err)
	}
}
