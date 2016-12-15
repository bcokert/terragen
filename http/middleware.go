package http

import (
	"github.com/julienschmidt/httprouter"
	"github.com/bcokert/terragen/timing"
	"time"
	"net/http"
)

func TimedRequest(handlerFunc httprouter.Handle, name string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer timing.Track(time.Now(), name)
		handlerFunc(w, r, p)
	}
}
