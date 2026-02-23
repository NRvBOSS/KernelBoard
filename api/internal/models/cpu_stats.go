package models

type CPUStats struct {
	Usage float64 `json:"usage"`
	Cores int     `json:"cores"`
}
