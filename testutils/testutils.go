package testutils

import (
	"io"
	"net/http"
	"net/http/httptest"
	"math"
)

// ExecuteTestRequest runs a request using the given router, and returns the response recorder
func ExecuteTestRequest(router http.Handler, method, url string, body io.Reader) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(method, url, body)

	router.ServeHTTP(response, request)

	return response
}

// IsFloatEqual compares two floats to see if they're close enough to be considered equal
func IsFloatEqual(a, b float64) bool {
	return math.Abs(a - b) <= 0.00000000000001
}
