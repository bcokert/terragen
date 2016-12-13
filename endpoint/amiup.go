package endpoint

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Amiup returns a 200 with a basic json response if the server is alive
func Amiup(response http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprintf(response, `{"amiup": true}`)
}
