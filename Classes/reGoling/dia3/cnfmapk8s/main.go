package main

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type ConfigMap struct {
	Name      string            `yaml:"name"`
	Namespace string            `yaml:"namespace"`
	Data      map[string]string `yaml:"data"`
}

func main() {
	// Simulando um ConfigMap que você leu de um arquivo
	yamlString := `
name: nginx-config
namespace: production
data:
  nginx.conf: |
    server {
        listen 80;
        server_name localhost;
    }
  timeout: 60s
`

	var cm ConfigMap
	err := yaml.Unmarshal([]byte(yamlString), &cm)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n=== ConfigMap lido ===")
	fmt.Printf("ConfigMap lido: %+v\n", cm)

	// Modificando e salvando de volta
	cm.Data["new-key"] = "new-value"

	novoYAML, _ := yaml.Marshal(cm)
	fmt.Println("\n=== ConfigMap modificado ===")
	fmt.Println("\nConfigMap modificado:")
	fmt.Println(string(novoYAML))
}
