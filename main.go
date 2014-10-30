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
	listen := getopt("listen", ":3000")

	// Manually registering resources for now.
	resources := []types.Resource{
		images.ImagePlugin{},
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		path := req.URL.Path
		parts := strings.Split(path, "/")
		resource := matchResource(parts[1], resources)

		if resource == nil {
			rw.WriteHeader(404)
			json.NewEncoder(rw).Encode("Not found")
			return
		}

		code, res := resource.Handle(*req)

		rw.WriteHeader(code)
		json.NewEncoder(rw).Encode(res)
	})

	http.ListenAndServe(listen, nil)
}
