package repository

import (
	"gorm.io/gorm"
	"time"
)

type Receipt struct {
	ID        uint            `json:"id" gorm:"primary_key"`
	OrderID   uint            `json:"orderId"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"-" gorm:"index"`
}
