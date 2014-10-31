package bitTorrent

import "net/http"

type BitTorrentPlugin struct {
}

func (p BitTorrentPlugin) Namespace() string {
	return "images.*.tracker"
}

func (p BitTorrentPlugin) Handle(r http.Request) (int, interface{}) {
	method := r.Method

	switch method {
	case "GET":
		return GetFile(&r)
	default:
		return 405, "Unsupported"
	}

}
