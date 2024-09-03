package main

import (
	"fmt"
	"log"
)

func generateError(isError bool) (int, error) {
	if isError {
		return 0, fmt.Errorf("simulated error")
	} else {
		return 1, nil
	}
}

func main() {
	fmt.Println("\nCall 1 with NO ERROR")
	retCode, err := generateError(false)
	if err != nil {
		log.Printf("Error processing request: %v", err)
		fmt.Println("Internal Server Error")
		return
	} else {
		fmt.Println("Request 1 processed successfully with retCode:", retCode)
	}

	fmt.Println("\nCall 2 with ERROR")
	retCode, err = generateError(true)
	if err != nil {
		log.Printf("Error processing request: %v", err)
		fmt.Println("Internal Server Error")
	} else {
		fmt.Println("Request 2 processed successfully with retCode:", retCode)
	}

	fmt.Println("\nRequests processed!")
}
