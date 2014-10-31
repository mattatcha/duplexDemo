package discovery

import "net/http"

type DiscoveryPlugin struct {
}

func (p DiscoveryPlugin) Namespace() string {
	return "discovery"
}

func (p DiscoveryPlugin) Handle(r http.Request) (int, interface{}) {
	method := r.Method

	switch method {
	case "GET":
		return GetDiscovery(&r)
	default:
		return 405, "Unsupported"
	}

}
