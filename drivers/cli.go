package drivers

import (
	"encoding/json"
	"fmt"
	"telemetry-logging/logger"
)

// CLIDriver outputs logs to the Command Line Interface.
type CLIDriver struct {
	Level logger.LogLevel
}

// NewCLIDriver creates a new CLIDriver with the specified log level.
func NewCLIDriver(level logger.LogLevel) *CLIDriver {
	return &CLIDriver{Level: level}
}

// Log writes the log entry to the CLI.
func (c *CLIDriver) Log(entry logger.LogEntry) error {
	if entry.Level < c.Level {
		return nil
	}
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
