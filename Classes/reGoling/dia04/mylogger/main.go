package main

import "fmt"

// ==================== INTERFACE ====================
type Logger interface {
	Log(message string)
	LogWithLevel(level string, message string)
}

// ==================== CONSOLE LOGGER ====================
type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	c.LogWithLevel("INFO", message)
}

func (c ConsoleLogger) LogWithLevel(level string, message string) {
	fmt.Printf("[%s] %s\n", level, message)
}

// ==================== FILE LOGGER ====================
type FileLogger struct {
	messages []string
}

func (f *FileLogger) Log(message string) {
	f.LogWithLevel("INFO", message)
}

func (f *FileLogger) LogWithLevel(level string, message string) {
	f.messages = append(f.messages, fmt.Sprintf("[%s] %s", level, message))
}

// Método extra (NÃO está na interface)
func (f *FileLogger) GetMessages() []string {
	return f.messages
}

// ==================== MULTI LOGGER (DESAFIO) ====================
type MultiLogger struct {
	loggers []Logger
}

func (m MultiLogger) Log(message string) {
	for _, logger := range m.loggers {
		logger.Log(message)
	}
}

func (m MultiLogger) LogWithLevel(level string, message string) {
	for _, logger := range m.loggers {
		logger.LogWithLevel(level, message)
	}
}

// ==================== FUNÇÃO QUE USA A INTERFACE ====================
func ProcessarOperacao(logger Logger, operacao string) {
	logger.LogWithLevel("START", fmt.Sprintf("Iniciando operação: %s", operacao))
	logger.Log(fmt.Sprintf("Processando etapa 1 de %s", operacao))
	logger.Log(fmt.Sprintf("Processando etapa 2 de %s", operacao))
	logger.LogWithLevel("END", fmt.Sprintf("Finalizando operação: %s", operacao))
}

// ==================== MAIN ====================
func main() {
	fmt.Println("=== CONSOLE LOGGER ===")
	console := ConsoleLogger{}
	ProcessarOperacao(console, "deploy-nginx")

	fmt.Println("\n=== FILE LOGGER ===")
	file := &FileLogger{}
	ProcessarOperacao(file, "scale-redis")

	// Acessando método extra (não disponível via interface)
	fmt.Println("\nMensagens no arquivo:")
	for i, msg := range file.GetMessages() {
		fmt.Printf("  %d: %s\n", i+1, msg)
	}

	fmt.Println("\n=== MULTI LOGGER ===")
	multi := MultiLogger{
		loggers: []Logger{
			ConsoleLogger{},
			&FileLogger{},
		},
	}
	ProcessarOperacao(multi, "rollback-api")
}
