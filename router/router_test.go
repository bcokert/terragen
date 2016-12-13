package router_test

import (
	"net/http"
	"testing"

	"encoding/json"
	"github.com/bcokert/terragen/controller"
	"github.com/bcokert/terragen/router"
	"time"
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
