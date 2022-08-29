package repository

import (
	"gorm.io/gorm"
	"time"
)

type Person struct {
	ID        uint `json:"id" gorm:"primarykey"`
	Name      string
	age       int
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"-" gorm:"index"`
}
