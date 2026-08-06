package main

import (
	"errors"
	"fmt"
)

// Função de baixo nível
func conectarServidor(addr string) error {
	// Simula erro de conexão
	return fmt.Errorf("servidor %s: conexão recusada", addr)
}

// Função de médio nível
func iniciarServico(servico string) error {
	err := conectarServidor("localhost:8080")
	if err != nil {
		// %w mantém a causa original
		return fmt.Errorf("iniciando serviço %s: %w", servico, err)
	}
	return nil
}

// Função de alto nível
func main() {
	err := iniciarServico("api-gateway")
	if err != nil {
		// Verifica se o erro contém uma causa específica
		if errors.Is(err, fmt.Errorf("conexão recusada")) {
			fmt.Println("Erro específico detectado!")
		}

		// Desempacota para ver a cadeia
		fmt.Printf("Erro completo: %v\n", err)

		// Desempacota para ver a causa raiz
		fmt.Printf("Causa raiz: %v\n", errors.Unwrap(err))
	}
}
