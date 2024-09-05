package drivers

import (
	"os"
	"telemetry-logging/logger"
	"testing"
)

// Test file-based logging
func TestFileDriver(t *testing.T) {
	logFile := "test.log"
	defer os.Remove(logFile)

	fileDriver, err := NewFileDriver(logger.DEBUG, logFile)
	if err != nil {
		t.Fatalf("Failed to create FileDriver: %v", err)
	}
	defer fileDriver.Close()

	log := logger.NewMultiLogger()
	log.AddDriver(fileDriver)

	// Log test message
	log.Debug("Debug message", nil)

	// Check if the log file is created and not empty
	fileInfo, err := os.Stat(logFile)
	if err != nil {
		t.Fatalf("Failed to find log file: %v", err)
	}

	if fileInfo.Size() == 0 {
		t.Error("Expected log file to contain entries")
	}
}
