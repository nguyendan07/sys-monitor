package collector

import (
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/disk"

	"github.com/nguyendan07/sys-monitor/internal/models"
)

func GetMetrics() (models.SystemMetrics, error) {
	cpuPercents, err := cpu.Percent(time.Second, false)
	if err != nil || len(cpuPercents) == 0 {
		return models.SystemMetrics{}, err
	}

	vm, err := mem.VirtualMemory()
	if err != nil {
		return models.SystemMetrics{}, err
	}

	d, err := disk.Usage("/")
	if err != nil {
		return models.SystemMetrics{}, err
	}

	return models.SystemMetrics{
		CPUUsage: cpuPercents[0],
		MemoryUsage: vm.UsedPercent,
		DiskUsage: d.UsedPercent,
	}, nil
}
