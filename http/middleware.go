package http

import (
	"net/http"
	"time"

	"github.com/bcokert/terragen/timing"
	"github.com/julienschmidt/httprouter"
)

func TimedRequest(handlerFunc httprouter.Handle, name string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer timing.Track(time.Now(), name)
		handlerFunc(w, r, p)
	}
}
