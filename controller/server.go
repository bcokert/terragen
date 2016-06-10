package controller

import (
	"encoding/json"
	"time"
)

// A Server is the instance that all handlers are implemented for
// The server instance allows you to pass dependencies to all controllers
type Server struct {
	Seed    int64
	Marshal func(interface{}) ([]byte, error)
}

// CreateDefaultServer creates the standard server
// It should always be used except in tests
func CreateDefaultServer() *Server {
	return &Server{
		Seed:    time.Now().Unix(),
		Marshal: json.Marshal,
	}
}
