package repository

import (
	"gorm.io/gorm"
	"time"
)

type ScheduleCast struct {
	Schedule     *Schedule       `json:"schedule" gorm:"foreignkey:ScheduleUUID"`
	ScheduleUUID string          `json:"-"`
	Cast         *Cast           `json:"cast" gorm:"foreignkey:CastID"`
	CastID       uint            `json:"-"`
	CreatedAt    time.Time       `json:"createdAt"`
	UpdatedAt    time.Time       `json:"updatedAt"`
	DeletedAt    *gorm.DeletedAt `json:"-" gorm:"index"`
}
