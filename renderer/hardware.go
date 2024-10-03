package renderer

import (
	"runtime"
)

// HardwareInfo provides information about the system hardware.
type HardwareInfo struct{}

// NewHardwareInfo returns an instance of HardwareInfo.
func NewHardwareInfo() *HardwareInfo {
	return &HardwareInfo{}
}

// Cores returns the number of CPU cores.
func (h *HardwareInfo) Cores() int {
	return runtime.NumCPU()
}

// Arch returns the system architecture.
func (h *HardwareInfo) Arch() string {
	return runtime.GOARCH
}
