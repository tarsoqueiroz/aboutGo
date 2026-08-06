package main

import (
	"errors"
	"fmt"
)

// error é uma interface
type error interface {
	Error() string
}

// Criando erros simples
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero não permitida")
	}
	return a / b, nil
}

func main() {
	// Padrão: sempre verifique o erro
	resultado, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Resultado:", resultado)

	// Erro esperado
	resultado, err = dividir(10, 0)
	if err != nil {
		fmt.Println("Erro esperado:", err)
		// Não retorna, continua execução
	}
}
