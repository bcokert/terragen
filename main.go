package main

import (
	stdLog "log"

	"github.com/bcokert/terragen/log"
	"os"
	"github.com/julienschmidt/httprouter"
	"github.com/bcokert/terragen/http"
)

func main() {
	port := os.Getenv("TERRAGEN_PORT")
	if port == "" {
		stdLog.Fatal("No legal port was provided. Please set the TERRAGEN_PORT variable.")
	}

	assetsDir := os.Getenv("TERRAGEN_STATIC_ASSETS")
	if assetsDir == "" {
		stdLog.Fatal("No static assets directory was provided. Please set the TERRAGEN_STATIC_ASSETS variable.")
	}

	bundleHash := os.Getenv("TERRAGEN_JAVASCRIPT_BUNDLE")
	if bundleHash == "" {
		stdLog.Fatal("No bundle file hash was specified for server. Please set the TERRAGEN_JAVASCRIPT_BUNDLE variable")
	}

	router := httprouter.New()

	router.GET("/static/*path", http.HandleStatic(assetsDir))

	router.GET("/", http.TimedRequest(http.HandleIndex(bundleHash), "Index"))

	router.GET("/amiup", http.TimedRequest(http.HandleStatus(), "Amiup"))

	router.GET("/noise", http.TimedRequest(http.HandleNoise(), "Noise"))

	log.Info("Starting Terragen Service on port %s and asset directory %s", port, assetsDir)

	stdLog.Fatal(http.ListenAndServe(":"+port, router))
}
