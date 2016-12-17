package http_test

import (
	"net/http/httptest"
	"testing"

	"net/http"

	tghttp "github.com/bcokert/terragen/http"
)

func TestHandleStatus(t *testing.T) {
	testCases := map[string]struct {
		ExpectedCode     int
		ExpectedResponse string
	}{
		"basic": {
			ExpectedCode:     http.StatusOK,
			ExpectedResponse: `{"amiup": true}`,
		},
	}

	for name, testCase := range testCases {
		w := httptest.NewRecorder()
		tghttp.HandleStatus()(w, nil, nil)

		if code := w.Code; code != testCase.ExpectedCode {
			t.Errorf("%s failed. Expected code %v, received %v", name, testCase.ExpectedCode, code)
		}

		if body := w.Body.String(); body != testCase.ExpectedResponse {
			t.Errorf("%s failed. Expected code %v, received %v", name, testCase.ExpectedResponse, body)
		}
	}
}
