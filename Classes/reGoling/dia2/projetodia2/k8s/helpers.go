package k8s

import "fmt"

// GetPodName é uma função EXPORTADA (maiúscula)
// Pode ser usada por outros pacotes
func GetPodName(namespace, name string) string {
	return fmt.Sprintf("%s/%s", namespace, name)
}

// parseYaml é uma função NÃO EXPORTADA (minúscula)
// Só pode ser usada dentro do pacote k8s
func parseYaml(yamlData string) string {
	// Simulação: num cenário real, usaria a biblioteca yaml
	return "Parsed: " + yamlData
}

// PodInfo é uma struct EXPORTADA
type PodInfo struct {
	Name      string
	Namespace string
	Ready     bool
}

// podCache é uma struct NÃO EXPORTADA (uso interno)
type podCache struct {
	pods []PodInfo
}
