package services

import (
	"bufio"
	"os"
	"strings"

	"github.com/NRvBOSS/KernelBoard/api/internal/models"
)

func GetMemoryStats() (models.MemoryStats, error) {

	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return models.MemoryStats{}, err
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

	stats := models.MemoryStats{
		Total: total,
		Free:  free,
	}

	return stats, nil
}
