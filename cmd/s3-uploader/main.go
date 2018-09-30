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

func main() {
	app = application.NewApp()
	err := app.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
