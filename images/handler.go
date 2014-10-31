package images

import "net/http"

var Images []Image

func GetImages(r *http.Request) (int, interface{}) {
	images := []Image{
		Image{
			ID:   "762g3yh4n",
			Name: "ubuntu",
		},
		Image{
			ID:   "76ag3yh4n",
			Name: "fedora",
		},
		Image{
			ID:   "762f3yh4n",
			Name: "rethinkdb",
		},
		Image{
			ID:          "76wg3yh4n",
			Name:        "progrium/registrator",
			Description: "Service registry bridge for Docker",
		},
	}
	return 200, images
}
