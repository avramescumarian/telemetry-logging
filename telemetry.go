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

	// Initialize the logger using the configuration
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

	// Start a transaction
	txn := logger.NewTransaction(map[string]interface{}{
		"CustomerID": "customer-123",
		"OrderID":    "order-456",
	})
	txnLogger := txn.LoggerWithTransaction(log)

	// Transaction logging
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
