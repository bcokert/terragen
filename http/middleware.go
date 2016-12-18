package http

import (
	"net/http"
	"time"

	"github.com/bcokert/terragen/log"
	"github.com/julienschmidt/httprouter"
)

func TimedRequest(handlerFunc httprouter.Handle, name string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer func(start time.Time, name string) {
			log.Info("%s took %s", name, time.Since(start))
		}(time.Now(), name)
		handlerFunc(w, r, p)
	}
}
