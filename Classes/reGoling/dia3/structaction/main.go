package main

import "fmt"

type Pod struct {
	Name      string
	Namespace string
	Replicas  int
	Ready     bool
}

// Método com RECEPTOR POR VALOR (cópia) - NÃO MODIFICA o original
func (p Pod) GetFullName() string {
	return p.Namespace + "/" + p.Name
}

// Método com RECEPTOR POR PONTEIRO (*) - PODE MODIFICAR o original
func (p *Pod) SetReplicas(n int) {
	p.Replicas = n // Modifica o campo
}

// Método com RECEPTOR POR PONTEIRO - útil para validação
func (p *Pod) IsValid() bool {
	return p.Name != "" && p.Namespace != ""
}

// Método que retorna uma cópia modificada
func (p Pod) WithReadyStatus(ready bool) Pod {
	p.Ready = ready
	return p // Retorna uma nova struct
}

func main() {
	pod := Pod{Name: "nginx", Namespace: "default", Replicas: 1, Ready: true}

	fmt.Printf("Pod: %+v\n", pod) // %+v mostra nomes dos campos

	// Método com receptor por valor (cópia)
	fullName := pod.GetFullName()
	fmt.Println("Nome completo:", fullName)

	// Método com receptor por ponteiro (modifica original)
	pod.SetReplicas(3)
	fmt.Printf("Pod após SetReplicas: %+v\n", pod)

	// Método que retorna nova cópia
	newPod := pod.WithReadyStatus(false)
	fmt.Printf("Novo Pod: %+v\n", newPod)
	fmt.Printf("Pod original ainda é: %+v\n", pod) // Não mudou
}
