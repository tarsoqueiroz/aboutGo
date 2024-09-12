package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Starting server on :8910")
	if err := http.ListenAndServe(":8910", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
