package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-zen-chu/s3-uploader/pkg/web"
	"github.com/pkg/errors"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	port       = kingpin.Arg("port", "Port of server").Envar("PORT").String()
	s3Endpoint = kingpin.Arg("s3-endpoint", "S3 Endpoint").Envar("S3_ENDPOINT").String()
	s3Bucket   = kingpin.Arg("s3-bucket", "S3 Bucket").Envar("S3_BUCKET").String()
	// these envvars are loaded by aws-sdk-go
	// [session - Amazon Web Services - Go SDK](https://docs.aws.amazon.com/sdk-for-go/api/aws/session/)
	// Environment Variables
	s3Region    = kingpin.Arg("s3-region", "S3 Region").Envar("AWS_REGION").String()
	s3AccessKey = kingpin.Arg("s3-access-key", "S3 Access Key").Envar("AWS_ACCESS_KEY").String()
	s3SecretKey = kingpin.Arg("s3-secret-key", "S3 Secret Key").Envar("AWS_SECRET_KEY").String()
)

func setEnvars() error {
	if s3Region != nil {
		err := os.Setenv("AWS_REGION", *s3Region)
		return errors.Wrap(err, "Error setting AWS_REGION")
	}
	if s3AccessKey != nil {
		err := os.Setenv("AWS_ACCESS_KEY", *s3AccessKey)
		return errors.Wrap(err, "Error setting AWS_ACCESS_KEY")
	}
	if s3SecretKey != nil {
		err := os.Setenv("AWS_SECRET_KEY", *s3SecretKey)
		return errors.Wrap(err, "Error setting AWS_SECRET_KEY")
	}
	return nil
}

func main() {
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
