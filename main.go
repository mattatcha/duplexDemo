package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

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

// Call namespace on each resource until we match the path.
// If we can't find anything then return nil.
func matchResource(part string, resources []types.Resource) types.Resource {
	for _, resource := range resources {
		if resource.Namespace() == part {
			return resource
		}
	}
	return nil
}

func main() {
	// ADDRESS:PORT the server should listen on.
	listen := getopt("listen", ":3000")

	// Manually registering resources for now.
	resources := []types.Resource{
		images.ImagePlugin{},
	}

	// Register a default handler which calls plugin functions.
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		path := req.URL.Path
		parts := strings.Split(path, "/")
		resource := matchResource(parts[1], resources)

		// We didn't match a resource so let's 404!
		if resource == nil {
			rw.WriteHeader(404)
			json.NewEncoder(rw).Encode("Not found")
			return
		}

		// Call handler for resource.
		code, res := resource.Handle(*req)

		// Write status code and then write response out as JSON.
		rw.WriteHeader(code)
		json.NewEncoder(rw).Encode(res)
	})

	// Startup HTTP server with a listening address.
	http.ListenAndServe(listen, nil)
}
