package models

import (
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"io"
)

type File interface {
	UploadFile(file io.Reader, path string) (*manager.UploadOutput, error)
}

type FileRepository interface {
	Save(file *repository.File) error
}
