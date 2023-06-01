package _1_simple

import (
	"fmt"
)

type OldLogger interface {
	Error(message string, args ...interface{})
}

func LegacyFunction(logger OldLogger) {
	// implementation removed
}

type NewLogger interface {
	Error(message string, tags ...Tag)
}

type Tag struct {
	Key, Value string
}

type oldLoggerAdapter struct {
	newLogger NewLogger
}

// implement OldLogger interface
func (o *oldLoggerAdapter) Error(message string, args ...interface{}) {
	// adapt from OldLogger requests to NewLogger format
	o.newLogger.Error(fmt.Sprintf(message, args...))
}

// Confirm the relationship between the adapter and the new interface.

func Usage() {
	var newLogger NewLogger = &newLoggerImpl{}
	adaptedNewLogger := &oldLoggerAdapter{newLogger: newLogger}

	LegacyFunction(adaptedNewLogger)
}

type newLoggerImpl struct{}

func (n *newLoggerImpl) Error(message string, tags ...Tag) {
	// implementation removed
}
