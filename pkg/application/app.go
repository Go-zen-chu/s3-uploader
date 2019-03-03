package application

import (
	"flag"
	"github.com/go-zen-chu/s3-uploader/pkg/web"
	"github.com/pkg/errors"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"net/http"
)

type App interface {
	Run () error
}

type app struct {
}

func NewApp () App {
	return &app {}
}


// Run : run application
func (a *app) Run () error {
	var (
		port = flag.String("port", "8080", "port for server")
		s3AccessKey = flag.String("s3-access-key", "8080", "port for server")
	)

	http.HandleFunc("/", web.IndexHandler)
	http.HandleFunc("/upload", web.FileUploadHandler)
	http.HandleFunc("/uploaded", web.ShowUploadedHandler)

	var portStr = "8080"
	if port != nil {
		portStr = *port
	}
	err = http.ListenAndServe(":"+portStr, nil)
	if err != nil {
		log.Fatal(errors.Wrap(err, "Error while running server"))
	}
}