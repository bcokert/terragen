package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/bcokert/terragen/log"
	"github.com/bcokert/terragen/middleware"
)

func TestTimedRequest(t *testing.T) {
	testCases := map[string]struct {
		Handler          http.HandlerFunc
		ExpectedLogRegex string
	}{
		"3ms Handler": {
			Handler: func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(3000000)
			},
			ExpectedLogRegex: `INFO: MyHandler took [3-4]\.\d*ms`,
		},
		"Failing Handler": {
			Handler: func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(1000000)
				w.WriteHeader(500)
			},
			ExpectedLogRegex: `INFO: MyHandler took [1-2]\.\d*ms`,
		},
	}

	log.UseTestLogger()

	for name, testCase := range testCases {
		log.FlushTestLogger()
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		recorder := httptest.NewRecorder()

		router := http.NewServeMux()
		router.HandleFunc("/", middleware.TimedRequest(testCase.Handler, "MyHandler"))
		router.ServeHTTP(recorder, request)

		output := log.FlushTestLogger()
		if matches, err := regexp.Match(testCase.ExpectedLogRegex, []byte(output)); !matches || err != nil {
			t.Errorf("%s failed. Expected %v, received %v", name, testCase.ExpectedLogRegex, output)
		}
	}
}
