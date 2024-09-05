// syslog.go

package drivers

import (
	"log/syslog"
	"telemetry-logging/logger"
)

type SyslogDriver struct {
	Level  logger.LogLevel
	Writer *syslog.Writer
}

func NewSyslogDriver(config map[string]interface{}) (logger.Driver, error) {
	writer, err := syslog.Dial("udp", "localhost:514", syslog.LOG_INFO, "telemetry-app")
	if err != nil {
		return nil, err
	}
	return &SyslogDriver{
		Level:  logger.INFO,
		Writer: writer,
	}, nil
}

func (s *SyslogDriver) Log(entry logger.LogEntry) error {
	if entry.Level < s.Level {
		return nil
	}
	return s.Writer.Info(entry.Message)
}
