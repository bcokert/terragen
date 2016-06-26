package log

import (
	"log"
	"os"
)

var logger = log.New(os.Stderr, "", log.LstdFlags)

// Debug logs a debug message
func Debug(format string, args ...interface{}) {
	logger.Printf("DEBUG: "+format, args...)
}

// Info logs an info message
func Info(format string, args ...interface{}) {
	logger.Printf("INFO: "+format, args...)
}

// Error logs an error message
func Error(format string, args ...interface{}) {
	logger.Printf("ERROR: "+format, args...)
}
