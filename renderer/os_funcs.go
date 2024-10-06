package renderer

import (
	"os/exec"
	"runtime"
	"strings"
)

// osInfo provides information about the OS.
type OsInfo struct{}

// newOSInfo returns an instance of osInfo.
func newOSInfo() *OsInfo {
	return &OsInfo{}
}

// Name returns the current operating system name.
func (o *OsInfo) Name() string {
	return runtime.GOOS
}

// Version returns the current OS version.
func (o *OsInfo) Version() string {
	switch runtime.GOOS {
	case "windows":
		return windowsVersion()
	case "darwin":
		return macVersion()
	case "linux":
		return linuxVersion()
	default:
		return "unknown"
	}
}

// macVersion returns the macOS version.
func macVersion() string {
	out, _ := exec.Command("sw_vers", "-productVersion").Output()
	return strings.TrimSpace(string(out))
}

// linuxVersion returns the Linux version.
func linuxVersion() string {
	out, _ := exec.Command("uname", "-r").Output()
	return strings.TrimSpace(string(out))
}

// windowsVersion returns the Windows version.
func windowsVersion() string {
	out, _ := exec.Command("cmd", "ver").Output()
	return strings.TrimSpace(string(out))
}
