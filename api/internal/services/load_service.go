package services

import (
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/NRvBOSS/KernelBoard/api/internal/models"
)

func getLoad() (float64, float64, float64, error) {
	data, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		return 0, 0, 0, err
	}

	fields := strings.Fields(string(data))
	if len(fields) < 3 {
		return 0, 0, 0, nil
	}

	load1, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return 0, 0, 0, err
	}

	load5, err := strconv.ParseFloat(fields[1], 64)
	if err != nil {
		return 0, 0, 0, err
	}

	load15, err := strconv.ParseFloat(fields[2], 64)
	if err != nil {
		return 0, 0, 0, err
	}

	return load1, load5, load15, nil

}
func GetLoadStats() (models.LoadStats, error) {
	l1, l5, l15, err := getLoad()
	if err != nil {
		return models.LoadStats{}, err
	}

	return models.LoadStats{
		Load1:  l1,
		Load5:  l5,
		Load15: l15,
		Cores:  runtime.NumCPU(),
	}, nil
}
