package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Tipos de erros
type TimeoutError struct {
	Duration time.Duration
}

func (e TimeoutError) Error() string {
	return fmt.Sprintf("timeout após %v", e.Duration)
}

type AuthError struct {
	Reason string
}

func (e AuthError) Error() string {
	return fmt.Sprintf("erro de autenticação: %s", e.Reason)
}

type RateLimitError struct {
	RetryAfter time.Duration
}

func (e RateLimitError) Error() string {
	return fmt.Sprintf("rate limit: aguarde %v", e.RetryAfter)
}

// Circuit Breaker
type CircuitBreaker struct {
	state           string
	failures        int
	maxFailures     int
	timeout         time.Duration
	lastFailureTime time.Time
	mutex           chan struct{} // Simples mutex com channel
}

func NewCircuitBreaker(maxFailures int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:       "closed",
		maxFailures: maxFailures,
		timeout:     timeout,
		mutex:       make(chan struct{}, 1),
	}
}

func (cb *CircuitBreaker) lock()   { cb.mutex <- struct{}{} }
func (cb *CircuitBreaker) unlock() { <-cb.mutex }

func (cb *CircuitBreaker) Call(fn func() error) error {
	cb.lock()
	defer cb.unlock()

	// Verifica estado do circuit breaker
	if cb.state == "open" {
		if time.Since(cb.lastFailureTime) > cb.timeout {
			fmt.Println("🔄 Circuit half-open: testando novamente...")
			cb.state = "half-open"
		} else {
			return fmt.Errorf("circuit breaker aberto (falhas: %d)", cb.failures)
		}
	}

	// Executa a função
	err := fn()

	if err != nil {
		cb.failures++
		cb.lastFailureTime = time.Now()

		// Verifica se deve abrir o circuito
		if cb.failures >= cb.maxFailures {
			cb.state = "open"
			fmt.Printf("🔴 Circuit breaker ABERTO (falhas: %d)\n", cb.failures)
		}
		return fmt.Errorf("chamada falhou: %w", err)
	}

	// Sucesso: reset
	if cb.state == "half-open" {
		fmt.Println("🟢 Circuit half-open: sucesso! Circuito fechado")
	}
	cb.state = "closed"
	cb.failures = 0
	return nil
}

// API simulada
func chamadaAPI() error {
	// Simula latência
	time.Sleep(50 * time.Millisecond)

	// Diferentes tipos de erro
	randVal := rand.Float32()
	switch {
	case randVal < 0.2:
		return TimeoutError{Duration: 5 * time.Second}
	case randVal < 0.35:
		return AuthError{Reason: "token expirado"}
	case randVal < 0.5:
		return RateLimitError{RetryAfter: 2 * time.Second}
	case randVal < 0.7:
		return fmt.Errorf("erro interno do servidor")
	default:
		return nil // Sucesso
	}
}

// Função com retry e circuit breaker
func chamadaComRetry(cb *CircuitBreaker, maxRetries int) error {
	var lastErr error
	baseDelay := 1 * time.Second

	for attempt := 0; attempt <= maxRetries; attempt++ {
		err := cb.Call(chamadaAPI)
		if err == nil {
			fmt.Printf("✅ Chamada bem-sucedida na tentativa %d\n", attempt+1)
			return nil
		}

		lastErr = err

		// Verifica se é erro de rate limit
		var rateErr RateLimitError
		if errors.As(err, &rateErr) {
			fmt.Printf("⏳ Rate limit: aguardando %v\n", rateErr.RetryAfter)
			time.Sleep(rateErr.RetryAfter)
			continue
		}

		if attempt == maxRetries {
			break
		}

		// Backoff exponencial com jitter
		backoff := baseDelay * time.Duration(1<<uint(attempt))
		jitter := time.Duration(rand.Intn(500)) * time.Millisecond
		waitTime := backoff + jitter

		fmt.Printf("⚠️  Tentativa %d falhou: %v. Aguardando %v...\n",
			attempt+1, err, waitTime)
		time.Sleep(waitTime)
	}

	return fmt.Errorf("falha após %d tentativas: %w", maxRetries+1, lastErr)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("=== SISTEMA DE CONEXÃO COM RETRY E CIRCUIT BREAKER ===")

	cb := NewCircuitBreaker(3, 5*time.Second)

	// Cenário 1: Múltiplas chamadas
	for i := 1; i <= 10; i++ {
		fmt.Printf("\n--- Chamada %d ---\n", i)
		err := chamadaComRetry(cb, 3)
		if err != nil {
			fmt.Printf("❌ Erro final: %v\n", err)
		}

		// Espera entre chamadas
		time.Sleep(1 * time.Second)
	}
}
