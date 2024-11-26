package main

import (
	"crypto/subtle"
	"fmt"
)

func main() {
	//
	// Constant-Time Boolean Operations: Secret Logic
	//
	fmt.Printf("===== Constant-Time Boolean Operations: Secret Logic\n=====\n")

	secretBit := 1
	guessedBit := 0

	// let's do some secret logic
	andResult := subtle.ConstantTimeByteEq(uint8(secretBit&guessedBit), 0)
	orResult := subtle.ConstantTimeByteEq(uint8(secretBit|guessedBit), 0)

	fmt.Printf("===== Constant-Time Boolean Operations: AND result is zero: %v\n", andResult == 1)
	fmt.Printf("===== Constant-Time Boolean Operations: OR  result is zero: %v\n", orResult == 1)
}
