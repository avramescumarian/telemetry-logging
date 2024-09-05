package drivers

import (
	"encoding/json"
	"os"
	"sync"
	"telemetry-logging/logger"
)

// FileDriver writes logs to a file in JSON format.
type FileDriver struct {
	Level    logger.LogLevel
	FilePath string
	file     *os.File
	mutex    sync.Mutex
}

// NewFileDriver creates a new FileDriver.
func NewFileDriver(level logger.LogLevel, filePath string) (*FileDriver, error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &FileDriver{
		Level:    level,
		FilePath: filePath,
		file:     file,
	}, nil
}

// Log writes the log entry to the file.
func (f *FileDriver) Log(entry logger.LogEntry) error {
	if entry.Level < f.Level {
		return nil
	}
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	f.mutex.Lock()
	defer f.mutex.Unlock()
	_, err = f.file.WriteString(string(data) + "\n")
	return err
}

// Close closes the file.
func (f *FileDriver) Close() error {
	return f.file.Close()
}
