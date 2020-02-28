// Package fqdn attempts to find a Fully Qualified Domain Name for a host.
// This is a fork of https://github.com/Showmax/go-fqdn with some API changes.
package fqdn

import (
	"errors"
	"net"
	"os"
	"strings"
)

// Get Fully Qualified Domain Name
// Returns a string containing an FQDN.  Returns an error if no FQDN could be found.
func Get() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}

	addrs, err := net.LookupIP(hostname)
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			ip, err := ipv4.MarshalText()
			if err != nil {
				return "", err
			}
			hosts, err := net.LookupAddr(string(ip))
			if err != nil {
				return "", err
			}
			if len(hosts) == 0 {
				return "", errors.New("No hosts found for ip: " + string(ip))
			}
			return strings.TrimSuffix(hosts[0], "."), nil
		}
	}
	return "", errors.New("no FQDN could be found")
}
