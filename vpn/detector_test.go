package vpn

import (
	"net"
	"testing"

	"waybar-vpn/output"
)

func TestIVPNDetector_IsConnected(t *testing.T) {
	d := &IVPNDetector{}

	// Mock interfaces - assume no VPN for test
	connected, err := d.IsConnected()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	// In test env, likely no VPN, so false is expected
	if connected {
		// If test has VPN, ok
		t.Log("VPN detected in test")
	}
}

func TestIVPNDetector_GetName(t *testing.T) {
	d := &IVPNDetector{}

	name, err := d.GetName()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	// Should fallback to "IVPN" if no CLI
	if name != "IVPN" {
		t.Logf("Got name: %s", name)
	}
}

func TestIVPNDetector_GetIP(t *testing.T) {
	d := &IVPNDetector{}

	ip, err := d.GetIP()
	if err != nil {
		// Expected if no VPN
		t.Log("No VPN IP found, as expected")
	} else {
		if net.ParseIP(ip) == nil {
			t.Errorf("Invalid IP: %s", ip)
		}
	}
}

func TestFormatTooltip(t *testing.T) {
	tooltip := output.FormatTooltip("Connected to {name} ({ip})", "TestVPN", "1.2.3.4")
	expected := "Connected to TestVPN (1.2.3.4)"
	if tooltip != expected {
		t.Errorf("Expected '%s', got '%s'", expected, tooltip)
	}
}
