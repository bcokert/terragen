package controller

// A Server is the instance that all handlers are implemented for
// The server instance allows you to pass dependencies to all controllers
type Server struct {
	Marshal func(interface{}) ([]byte, error)
}
