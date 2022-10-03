package repository

import (
	"gorm.io/gorm"
	"time"
)

type Person struct {
	ID        uint `json:"id" gorm:"primarykey"`
	Name      string
	Birth     time.Time       `json:"birth"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"-" gorm:"index"`
}
