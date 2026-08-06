package main

import (
	"fmt"
	"time"
)

// Erro customizado com mais informações
type TimeoutError struct {
	Operation string
	Timeout   time.Duration
	Cause     error
}

func (e TimeoutError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("timeout na operação '%s' após %v: %v",
			e.Operation, e.Timeout, e.Cause)
	}
	return fmt.Sprintf("timeout na operação '%s' após %v",
		e.Operation, e.Timeout)
}

// Função que pode retornar erro customizado
func operacaoComTimeout(operation string, timeout time.Duration) error {
	// Simula uma operação que pode timeout
	if timeout < time.Second {
		return TimeoutError{
			Operation: operation,
			Timeout:   timeout,
			Cause:     fmt.Errorf("timeout muito curto"),
		}
	}
	return nil
}

func main() {
	err := operacaoComTimeout("conectar", 100*time.Millisecond)
	if err != nil {
		// Type assertion para acessar campos específicos
		if timeoutErr, ok := err.(TimeoutError); ok {
			fmt.Printf("⚠️  Erro customizado: %s\n", timeoutErr.Error())
			fmt.Printf("   Operação: %s\n", timeoutErr.Operation)
			fmt.Printf("   Timeout: %v\n", timeoutErr.Timeout)
		}
	}
}
