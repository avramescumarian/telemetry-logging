package drivers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"telemetry-logging/logger"

	"github.com/elastic/go-elasticsearch/v7"
)

type ElasticsearchDriver struct {
	Level  logger.LogLevel
	Client *elasticsearch.Client
	Index  string
}

func NewElasticsearchDriver(config map[string]interface{}) (logger.Driver, error) {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, err
	}
	return &ElasticsearchDriver{
		Level:  logger.DEBUG,
		Client: es,
		Index:  config["index"].(string),
	}, nil
}

func (e *ElasticsearchDriver) Log(entry logger.LogEntry) error {
	if entry.Level < e.Level {
		return nil
	}

	// Convert the log entry into a JSON document
	doc := map[string]interface{}{
		"timestamp": entry.Timestamp,
		"level":     entry.Level.String(),
		"message":   entry.Message,
		"metadata":  entry.Metadata,
		"trace_id":  entry.TraceID,
	}

	// Marshal the doc to JSON
	jsonData, err := json.Marshal(doc)
	if err != nil {
		return fmt.Errorf("failed to marshal log entry to JSON: %w", err)
	}

	// Use a bytes.Reader to pass the JSON data as an io.Reader
	res, err := e.Client.Index(
		e.Index,
		bytes.NewReader(jsonData),
		e.Client.Index.WithContext(context.Background()),
	)
	if err != nil {
		return fmt.Errorf("failed to send log entry to Elasticsearch: %w", err)
	}
	defer res.Body.Close()

	return nil
}
