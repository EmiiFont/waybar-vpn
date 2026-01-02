// Package config
package config

import "flag"

type Config struct {
	Client           string
	Action           string
	IconConnected    string
	IconDisconnected string
	TooltipFormat    string
}

func Parse() *Config {
	cfg := &Config{}

	flag.StringVar(&cfg.Client, "client", "ivpn", "VPN client type (e.g., ivpn)")
	flag.StringVar(&cfg.Action, "action", "status", "Action to perform: status or disconnect")
	flag.StringVar(&cfg.IconConnected, "icon-connected", "\uf510", "Icon for connected state (Nerd Font)")
	flag.StringVar(&cfg.IconDisconnected, "icon-disconnected", "\uf512", "Icon for disconnected state (Nerd Font)")
	flag.StringVar(&cfg.TooltipFormat, "tooltip-format", "Connected to {name} ({ip})", "Tooltip format: use {name} and {ip} placeholders")

	flag.Parse()
	return cfg
}
