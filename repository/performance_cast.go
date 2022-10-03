package repository

import (
	"gorm.io/gorm"
)

type PerformanceCast struct {
	Performance   *Performance `json:"performance" gorm:"foreignkey:PerformanceID"`
	PerformanceID uint         `json:"-"`
	Cast          *Cast        `json:"cast" gorm:"foreignkey:CastID"`
	CastID        uint         `json:"-"`
}

type PerformanceCastRepository struct {
	DB *gorm.DB
}
