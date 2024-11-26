package main

import (
	"crypto/subtle"
	"fmt"
)

func main() {
	//
	// Constant-Time Selection: The Invisible Choice
	//
	fmt.Printf("===== Constant-Time Selection: The Invisible Choice\n=====\n")

	secretDoor := uint32(5)
	fakeDoor := uint32(8)
	condition := 1 // This could be the result of a secret operation

	chosenDoor := subtle.ConstantTimeSelect(condition, int(secretDoor), int(fakeDoor))
	fmt.Printf("===== Constant-Time Selection: the chosen door is: %d\n", chosenDoor)
}
