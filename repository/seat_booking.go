package repository

import (
	"gorm.io/gorm"
	"time"
)

type SeatBooking struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Seat         *Seat          `json:"seat" gorm:"foreignKey:SeatUUID"`
	SeatUUID     string         `json:"seatUUID" gorm:"size:36"`
	Schedule     *Schedule      `json:"schedule" gorm:"foreignKey:ScheduleUUID"`
	ScheduleUUID string         `json:"scheduleUUID" gorm:"size:36"`
	Booked       bool           `json:"booked"`
	Canceled     bool           `json:"canceled"`
	BookedAt     time.Time      `json:"bookedAt"`
	CanceledAt   time.Time      `json:"canceledAt"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}
