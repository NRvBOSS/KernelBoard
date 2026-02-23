package services

import (
	"os"
	"strconv"
	"strings"

	"github.com/NRvBOSS/KernelBoard/api/internal/models"
)

func GetUptStats() (models.UptStats, error) {

	data, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return models.UptStats{}, err
	}

	fields := strings.Fields(string(data))
	if len(fields) < 1 {
		return models.UptStats{}, nil
	}

	uptimeSeconds, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return models.UptStats{}, err
	}

	uptimeHours := uptimeSeconds / 3600

	return models.UptStats{
		UptimeSeconds: uptimeSeconds,
		UptimeHours:   uptimeHours,
	}, nil
}
