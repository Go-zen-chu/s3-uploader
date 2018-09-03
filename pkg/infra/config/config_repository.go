package config

import (
	"github.com/go-zen-chu/s3-uploader/pkg/domain/server"
	"os"
)

type configRepository struct {

}

func (c *configRepository) LoadFromEnv() *server.Config {
	return &server.Config{
		S3Endpoint: os.Getenv("S3_ENDPOINT"),
		S3Bucket: os.Getenv("S3_BUCKET"),
		S3Region: os.Getenv("AWS_REGION"),
		S3AccessKey: os.Getenv("AWS_ACCESS_KEY"),
		S3SecretKey: os.Getenv("AWS_SECRET_KEY"),
	}
}
