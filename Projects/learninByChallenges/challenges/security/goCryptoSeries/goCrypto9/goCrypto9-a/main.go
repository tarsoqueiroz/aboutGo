package main

import (
	"crypto/subtle"
	"fmt"
)

func main() {
	//
	// Constant-Time Comparison: The Secret Handshake
	//
	fmt.Printf("===== Constant-Time Comparison: The Secret Handshake\n=====\n")

	secretHandshake := []byte("up-down-left-right-a-b-start")
	attemptedHandshake := []byte("down-up-right-left-b-a-stop")
	correctAttempt := []byte("up-down-left-right-a-b-start")

	// let's check the wrong attempt
	if subtle.ConstantTimeCompare(secretHandshake, attemptedHandshake) == 1 {
		fmt.Printf("===== Constant-Time Comparison: you're in the club!\n")
	} else {
		fmt.Printf("===== Constant-Time Comparison: sorry, that's not the secret handshake!\n")
	}
	// now the correct attempt
	if subtle.ConstantTimeCompare(secretHandshake, correctAttempt) == 1 {
		fmt.Printf("===== Constant-Time Comparison: welcome to the club!\n")
	} else {
		fmt.Printf("===== Constant-Time Comparison: nope, still not right!\n")
	}
}
