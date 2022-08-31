package repository

import (
	"gorm.io/gorm"
	"time"
)

type PerformanceCast struct {
	ID            uint            `json:"id" gorm:"primarykey"`
	Performance   *Performance    `json:"performance" gorm:"foreignkey:PerformanceID"`
	PerformanceID uint            `json:"-"`
	Cast          *Cast           `json:"cast" gorm:"foreignkey:CastID"`
	CastID        uint            `json:"-"`
	CreatedAt     time.Time       `json:"createdAt"`
	UpdatedAt     time.Time       `json:"updatedAt"`
	DeletedAt     *gorm.DeletedAt `json:"-" gorm:"index"`
}

type PerformanceCastRepository struct {
	DB *gorm.DB
}
