package images

import "net/http"

var Images []Image

func GetImages(r *http.Request) (int, interface{}) {
	images := []Image{
		Image{
			Name: "Test",
		},
	}
	return 200, images
}
