package main

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Timeout int `yaml:"timeout"`
	Retries int `yaml:"retries"`
}

func main() {
	yamlData := `
timeout: 30
retries: 5
`

	var config Config
	err := yaml.Unmarshal([]byte(yamlData), &config)
	if err != nil {
		fmt.Println("Erro ao parsear YAML:", err)
		return
	}
	fmt.Printf("Configuração carregada: Timeout=%d, Retries=%d\n", config.Timeout, config.Retries)
}
