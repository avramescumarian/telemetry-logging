
# Telemetry Logging Package

## Overview
The Telemetry Logging Package is designed to simplify logging across multiple platforms and provide a unified, extensible interface for logging in Go applications. It supports multiple drivers, transaction-style logs, log levels, and metadata for enhanced logging.

### Features
- **Multiple Log Levels**: Supports `DEBUG`, `INFO`, `WARN`, and `ERROR` log levels.
- **Transaction Logging**: Group related log entries with a unique `TraceID` for easier tracking.
- **Multiple Drivers**: Logs can be output to the CLI, file, Syslog, database, Elasticsearch, HTTP, Kafka, and Graylog.
- **Configurable**: Logging is configurable via a JSON config file, allowing easy setup without modifying the core code.
- **Extensible**: New drivers can be added easily without changing the core logging logic.

## Installation

1. Clone the repository or add it as a module to your project:
   
   ```bash
   git clone https://github.com/username/telemetry-logging.git
   ```

2. Initialize Go modules (if not done yet) by running:

   ```bash
   go mod init telemetry-logging
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

## Usage

### Configuration

Create a configuration file (e.g., `config/config.json`) to define log levels and drivers. The package supports several built-in drivers including `cli`, `file`, `syslog`, `db` (database), `elasticsearch`, `http`, `kafka`, and `graylog`.

```json
{
    "log_level": "DEBUG",
    "drivers": [
        {
            "type": "cli",
            "settings": {}
        },
        {
            "type": "file",
            "settings": {
                "file_path": "logs/app.log"
            }
        },
        {
            "type": "syslog",
            "settings": {
                "address": "localhost:514"
            }
        },
        {
            "type": "db",
            "settings": {
                "dsn": "user:password@tcp(localhost:3306)/logdb"
            }
        },
        {
            "type": "elasticsearch",
            "settings": {
                "index": "logs-index"
            }
        },
        {
            "type": "http",
            "settings": {
                "endpoint": "https://example.com/api/logs"
            }
        },
        {
            "type": "kafka",
            "settings": {
                "broker": "localhost:9092",
                "topic": "logs"
            }
        }
    ]
}
```

### Example Code

The following code demonstrates basic and transaction-based logging using the package:

```go
package main

import (
    "fmt"
    "telemetry-logging/config"
    "telemetry-logging/logger"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig("config/config.json")
    if err != nil {
        fmt.Println("Error loading config:", err)
        return
    }

    // Initialize the logger
    log, err := config.InitializeLogger(cfg)
    if err != nil {
        fmt.Println("Error initializing logger:", err)
        return
    }

    // Basic logging
    log.Debug("Debug message: Starting application", map[string]interface{}{
        "environment": "development",
    })
    log.Info("Info message: Application started", nil)
    log.Warn("Warning message: High memory usage", map[string]interface{}{
        "usage": "512MB",
    })
    log.Error("Error message: Application crashed", map[string]interface{}{
        "error_code": 1234,
    })

    // Transaction logging
    txn := logger.NewTransaction(map[string]interface{}{
        "CustomerID": "customer-123",
        "OrderID":    "order-456",
    })
    txnLogger := txn.LoggerWithTransaction(log)

    txnLogger.Info("Transaction started", nil)
    txnLogger.Debug("Processing payment", map[string]interface{}{
        "payment_method": "credit_card",
    })
    txnLogger.Error("Payment failed", map[string]interface{}{
        "error": "Insufficient funds",
    })
    txnLogger.Info("Transaction ended", nil)

    // Complete
    fmt.Println("Telemetry logging example completed.")
}
```

### Running the Example

1. **Create the configuration**: Place the `config/config.json` file in your project root.
2. **Run the Go program**:

   ```bash
   go run telemetry.go
   ```

3. The logs will be output to both the CLI and the file specified in the configuration (`logs/app.log`), as well as to the configured remote drivers (if applicable).

## Built-in Drivers

### 1. **CLI Driver**
Outputs logs to the command line interface (CLI) for real-time monitoring.

```json
{
    "type": "cli",
    "settings": {}
}
```

### 2. **File-based Driver**
Saves logs to a file in JSON format for future review.

```json
{
    "type": "file",
    "settings": {
        "file_path": "logs/app.log"
    }
}
```

### 3. **Syslog Driver**
Sends logs to a Syslog server, commonly used in Linux environments for centralized logging.

```json
{
    "type": "syslog",
    "settings": {
        "address": "localhost:514"
    }
}
```

### 4. **Database Driver**
Stores logs in a relational database such as MySQL or PostgreSQL for centralized storage and querying.

```json
{
    "type": "db",
    "settings": {
        "dsn": "user:password@tcp(localhost:3306)/logdb"
    }
}
```

### 5. **Elasticsearch Driver**
Sends logs to an Elasticsearch cluster for efficient storage, search, and analytics.

```json
{
    "type": "elasticsearch",
    "settings": {
        "index": "logs-index"
    }
}
```

### 6. **HTTP Driver**
Sends logs to a remote server via an HTTP POST request.

```json
{
    "type": "http",
    "settings": {
        "endpoint": "https://example.com/api/logs"
    }
}
```

### 7. **Kafka Driver**
Sends logs to an Apache Kafka topic for distributed, real-time logging and analytics.

```json
{
    "type": "kafka",
    "settings": {
        "broker": "localhost:9092",
        "topic": "logs"
    }
}
```

### 8. **Graylog Driver**
Sends logs to a Graylog instance over UDP for log aggregation and management.

```json
{
    "type": "graylog",
    "settings": {
        "address": "localhost:12201"
    }
}
```

## Project Structure

```
telemetry-logging/
├── config/
│   └── config.go              # Handles configuration loading
├── drivers/
│   ├── cli.go                 # CLI driver
│   ├── file.go                # File-based driver
│   ├── syslog.go              # Syslog driver
│   ├── db.go                  # Database driver
│   ├── elasticsearch.go       # Elasticsearch driver
│   ├── http.go                # HTTP/REST API driver
│   ├── kafka.go               # Kafka driver
├── logger/
│   ├── logger.go              # Main logger implementation
│   ├── log_level.go           # Log level definitions
│   └── transaction.go         # Transaction-based logging
├── telemetry.go               # Example usage of the logging package
├── telemetry_test.go          # Unit tests for the logging package
├── go.mod                     # Go module file
└── README.md                  # Project README file
```

## Extending the Package

### Adding a New Driver

To add a new logging driver (e.g., for a remote logging service), implement the `Driver` interface and register it using the `RegisterDriver` function in `logger/driver_registry.go`.

```go
type Driver interface {
    Log(entry LogEntry) error
}
```

#### Example of Adding a Remote Driver:

```go
package drivers

import (
    "telemetry-logging/logger"
    "fmt"
    "net/http"
    "bytes"
)

type RemoteDriver struct {
    Level    logger.LogLevel
    Endpoint string
}

func NewRemoteDriver(config map[string]interface{}) (logger.Driver, error) {
    endpoint := config["endpoint"].(string)
    return &RemoteDriver{
        Level:    logger.DEBUG,
        Endpoint: endpoint,
    }, nil
}

func (r *RemoteDriver) Log(entry logger.LogEntry) error {
    if entry.Level < r.Level {
        return nil
    }

    payload, err := json.Marshal(entry)
    if err != nil {
        return err
    }

    _, err = http.Post(r.Endpoint, "application/json", bytes.NewBuffer(payload))
    if err != nil {
        return err
    }
    return nil
}
```

## Running Unit Tests

Unit tests are included for different parts of the package. To run the tests, use:

```bash
go test -v
```

This will run the tests in `telemetry_test.go` and output the results.
