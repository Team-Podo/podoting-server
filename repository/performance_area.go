package repository

import (
	"gorm.io/gorm"
	"time"
)

type PerformanceArea struct {
	ID            uint            `json:"id" gorm:"primarykey"`
	Performance   *Performance    `json:"performance" gorm:"foreignkey:PerformanceID"`
	PerformanceID uint            `json:"-"`
	Area          *Cast           `json:"area" gorm:"foreignkey:AreaID"`
	AreaID        uint            `json:"-"`
	CreatedAt     time.Time       `json:"createdAt"`
	UpdatedAt     time.Time       `json:"updatedAt"`
	DeletedAt     *gorm.DeletedAt `json:"-" gorm:"index"`
}
