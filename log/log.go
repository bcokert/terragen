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

// Fatal logs a fatal message, and then exits via os.Exit(1)
func Fatal(e error) {
	logger.Fatal("FATAL: " + e.Error())
}
