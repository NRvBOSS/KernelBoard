package services

import (
	"syscall"

	"github.com/NRvBOSS/KernelBoard/api/internal/models"
)

func GetDiskStats() (models.DiskStats, error) {
	var stat syscall.Statfs_t
	err := syscall.Statfs("/", &stat)
	if err != nil {
		return models.DiskStats{}, err
	}

	totalBytes := stat.Blocks * uint64(stat.Bsize)
	freeBytes := stat.Bfree * uint64(stat.Bsize)
	usedBytes := totalBytes - freeBytes

	totalGB := float64(totalBytes) / (1024 * 1024 * 1024)
	freeGB := float64(freeBytes) / (1024 * 1024 * 1024)
	usedGB := float64(usedBytes) / (1024 * 1024 * 1024)

	var usagePercent float64
	if totalGB > 0 {
		usagePercent = (usedGB / totalGB) * 100
	}

	return models.DiskStats{
		TotalGB:      totalGB,
		UsedGB:       usedGB,
		FreeGB:       freeGB,
		UsagePercent: usagePercent,
	}, nil

}
