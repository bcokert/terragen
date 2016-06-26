package log

import (
	"log"
	"os"
)

var logger = log.New(os.Stderr, "", log.LstdFlags)

// Info logs an info message
func Info(format string, args ...interface{}) {
	logger.Printf("I: "+format, args...)
}

// Error logs an error message
func Error(format string, args ...interface{}) {
	logger.Printf("ERROR: "+format, args...)
}
