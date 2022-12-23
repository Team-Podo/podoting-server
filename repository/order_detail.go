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

type OrderDetailRepository struct {
	DB *gorm.DB
}

func (r *OrderDetailRepository) Save(detail *OrderDetail) error {
	err := r.DB.Save(detail).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderDetailRepository) FindByID(id uint) *OrderDetail {
	var orderDetail *OrderDetail
	err := r.DB.
		Joins("SeatBooking").
		First(&orderDetail, id).
		Error
	if err != nil {
		return nil
	}

	return orderDetail
}

func (r *OrderDetailRepository) CancelOrderDetail(detail *OrderDetail) error {
	now := time.Now()

	err := r.DB.
		Model(OrderDetail{}).
		Where("id = ?", detail.ID).
		Updates(OrderDetail{
			Canceled:   true,
			CanceledAt: &now,
		}).Error

	if err != nil {
		return err
	}

	err = r.DB.Model(SeatBooking{}).
		Where("id = ?", detail.SeatBookingID).
		Updates(SeatBooking{
			Booked:     false,
			Canceled:   true,
			CanceledAt: &now,
		}).Error

	if err != nil {
		return err
	}

	return nil
}
