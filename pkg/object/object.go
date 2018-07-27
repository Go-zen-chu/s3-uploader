package object

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func PutFileToS3(s *session.Session, bucket string, acl string, file *os.File) error {
	pi := &s3.PutObjectInput{
		Bucket:               aws.String(bucket),
		Key:                  aws.String(file.Name()),
		ACL:                  aws.String(acl),
		Body:                 file,
		ServerSideEncryption: aws.String("AES256"),
	}
	_, err := s3.New(s).PutObject(pi)
	return err
}
