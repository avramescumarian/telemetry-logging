
# Telemetry Logging Package

## Overview
The Telemetry Logging Package is designed to simplify logging across multiple platforms and provide a unified, extensible interface for logging in Go applications. It supports multiple drivers, transaction-style logs, log levels, and metadata for enhanced logging.

### Features
- **Multiple Log Levels**: Supports `DEBUG`, `INFO`, `WARN`, and `ERROR` log levels.
- **Transaction Logging**: Group related log entries with a unique `TraceID` for easier tracking.
- **Multiple Drivers**: Logs can be output to the CLI or a file (JSON format).
- **Configurable**: Logging is configurable via a JSON config file, allowing easy setup without modifying the core code.
- **Extensible**: New drivers can be added easily without changing the core logging logic.

## Installation

1. Clone the repository or add it as a module to your project.
   
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

Create a configuration file (e.g., `config/config.json`) to define log levels and drivers.

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
        }
    ]
}
```

- **log_level**: Sets the minimum log level for logging (`DEBUG`, `INFO`, `WARN`, `ERROR`).
- **drivers**: Specifies where the logs will be output. This example outputs logs to both the CLI and a file.

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

3. The logs will be output to both the CLI and the file specified in the configuration (`logs/app.log`).

## Project Structure

```
telemetry-logging/
├── config/
│   └── config.go              # Handles configuration loading
├── drivers/
│   ├── cli.go                 # CLI driver for logging
│   └── file.go                # File-based driver for logging
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

To add a new logging driver (e.g., for a remote logging service), implement the `Driver` interface and add it to the `MultiLogger` using `AddDriver`.

```go
type Driver interface {
    Log(entry LogEntry) error
}
```

### Example of Adding a Remote Driver:

```go
package drivers

import (
    "telemetry-logging/logger"
)

type RemoteDriver struct {
    Level logger.LogLevel
}

func NewRemoteDriver(level logger.LogLevel) *RemoteDriver {
    return &RemoteDriver{Level: level}
}

func (r *RemoteDriver) Log(entry logger.LogEntry) error {
    if entry.Level < r.Level {
        return nil
    }
    // Code to send log entry to a remote server
    return nil
}
```

## Running Unit Tests

Unit tests are included for different parts of the package. To run the tests, use:

```bash
go test -v
```

This will run the tests in `telemetry_test.go` and output the results.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
