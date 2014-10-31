package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/MattAitchison/duplexDemo/bitTorrent"
	"github.com/MattAitchison/duplexDemo/discovery"
	"github.com/MattAitchison/duplexDemo/images"
	"github.com/MattAitchison/duplexDemo/types"
)

func getopt(name, dfault string) string {
	value := os.Getenv(name)
	if value == "" {
		value = dfault
	}
	return value
}
func main() {
	// ADDRESS:PORT the server should listen on.
	listen := getopt("listen", ":3000")

	// Manually registering resources for now.
	resources := []types.Resource{
		bitTorrent.BitTorrentPlugin{},
		images.ImagePlugin{},
		discovery.DiscoveryPlugin{},
	}

	// Register a default handler which calls plugin functions.
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		path := req.URL.Path

		for _, resource := range resources {
			pattern := resource.Namespace()
			matched, err := regexp.MatchString(pattern, path)
			if err != nil {
				fmt.Println(err)
			}
			if matched {
				code, res := resource.Handle(*req)
				rw.WriteHeader(code)
				json.NewEncoder(rw).Encode(res)

				return
			}
		}
		// We didn't match a resource so let's 404!

		rw.WriteHeader(404)
		json.NewEncoder(rw).Encode("Not found")
		return

	})

	// Startup HTTP server with a listening address.
	http.ListenAndServe(listen, nil)
}
