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

func hasInterfaceWithPrefix(prefixes ...string) (bool, error) {
	iface, err := findInterfaceByPrefix(prefixes...)
	if err != nil {
		return false, err
	}

	return iface != nil, nil
}

func getIPv4ByInterfacePrefix(prefixes ...string) (string, error) {
	iface, err := findInterfaceByPrefix(prefixes...)
	if err != nil {
		return "", err
	}
	if iface == nil {
		return "", fmt.Errorf("no VPN interface found")
	}

	return getInterfaceIPv4(iface)
}

func findInterfaceByPrefix(prefixes ...string) (*net.Interface, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range ifaces {
		for _, prefix := range prefixes {
			if strings.HasPrefix(iface.Name, prefix) {
				return &iface, nil
			}
		}
	}

	return nil, nil
}

func getInterfaceIPv4(iface *net.Interface) (string, error) {
	addrs, err := iface.Addrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.To4() != nil {
			return ipnet.IP.String(), nil
		}
	}

	return "", fmt.Errorf("no VPN IP found")
}

func runCommand(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	cmd.Env = append(cmd.Env, "PATH=/usr/local/bin:/usr/bin:/bin")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s command failed: %v, output: %s", name, err, string(output))
	}

	return string(output), nil
}

type IVPNDetector struct{}

func (d *IVPNDetector) IsConnected() (bool, error) {
	return hasInterfaceWithPrefix("tun", "wg")
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
	return getIPv4ByInterfacePrefix("tun", "wg")
}

func (d *IVPNDetector) Disconnect() error {
	_, err := d.runIVPNCmd("disconnect")
	return err
}

func (d *IVPNDetector) runIVPNCmd(args ...string) (string, error) {
	return runCommand("ivpn", args...)
}

func NewDetector(client string) VPNDetector {
	switch client {
	case "ivpn":
		return &IVPNDetector{}
	case "nordvpn":
		return &NordVPNDetector{}
	case "gpclient":
		return &GPClientDetector{}
	default:
		return &IVPNDetector{}
	}
}
