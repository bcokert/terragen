package http_test

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/bcokert/terragen/log"
	tghttp "github.com/bcokert/terragen/http"
	"github.com/julienschmidt/httprouter"
)

func TestTimedRequest(t *testing.T) {
	testCases := map[string]struct {
		Handler          httprouter.Handle
		ExpectedLogRegex string
	}{
		"3ms Handler": {
			Handler: func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
				time.Sleep(3000000)
			},
			ExpectedLogRegex: `INFO: MyHandler took [3-4]\.\d*ms`,
		},
		"Failing Handler": {
			Handler: func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

		router := httprouter.New()
		router.GET("/", tghttp.TimedRequest(testCase.Handler, "MyHandler"))
		router.ServeHTTP(recorder, request)

		output := log.FlushTestLogger()
		if matches, err := regexp.Match(testCase.ExpectedLogRegex, []byte(output)); !matches || err != nil {
			t.Errorf("%s failed. Expected %v, received %v", name, testCase.ExpectedLogRegex, output)
		}
	}
}
