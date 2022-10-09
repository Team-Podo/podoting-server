package repository

import (
	"gorm.io/gorm"
	"time"
)

type OrderDetail struct {
	ID             uint            `json:"id" gorm:"primary_key"`
	OrderDetailKey string          `json:"orderDetailKey" gorm:"unique_index;not null;size:12"`
	OrderID        uint            `json:"orderId"`
	SeatBooking    *SeatBooking    `json:"-"`
	SeatBookingID  uint            `json:"seatBookingId"`
	OriginalPrice  uint            `json:"originalPrice"`
	Discount       uint            `json:"discount" gorm:"default:0"`
	Canceled       bool            `json:"canceled" gorm:"default:false"`
	CanceledAt     *time.Time      `json:"canceledAt"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      time.Time       `json:"updatedAt"`
	DeletedAt      *gorm.DeletedAt `json:"-" gorm:"index"`
}
