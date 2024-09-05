package logger

import (
	"testing"
)

// Test logging with mock driver
func TestBasicLogging(t *testing.T) {
	log := NewMultiLogger()
	log.SetLevel(DEBUG)

	// Mock driver for testing
	mockDriver := &MockDriver{}
	log.AddDriver(mockDriver)

	log.Debug("Test debug message", nil)

	if len(mockDriver.GetEntries()) != 1 {
		t.Errorf("Expected 1 log entry, got %d", len(mockDriver.GetEntries()))
	}

	entry := mockDriver.GetEntries()[0]
	if entry.Message != "Test debug message" {
		t.Errorf("Expected 'Test debug message', got '%s'", entry.Message)
	}
}

// MockDriver is a mock implementation of the Driver interface for testing.
type MockDriver struct {
	entries []LogEntry
}

func (m *MockDriver) Log(entry LogEntry) error {
	m.entries = append(m.entries, entry)
	return nil
}

func (m *MockDriver) GetEntries() []LogEntry {
	return m.entries
}
