package http

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func HandleStatus() httprouter.Handle {
	return Handle(func(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params) (interface{}, int) {
		return `{"amiup": true}`, http.StatusOK
	})
}
