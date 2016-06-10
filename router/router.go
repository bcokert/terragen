package router

import (
	"net/http"

	"github.com/bcokert/terragen/controller"
	"github.com/bcokert/terragen/endpoint"
)

// CreateDefaultRouter returns a router with all the default routes configured
func CreateDefaultRouter(server *controller.Server) http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/amiup", endpoint.Amiup)

	router.HandleFunc("/noise", server.Noise)

	return router
}
