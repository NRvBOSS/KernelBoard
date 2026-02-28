package services

import (
	"os"
	"runtime"
	"strings"

	"github.com/NRvBOSS/KernelBoard/api/internal/models"
)

func GetSystemStats() (models.SystemStats, error) {

	// Hostname
	hostname, err := os.Hostname()
	if err != nil {
		return models.SystemStats{}, err
	}

	// Architecture
	arch := runtime.GOARCH

	// Go version
	goVersion := runtime.Version()

	// Kernel version
	data, err := os.ReadFile("/proc/version")
	if err != nil {
		return models.SystemStats{}, err
	}

	fields := strings.Fields(string(data))

	var kernelVersion string
	if len(fields) >= 3 {
		kernelVersion = fields[2]
	}

	return models.SystemStats{
		Hostname:      hostname,
		KernelVersion: kernelVersion,
		Architecture:  arch,
		GoVersion:     goVersion,
	}, nil
}
