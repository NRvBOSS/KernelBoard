package main

import (
	"log"
	"net/http"

	"github.com/NRvBOSS/KernelBoard/api/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/memory", handlers.MemoryHandler)
	mux.HandleFunc("/api/cpu", handlers.CPUHandler)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", mux)
}
