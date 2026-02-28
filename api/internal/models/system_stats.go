package models

type SystemStats struct {
	Hostname      string `json:"hostname"`
	KernelVersion string `json:"kernel_version"`
	Architecture  string `json:"architecture"`
	GoVersion     string `json:"go_version"`
}
