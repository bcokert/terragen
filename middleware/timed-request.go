package middleware

import (
	"net/http"
	"time"

	"github.com/bcokert/terragen/timing"
)

func TimedRequest(handlerFunc http.HandlerFunc, name string) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		defer timing.Track(time.Now(), name)
		handlerFunc(response, request)
	}
}
