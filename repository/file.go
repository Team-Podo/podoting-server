package repository

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type File struct {
	ID        uint `json:"id" gorm:"primarykey"`
	Title     string
	Size      int64
	Path      string
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" gorm:"index"`
}

type FileRepository struct {
	DB *gorm.DB
}

func (file *File) FullPath() {
	file.Path = fmt.Sprintf("%s/%s", file.Path, file.Title)
}

func (f *FileRepository) Get(id uint) (*File, error) {
	var file File
	if err := f.DB.First(&file, id).Error; err != nil {
		return nil, err
	}

	return &file, nil
}

func (f *FileRepository) Save(file *File) error {
	if err := f.DB.Create(file).Error; err != nil {
		return err
	}

	return nil
}
