package endpoint

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/bcokert/terragen/log"
)

var _ http.HandlerFunc = Amiup

// Index returns the browser webpage that acts as a basic client of the rest of the service
func Index(response http.ResponseWriter, request *http.Request) {
	var files []os.FileInfo
	var err error

	html := `
		<!DOCTYPE html>
		<html><head><meta charset="utf-8"><title>Terragen</title></head>
		<body>
			<div id="app"></div>
			<script src="static/%s" type="text/javascript"></script>
		</body>
		</html>
		`

	if files, err = ioutil.ReadDir("./build/static/"); err != nil {
		fmt.Fprintf(response, "Error while searching for bundle file: %s", err.Error())
		return
	}

	for _, f := range files {
		log.Info("Found file: %s", f.Name())
		if strings.HasSuffix(f.Name(), ".js") {
			fmt.Fprintf(response, html, f.Name())
			return
		}
	}

	fmt.Fprint(response, "Unable to find bundle file. Time to panic.")
}
