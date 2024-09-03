package main

import (
	"fmt"
	"log"
)

func processRequest() error {
	// Simulate an error
	return fmt.Errorf("simulated error")
}

func main() {
	if err := processRequest(); err != nil {
		log.Printf("Error processing request: %v", err)
		fmt.Println("Internal Server Error")
		return
	}
	fmt.Println("Request processed successfully!")
}
