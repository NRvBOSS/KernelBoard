package models

type SystemStats struct {
	Hostname      string `json:"hostname"`
	KernelVersion string `json:"kernelversion"`
	Architecture  string `json:"architecture"`
	GoVersion     string `json:"goversion"`
}
