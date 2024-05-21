package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("FullCycle API")

	mux := http.NewServeMux()

	mux.HandleFunc("/users", listUsersHandler)

	http.ListenAndServe(":3000", mux)
}

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of users"))
}
