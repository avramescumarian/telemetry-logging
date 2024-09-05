package drivers

import (
	"telemetry-logging/logger"
	"testing"
)

// Test CLI driver
func TestCLIDriver(t *testing.T) {
	cli := NewCLIDriver(logger.DEBUG)
	log := logger.NewMultiLogger()
	log.AddDriver(cli)

	// Test CLI output
	err := cli.Log(logger.LogEntry{
		Level:   logger.DEBUG,
		Message: "CLI driver test",
	})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
