package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/NRvBOSS/KernelBoard/api/internal/services"
)

func DiskHandler(w http.ResponseWriter, r *http.Request) {
	stats, err := services.GetDiskStats()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
