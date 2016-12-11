package endpoint_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bcokert/terragen/controller"
	"github.com/bcokert/terragen/router"
)

func TestAmiup(t *testing.T) {
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
		request, _ := http.NewRequest(http.MethodGet, "/amiup", nil)
		recorder := httptest.NewRecorder()

		router := router.CreateDefaultRouter(&controller.Server{}, ".")
		router.ServeHTTP(recorder, request)

		if code := recorder.Code; code != testCase.ExpectedCode {
			t.Errorf("%s failed. Expected code %v, received %v", name, testCase.ExpectedCode, code)
		}

		if body := recorder.Body.String(); body != testCase.ExpectedResponse {
			t.Errorf("%s failed. Expected code %v, received %v", name, testCase.ExpectedResponse, body)
		}
	}
}
