package testutils_test

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/bcokert/terragen/testutils"
)

func TestExecuteTestRequest(t *testing.T) {
	testCases := map[string]struct {
		Method       string
		Url          string
		Body         io.Reader
		Router       *http.ServeMux
		HandlerFunc  http.HandlerFunc
		ExpectedBody string
		ExpectedCode int
	}{
		"basic, no body": {
			Method: http.MethodGet,
			Url:    "/tttest",
			Body:   nil,
			Router: http.NewServeMux(),
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{"good": true}`)
			},
			ExpectedBody: `{"good": true}`,
			ExpectedCode: http.StatusForbidden,
		},
	}

	for name, testCase := range testCases {
		testCase.Router.HandleFunc(testCase.Url, testCase.HandlerFunc)
		recorder := testutils.ExecuteTestRequest(testCase.Router, testCase.Method, testCase.Url, testCase.Body)

		if recorder.Body.String() != testCase.ExpectedBody {
			t.Errorf("%s failed. Expected body '%s', received '%s'", name, testCase.ExpectedBody, recorder.Body.String())
		}

		if recorder.Code != testCase.ExpectedCode {
			t.Errorf("%s failed. Expected code %d, received %d", name, testCase.ExpectedCode, recorder.Code)
		}
	}
}

func TestIsFloatEqual(t *testing.T) {
	testCases := map[string]struct {
		A, B     float64
		Expected bool
	}{
		"zero": {
			A:        0,
			B:        0,
			Expected: true,
		},
		"large": {
			A:        623462362.23452345,
			B:        623462362.23452345,
			Expected: true,
		},
		"almost": {
			A:        623462362.2345234,
			B:        623462362.2345235,
			Expected: false,
		},
	}

	for name, testCase := range testCases {
		result := testutils.IsFloatEqual(testCase.A, testCase.B)
		if result != testCase.Expected {
			t.Errorf("'%s' failed. Expected %v, received %v", name, testCase.Expected, result)
		}
	}
}
