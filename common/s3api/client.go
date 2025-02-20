package s3api

import (
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/celer-network/goutils/log"
)

//go:generate mockery --name IClient
type IClient interface {
	ListObjects(prefix string) ([]*s3.Object, error)
	Upload(key string, body io.Reader) error
	Download(key string, w io.WriterAt) error
	Delete(key string) error
}

type Client struct {
	s3         *s3.S3
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader

	bucket string
}

func NewClient(bucket string) (*Client, error) {
	var cre *credentials.Credentials

	// from aws
	id := os.Getenv("AWS_ACCESS_KEY_ID")
	if id == "" {
		id = os.Getenv("AWS_ACCESS_KEY")
	}
	secret := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if secret == "" {
		secret = os.Getenv("AWS_SECRET_KEY")
	}

	if id != "" || secret != "" {
		log.Infof("find ak sk in env, use it")
		cre = credentials.NewCredentials(&credentials.EnvProvider{})
	} else {
		log.Infof("do not find ak sk in env, use local .aws cre or ec2 role")
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: cre,
	})
	if err != nil {
		return nil, err
	}
	return &Client{
		s3:         s3.New(sess),
		uploader:   s3manager.NewUploader(sess),
		downloader: s3manager.NewDownloader(sess),
		bucket:     bucket,
	}, nil
}

func (s *Client) ListObjects(prefix string) ([]*s3.Object, error) {
	res, err := s.s3.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(s.bucket),
		Prefix: aws.String(prefix),
	})
	if err != nil {
		return nil, err
	}
	return res.Contents, nil
}

func (s *Client) Upload(key string, body io.Reader) error {
	log.Debugf("uploading file: bucket %s, key %s", s.bucket, key)
	_, err := s.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
		Body:   body,
	})
	return err
}

func (s *Client) HeadObject(filePath string) error {
	// Create a HeadObject request to check if the file exists
	params := &s3.HeadObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(filePath), // file path
	}
	s.s3.HeadObject(params)
	_, err := s.s3.HeadObject(params)
	return err
}

func (s *Client) Download(key string, w io.WriterAt) error {
	log.Debugf("downloading file: bucket %s, key %s", s.bucket, key)
	_, err := s.downloader.Download(w, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	return err
}

func (s *Client) Delete(key string) error {
	log.Debugf("downloading file: bucket %s, key %s", s.bucket, key)
	_, err := s.s3.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	return err
}
