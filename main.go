package main

import (
	"net/http"

	stdLog "log"

	"encoding/json"
	"github.com/bcokert/terragen/controller"
	"github.com/bcokert/terragen/log"
	"github.com/bcokert/terragen/router"
	"time"
	"os"
)

func main() {
	port := os.Getenv("TERRAGEN_PORT")
	if port == "" {
		stdLog.Fatal("No legal port was provided. Please set the TERRAGEN_PORT variable.")
	}

	assetsDir := os.Getenv("TERRAGEN_STATIC_ASSETS")
	if port == "" {
		stdLog.Fatal("No static assets directory was provided. Please set the TERRAGEN_STATIC_ASSETS variable.")
	}

	server := controller.Server{
		Seed:    time.Now().Unix(),
		Marshal: json.Marshal,
	}

	log.Info("Starting Terragen Service on port %s and asset directory %s", port, assetsDir)

	r := router.CreateDefaultRouter(&server, assetsDir)
	stdLog.Fatal(http.ListenAndServe(":" + port, r))
}
