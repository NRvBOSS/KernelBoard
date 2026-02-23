package services

import (
	"bufio"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/NRvBOSS/KernelBoard/api/internal/models"
)

func readCPU() (int64, int64, error) {
	file, err := os.Open("/proc/stat")
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	var total, idle int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "cpu ") {
			fields := strings.Fields(line)

			for i := 1; i < len(fields); i++ {
				value, err := strconv.ParseInt(fields[i], 10, 64)
				if err != nil {
					continue
				}

				total += value

				// idle + iowait
				if i == 4 || i == 5 {
					idle += value
				}
			}
			break
		}
	}

	return total, idle, nil
}

func GetCPUStats() (models.CPUStats, error) {
	t1, i1, err := readCPU()
	if err != nil {
		return models.CPUStats{}, err
	}

	time.Sleep(100 * time.Millisecond)

	t2, i2, err := readCPU()
	if err != nil {
		return models.CPUStats{}, err
	}

	deltaTotal := t2 - t1
	deltaIdle := i2 - i1

	if deltaTotal == 0 {
		return models.CPUStats{}, nil
	}

	usage := float64(deltaTotal-deltaIdle) / float64(deltaTotal) * 100

	return models.CPUStats{
		Usage: usage,
		Cores: runtime.NumCPU(),
	}, nil
}
