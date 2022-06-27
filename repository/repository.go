package repository

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint            `gorm:"primarykey"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
	DeletedAt *gorm.DeletedAt `json:"-" gorm:"index"`
}
