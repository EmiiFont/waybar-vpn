package main

import (
	"fmt"
	"log"
	"os"

	"waybar-vpn/config"
	"waybar-vpn/output"
	"waybar-vpn/vpn"
)

func main() {
	cfg := config.Parse()

	detector := vpn.NewDetector(cfg.Client)

	switch cfg.Action {
	case "status":
		handleStatus(detector, cfg)
	case "disconnect":
		handleDisconnect(detector)
	default:
		log.Fatalf("Unknown action: %s", cfg.Action)
	}
}

func handleStatus(detector vpn.VPNDetector, cfg *config.Config) {
	connected, err := detector.IsConnected()
	if err != nil {
		log.Printf("Error checking connection: %v", err)
		connected = false
	}

	var text, class, tooltip string
	if connected {
		name, _ := detector.GetName()
		ip, _ := detector.GetIP()
		text = cfg.IconConnected
		class = "connected"
		tooltip = output.FormatTooltip(cfg.TooltipFormat, name, ip)
	} else {
		text = cfg.IconDisconnected
		class = "disconnected"
		tooltip = "Not connected to VPN"
	}

	out := output.WaybarOutput{
		Text:    text,
		Tooltip: tooltip,
		Class:   class,
	}

	if err := out.Print(); err != nil {
		log.Fatalf("Error outputting JSON: %v", err)
	}
}

func handleDisconnect(detector vpn.VPNDetector) {
	err := detector.Disconnect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error disconnecting: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Disconnected successfully")
}
