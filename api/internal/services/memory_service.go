package services

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/NRvBOSS/KernelBoard/api/internal/models"
)

func GetMemoryStats() (models.MemoryStats, error) {

	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return models.MemoryStats{}, err
	}
	defer file.Close()

	var totalStr, freeStr string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "MemTotal") {
			totalStr = strings.Fields(line)[1]
		}

		if strings.HasPrefix(line, "MemFree") {
			freeStr = strings.Fields(line)[1]
		}
	}

	totalKB, err := strconv.ParseInt(totalStr, 10, 64)
	if err != nil {
		return models.MemoryStats{}, err
	}

	freeKB, err := strconv.ParseInt(freeStr, 10, 64)
	if err != nil {
		return models.MemoryStats{}, err
	}

	usedKB := totalKB - freeKB

	totalGB := float64(totalKB) / 1024 / 1024
	freeGB := float64(freeKB) / 1024 / 1024
	usedGB := float64(usedKB) / 1024 / 1024

	usagePercent := (float64(usedKB) / float64(totalKB)) * 100

	stats := models.MemoryStats{
		TotalGB:      strconv.FormatFloat(totalGB, 'f', 2, 64),
		UsedGB:       strconv.FormatFloat(usedGB, 'f', 2, 64),
		FreeGB:       strconv.FormatFloat(freeGB, 'f', 2, 64),
		UsagePercent: strconv.FormatFloat(usagePercent, 'f', 2, 64),
	}

	return stats, nil
}
