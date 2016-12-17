package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// ListendAndServe starts the server on the specified port with the specified handler
func ListenAndServe(addr string, handler http.Handler) error {
	return http.ListenAndServe(addr, handler)
}

// A handlerFunc is exported for each endoint, and it always returns the http code of the result
// to return no response, or if the handler managed the response itself, a nil json should be sent
// if json is an error, then a specific json error will be written
// if a string is returned, it is assumed to be pre-marshalled json
// otherwise, the json will be marshaled and written
type HandlerFunc func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) (interface{}, int)

// Handle converts a handlerFunc into a httprouter.Handle, so that it can be easily used with the httprouter.Router
func Handle(h HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		response, code := h(w, r, p)
		output, finalCode := marshalOutput(response, code)

		w.Header().Add("Access-Control-Allow-Origin", "*") //TODO: Find out if this is necessary for other domains than terragen.brandonokert.com
		if len(output) > 0 {
			w.Header().Add("Content-Type", "application/json")
		}
		if finalCode != http.StatusOK {
			w.WriteHeader(finalCode)
		}
		fmt.Fprintf(w, output)
	}
}

// Converts the given object into a writable string with the correct code
// eg: if the given response is not actually marshalable, the code may change, and the output will be a valid error response
func marshalOutput(response interface{}, code int) (string, int) {
	if response == nil {
		return "", code
	}

	if str, ok := response.(string); ok {
		return str, code
	}

	if err, ok := response.(error); ok {
		return fmt.Sprintf(`{"error": "%s"}`, err.Error()), code
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		return fmt.Sprintf(`{"error": "An error occurred while marshaling a response: %s"}`, err.Error()), http.StatusInternalServerError
	}

	return string(bytes), code
}
