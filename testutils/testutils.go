package testutils

import (
	"io"
	"net/http"
	"net/http/httptest"
)

// ExecuteTestRequest runs a request using the given router, and returns the response recorder
func ExecuteTestRequest(router http.Handler, method, url string, body io.Reader) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(method, url, body)

	router.ServeHTTP(response, request)

	return response
}
