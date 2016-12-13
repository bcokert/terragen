package router

import (
	"net/http"

	"github.com/bcokert/terragen/controller"
	"github.com/bcokert/terragen/endpoint"
	"github.com/bcokert/terragen/middleware"
	"github.com/julienschmidt/httprouter"
)

// CreateDefaultRouter returns a router with all the default routes configured
func CreateDefaultRouter(server *controller.Server, assetsDir string) *httprouter.Router {
	router := httprouter.New()

	router.GET("/static/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		http.ServeFile(w, r, assetsDir+r.URL.Path)
	})

	router.GET("/", middleware.TimedRequest(endpoint.Index, "Index"))

	router.GET("/amiup", middleware.TimedRequest(endpoint.Amiup, "Amiup"))

	router.GET("/noise", middleware.TimedRequest(server.Noise, "Noise"))

	return router
}
