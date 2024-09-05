package drivers

import (
	"telemetry-logging/logger"
	"testing"
)

// MockKafkaDriver simulates the behavior of a Kafka driver
type MockKafkaDriver struct {
	entries []logger.LogEntry
}

func (m *MockKafkaDriver) Log(entry logger.LogEntry) error {
	m.entries = append(m.entries, entry)
	return nil
}

func TestKafkaDriver(t *testing.T) {
	mockDriver := &MockKafkaDriver{}
	log := logger.NewMultiLogger()
	log.AddDriver(mockDriver)

	log.Debug("Test log entry for Kafka", nil)

	if len(mockDriver.entries) != 1 {
		t.Errorf("Expected 1 log entry, got %d", len(mockDriver.entries))
	}

	entry := mockDriver.entries[0]
	if entry.Message != "Test log entry for Kafka" {
		t.Errorf("Expected 'Test log entry for Kafka', got '%s'", entry.Message)
	}
}
