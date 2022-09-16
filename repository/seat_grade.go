package repository

import (
	"gorm.io/gorm"
	"time"
)

type SeatGrade struct {
	ID        uint            `json:"id" gorm:"primaryKey"`
	Name      string          `json:"name"`
	Price     int             `json:"price"`
	Color     string          `json:"color"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"-" gorm:"index"`
}
