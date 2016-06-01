package endpoint

import (
	"fmt"
	"net/http"
)

var _ http.HandlerFunc = Amiup

// Amiup returns a 200 with a basic json response if the server is alive
func Amiup(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, `{"amiup": true}`)
}
