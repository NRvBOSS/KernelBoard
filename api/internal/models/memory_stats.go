package models

type MemoryStats struct {
	TotalGB      string `json:"total_gb"`
	UsedGB       string `json:"used_gb"`
	FreeGB       string `json:"free_gb"`
	UsagePercent string `json:"usage_percent"`
}
