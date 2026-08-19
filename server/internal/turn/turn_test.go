package turn

import (
	"net"
	"testing"
)

func TestGetPublicIP_WithValidIPv4(t *testing.T) {
	const expected = "203.0.113.195"
	t.Setenv("PUBLIC_IP", expected)

	ip := getPublicIP()
	if ip == nil {
		t.Fatalf("expected non-nil IP, got nil")
	}
	if ip.String() != expected {
		t.Errorf("expected %s, got %s", expected, ip.String())
	}
}

func TestGetPublicIP_WithValidIPv6(t *testing.T) {
	const input = "2001:db8::1"
	t.Setenv("PUBLIC_IP", input)

	ip := getPublicIP()
	if ip == nil {
		t.Fatalf("expected non-nil IP, got nil")
	}
	expectedIP := net.ParseIP(input)
	if !ip.Equal(expectedIP) {
		t.Errorf("expected %s, got %s", expectedIP.String(), ip.String())
	}
}

func TestGetPublicIP_WithWhitespace(t *testing.T) {
	const expected = "198.51.100.22"
	t.Setenv("PUBLIC_IP", "  "+expected+" \n\t")

	ip := getPublicIP()
	if ip == nil {
		t.Fatalf("expected non-nil IP, got nil")
	}
	if ip.String() != expected {
		t.Errorf("expected %s, got %s", expected, ip.String())
	}
}

func TestGetPublicIP_WithInvalidIP(t *testing.T) {
	t.Setenv("PUBLIC_IP", "invalid.ip.address.format")

	// Invalid PUBLIC_IP should not panic, but fall back to auto-detection.
	// We verify that getPublicIP() executes gracefully.
	_ = getPublicIP()
}

func TestGetLocalIP(t *testing.T) {
	ip := getLocalIP()
	if ip == nil {
		t.Fatalf("expected non-nil local IP, got nil")
	}
}
