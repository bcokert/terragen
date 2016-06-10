package testutils_test

import (
	"fmt"
	"io"
	"math"
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

func TestFloatSlice(t *testing.T) {
	testCases := map[string]struct {
		From       float64
		To         float64
		Resolution int
		Expected   []float64
	}{
		"basic": {
			From:       1,
			To:         6,
			Resolution: 1,
			Expected:   []float64{1, 2, 3, 4, 5},
		},
		"higher resolution": {
			From:       1,
			To:         2,
			Resolution: 5,
			Expected:   []float64{1, 1.2, 1.4, 1.6, 1.8},
		},
	}

	for name, testCase := range testCases {
		result := testutils.FloatSlice(testCase.From, testCase.To, testCase.Resolution)

		if len(result) != len(testCase.Expected) {
			t.Logf("%v", result)
			t.Errorf("%s failed. Expected slice with length %d, found length %d", name, len(testCase.Expected), len(result))
			continue
		}

		for i, val := range result {
			if math.Abs(val-testCase.Expected[i]) > 0.000000000001 {
				t.Errorf("%s failed. Expected %d to be %v, received %v", name, i, testCase.Expected[i], val)
			}
		}
	}
}
