package router_test

import (
	"net/http"
	"testing"

	"github.com/bcokert/terragen/controller"
	"github.com/bcokert/terragen/router"
	"time"
	"encoding/json"
)

func TestCreateDefaultRouter(t *testing.T) {
	var r http.Handler = router.CreateDefaultRouter(&controller.Server{
		Seed:    time.Now().Unix(),
		Marshal: json.Marshal,
	}, ".")

	if r == nil {
		t.Errorf("Did not create a router")
	}
}
