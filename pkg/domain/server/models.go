package server

import "os"

// models are structs of entities

type Config struct {
	S3Endpoint string
	S3Bucket string
	S3Region string
	S3AccessKey string
	S3SecretKey string
}

func (c *Config) SetToEnvVar () error {
	// these env vars are used by aws-sdk-go
	// https://docs.aws.amazon.com/sdk-for-go/api/aws/session/
	if err := os.Setenv("AWS_REGION", c.S3Region); err != nil {
		return err
	}
	if err := os.Setenv("AWS_ACCESS_KEY", c.S3AccessKey); err != nil {
		return err
	}
	if err := os.Setenv("AWS_SECRET_KEY", c.S3SecretKey); err != nil {
		return err
	}
}

