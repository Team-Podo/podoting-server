package models

import (
	"github.com/Team-Podo/podoting-server/repository"
)

type FileRepository interface {
	Save(file *repository.File) error
}
