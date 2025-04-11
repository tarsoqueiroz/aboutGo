package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	ready bool = false
	mu    sync.Mutex
)

func main() {
	// Simulate service readiness after 10 seconds
	go func() {
		time.Sleep(10 * time.Second)
		mu.Lock()
		ready = true
		mu.Unlock()
	}()

	// Handle root endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, Kubernetes!")
	})

	// Handle liveness probe
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})

	// Handle readiness probe
	http.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()

		if ready {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "Ready")
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(w, "Not Ready")
		}
	})

	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
