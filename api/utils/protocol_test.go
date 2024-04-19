package utils

import (
	"testing"
)

func TestIsValidIPv4(t *testing.T) {
	tests := []struct {
		name     string
		ip       string
		expected bool
	}{
		{"valid IPv4", "10.5.2.6", true},
		{"invalid IPv4", "256.256.256.256", false},
		{"valid IPv6", "::1", false},
		{"invalid IPv6", "2001:0db8:85a3:0000:0000:8a2e:0370:7334", false},
		{"invalid format", "not an IP", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := IsValidIPv4(test.ip)
			if got != test.expected {
				t.Errorf("got %v, want %v", got, test.expected)
			}
		})
	}
}

func TestIsValidIPv6(t *testing.T) {
	tests := []struct {
		name     string
		ip       string
		expected bool
	}{
		{"valid IPv4", "192.168.0.1", false},
		{"invalid IPv4", "256.256.256.256", false},
		{"valid IPv6", "::1", true},
		{"invalid IPv6", "2001:0db8:85a3:0000:0000:8a2e:0370:7334", true},
		{"invalid format", "not an IP", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := IsValidIPv6(test.ip)
			if got != test.expected {
				t.Errorf("got %v, want %v", got, test.expected)
			}
		})
	}
}
