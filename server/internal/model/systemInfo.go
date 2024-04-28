package model

type DiskInfo struct {
	Total   uint64
	Percent float64
	Used    uint64
}

type SystemInfo struct {
	Uptime          string
	Hostname        string
	Os              string
	Platform        string
	PlatformVersion string
	KernelVersion   string
	KernelArch      string
}

type MemoryInfo struct {
	Total   uint64
	Percent float64
	Used    uint64
}
