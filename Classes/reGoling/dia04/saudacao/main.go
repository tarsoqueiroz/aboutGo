package main

import "fmt"

// Interface
type Saudacao interface {
	Ola() string
}

// Tipo 1 - implementa automaticamente
type Portugues struct{}

func (p Portugues) Ola() string { return "Olá!" }

// Tipo 2 - também implementa
type Ingles struct{}

func (i Ingles) Ola() string { return "Hello!" }

// Função que aceita QUALQUER tipo que implemente Saudacao
func Cumprimentar(s Saudacao) {
	fmt.Println(s.Ola())
}

func main() {
	// Nenhuma declaração "implements" necessária!
	p := Portugues{}
	i := Ingles{}

	Cumprimentar(p) // "Olá!"
	Cumprimentar(i) // "Hello!"
}
