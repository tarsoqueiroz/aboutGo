package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("FullCycle API")

	mux := http.NewServeMux()
	http.ListenAndServe(":3000", mux)
}
