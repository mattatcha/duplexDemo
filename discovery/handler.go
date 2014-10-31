package discovery

import "net/http"

func GetDiscovery(r *http.Request) (int, interface{}) {
	ips := []string{
		"10.10.10.1",
		"10.10.10.4",
		"10.10.10.5",
		"10.10.10.2",
	}
	return 200, ips
}
