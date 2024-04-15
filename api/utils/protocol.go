package utils

import (
	"net"
	"strings"
)

func IsValidIPv4(ip string) bool {
	parts := strings.Split(ip, "/")
	if len(parts) != 2 {
		return false
	}
	return net.ParseIP(parts[0]).To4() != nil
}

func IsValidIPv6(ip string) bool {
	parts := strings.Split(ip, "/")
	if len(parts) != 2 {
		return false
	}
	return net.ParseIP(parts[0]).To16() != nil && !IsValidIPv4(ip)
}
