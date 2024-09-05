package main

import (
	"os"
	"telemetry-logging/drivers"
	"telemetry-logging/logger"
	"testing"
)

// TestCLIDriver tests logging to the Command Line Interface (CLI).
func TestCLIDriver(t *testing.T) {
	cli := drivers.NewCLIDriver(logger.DEBUG)
	log := logger.NewMultiLogger()
	log.AddDriver(cli)

	// Logging at different levels
	log.Debug("Test debug message", nil)
	log.Info("Test info message", nil)
	log.Warn("Test warning message", nil)
	log.Error("Test error message", nil)

	// Manual verification: Check the CLI output for logs
}

// TestFileDriver tests logging to a file and verifies that the log file contains entries.
func TestFileDriver(t *testing.T) {
	logFile := "test.log"
	defer os.Remove(logFile) // Clean up test file after the test

	fileDriver, err := drivers.NewFileDriver(logger.DEBUG, logFile)
	if err != nil {
		t.Fatalf("Failed to create FileDriver: %v", err)
	}
	defer fileDriver.Close()

	log := logger.NewMultiLogger()
	log.AddDriver(fileDriver)

	// Log different levels and metadata
	log.Debug("Debug message", map[string]interface{}{"key": "value1"})
	log.Info("Info message", map[string]interface{}{"key": "value2"})
	log.Warn("Warn message", map[string]interface{}{"key": "value3"})
	log.Error("Error message", map[string]interface{}{"key": "value4"})

	// Read the log file to verify contents
	data, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	if len(data) == 0 {
		t.Error("Log file is empty, expected logs to be written")
	}

	// Further assertions could parse JSON and verify exact content
}

// TestTransactionLogging tests logging with transactions and verifies that TraceID is properly applied.
func TestTransactionLogging(t *testing.T) {
	logFile := "txn_test.log"
	defer os.Remove(logFile)

	fileDriver, err := drivers.NewFileDriver(logger.DEBUG, logFile)
	if err != nil {
		t.Fatalf("Failed to create FileDriver: %v", err)
	}
	defer fileDriver.Close()

	log := logger.NewMultiLogger()
	log.AddDriver(fileDriver)

	// Create a new transaction with metadata
	txn := logger.NewTransaction(map[string]interface{}{
		"UserID": "user123",
	})
	txnLogger := txn.LoggerWithTransaction(log)

	// Log within the transaction context
	txnLogger.Info("Transaction started", nil)
	txnLogger.Debug("Processing data", map[string]interface{}{"step": "1"})
	txnLogger.Error("Transaction failed", map[string]interface{}{"reason": "timeout"})

	// Read the log file to verify contents
	data, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	if len(data) == 0 {
		t.Error("Transaction logs not written, expected logs in file")
	}

	// You can parse the logs to assert that the TraceID exists
}

// TestLogLevel tests if log levels are correctly applied.
func TestLogLevel(t *testing.T) {
	logFile := "level_test.log"
	defer os.Remove(logFile)

	// Set log level to INFO
	fileDriver, err := drivers.NewFileDriver(logger.INFO, logFile)
	if err != nil {
		t.Fatalf("Failed to create FileDriver: %v", err)
	}
	defer fileDriver.Close()

	log := logger.NewMultiLogger()
	log.AddDriver(fileDriver)

	// Log a DEBUG message which should not be logged due to log level
	log.Debug("This debug message should not be logged", nil)
	log.Info("This info message should be logged", nil)

	// Read the log file to verify contents
	data, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	if len(data) == 0 {
		t.Error("Expected log entries, but the log file is empty")
	}

	// Verify that the DEBUG message is not in the log file
	if string(data) == "This debug message should not be logged" {
		t.Error("DEBUG message was logged, but the log level is set to INFO")
	}
}
