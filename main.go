package main

import (
	"net/http"

	"github.com/bcokert/terragen/log"
	"github.com/bcokert/terragen/router"
)

func main() {
	log.Info("Starting Terragen on localhost:8080/...")

	r := router.CreateDefaultRouter()
	log.Fatal(http.ListenAndServe("localhost:8080/", r))
}
