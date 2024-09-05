package config

import (
	"os"
	"testing"
)

// Test loading configuration file
func TestLoadConfig(t *testing.T) {
	// Create a temporary config file for testing
	configData := `{
        "log_level": "DEBUG",
        "drivers": [
            {"type": "cli", "settings": {}},
            {"type": "file", "settings": {"file_path": "logs/app.log"}}
        ]
    }`

	file, err := os.CreateTemp("", "config*.json")
	if err != nil {
		t.Fatalf("Failed to create temp config file: %v", err)
	}
	defer os.Remove(file.Name())

	file.WriteString(configData)
	file.Close()

	// Test loading the configuration
	config, err := LoadConfig(file.Name())
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if config.LogLevel != "DEBUG" {
		t.Errorf("Expected log level 'DEBUG', got %s", config.LogLevel)
	}
}
