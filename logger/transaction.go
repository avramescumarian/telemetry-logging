package logger

import (
	"github.com/google/uuid"
)

// Transaction represents a set of related logs grouped by a TraceID.
type Transaction struct {
	TraceID    string
	Attributes map[string]interface{}
}

// NewTransaction creates a new transaction with a unique TraceID.
func NewTransaction(attributes map[string]interface{}) *Transaction {
	return &Transaction{
		TraceID:    uuid.New().String(),
		Attributes: attributes,
	}
}

// LoggerWithTransaction returns a Logger bound to the transaction's TraceID.
func (t *Transaction) LoggerWithTransaction(base Logger) Logger {
	return base.WithTrace(t.TraceID)
}
