package vpn

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
)

type NordVPNDetector struct{}

func (d *NordVPNDetector) IsConnected() (bool, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return false, err
	}

	for _, iface := range ifaces {
		if iface.Name == "nordlynx" {
			return true, nil
		}
	}
	return false, nil
}

func (d *NordVPNDetector) GetName() (string, error) {
	output, err := d.runNordVPNCmd("status")
	if err != nil {
		return "NordVPN", nil // Fallback
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
	return "NordVPN", nil
}

func (d *NordVPNDetector) GetIP() (string, error) {
	iface, err := net.InterfaceByName("nordlynx")
	if err != nil {
		return "", err
	}

	addrs, err := iface.Addrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.To4() != nil {
			return ipnet.IP.String(), nil
		}
	}
	return "", fmt.Errorf("no NordVPN IP found")
}

func (d *NordVPNDetector) Disconnect() error {
	_, err := d.runNordVPNCmd("disconnect")
	return err
}

func (d *NordVPNDetector) runNordVPNCmd(args ...string) (string, error) {
	cmd := exec.Command("nordvpn", args...)
	cmd.Env = append(cmd.Env, "PATH=/usr/local/bin:/usr/bin:/bin")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("nordvpn command failed: %v, output: %s", err, string(output))
	}

	return string(output), nil
}
