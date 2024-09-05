package config

import (
	"encoding/json"
	"os"
	"telemetry-logging/drivers"
	"telemetry-logging/logger"
)

type Config struct {
	LogLevel string `json:"log_level"`
	Drivers  []struct {
		Type     string                 `json:"type"`
		Settings map[string]interface{} `json:"settings"`
	} `json:"drivers"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var config Config
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func InitializeLogger(config *Config) (*logger.MultiLogger, error) {
	logLevel, err := logger.ParseLogLevel(config.LogLevel)
	if err != nil {
		return nil, err
	}
	ml := logger.NewMultiLogger()
	ml.SetLevel(logLevel)

	for _, drv := range config.Drivers {
		switch drv.Type {
		case "cli":
			cliDriver := drivers.NewCLIDriver(logLevel)
			ml.AddDriver(cliDriver)
		case "file":
			filePath, ok := drv.Settings["file_path"].(string)
			if !ok {
				return nil, nil
			}
			fileDriver, err := drivers.NewFileDriver(logLevel, filePath)
			if err != nil {
				return nil, err
			}
			ml.AddDriver(fileDriver)
		default:
			// Unknown driver type
		}
	}
	return ml, nil
}
