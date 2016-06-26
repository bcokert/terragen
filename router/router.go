package router

import (
	"net/http"

	"github.com/bcokert/terragen/controller"
	"github.com/bcokert/terragen/endpoint"
	"github.com/bcokert/terragen/middleware"
)

// CreateDefaultRouter returns a router with all the default routes configured
func CreateDefaultRouter(server *controller.Server) http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/amiup", middleware.TimedRequest(endpoint.Amiup, "Amiup"))

	router.HandleFunc("/noise", middleware.TimedRequest(server.Noise, "Noise"))

	return router
}
