package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

// Pod representa um pod Kubernetes
type Pod struct {
	Name  string `json:"name" yaml:"name"`
	Image string `json:"image" yaml:"image"`
	Ports []int  `json:"ports,omitempty" yaml:"ports,omitempty"`
}

// Deployment representa um deployment Kubernetes
type Deployment struct {
	Name        string            `json:"name" yaml:"name"`
	Replicas    int               `json:"replicas" yaml:"replicas"`
	PodTemplate Pod               `json:"podTemplate" yaml:"podTemplate"`
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
}

// DeploymentStatus representa o status de um deployment
type DeploymentStatus struct {
	Ready             bool `json:"ready" yaml:"ready"`
	AvailableReplicas int  `json:"availableReplicas" yaml:"availableReplicas"`
}

// Scale - modifica o número de réplicas (PONTEIRO)
func (d *Deployment) Scale(n int) {
	if n > 0 {
		d.Replicas = n
	}
}

// GetFullName - retorna nome completo (VALOR - apenas leitura)
func (d Deployment) GetFullName() string {
	return "deployment/" + d.Name
}

// Validate - verifica se o deployment é válido (PONTEIRO)
func (d *Deployment) Validate() bool {
	return d.Name != "" && d.Replicas > 0
}

// Status - retorna o status atual (VALOR - cria cópia)
func (d Deployment) Status() DeploymentStatus {
	return DeploymentStatus{
		Ready:             d.Replicas > 0,
		AvailableReplicas: d.Replicas,
	}
}

// String - implementa a interface fmt.Stringer (método especial)
func (d Deployment) String() string {
	return fmt.Sprintf("Deployment: %s (replicas: %d)", d.Name, d.Replicas)
}

func main() {
	// Criando um deployment
	deploy := Deployment{
		Name:     "nginx-deployment",
		Replicas: 3,
		PodTemplate: Pod{
			Name:  "nginx-pod",
			Image: "nginx:1.21",
			Ports: []int{80},
		},
		Annotations: map[string]string{
			"description": "Test deployment for Day 3",
			"owner":       "arquiteto",
			"environment": "dev",
		},
	}

	fmt.Println("=== INFORMAÇÕES DO DEPLOYMENT ===")
	fmt.Println("Criado:", deploy.String())
	fmt.Println("Nome completo:", deploy.GetFullName())
	fmt.Println("Válido?", deploy.Validate())

	// Status inicial
	status := deploy.Status()
	fmt.Printf("Status inicial: %+v\n", status)

	// Escalando
	deploy.Scale(5)
	fmt.Println("Após scale:", deploy.String())

	// Novo status
	newStatus := deploy.Status()
	fmt.Printf("Novo status: %+v\n", newStatus)

	// Teste com deployment inválido
	invalid := Deployment{Name: ""}
	fmt.Println("Deployment inválido é válido?", invalid.Validate())

	// ============================================
	// SERIALIZAÇÃO EM JSON E YAML (conforme solicitado)
	// ============================================
	fmt.Println("\n=== SERIALIZAÇÃO ===")

	// Serializando para JSON
	jsonData, err := json.MarshalIndent(deploy, "", "  ")
	if err != nil {
		fmt.Println("Erro ao serializar para JSON:", err)
	} else {
		fmt.Println("JSON:")
		fmt.Println(string(jsonData))
	}

	// Serializando para YAML
	yamlData, err := yaml.Marshal(deploy)
	if err != nil {
		fmt.Println("Erro ao serializar para YAML:", err)
	} else {
		fmt.Println("\nYAML:")
		fmt.Println(string(yamlData))
	}

	// ============================================
	// DEMONSTRAÇÃO DE DESERIALIZAÇÃO (bônus)
	// ============================================
	fmt.Println("\n=== DESERIALIZAÇÃO A PARTIR DO YAML ===")
	yamlString := `
name: redis-deployment
replicas: 2
podTemplate:
  name: redis-pod
  image: redis:alpine
  ports:
  - 6379
annotations:
  description: Redis cache
  team: database
`
	var newDeploy Deployment
	err = yaml.Unmarshal([]byte(yamlString), &newDeploy)
	if err != nil {
		fmt.Println("Erro ao desserializar YAML:", err)
	} else {
		fmt.Println("Deployment criado a partir do YAML:")
		fmt.Printf("  Nome: %s\n", newDeploy.Name)
		fmt.Printf("  Réplicas: %d\n", newDeploy.Replicas)
		fmt.Printf("  Pod: %s (%s)\n", newDeploy.PodTemplate.Name, newDeploy.PodTemplate.Image)
		fmt.Printf("  Portas: %v\n", newDeploy.PodTemplate.Ports)
		fmt.Printf("  Anotações: %v\n", newDeploy.Annotations)
	}
}
