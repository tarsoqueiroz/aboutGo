package main

import "fmt"

// Definindo uma struct
type Pod struct {
	Name      string
	Namespace string
	Replicas  int
	Ready     bool
}

// Struct aninhada (comum em Kubernetes)
type Deployment struct {
	Name   string
	Pod    Pod // Composição (não herança)
	Labels map[string]string
}

func main() {
	// Forma 1: Declarativa (recomendada)
	pod1 := Pod{
		Name:      "nginx",
		Namespace: "default", // Vírgula obrigatória mesmo na última linha
		Replicas:  1,
		Ready:     true,
	}

	// Forma 2: Posicional (NÃO RECOMENDADO - frágil)
	pod2 := Pod{"redis", "cache", 3, false}

	// Forma 3: Vazia (campos com valores zero)
	pod3 := Pod{}

	// Forma 4: Parcial (campos não especificados ficam com valor zero)
	pod4 := Pod{Name: "api", Namespace: "prod"}

	fmt.Println(pod1)
	fmt.Printf("Pod 1: %+v\n", pod1) // %+v mostra nomes dos campos
	fmt.Printf("Pod 2: %+v\n", pod2) // %+v mostra nomes dos campos
	fmt.Printf("Pod 3: %+v\n", pod3) // %+v mostra nomes dos campos
	fmt.Printf("Pod 4: %+v\n", pod4) // %+v mostra nomes dos campos
}
