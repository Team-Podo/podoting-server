package repository

import (
	"gorm.io/gorm"
	"time"
)

type Location struct {
	ID        uint    `json:"id" gorm:"primarykey"`
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt `json:"-" gorm:"index"`
}
