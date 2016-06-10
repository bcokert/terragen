package controller_test

import (
	"testing"
	"time"

	"github.com/bcokert/terragen/controller"
)

func TestCreateDefaultServer(t *testing.T) {
	server := controller.CreateDefaultServer()

	if time.Now().Unix()-server.Seed > 10000000000 {
		t.Errorf("Expected server seed to be current time")
	}

	model := struct {
		Stuff     int
		MoreStuff string
	}{
		Stuff:     45,
		MoreStuff: "hello",
	}

	json, err := server.Marshal(model)

	if err != nil {
		t.Errorf("Expected server.Marshal to succeed in marshalling a model")
	}

	expectedEncoding := `{"Stuff":45,"MoreStuff":"hello"}`
	if string(json) != expectedEncoding {
		t.Errorf("Expected encoding %s, received %s", expectedEncoding, string(json))
	}
}
