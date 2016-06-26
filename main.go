package main

import (
	"net/http"

	stdLog "log"

	"github.com/bcokert/terragen/controller"
	"github.com/bcokert/terragen/log"
	"github.com/bcokert/terragen/router"
)

func main() {
	log.Info("Starting Terragen on localhost:8080/...")

	server := controller.CreateDefaultServer()
	r := router.CreateDefaultRouter(server)
	stdLog.Fatal(http.ListenAndServe("localhost:8080/", r))
}
