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

const (
	defaultPort = "8080"
)

var (
	port = flag.String("port", defaultPort, "Port of server")
	s3Endpoint = flag.String("s3-endpoint", "", "S3 Endpoint")
	s3Bucket = flag.String("s3-bucket", "", "S3 Endpoint")
	s3Region = flag.String("s3-region", "", "S3 Endpoint")
	s3AccessKey = flag.String("s3-access-key", "", "S3 Endpoint")
	s3SecretKey = flag.String("s3-secret-key", "", "S3 Endpoint")
)

// Command :
type Command interface {
	Execute() []string
}

// Run : run application
func (a *app) Run () error {
	flag.Parse()



	kingpin.Parse()

	err := setEnvars()
	if err != nil {
		log.Fatalln(err)
	}

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