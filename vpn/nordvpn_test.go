package vpn

import (
	"testing"
)

func TestNordVPNDetector_IsConnected(t *testing.T) {
	d := &NordVPNDetector{}

	connected, err := d.IsConnected()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	// In test env, likely no VPN
	if connected {
		t.Log("NordVPN detected in test")
	}
}

func TestNordVPNDetector_GetName(t *testing.T) {
	d := &NordVPNDetector{}

	name, err := d.GetName()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	// Fallback to "NordVPN"
	if name != "NordVPN" {
		t.Logf("Got name: %s", name)
	}
}

func TestNordVPNDetector_GetIP(t *testing.T) {
	d := &NordVPNDetector{}

	ip, err := d.GetIP()
	if err != nil {
		// Expected if no VPN
		t.Log("No NordVPN IP found, as expected")
	} else {
		t.Logf("Got IP: %s", ip)
	}
}
