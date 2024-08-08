package main

import (
	"fmt"

	"github.com/tyler-smith/go-bip39"
)

func main() {
	// Gera um mnemônico aleatório com 12 palavras
	entropy, err := bip39.NewEntropy(128) // 128 bits para 12 palavras
	if err != nil {
		panic(err)
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		panic(err)
	}

	fmt.Println("Mnemônico gerado (12 palavras):", mnemonic)

	// Gera um mnemônico aleatório com 24 palavras
	entropy, err = bip39.NewEntropy(256) // 128 bits para 12 palavras
	if err != nil {
		panic(err)
	}

	mnemonic, err = bip39.NewMnemonic(entropy)
	if err != nil {
		panic(err)
	}

	fmt.Println("Mnemônico gerado (24 palavras):", mnemonic)
}
