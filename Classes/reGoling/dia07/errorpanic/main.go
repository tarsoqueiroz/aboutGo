package main

import "fmt"

// ❌ NUNCA faça isso para fluxo normal
func exemploErrado() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado:", r)
		}
	}()

	panic("algo deu errado") // Isso é um anti-pattern!
}

// ✅ Use panic apenas para erros que não podem ser recuperados
func exemploCorreto() {
	// Exemplo: arquivo de configuração obrigatório não existe
	// Aqui panic é aceitável porque o programa não pode continuar
	config := lerConfiguracaoObrigatoria()
	// ... continua
}

func lerConfiguracaoObrigatoria() string {
	// Se não encontrar, panic é aceitável
	// ...
	return "config"
}

// ✅ Em serviços long-running, use retry com erros, não panic
func operacaoComRetry() error {
	for i := 0; i < 3; i++ {
		err := operacaoArriscada()
		if err == nil {
			return nil
		}
		// Log do erro e tenta novamente
	}
	return fmt.Errorf("falha após 3 tentativas")
}

func operacaoArriscada() error {
	return fmt.Errorf("erro simulado")
}
