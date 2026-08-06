package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validação falhou para %s (%v): %s",
		e.Field, e.Value, e.Message)
}

func validarCampo(field string, value interface{}) error {
	if value == nil || value == "" {
		return ValidationError{
			Field:   field,
			Value:   value,
			Message: "valor não pode ser vazio",
		}
	}
	return nil
}

func processarDados() error {
	err := validarCampo("nome", "")
	if err != nil {
		return fmt.Errorf("processando dados: %w", err)
	}
	return nil
}

func main() {
	err := processarDados()
	if err != nil {
		// errors.Is: verifica se o erro (ou sua causa) é de um tipo específico
		if errors.Is(err, ValidationError{}) {
			fmt.Println("❌ Erro de validação detectado")
		}

		// errors.As: extrai o erro para um tipo específico
		var valErr ValidationError
		if errors.As(err, &valErr) {
			fmt.Printf("⚠️  Campo: %s, Mensagem: %s\n",
				valErr.Field, valErr.Message)
		}
	}
}
