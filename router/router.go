package router

import (
	"net/http"

	"github.com/bcokert/terragen/endpoint"
)

// CreateDefaultRouter returns a router with all the default routes configured
func CreateDefaultRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/amiup", endpoint.Amiup)

	return router
}
