package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/MattAitchison/duplexDemo/plugins/bitTorrent"
	"github.com/MattAitchison/duplexDemo/plugins/discovery"
	"github.com/MattAitchison/duplexDemo/types"

	dplx "github.com/progrium/duplex/prototype"
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
	rpcListen := getopt("rpc_listen", "127.0.0.1:9877")

	// Manually registering resources for now.
	resources := []types.Resource{
		bitTorrent.BitTorrentPlugin{},
		// images.ImagePlugin{},
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
	go http.ListenAndServe(listen, nil)

	client := dplx.NewPeer()

	if err := client.Bind(rpcListen); err != nil {
		log.Fatal(err)
	}

	// defer client.Close()
	reply := new(string)
	err := client.Call("ImagePlugin.Namespace", map[string]interface{}{"f": "a"}, reply)
	if err != nil {
		log.Fatal(err)

	}

	log.Println(*reply)

	// Startup HTTP server with a listening address.

}
