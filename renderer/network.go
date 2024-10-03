package renderer

import (
	"net"
	"os"
)

// NetworkInfo provides information about the network.
type NetworkInfo struct{}

// NewNetworkInfo returns an instance of NetworkInfo.
func NewNetworkInfo() *NetworkInfo {
	return &NetworkInfo{}
}

// Hostname returns the hostname of the machine.
func (n *NetworkInfo) Hostname() (string, error) {
	return os.Hostname()
}

// IPAddresses returns the list of IP addresses.
func (n *NetworkInfo) IPAddresses() ([]string, error) {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}
	return ips, nil
}
