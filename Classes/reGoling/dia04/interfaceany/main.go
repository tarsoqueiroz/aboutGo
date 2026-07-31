package main

// Interface vazia - aceita QUALQUER tipo
var qualquer interface{}

qualquer = 42
qualquer = "string"
qualquer = struct{Nome string}{"João"}

// Em Go 1.18+, pode usar 'any' (alias para interface{})
var qualquer2 any

qualquer2 = true
qualquer2 = "string"
qualquer2 = 1

// Útil para funções genéricas (mas prefira tipos específicos quando possível)
func Imprimir(v any) {
    fmt.Printf("Valor: %v, Tipo: %T\n", v, v)
}

func main() {
    Imprimir(42)           // Valor: 42, Tipo: int
    Imprimir("teste")      // Valor: teste, Tipo: string
    Imprimir([]int{1,2,3}) // Valor: [1 2 3], Tipo: []int
}
