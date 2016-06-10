package errors_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	stderrors "errors"

	"github.com/bcokert/terragen/errors"
)

func TestErrorWithCause(t *testing.T) {
	testCases := map[string]struct {
		Message  string
		Error    error
		Expected string
	}{
		"simple": {
			Message:  "the thing failed",
			Error:    stderrors.New("Another Thing"),
			Expected: "the thing failed: (Another Thing)",
		},
		"format strings in message": {
			Message:  "the %s thing %s failed %s",
			Error:    stderrors.New("Another Thing %s"),
			Expected: "the %s thing %s failed %s: (Another Thing %s)",
		},
	}

	for name, testCase := range testCases {
		if result := errors.ErrorWithCause(testCase.Message, testCase.Error).Error(); result != testCase.Expected {
			t.Errorf("%s failed. Expected %s, received %s", name, testCase.Expected, result)
		}
	}
}

func TestUnsupportedError(t *testing.T) {
	testCases := map[string]struct {
		Message  string
		Expected string
	}{
		"simple": {
			Message:  "doing things",
			Expected: "'doing things' is not currently supported",
		},
	}

	for name, testCase := range testCases {
		if result := errors.UnsupportedError(testCase.Message).Error(); result != testCase.Expected {
			t.Errorf("%s failed. Expected %s, received %s", name, testCase.Expected, result)
		}
	}
}

func TestWriteError(t *testing.T) {
	testCases := map[string]struct {
		Error              error
		ExpectedHttpStatus int
		ExpectedBody       string
	}{
		"simple 500": {
			Error:              stderrors.New("it dun broked"),
			ExpectedHttpStatus: http.StatusInternalServerError,
			ExpectedBody:       `{"error": "it dun broked"}`,
		},
		"escaped 404": {
			Error:              stderrors.New(`it dun \\"kk" \n +_9*& \t& @#broked`),
			ExpectedHttpStatus: http.StatusNotFound,
			ExpectedBody:       `{"error": "it dun \\"kk" \n +_9*& \t& @#broked"}`,
		},
	}

	for name, testCase := range testCases {
		response := httptest.NewRecorder()
		errors.WriteError(response, testCase.Error, testCase.ExpectedHttpStatus)

		if response.Code != testCase.ExpectedHttpStatus {
			t.Errorf("%s failed. Expected status code %d, received %d", name, testCase.ExpectedHttpStatus, response.Code)
		}
		if response.Body.String() != testCase.ExpectedBody {
			t.Errorf("%s failed. Expected body '%s', received '%s'", name, testCase.ExpectedBody, response.Body.String())
		}
	}
}
