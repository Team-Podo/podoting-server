package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"io"
	"log"
	"os"
)

type S3Config struct {
	bucket string
	key    string
	secret string
	region string
	Client *s3.Client
}

var S3 *S3Config

func Init() {
	S3 = newS3Config()
	S3.setS3Client()
}

func newS3Config() *S3Config {
	return &S3Config{
		bucket: os.Getenv("S3_BUCKET"),
		key:    os.Getenv("S3_KEY"),
		secret: os.Getenv("S3_SECRET"),
		region: os.Getenv("S3_REGION"),
	}
}

func (s *S3Config) setS3Client() {
	credential := credentials.NewStaticCredentialsProvider(s.key, s.secret, "")
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(credential),
		config.WithRegion(s.region))

	if err != nil {
		log.Fatal(err)
	}

	s.Client = s3.NewFromConfig(cfg)
}

func (s *S3Config) UploadFile(file io.Reader, path string, fileExtension string) (*manager.UploadOutput, error) {
	uploader := manager.NewUploader(s.Client)
	newUUID := uuid.New()

	s3Object := &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String("uploads" + path + "/" + newUUID.String() + fileExtension),
		Body:   file,
	}

	result, err := uploader.Upload(context.TODO(), s3Object)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return result, nil
}
