package drivers

import (
	"telemetry-logging/logger"
	"testing"
)

// MockElasticsearchDriver simulates the behavior of an Elasticsearch driver
type MockElasticsearchDriver struct {
	entries []logger.LogEntry
}

func (m *MockElasticsearchDriver) Log(entry logger.LogEntry) error {
	m.entries = append(m.entries, entry)
	return nil
}

func TestElasticsearchDriver(t *testing.T) {
	mockDriver := &MockElasticsearchDriver{}
	log := logger.NewMultiLogger()
	log.AddDriver(mockDriver)

	log.Debug("Test log entry for Elasticsearch", nil)

	if len(mockDriver.entries) != 1 {
		t.Errorf("Expected 1 log entry, got %d", len(mockDriver.entries))
	}

	entry := mockDriver.entries[0]
	if entry.Message != "Test log entry for Elasticsearch" {
		t.Errorf("Expected 'Test log entry for Elasticsearch', got '%s'", entry.Message)
	}
}
