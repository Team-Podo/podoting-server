package repository

import (
	"gorm.io/gorm"
	"time"
)

type Character struct {
	ID        uint     `json:"id" gorm:"primarykey"`
	Product   *Product `json:"product" gorm:"foreignkey:ProductID"`
	ProductID uint     `json:"-"`
	Name      string
	age       int
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"-" gorm:"index"`
}
