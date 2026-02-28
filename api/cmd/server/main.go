package main

import (
	"log"
	"net/http"

	"github.com/NRvBOSS/KernelBoard/api/internal/handlers"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/memory", handlers.MemoryHandler)
	mux.HandleFunc("/api/cpu", handlers.CPUHandler)
	mux.HandleFunc("/api/load", handlers.LoadHandler)
	mux.HandleFunc("/api/disk", handlers.DiskHandler)
	mux.HandleFunc("/api/uptime", handlers.UptHandler)
	mux.HandleFunc("/api/system", handlers.SystemHandler)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", enableCORS(mux))
}
