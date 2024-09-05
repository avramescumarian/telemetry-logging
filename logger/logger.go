package logger

import (
	"sync"
	"time"
)

// LogEntry represents a single log entry.
type LogEntry struct {
	Timestamp time.Time              `json:"timestamp"`
	Level     LogLevel               `json:"level"`
	Message   string                 `json:"message"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	TraceID   string                 `json:"trace_id,omitempty"`
}

// Logger defines the interface for logging.
type Logger interface {
	Debug(message string, metadata map[string]interface{})
	Info(message string, metadata map[string]interface{})
	Warn(message string, metadata map[string]interface{})
	Error(message string, metadata map[string]interface{})
	Log(level LogLevel, message string, metadata map[string]interface{})
	WithTrace(traceID string) Logger
}

// MultiLogger manages multiple logging drivers.
type MultiLogger struct {
	drivers []Driver
	level   LogLevel
	traceID string
	mutex   sync.Mutex
}

// Driver interface that all drivers must implement.
type Driver interface {
	Log(entry LogEntry) error
}

// NewMultiLogger creates a new MultiLogger.
func NewMultiLogger() *MultiLogger {
	return &MultiLogger{
		drivers: make([]Driver, 0),
		level:   DEBUG,
	}
}

// AddDriver adds a new driver to the logger.
func (m *MultiLogger) AddDriver(driver Driver) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.drivers = append(m.drivers, driver)
}

// SetLevel sets the minimum log level.
func (m *MultiLogger) SetLevel(level LogLevel) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.level = level
}

// WithTrace sets the TraceID for transaction logs.
func (m *MultiLogger) WithTrace(traceID string) Logger {
	newLogger := &MultiLogger{
		drivers: m.drivers, // Copy reference to the same drivers
		level:   m.level,   // Copy the log level
		traceID: traceID,   // Set the new TraceID
	}
	return newLogger
}

// Log creates a log entry and dispatches it to all drivers.
func (m *MultiLogger) Log(level LogLevel, message string, metadata map[string]interface{}) {
	if level < m.level {
		return
	}
	entry := LogEntry{
		Timestamp: time.Now(),
		Level:     level,
		Message:   message,
		Metadata:  metadata,
		TraceID:   m.traceID,
	}
	for _, driver := range m.drivers {
		driver.Log(entry)
	}
}

func (m *MultiLogger) Debug(message string, metadata map[string]interface{}) {
	m.Log(DEBUG, message, metadata)
}

func (m *MultiLogger) Info(message string, metadata map[string]interface{}) {
	m.Log(INFO, message, metadata)
}

func (m *MultiLogger) Warn(message string, metadata map[string]interface{}) {
	m.Log(WARN, message, metadata)
}

func (m *MultiLogger) Error(message string, metadata map[string]interface{}) {
	m.Log(ERROR, message, metadata)
}
