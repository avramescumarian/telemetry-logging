package drivers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"telemetry-logging/logger"
)

type HTTPDriver struct {
	Level    logger.LogLevel
	Endpoint string
}

func NewHTTPDriver(config map[string]interface{}) (logger.Driver, error) {
	return &HTTPDriver{
		Level:    logger.INFO,
		Endpoint: config["endpoint"].(string),
	}, nil
}

func (h *HTTPDriver) Log(entry logger.LogEntry) error {
	if entry.Level < h.Level {
		return nil
	}

	payload, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	_, err = http.Post(h.Endpoint, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	return nil
}
