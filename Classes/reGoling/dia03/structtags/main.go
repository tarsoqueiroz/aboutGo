package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

type ConfigMap struct {
	Name      string            `json:"name" yaml:"name"`
	Namespace string            `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Data      map[string]string `json:"data" yaml:"data"`
	Internal  string            `json:"-" yaml:"-"` // Ignorado em ambos
}

func main() {
	cm := ConfigMap{
		Name:      "app-config",
		Namespace: "default",
		Data:      map[string]string{"key": "value", "timeout": "30s"},
		Internal:  "segredo",
	}

	// Serializando para JSON
	jsonData, _ := json.MarshalIndent(cm, "", "  ")
	fmt.Println("=== JSON ===")
	fmt.Println(string(jsonData))

	// Serializando para YAML
	yamlData, _ := yaml.Marshal(cm)
	fmt.Println("\n=== YAML ===")
	fmt.Println(string(yamlData))
}
