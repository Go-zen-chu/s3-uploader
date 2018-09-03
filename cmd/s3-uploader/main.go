package main

import (
	"github.com/go-zen-chu/s3-uploader/pkg/application"
	"log"
	"net/http"
	"os"

	"github.com/go-zen-chu/s3-uploader/pkg/web"
	"github.com/pkg/errors"
	"gopkg.in/alecthomas/kingpin.v2"
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
	app = application.NewApp()
	app.Run()

}
