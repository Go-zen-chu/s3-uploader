package config

import (
	"flag"
	"github.com/go-zen-chu/s3-uploader/pkg/domain/server"
	"os"
)

const (
	defaultPort = "8080"
)

type configRepository struct {
}

func (c *configRepository) LoadFromEnv() *server.Config {
	return &server.Config{
		Port: os.Getenv("PORT"),
		S3Endpoint: os.Getenv("S3_ENDPOINT"),
		S3Bucket: os.Getenv("S3_BUCKET"),
		S3Region: os.Getenv("AWS_REGION"),
		S3AccessKey: os.Getenv("AWS_ACCESS_KEY"),
		S3SecretKey: os.Getenv("AWS_SECRET_KEY"),
	}
}

func (c *configRepository) LoadFromArgs() *server.Config {
	conf := c.LoadFromEnv()
	var (
		port = flag.String("port", defaultPort, "Port of server")
		s3Endpoint = flag.String("s3-endpoint", "", "S3 Endpoint")
		s3Bucket = flag.String("s3-bucket", "", "S3 Endpoint")
		s3Region = flag.String("s3-region", "", "S3 Endpoint")
		s3AccessKey = flag.String("s3-access-key", "", "S3 Endpoint")
		s3SecretKey = flag.String("s3-secret-key", "", "S3 Endpoint")
	)
}