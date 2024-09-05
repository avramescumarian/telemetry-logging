package drivers

import (
	"telemetry-logging/logger"
	"testing"
)

// MockSyslogDriver simulates the behavior of a Syslog driver
type MockSyslogDriver struct {
	entries []logger.LogEntry
}

func (m *MockSyslogDriver) Log(entry logger.LogEntry) error {
	m.entries = append(m.entries, entry)
	return nil
}

func TestSyslogDriver(t *testing.T) {
	mockDriver := &MockSyslogDriver{}
	log := logger.NewMultiLogger()
	log.AddDriver(mockDriver)

	log.Debug("Test log entry for Syslog", nil)

	if len(mockDriver.entries) != 1 {
		t.Errorf("Expected 1 log entry, got %d", len(mockDriver.entries))
	}

	entry := mockDriver.entries[0]
	if entry.Message != "Test log entry for Syslog" {
		t.Errorf("Expected 'Test log entry for Syslog', got '%s'", entry.Message)
	}
}
