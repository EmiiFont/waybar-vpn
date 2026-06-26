package vpn

import "testing"

func TestGPClientDetector_GetName(t *testing.T) {
	d := &GPClientDetector{}

	name, err := d.GetName()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if name != "GlobalProtect" {
		t.Fatalf("expected GlobalProtect, got %q", name)
	}
}

func TestGPClientDetector_IsConnected(t *testing.T) {
	d := &GPClientDetector{}

	connected, err := d.IsConnected()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if connected {
		t.Log("GlobalProtect detected in test")
	}
}

func TestGPClientDetector_GetIP(t *testing.T) {
	d := &GPClientDetector{}

	ip, err := d.GetIP()
	if err != nil {
		t.Log("No GlobalProtect IP found, as expected")
		return
	}

	t.Logf("Got IP: %s", ip)
}
