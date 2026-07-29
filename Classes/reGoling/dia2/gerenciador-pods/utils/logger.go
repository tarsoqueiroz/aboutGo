package utils

import "fmt"

// Info - exportada
func Info(msg string) {
	logWithLevel("INFO", msg)
}

// logWithLevel - NÃO exportada
func logWithLevel(level, msg string) {
	fmt.Printf("[%s] %s\n", level, msg)
}
