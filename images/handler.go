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
	// json.NewEncoder(w).Encode(images)
}

func DeleteImage(w http.ResponseWriter, r *http.Request) {

}

func PutImage(w http.ResponseWriter, r *http.Request) {

}
func PostImage(w http.ResponseWriter, r *http.Request) {

}
