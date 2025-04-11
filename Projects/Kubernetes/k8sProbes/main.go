package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type HealthStatus struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

type AppInfo struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	Hostname  string `json:"hostname"`
	Timestamp string `json:"timestamp"`
}

var (
	isReady     bool
	isHealthy   bool
	startupTime = time.Now()
)

func main() {
	// Inicialmente a aplicação não está pronta
	isReady = false
	isHealthy = true

	// Simula um tempo de inicialização
	go func() {
		time.Sleep(5 * time.Second)
		isReady = true
		log.Println("Aplicação pronta para receber tráfego")
	}()

	// Configura os handlers HTTP
	http.HandleFunc("/", infoHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/ready", readyHandler)
	http.HandleFunc("/startup", startupHandler)

	// Inicia o servidor
	port := ":8081"
	log.Printf("Servidor iniciado na porta %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// Handler principal que retorna informações da aplicação
func infoHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Call 'info' at %s", time.Now().Format(time.RFC3339))

	hostname, _ := os.Hostname()

	info := AppInfo{
		Name:      "API Demo Kubernetes Probes",
		Version:   "1.0.0",
		Hostname:  hostname,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

// Liveness probe - indica se a aplicação está rodando
func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Call 'health' at %s", time.Now().Format(time.RFC3339))

	w.Header().Set("Content-Type", "application/json")

	if !isHealthy {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(HealthStatus{
			Status:  "DOWN",
			Message: "Aplicação não está saudável",
		})
		return
	}

	json.NewEncoder(w).Encode(HealthStatus{
		Status: "UP",
	})
}

// Readiness probe - indica se a aplicação está pronta para receber tráfego
func readyHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Call 'ready' at %s", time.Now().Format(time.RFC3339))

	w.Header().Set("Content-Type", "application/json")

	if !isReady {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(HealthStatus{
			Status:  "DOWN",
			Message: "Aplicação não está pronta",
		})
		return
	}

	json.NewEncoder(w).Encode(HealthStatus{
		Status: "UP",
	})
}

// Startup probe - indica se a aplicação terminou sua inicialização
func startupHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Call 'start' at %s", time.Now().Format(time.RFC3339))

	w.Header().Set("Content-Type", "application/json")

	// Considera que a inicialização terminou após 10 segundos
	if time.Since(startupTime) < 10*time.Second {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(HealthStatus{
			Status:  "DOWN",
			Message: "Aplicação em inicialização",
		})
		return
	}

	json.NewEncoder(w).Encode(HealthStatus{
		Status: "UP",
	})
}
