// Package vpn
package vpn

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
)

type VPNDetector interface {
	IsConnected() (bool, error)
	GetName() (string, error)
	GetIP() (string, error)
	Disconnect() error
}

type IVPNDetector struct{}

func (d *IVPNDetector) IsConnected() (bool, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return false, err
	}

	for _, iface := range ifaces {
		if strings.HasPrefix(iface.Name, "tun") || strings.HasPrefix(iface.Name, "wg") {
			return true, nil
		}
	}
	return false, nil
}

func (d *IVPNDetector) GetName() (string, error) {
	output, err := d.runIVPNCmd("status")
	if err != nil {
		return "IVPN", nil
	}

	lines := strings.SplitSeq(output, "\n")
	for line := range lines {
		if strings.Contains(line, "Server:") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				return strings.TrimSpace(parts[1]), nil
			}
		}
	}
	return "IVPN", nil
}

func (d *IVPNDetector) GetIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range ifaces {
		if strings.HasPrefix(iface.Name, "tun") || strings.HasPrefix(iface.Name, "wg") {
			addrs, err := iface.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.To4() != nil {
					return ipnet.IP.String(), nil
				}
			}
		}
	}
	return "", fmt.Errorf("no VPN IP found")
}

func (d *IVPNDetector) Disconnect() error {
	_, err := d.runIVPNCmd("disconnect")
	return err
}

func (d *IVPNDetector) runIVPNCmd(args ...string) (string, error) {
	cmd := exec.Command("ivpn", args...)
	cmd.Env = append(cmd.Env, "PATH=/usr/local/bin:/usr/bin:/bin")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("ivpn command failed: %v, output: %s", err, string(output))
	}

	return string(output), nil
}

func NewDetector(client string) VPNDetector {
	switch client {
	case "ivpn":
		return &IVPNDetector{}
	case "nordvpn":
		return &NordVPNDetector{}
	default:
		return &IVPNDetector{}
	}
}
