package log

import (
	"bytes"
	"log"
)

var testBuffer = &bytes.Buffer{}
var testLogger = log.New(testBuffer, "", 0)

// UseTestLogger replaces the current logger with one that logs to a buffer.
// The data can retrieved via FlushTestLogger
func UseTestLogger() {
	logger = testLogger
}

// FlushTestLogger empties the test logger and returns whatever has been logged to it since calling this function
func FlushTestLogger() string {
	output := testBuffer.String()
	testBuffer.Reset()
	return output
}
