package types

import "net/http"

// Resource an API resource.
type Resource interface {
	// Namespace returns an identifier to use to route requests to the correct
	// endpoint.
	Namespace() string

	// Handle is a method that is implimented to handle request routing within
	// a plugin.
	Handle(r http.Request) (int, interface{})
}
