package main

import (
	"log"

	"github.com/MattAitchison/duplexDemo/plugins/images"
	dplx "github.com/progrium/duplex/prototype"
)

func main() {
	server := dplx.NewPeer()
	defer server.Close()

	if err := server.Connect("127.0.0.1:9877"); err != nil {
		log.Fatal(err)
	}

	server.Register(new(images.ImagePlugin))

	go server.Serve()

	// hack to keep server up
	select {}
}
