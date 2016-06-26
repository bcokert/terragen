package errors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/bcokert/terragen/log"
)

// ErrorWithCause creates a new error that encodes the old error within it
func ErrorWithCause(msg string, err error) error {
	errorMsg := msg + fmt.Sprintf(": (%s)", err.Error())
	return errors.New(errorMsg)
}

// UnsupportedError returns standard error about the feature not being supported
func UnsupportedError(feature string) error {
	return fmt.Errorf("'%s' is not currently supported", feature)
}

// WriteError writes a standard error response to the response writer
func WriteError(response http.ResponseWriter, e error, httpStatus int) {
	log.Debug("Request Error (%d): %s", httpStatus, e.Error())
	response.WriteHeader(httpStatus)
	fmt.Fprintf(response, `{"error": "%s"}`, e.Error())
}
