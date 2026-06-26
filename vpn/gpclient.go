package vpn

type GPClientDetector struct{}

func (d *GPClientDetector) IsConnected() (bool, error) {
	return hasInterfaceWithPrefix("tun", "gp")
}

func (d *GPClientDetector) GetName() (string, error) {
	return "GlobalProtect", nil
}

func (d *GPClientDetector) GetIP() (string, error) {
	return getIPv4ByInterfacePrefix("tun", "gp")
}

func (d *GPClientDetector) Disconnect() error {
	_, err := runCommand("gpclient", "disconnect")
	return err
}
