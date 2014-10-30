package types

import "net/http"

type Resource interface {
	Namespace() string
	Handle(r http.Request) (int, interface{})
}
