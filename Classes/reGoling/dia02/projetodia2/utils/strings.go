package utils

import "strings"

// UpperExportada - função exportada
func UpperExportada(texto string) string {
	return strings.ToUpper(texto)
}

// lowerInterna - função não exportada
func lowerInterna(texto string) string {
	return strings.ToLower(texto)
}
