package config

import (
	"flag"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	// Reset flags
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	// Set test args
	os.Args = []string{"test", "--client=nord", "--action=disconnect", "--icon-connected=test"}

	cfg := Parse()

	if cfg.Client != "nord" {
		t.Errorf("Expected client 'nord', got '%s'", cfg.Client)
	}
	if cfg.Action != "disconnect" {
		t.Errorf("Expected action 'disconnect', got '%s'", cfg.Action)
	}
	if cfg.IconConnected != "test" {
		t.Errorf("Expected icon 'test', got '%s'", cfg.IconConnected)
	}
}
