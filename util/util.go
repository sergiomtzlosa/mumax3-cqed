// package util provides common utilities for all other packages.
package util

import (
	"net"
	"path"
	"strings"
)

// var Brms_vector [3]float64
// var Wc float64 = 0.0
// var TimeEvolution = false

// Remove extension from file name.
func NoExt(file string) string {
	ext := path.Ext(file)
	return file[:len(file)-len(ext)]
}

// returns all network interface addresses, without CIDR mask
func InterfaceAddrs() []string {
	addrs, _ := net.InterfaceAddrs()
	ips := make([]string, 0, len(addrs))
	for _, addr := range addrs {
		IpCidr := strings.Split(addr.String(), "/")
		ips = append(ips, IpCidr[0])
	}
	return ips
}
