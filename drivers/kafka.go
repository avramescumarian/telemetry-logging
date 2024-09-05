package drivers

import (
	"context"
	"telemetry-logging/logger"

	"github.com/segmentio/kafka-go"
)

type KafkaDriver struct {
	Level       logger.LogLevel
	KafkaWriter *kafka.Writer
}

func NewKafkaDriver(config map[string]interface{}) (logger.Driver, error) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{config["broker"].(string)},
		Topic:   config["topic"].(string),
	})

	return &KafkaDriver{
		Level:       logger.DEBUG,
		KafkaWriter: writer,
	}, nil
}

func (k *KafkaDriver) Log(entry logger.LogEntry) error {
	if entry.Level < k.Level {
		return nil
	}

	message := kafka.Message{
		Key:   []byte(entry.TraceID),
		Value: []byte(entry.Message),
	}

	err := k.KafkaWriter.WriteMessages(context.Background(), message)
	if err != nil {
		return err
	}

	return nil
}
