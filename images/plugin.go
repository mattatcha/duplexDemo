package images

import "net/http"

type ImagePlugin struct {
}

func (p ImagePlugin) Namespace() string {
	return "images"
}

func (p ImagePlugin) Handle(r http.Request) (int, interface{}) {
	method := r.Method

	switch method {
	case "GET":
		return GetImages(&r)
	default:
		return 405, "Unsupported"
	}

}

// func (p ImagePlugin) Endpoints() []types.Endpoint {
// 	return []types.Endpoint{
// 		types.Endpoint{
// 			Method:  "GET",
// 			Match:   "/",
// 			Handler: GetImages,
// 		},
// 	}
// }
