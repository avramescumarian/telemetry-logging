package drivers

import (
	"database/sql"
	"fmt"
	"telemetry-logging/logger"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

type DBDriver struct {
	Level logger.LogLevel
	DB    *sql.DB
}

func NewDBDriver(config map[string]interface{}) (logger.Driver, error) {
	dsn := config["dsn"].(string) // Data Source Name (e.g., "user:password@tcp(localhost:3306)/logdb")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return &DBDriver{
		Level: logger.INFO,
		DB:    db,
	}, nil
}

func (d *DBDriver) Log(entry logger.LogEntry) error {
	if entry.Level < d.Level {
		return nil
	}

	query := "INSERT INTO logs (timestamp, level, message, trace_id) VALUES (?, ?, ?, ?)"
	_, err := d.DB.Exec(query, entry.Timestamp, entry.Level.String(), entry.Message, entry.TraceID)
	if err != nil {
		return fmt.Errorf("failed to insert log into database: %w", err)
	}
	return nil
}
