package endpoint

import (
	"fmt"
	"net/http"
	"os"
)

var _ http.HandlerFunc = Amiup

// Index returns the browser webpage that acts as a basic client of the rest of the service
func Index(response http.ResponseWriter, request *http.Request) {
	html := `
		<!DOCTYPE html>
		<html><head><meta charset="utf-8"><title>Terragen</title></head>
		<body>
			<div id="app"></div>
			<script src="static/%s" type="text/javascript"></script>
		</body>
		</html>
		`

	bundleHash := os.Getenv("TERRAGEN_JAVASCRIPT_BUNDLE")
	if bundleHash == "" {
		fmt.Fprint(response, "No bundle file hash was specified for server.")
		return
	}

	fmt.Fprintf(response, html, bundleHash + ".js")
}
