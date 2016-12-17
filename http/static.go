package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HandleStatic handles service static files
// unlike all other handlers, this one does NOT produce a HandlerFunc, but a fully ready httprouter.Handle
// this is because http.ServeFile already manages the output and headers and codes
func HandleStatic(assetsDir string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		http.ServeFile(w, r, assetsDir+r.URL.Path)
	}
}
