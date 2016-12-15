package http

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func HandleIndex(bundleHash string) httprouter.Handle {
	return Handle(func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) (interface{}, int) {
		html := `
		<!DOCTYPE html>
		<html><head><meta charset="utf-8"><title>Terragen</title></head>
		<body>
			<div id="app"></div>
			<script src="static/%s" type="text/javascript"></script>
		</body>
		</html>
		`
		w.Header().Add("Content-Type", "text/html")
		fmt.Fprintf(w, html, bundleHash+".js")
		return nil, http.StatusOK
	})
}
