package images

import "net/http"

type ImagePlugin struct {
}
type Reply struct {
	Code int
	Data interface{}
}

// func (p ImagePlugin) Namespace() string {
func (p *ImagePlugin) Namespace(args interface{}, reply *string) error {
	*reply = "images"
	return nil
}

// func (p ImagePlugin) Handle(r http.Request) (int, interface{}) {
func (p *ImagePlugin) Handle(r http.Request, reply *Reply) error {
	method := r.Method

	switch method {
	case "GET":
		reply.Code, reply.Data = GetImages(&r)
		// return GetImages(&r)
		return nil
	default:
		reply.Code, reply.Data = 405, "Unsupported"
		return nil
	}

}
