package main

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type MemoryStats struct {
	Total string `json:"total"`
	Free  string `json:"free"`
}

func memoryHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()

	var total, free string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "MemTotal") {
			total = strings.Fields(line)[1] + " kB"
		}
		if strings.HasPrefix(line, "MemFree") {
			free = strings.Fields(line)[1] + " kB"
		}
	}

	stats := MemoryStats{
		Total: total,
		Free:  free,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func main() {
	http.HandleFunc("/api/memory", memoryHandler)

	http.ListenAndServe(":8080", nil)
}
