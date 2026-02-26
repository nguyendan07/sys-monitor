package models

type SystemMetrics struct {
	CPUUsage	float64	`json:"cpuUsage"`
	MemoryUsage	float64	`json:"memoryUsage"`
	DiskUsage	float64	`json:"diskUsage"`
}
