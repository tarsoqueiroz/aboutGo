package main

import "fmt"

// Pod - struct simples para demonstração
type Pod struct {
	Name      string
	Namespace string
	Ready     bool
}

// Interface do cliente Kubernetes
type K8sClient interface {
	GetPods(namespace string) ([]Pod, error)
	DeletePod(name string, namespace string) error
	GetPodLogs(name string, namespace string) (string, error)
}

// Implementação REAL (simulada)
type RealK8sClient struct {
	serverURL string
}

func (c RealK8sClient) GetPods(namespace string) ([]Pod, error) {
	// Simulação: em produção, chamaria a API do K8s
	fmt.Printf("🔴 Conectando a %s para buscar pods\n", c.serverURL)
	return []Pod{
		{Name: "nginx", Namespace: namespace, Ready: true},
		{Name: "redis", Namespace: namespace, Ready: false},
	}, nil
}

func (c RealK8sClient) DeletePod(name string, namespace string) error {
	fmt.Printf("🔴 Deletando pod %s/%s\n", namespace, name)
	return nil
}

func (c RealK8sClient) GetPodLogs(name string, namespace string) (string, error) {
	fmt.Printf("🔴 Buscando logs de %s/%s\n", namespace, name)
	return "Logs simulados...", nil
}

// Implementação MOCK (para testes)
type MockK8sClient struct {
	ShouldFail bool
}

func (m MockK8sClient) GetPods(namespace string) ([]Pod, error) {
	if m.ShouldFail {
		return nil, fmt.Errorf("erro simulado")
	}
	return []Pod{
		{Name: "mock-pod-1", Namespace: namespace, Ready: true},
		{Name: "mock-pod-2", Namespace: namespace, Ready: true},
	}, nil
}

func (m MockK8sClient) DeletePod(name string, namespace string) error {
	if m.ShouldFail {
		return fmt.Errorf("erro ao deletar mock")
	}
	fmt.Printf("✅ Mock: Pod %s/%s deletado\n", namespace, name)
	return nil
}

func (m MockK8sClient) GetPodLogs(name string, namespace string) (string, error) {
	if m.ShouldFail {
		return "", fmt.Errorf("erro ao buscar logs mock")
	}
	return "✅ Logs mockados...", nil
}

// Função que usa a interface (desacoplada da implementação)
func ProcessarPods(cliente K8sClient, namespace string) {
	pods, err := cliente.GetPods(namespace)
	if err != nil {
		fmt.Printf("❌ Erro ao buscar pods: %v\n", err)
		return
	}

	fmt.Printf("📦 Encontrados %d pods no namespace %s\n", len(pods), namespace)
	for _, pod := range pods {
		status := "✅"
		if !pod.Ready {
			status = "❌"
		}
		fmt.Printf("  %s Pod: %s (Ready: %v)\n", status, pod.Name, pod.Ready)
	}
}

func main() {
	// Usando o cliente REAL
	realClient := RealK8sClient{serverURL: "https://k8s-api.cluster.local"}
	fmt.Println("=== Teste com Cliente REAL ===")
	ProcessarPods(realClient, "default")

	// Usando o cliente MOCK (para testes)
	mockClient := MockK8sClient{ShouldFail: false}
	fmt.Println("\n=== Teste com Cliente MOCK ===")
	ProcessarPods(mockClient, "test")

	// Simulando falha no mock
	mockClientFail := MockK8sClient{ShouldFail: true}
	fmt.Println("\n=== Teste com MOCK que FALHA ===")
	ProcessarPods(mockClientFail, "test")
}
