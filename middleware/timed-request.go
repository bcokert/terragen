package middleware

import (
	"net/http"
	"time"

	"github.com/bcokert/terragen/timing"
	"github.com/julienschmidt/httprouter"
)

func TimedRequest(handlerFunc httprouter.Handle, name string) httprouter.Handle {
	return func(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
		defer timing.Track(time.Now(), name)
		handlerFunc(response, request, params)
	}
}
