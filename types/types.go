package types

import "net/http"

//
// type APIPlugin struct {
// 	Register func()
// 	Get      func()
// 	Put      func()
// 	Delete   func()
// 	Post     func()
// }

// type Endpoint struct {
// 	Method  string
// 	Match   string
// 	Handler http.HandlerFunc
// }

type Resource interface {
	Namespace() string
	// Endpoints() []Endpoint
	Handle(r http.Request) (int, interface{})
}
