package main

import (
	"net/http"

	stdLog "log"

	"encoding/json"
	"github.com/bcokert/terragen/controller"
	"github.com/bcokert/terragen/log"
	"github.com/bcokert/terragen/router"
	"time"
)

func main() {
	log.Info("Starting Terragen on localhost:8080/...")

	server := controller.Server{
		Seed:    time.Now().Unix(),
		Marshal: json.Marshal,
	}
	r := router.CreateDefaultRouter(&server)
	stdLog.Fatal(http.ListenAndServe("localhost:8080", r))
}
