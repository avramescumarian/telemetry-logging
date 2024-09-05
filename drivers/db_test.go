package drivers

import (
	"telemetry-logging/logger"
	"testing"
)

// MockDBDriver simulates the behavior of a Database driver.
type MockDBDriver struct {
	entries []logger.LogEntry
}

func (m *MockDBDriver) Log(entry logger.LogEntry) error {
	m.entries = append(m.entries, entry)
	return nil
}

func TestDBDriver(t *testing.T) {
	mockDriver := &MockDBDriver{}
	log := logger.NewMultiLogger()
	log.AddDriver(mockDriver)

	// Log a test message
	log.Debug("Test log entry for DB", nil)

	if len(mockDriver.entries) != 1 {
		t.Errorf("Expected 1 log entry, got %d", len(mockDriver.entries))
	}

	entry := mockDriver.entries[0]
	if entry.Message != "Test log entry for DB" {
		t.Errorf("Expected 'Test log entry for DB', got '%s'", entry.Message)
	}
}
